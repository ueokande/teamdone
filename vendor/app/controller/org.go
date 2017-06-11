package controller

import (
	"net/http"
)

func (c *Context) OrgGet(key string, w http.ResponseWriter, r *http.Request) {
	c.r.Render(w, "org.html", nil)
}
