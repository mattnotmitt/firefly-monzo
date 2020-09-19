package main

import (
	"log"

	"github.com/mattnotmitt/firefly-monzo/internal/serve"
)

func main() {
	log.Print("firefly-monzo sync daemon starting")

	serve.Serve()
	
}
