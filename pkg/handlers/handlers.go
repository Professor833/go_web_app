package handlers

import (
	"net/http"

	"github.com/Professor833/go_web_app/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page!")
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
