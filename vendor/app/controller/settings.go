package controller

import (
	"app/render"
	"net/http"
)

func (c *Context) SettingsGet(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "settings.html", nil)
}
