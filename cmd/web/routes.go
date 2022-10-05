package main

import (
	"net/http"

	"github.com/Professor833/go_web_app/pkg/config"
	"github.com/Professor833/go_web_app/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New() // Create a new multiplexer
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
