package main

import (
	"devquest-server/config"
	"devquest-server/devquest/server"
	"devquest-server/devquest/server/chiServer"
	"log"
)

func main() {
	var server server.Server

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	server = chiServer.NewChiServer(conf)
	server.Start()
}