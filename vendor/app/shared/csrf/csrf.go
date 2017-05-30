package csrf

import (
	"crypto/subtle"
	"math/rand"
	"net/http"
	"time"
)

type CSRFHandler struct {
	HeaderName   string
	CookieName   string
	TokenLength  int
	CookieMaxAge int
}

var DefaultCSRFHandler = CSRFHandler{
	TokenLength:  32,
	HeaderName:   "X-Request-Token",
	CookieName:   "_request_token",
	CookieMaxAge: 86400,
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func randomToken(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func DefaultCSRF(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		DefaultCSRFHandler.ServeHTTP(w, r)
		h.ServeHTTP(w, r)
	})
}

func (h *CSRFHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(h.CookieName)
	var token string
	if err != nil {
		token = randomToken(h.TokenLength)
	} else {
		token = cookie.Value
	}

	switch r.Method {
	case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
	default:
		reqtoken := r.Header.Get(h.HeaderName)
		if subtle.ConstantTimeCompare([]byte(token), []byte(reqtoken)) != 1 {
			http.Error(w, "invalid request token", http.StatusForbidden)
			return
		}
	}

	cookie = &http.Cookie{
		Name:    h.CookieName,
		Value:   token,
		Expires: time.Now().Add(time.Duration(h.CookieMaxAge) * time.Second),
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Vary", "Cookie")
}
