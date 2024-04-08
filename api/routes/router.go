package routes

import (
	"github.com/Zubayear/holiday/api/handlers"
	"github.com/Zubayear/holiday/api/presenter"
	jobs2 "github.com/Zubayear/holiday/pkg/jobs"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)

func Router(service jobs2.Service) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))

	r.Get("/", templ.Handler(presenter.Home()).ServeHTTP)
	r.Get("/jobs", handlers.GetJobs(service))

	return r
}
