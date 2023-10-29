package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vibhushit/webapp/pkg/config"
	"github.com/vibhushit/webapp/pkg/handlers"
	"github.com/vibhushit/webapp/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Server listening on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
