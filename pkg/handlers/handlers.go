package handlers

import (
	"net/http"

	"github.com/camiloa17/gowebapp/pkg/config"
	"github.com/camiloa17/gowebapp/pkg/models"
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
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (r *Repository) Home(responseWriter http.ResponseWriter, request *http.Request) {

	render.RenderTemplate(responseWriter, "home.page.gohtml", &models.TemplateData{})
}

// About is the about page handler
func (r *Repository) About(responseWriter http.ResponseWriter, request *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	// send the template to the front end
	render.RenderTemplate(responseWriter, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})
}
