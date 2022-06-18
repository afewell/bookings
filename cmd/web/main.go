package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/afewell/bookings/pkg/config"
	"github.com/afewell/bookings/pkg/handlers"
	"github.com/afewell/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// The OsInterruptHandler function is used to handle the interrupt signal to ensure our program exits cleanly.
	// This is the first line in our program to ensure interrupts can be handled at any point.
	handlers.OsInterruptHandler()
	
	

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session


	// Initialize the template cache at app launch
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create templae cache: ", err)
	}

	app.TemplateCache = tc

	app.UseCache = true
	
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Server is listening on port%s", portNumber)

	srv := &http.Server{Addr: portNumber, Handler: routes(&app)}

	err = srv.ListenAndServe()

	log.Fatal(err)

}