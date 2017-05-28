package controller

import (
	"app/render"
	"fmt"
	"net/http"
)

func TopGet(w http.ResponseWriter, r *http.Request) {
	err := render.Render(w, "top.html", nil)
	if err != nil {
		fmt.Fprintf(w, "error : %s", err)
	}
}
