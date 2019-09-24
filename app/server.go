package app

import (
	"fmt"
	"github.com/go-chi/chi"
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

	router.Get("/", rootHandler)

	return router
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello World")
	if err != nil {
		panic(err)
	}
}
