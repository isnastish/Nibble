package main

import (
	"flag"

	"github.com/isnastish/nibble/pkg/api"
	"github.com/isnastish/nibble/pkg/log"
)

func main() {
	port := flag.Int("port", 3030, "Listening port")
	flag.Parse()

	apiServer, err := api.NewServer(*port)
	if err != nil {
		log.Logger.Fatal("Faied to create api server: %s", err.Error())
	}

	go func() {
		if err := apiServer.Serve(); err != nil {
			log.Logger.Fatal("Failed to server: %s", err.Error())
		}
	}()

	apiServer.Shutdown() // closes db connection, as well as shutsdown the http server
}
