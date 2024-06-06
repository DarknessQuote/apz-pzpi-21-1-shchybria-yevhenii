package chiServer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getRoutes() http.Handler {
	mux := chi.NewMux()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	return mux
}