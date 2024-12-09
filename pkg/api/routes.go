package api

import (
	"io"
	"net/http"
)

func (s *Server) helloRoute(respWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(respWriter, "Hello world!")
}
