package controller

import (
	"app/render"
	"fmt"
	"net/http"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
	err := render.Render(w, "home.html", nil)
	if err != nil {
		fmt.Fprintf(w, "error : %s", err)
	}
}
