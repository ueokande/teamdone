package main

import (
	"app/route"
	"app/shared/database"
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

	mux := http.NewServeMux()
	mux.Handle("/", route.WebHandler{})
	mux.Handle("/i/", http.StripPrefix("/i", route.ApiHandler{}))

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