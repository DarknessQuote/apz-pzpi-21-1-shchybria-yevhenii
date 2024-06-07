package chiServer

import (
	"devquest-server/devquest/handlers"
	"devquest-server/devquest/infrastructure/test"
	"devquest-server/devquest/usecases"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiMux struct {
	*chi.Mux
}

func getRoutes() http.Handler {
	mux := &chiMux{chi.NewMux()}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	mux.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		authSettings := &handlers.Auth{Auth: GetChiServer().AuthSettings}
		authSettings.Login(w, r)
	})
	mux.InitializeCompanyHttpHandler()

	return mux
}

func (m *chiMux) InitializeCompanyHttpHandler() {
	companyRepository := test.NewCompanyTestRepo()
	companyUsecase := usecases.NewCompanyUsecase(companyRepository)
	companyHandler := handlers.NewCompanyHttpHandler(*companyUsecase)

	m.Route("/companies", func(r chi.Router) {
		r.Get("/", companyHandler.GetAllCompanies)
		r.Get("/{id}", companyHandler.GetCompanyByID)
		r.Post("/", companyHandler.AddCompany)
		r.Put("/{id}", companyHandler.UpdateCompany)
		r.Delete("/{id}", companyHandler.DeleteCompany)
	})
}