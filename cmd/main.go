package main

import (
	"log"

	"github.com/go-list-templ/grpc/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
