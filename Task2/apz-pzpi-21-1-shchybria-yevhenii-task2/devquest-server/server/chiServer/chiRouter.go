package chiServer

import (
	"devquest-server/devquest/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getRoutes() http.Handler {
	mux := chi.NewMux()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	mux.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		authSettings := &handlers.Auth{Auth: GetChiServer().AuthSettings}
		authSettings.Login(w, r)
	})

	return mux
}