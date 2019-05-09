package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

type Application struct {
	Title       string       `yaml:"title"`       // Title stores the name of the application
	Version     string       `yaml:"version"`     // Version stores a semantic version of the application
	Maintainers []Maintainer `yaml:"maintainers"` // Maintainers stores the application's maintainers
	Company     string       `yaml:"company"`     // Company stores the company that produced the application
	Website     string       `yaml:"website"`     // Websites stores the http://some-url-here.com of the application
	Source      string       `yaml:"source"`      // Source stores the GIT/VCS repository of the application
	License     string       `yaml:"license"`     // License stores the license under which the application was released
	Description string       `yaml:"description"` // Description stores a brief explanation of the application's intended uses and objectives
}

type ApplicationWithID struct {
	ID          string
	Application Application
}

type Maintainer struct {
	Name  string `yaml:"name"`  // Name stores the full name of the maintainer
	Email string `yaml:"email"` // Email stores the email address of the maintainer
}

func (app Application) Validate() error {
	return validation.ValidateStruct(&app,
		// Title is required and cannot be empty
		validation.Field(&app.Title, validation.Required, validation.NilOrNotEmpty),
		// Version is required and cannot be empty, and must be a semantic version
		validation.Field(&app.Version, validation.Required, validation.NilOrNotEmpty, is.Semver),
		// Maintainers is required and cannot be empty
		validation.Field(&app.Maintainers, validation.Required, validation.NilOrNotEmpty),
		//  Company is required and cannot be empty
		validation.Field(&app.Company, validation.Required, validation.NilOrNotEmpty),
		// Website is required and cannot be empty and must be a URL
		validation.Field(&app.Website, validation.Required, validation.NilOrNotEmpty, is.URL),
		// Source is required and cannot be empty and must be a URL
		validation.Field(&app.Source, validation.Required, validation.NilOrNotEmpty, is.URL),
		// License is required and cannot be empty
		validation.Field(&app.License, validation.Required, validation.NilOrNotEmpty),
		// Description is required and cannot be empty
		validation.Field(&app.Description, validation.Required, validation.NilOrNotEmpty),
	)
}

func (maintainer Maintainer) Validate() error {
	return validation.ValidateStruct(&maintainer,
		// Name is required and cannot be empty
		validation.Field(&maintainer.Name, validation.Required, validation.NilOrNotEmpty),
		// Email is required and cannot be empty and must be an Email
		validation.Field(&maintainer.Email, validation.Required, validation.NilOrNotEmpty, is.Email),
	)
}

func Unmarshal(w http.ResponseWriter, r *http.Request) (app Application, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return app, err
	}

	yaml.Unmarshal(body, &app)

	err = validation.Validate(app)
	if err != nil {
		return app, err
	}
	log.Print(app)
	return app, nil
}
