package controller

import (
	"net/http"
)

func (c *Context) SettingsGet(w http.ResponseWriter, r *http.Request) {
	c.r.Render(w, "settings.html", nil)
}
