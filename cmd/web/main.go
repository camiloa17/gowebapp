package main

import (
	"log"
	"net/http"

	"github.com/camiloa17/gowebapp/pkg/config"
	"github.com/camiloa17/gowebapp/pkg/handlers"
	"github.com/camiloa17/gowebapp/pkg/render"
)

const portNumber = ":8080"

// main is the main application handler
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Listening on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
