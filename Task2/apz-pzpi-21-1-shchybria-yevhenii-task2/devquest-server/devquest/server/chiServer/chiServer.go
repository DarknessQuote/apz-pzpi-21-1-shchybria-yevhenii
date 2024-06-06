package chiServer

import (
	"devquest-server/config"
	"fmt"
	"log"
	"net/http"
)

type chiServer struct {
	conf *config.Config
}

func NewChiServer(config *config.Config) *chiServer {
	return &chiServer {
		conf: config,
	}
}

func (s *chiServer) Start() {
	port := s.conf.Server.Port
	serverUrl := fmt.Sprintf(":%d", port)
	router := getRoutes()
	
	log.Printf("Starting application on port %d", port)
	if err := http.ListenAndServe(serverUrl, router); err != nil {
		log.Fatal(err)
	}
}