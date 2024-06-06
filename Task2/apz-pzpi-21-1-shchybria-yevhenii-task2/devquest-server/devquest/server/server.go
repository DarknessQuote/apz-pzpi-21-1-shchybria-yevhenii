package server

import (
	"log"
	"net/http"
)

func Start() {
	err := http.ListenAndServe(":8080", getRoutes())
	if err != nil {
		log.Fatal(err)
	}
}