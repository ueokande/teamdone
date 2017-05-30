package main

import (
	"app/render"
	"app/route"
	"app/shared/csrf"
	"app/shared/database"
	"html/template"
	"log"
	"net/http"
	"os"
)

func run() int {
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	db, err := database.LoadConfig("config/development.json")
	if err != nil {
		logger.Print("Failed to load db config:", err)
	}
	err = database.Connect(db)
	if err != nil {
		logger.Print("Failed to connect to db:", err)
		return 1
	}

	render.InitTemplateRenderer(template.Must(template.ParseGlob("template/*.html")))

	mux := http.NewServeMux()
	mux.Handle("/", csrf.DefaultCSRF(route.WebHandler{}))
	mux.Handle("/i/", csrf.DefaultCSRF(http.StripPrefix("/i", route.ApiHandler{})))
	mux.HandleFunc("/assets/index.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "build/index.js")
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Print(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
