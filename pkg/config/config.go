package config

// App configuration is used for application global configuration variables and settings, and should replace using global variables as a superior method.
// This package should not import anything other than things from the standard library.
// It should only import the absolute minimum necessary to get the job done.
// It should NOT import any local packages

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	InProduction bool
	Session *scs.SessionManager
}