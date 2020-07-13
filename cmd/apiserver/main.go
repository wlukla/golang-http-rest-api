package main

import (
	"golang-http-rest-api/internal/app/apiserver"
	"log"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
