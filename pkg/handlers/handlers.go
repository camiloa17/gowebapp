package handlers

import (
	"net/http"

	"github.com/camiloa17/gowebapp/pkg/config"
	"github.com/camiloa17/gowebapp/pkg/render"
)

// Repository pattern
var Repo *Repository

// Repository is teh repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repository *Repository) {
	Repo = repository
}

// Home is the home page handler
func (repo *Repository) Home(responseWriter http.ResponseWriter, request *http.Request) {

	render.RenderTemplate(responseWriter, "home.page.gohtml")
}

// About is the about page handler
func (repo *Repository) About(responseWriter http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(responseWriter, "about.page.gohtml")
}
