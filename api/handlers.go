package api

import (
	"github.com/apaz037/go-metadata-api/api/handlers"
	"net/http"
)

// Handlers are closures to encourage testability
func getApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.GetApplicationHandler(w, r)
	}
}

func createApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateApplicationHandler(w, r)
	}
}

func deleteApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteApplicationHandler(w, r)
	}
}

func updateApplication() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateApplicationHandler(w, r)
	}
}

func getAllApplications() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllApplicationsHandler(w, r)
	}
}

func searchApplications() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlers.SearchApplicationHandler(w, r)
	}
}
