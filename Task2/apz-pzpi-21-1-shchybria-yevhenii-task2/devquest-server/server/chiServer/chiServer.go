package chiServer

import (
	"devquest-server/config"
	"devquest-server/devquest/handlers"
	"devquest-server/devquest/infrastructure"
	"devquest-server/devquest/infrastructure/postgres"
	"devquest-server/devquest/usecases"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type chiServer struct {
	Config *config.Config
	Database *infrastructure.Database
	AuthSettings *infrastructure.Auth
}

var (
	once sync.Once;
	serverInstance *chiServer
	authHttpHandler *handlers.AuthHttpHandler
	companyHttpHandler *handlers.CompanyHttpHandler
	projectHttpHandler *handlers.ProjectHttpHandler
)

func NewChiServer(conf *config.Config, db *infrastructure.Database, auth *infrastructure.Auth) *chiServer {
	once.Do( func() {
		serverInstance = &chiServer {
		Config: conf,
		Database: db,
		AuthSettings: auth,
	}
	})
	return serverInstance
}

func GetChiServer() *chiServer {
	return serverInstance
}

func (s *chiServer) Start() {
	initializeHttpHandlers()

	port := s.Config.Server.Port
	serverUrl := fmt.Sprintf(":%d", port)	
	router := getRoutes()
	
	log.Printf("Starting application on port %d", port)
	if err := http.ListenAndServe(serverUrl, router); err != nil {
		log.Fatal(err)
	}
}

func initializeHttpHandlers() {
	userRepository := postgres.NewUserPostgresRepo(*serverInstance.Database)
	companyRepository := postgres.NewCompanyPostgresRepo(*serverInstance.Database)
	projectRepository := postgres.NewProjectPostgresRepo(*serverInstance.Database)

	userUsecase := usecases.NewUserUsecase(userRepository, companyRepository)
	companyUsecase := usecases.NewCompanyUsecase(companyRepository)
	projectUsecase := usecases.NewProjectUsecase(projectRepository, userRepository, companyRepository)

	authHttpHandler = handlers.NewAuthHttpHandler(*userUsecase)
	companyHttpHandler = handlers.NewCompanyHttpHandler(*companyUsecase)
	projectHttpHandler = handlers.NewProjectHttpHandler(*projectUsecase)
}