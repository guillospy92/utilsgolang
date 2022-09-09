package celeritas

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (a *Accelerator) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	if a.Config.debug {
		mux.Use(middleware.Logger)
	}

	mux.Use(middleware.Recoverer)
	mux.Use(a.SessionLoad)

	return mux
}
