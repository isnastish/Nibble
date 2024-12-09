package main

import (
	"github.com/isnastish/nibble/pkg/api"
	"github.com/isnastish/nibble/pkg/log"
)

func main() {
	apiServer, err := api.NewServer(3030)
	if err != nil {
		log.Logger.Fatal("Faied to create api server: %s", err.Error())
	}

	go func() {
		// TODO: Pass port to Serve or NewServer?
		if err := apiServer.Serve(); err != nil {
			log.Logger.Fatal("Failed to server: %s", err.Error())
		}
	}()
}
