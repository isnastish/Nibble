package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/isnastish/nibble/pkg/db"
	"github.com/isnastish/nibble/pkg/ipresolver"
	"github.com/isnastish/nibble/pkg/log"
)

type Server struct {
	// http server
	*http.Server
	// ip resolver client
	ipResolverClient *ipresolver.Client
	// database connector
	db *db.PostgresDB
	// port, addr, other settings ...
}

func NewServer(port int) (*Server, error) {
	// TODO: Create http server
	db, err := db.NewPostgresDB()
	if err != nil {
	}

	server := &Server{
		db: db,
	}

	router := mux.NewRouter()

	// add logging middleware
	router.Use(loggingMiddleware)

	http.HandleFunc("/hello", server.helloRoute)

	return server, nil
}

func (s *Server) Serve() error {
	log.Logger.Info("Listening on port 3030")

	// TODO: Bind endpoints here or on the server creation?

	if err := http.ListenAndServe(":3030", nil); err != http.ErrServerClosed {
		// TODO: More robust error message
		return fmt.Errorf("Failed to listen")
	}

	return nil
}

func (s *Server) Shutdown() error {
	// Close db connection
	s.db.Close()
	// Shutdown the http server
	return nil
}
