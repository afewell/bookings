package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole writes session details to the console
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})


	return csrfHandler
}

// SessionLoad loads the session from the request and stores it in the request context.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}