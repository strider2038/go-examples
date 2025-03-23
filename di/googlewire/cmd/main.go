package main

import (
	"log"

	"googlewire/di"
)

func main() {
	server, err := di.InitializeServer()
	if err != nil {
		log.Fatal(err)
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
