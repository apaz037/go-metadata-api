package api

import (
	"github.com/apaz037/go-metadata-api/db"
	"github.com/apaz037/go-metadata-api/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
	"github.com/unrolled/secure"
	"time"
)

func New() *chi.Mux {
	router := chi.NewRouter()                         // instantiate our router
	db.Database = make(map[string]models.Application) // in memory database of applications
	security := secure.New(secure.Options{
		FrameDeny:   true,
		SSLRedirect: true,
		SSLHost:     viper.GetString("port"),
	})

	// Middlewares
	router.Use(security.Handler)
	router.Use(middleware.RequestID)                          // assign each request a unique ID
	router.Use(middleware.RealIP)                             // monitor forwarded requests from origin IP
	router.Use(middleware.Recoverer)                          // restart on panic
	router.Use(middleware.Logger)                             // log requests with detail in console
	router.Use(middleware.Timeout(15 * time.Second))          // default request timeout
	router.Use(render.SetContentType(render.ContentTypeJSON)) // response type will always be json
	router.Use(middleware.Heartbeat("/health"))               // health endpoint for monitoring

	// Routes
	router.Get("/application", getAllApplications())         // retrieve all applications
	router.Post("/application", createApplication())         // create a new application
	router.Get("/application/{id}", getApplication())        // retrieve an application
	router.Delete("/application/{id}", deleteApplication())  // delete an application
	router.Put("/application/{id}", updateApplication())     // update an application
	router.Get("/application/search/", searchApplications()) // search applications

	return router
}
