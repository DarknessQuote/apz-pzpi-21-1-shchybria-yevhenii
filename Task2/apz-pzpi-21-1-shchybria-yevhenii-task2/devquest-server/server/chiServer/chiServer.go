package chiServer

import (
	"devquest-server/config"
	"devquest-server/devquest/infrastructure"
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
	port := s.Config.Server.Port
	serverUrl := fmt.Sprintf(":%d", port)
	router := getRoutes()
	
	log.Printf("Starting application on port %d", port)
	if err := http.ListenAndServe(serverUrl, router); err != nil {
		log.Fatal(err)
	}
}