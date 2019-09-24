package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func StartServer(port string) error {
	router := routes()

	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
	return nil
}

func routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", rootHandler)

	return router
}
