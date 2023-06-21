package main

import (
	"UmbrellaTask/internal/app/server"
	"UmbrellaTask/pkg/tools"
)

func main() {
	// TODO defer recovery()

	httpServer, err := server.New()
	tools.FatalError(err)

	httpServer.Start()
}
