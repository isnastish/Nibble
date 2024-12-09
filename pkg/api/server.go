package api

import (
	"fmt"
	"net/http"

	"github.com/isnastish/nibble/pkg/ipresolver"
	"github.com/isnastish/nibble/pkg/log"
)

type Server struct {
	// http server
	*http.Server
	// ip resolver client
	IpResolverClient *ipresolver.Client
	// database connector
	// port, addr, other settings ...
}

func NewServer(port int) (*Server, error) {
	// TODO: Create http server
	return &Server{}, nil
}

func (s *Server) Serve() error {
	log.Logger.Info("Listening on port 3030")

	if err := http.ListenAndServe(":3030", nil); err != http.ErrServerClosed {
		// TODO: More robust error message
		return fmt.Errorf("Failed to listen")
	}

	return nil
}

func (s *Server) Shutdown() error {
	return nil
}
