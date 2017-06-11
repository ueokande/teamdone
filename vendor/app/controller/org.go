package controller

import (
	"app/render"
	"net/http"
)

func (c *Context) OrgGet(key string, w http.ResponseWriter, r *http.Request) {
	render.Render(w, "org.html", nil)
}
