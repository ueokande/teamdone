package main

import (
	"app/controller"
	"app/render"
	"app/route"
	"app/session"
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

	r := &render.TemplateRenderer{
		Template: template.Must(template.ParseGlob("template/*.html")),
	}
	sm := &session.Manager{
		CookieName: "session",
		Storage:    session.NewMembachedSessionStorage("localhost:11211", 30*24*time.Hour),
		LifeTime:   30 * 24 * time.Hour,
	}
	cc := controller.NewContext(db, sm, r)

	mux := http.NewServeMux()
	mux.Handle("/", csrf.DefaultCSRF(route.Web(cc)))
	mux.Handle("/i/", csrf.DefaultCSRF(http.StripPrefix("/i", route.Api(cc))))
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
