package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

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

	osSigChan := make(chan os.Signal, 1)
	signal.Notify(osSigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := apiServer.Serve(); err != nil {
			log.Logger.Fatal("Failed to server: %s", err.Error())
		}
	}()

	<-osSigChan
	if err := apiServer.Shutdown(); err != nil {
		log.Logger.Fatal("Failed to shutdown the server: %s", err.Error())
	}

	log.Logger.Info("Gracefully shutdown the server")
}
