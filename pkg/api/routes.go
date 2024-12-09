package api

import (
	"io"
	"net/http"
)

func helloRoute(respWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(respWriter, "Hello world!")
}
