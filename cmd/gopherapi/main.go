package main

import (
	"log"
	"net/http"

	server "./pkg/server"
)

func main() {

	s := server.New()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
