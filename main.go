package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/isnastish/nibble/pkg/log"
)

func helloRoute(respWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(respWriter, "Hello world!")
}

func main() {
	router := mux.NewRouter()
	_ = router

	http.HandleFunc("/hello", helloRoute)

	log.Logger.Info("Listeing on port %s", ":3030")

	if err := http.ListenAndServe(":3030", nil); err != http.ErrServerClosed {
		log.Logger.Fatal("Failed to server on port: %s", ":3030")
	}
}
