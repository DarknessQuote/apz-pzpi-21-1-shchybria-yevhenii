package chiServer

import (
	"devquest-server/server/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)
func getRoutes() http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.EnableCORS)
	
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	mux.Route("/auth", func(r chi.Router) {
		authSettings := GetChiServer().AuthSettings

		r.Post("/login", authHttpHandler.Login(authSettings))
		r.Post("/register", authHttpHandler.Register(authSettings))
		r.Delete("/logout", authHttpHandler.Logout(authSettings))
	})
	
	mux.Route("/companies", func(r chi.Router) {
		r.Get("/", companyHttpHandler.GetAllCompanies)
		r.Get("/{id}", companyHttpHandler.GetCompanyByID)
		r.Post("/", companyHttpHandler.AddCompany)
		r.Put("/{id}", companyHttpHandler.UpdateCompany)
		r.Delete("/{id}", companyHttpHandler.DeleteCompany)
	})

	mux.Route("/projects", func(r chi.Router) {
		r.Get("/manager/{manager_id}", projectHttpHandler.GetProjectsOfManager)
		r.Get("/developer/{developer_id}", projectHttpHandler.GetProjectsOfDeveloper)
		
		r.Put("/{id}", projectHttpHandler.UpdateProject)
		r.Post("/", projectHttpHandler.AddProject)
		r.Delete("/{id}", projectHttpHandler.DeleteProject)

		r.Get("/developers/{project_id}", projectHttpHandler.GetProjectDevelopers)
		r.Post("/developers", projectHttpHandler.AddDeveloperToProject)
		r.Delete("/developers", projectHttpHandler.RemoveDeveloperFromProject)
	})

	mux.Route("/achievements", func(r chi.Router) {
		r.Get("/project/{project_id}", achievementHttpHandler.GetProjectAchievements)
		r.Get("/developer/{developer_id}", achievementHttpHandler.GetDeveloperAchievements)
		r.Post("/{project_id}", achievementHttpHandler.AddAchievementToProject)
		r.Put("/{id}", achievementHttpHandler.UpdateAchievement)
		r.Delete("/{id}", achievementHttpHandler.DeleteAchievement)
		r.Post("/give", achievementHttpHandler.GiveAchievementToDeveloper)
	})

	return mux
}