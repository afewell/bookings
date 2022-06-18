package handlers

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/afewell/bookings/pkg/config"
	"github.com/afewell/bookings/pkg/models"
	"github.com/afewell/bookings/pkg/render"
)

// Repo is the repository for handlers
var Repo *Repository


// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is a handler function that returns the home page.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)		
	
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is a handler function that returns the about page.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "You Wish!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP


	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}







func OsInterruptHandler() {
	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt)
    go func() {
        //select {
        sig := <-c
            fmt.Printf("Got %s signal.  Aborting...\n", sig)
            os.Exit(1)
        //}
    }()
}