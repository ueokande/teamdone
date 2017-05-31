package csrf

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var empty = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

func cookieByName(cookies []*http.Cookie, name string) *http.Cookie {
	for _, c := range cookies {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func TestServeHTTP_Get(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	DefaultCSRF(empty).ServeHTTP(rec, req)

	res := rec.Result()
	cookie := cookieByName(res.Cookies(), "_request_token")
	if cookie == nil {
		t.Fatal("Unexpected response:", *res)
	}
	if len(cookie.Value) != 32 {
		t.Fatal("Unexpected CSRF token:", cookie.Value)
	}
}

func TestServeHTTP_ValidPost(t *testing.T) {
	cases := []struct {
		token  string
		set    bool
		status int
	}{
		{"secret-csrf-token", true, http.StatusOK},
		{"", true, http.StatusForbidden},
		{"", false, http.StatusForbidden},
	}

	for _, c := range cases {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Add("Cookie", "_request_token=secret-csrf-token")
		if c.set {
			req.Header.Set("X-Request-Token", c.token)
		}
		rec := httptest.NewRecorder()

		DefaultCSRF(empty).ServeHTTP(rec, req)

		res := rec.Result()
		if res.StatusCode != c.status {
			t.Fatal("Unexpected response:", *res)
		}
	}
}
