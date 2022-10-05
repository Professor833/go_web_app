package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Professor833/go_web_app/pkg/config"
	"github.com/Professor833/go_web_app/pkg/handlers"
	"github.com/Professor833/go_web_app/pkg/render"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
