package main

import (
	"net/http"

	"github.com/isnastish/nibble/pkg/log"
)

func helloRoute() {

}

func main() {
	log.Logger.Info("Listeing on port %s", ":3030")

	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Logger.Fatal("Failed to server on port: %s", ":3030")
	}
}
