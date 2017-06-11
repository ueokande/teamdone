package main

import (
	"app/controller"
	"app/render"
	"app/route"
	"app/shared/csrf"
	"app/shared/database"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func run() int {
	rand.Seed(time.Now().UnixNano())

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	conf, err := database.LoadConfig("config/development.json")
	if err != nil {
		logger.Print("Failed to load db config:", err)
	}
	db, err := database.Connect(conf)
	if err != nil {
		logger.Print("Failed to connect to db:", err)
		return 1
	}

	cc := controller.NewContext(db)
	web := route.WebHandler{C: cc}
	api := route.ApiHandler{C: cc}

	render.InitTemplateRenderer(template.Must(template.ParseGlob("template/*.html")))

	mux := http.NewServeMux()
	mux.Handle("/", csrf.DefaultCSRF(web))
	mux.Handle("/i/", csrf.DefaultCSRF(http.StripPrefix("/i", api)))
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
