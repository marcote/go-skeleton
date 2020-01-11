package main

import (
	"github.com/marcote/go-skeleton/routing"
	"log"
)

func main() {
	routing.Configure()
	routing.Register()

	if err := run("8080"); err != nil {
		log.Println("error running server", err)
	}
}

func run(port string) error {
	return routing.Router.Run(":" + port)
}
