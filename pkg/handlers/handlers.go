package handlers

import (
	"net/http"

	"github.com/camiloa17/gowebapp/pkg/render"
)

// Home is the home page handler
func Home(responseWriter http.ResponseWriter, request *http.Request) {

	render.RenderTemplate(responseWriter, "home.page.gohtml")
}

// About is the about page handler
func About(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, "about.page.gohtml")
}
