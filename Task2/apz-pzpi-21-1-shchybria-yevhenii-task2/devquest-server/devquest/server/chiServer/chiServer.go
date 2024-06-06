package chiServer

import (
	"devquest-server/config"
	"devquest-server/devquest/infrastructure"
	"fmt"
	"log"
	"net/http"
)

type chiServer struct {
	config *config.Config
	database *infrastructure.Database
}

func NewChiServer(conf *config.Config, db *infrastructure.Database) *chiServer {
	return &chiServer {
		config: conf,
		database: db,
	}
}

func (s *chiServer) Start() {
	port := s.config.Server.Port
	serverUrl := fmt.Sprintf(":%d", port)
	router := getRoutes()
	
	log.Printf("Starting application on port %d", port)
	if err := http.ListenAndServe(serverUrl, router); err != nil {
		log.Fatal(err)
	}
}