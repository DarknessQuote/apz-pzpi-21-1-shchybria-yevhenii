package main

import (
	"devquest-server/config"
	"devquest-server/devquest/infrastructure"
	"devquest-server/devquest/infrastructure/postgres"
	"devquest-server/devquest/server"
	"devquest-server/devquest/server/chiServer"
	"log"
)

func main() {
	var server server.Server
	var db infrastructure.Database

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err = postgres.NewPostgresDB(conf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Postgres")
	defer db.GetDB().Close()

	server = chiServer.NewChiServer(conf, &db)
	server.Start()
}