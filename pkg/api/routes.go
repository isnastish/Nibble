package api

import (
	"io"
	"net/http"
)

func (s *Server) signupRoute(respWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(respWriter, "Hello from signup route!")
}

func (s *Server) loginRoute(respWriter http.ResponseWriter, req *http.Request) {
	io.WriteString(respWriter, "Hello from login route")
}
