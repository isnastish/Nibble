package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/isnastish/nibble/pkg/db"
	"github.com/isnastish/nibble/pkg/ipresolver"
	"github.com/isnastish/nibble/pkg/log"
)

// Server for handling http request.
// It bundles all the pieces together, the ip resolver client,
// and postgres client to store user's data in a persistent storage.
type Server struct {
	// http server
	*http.Server
	// ip resolver client
	ipResolverClient *ipresolver.Client
	// database connector
	db *db.PostgresDB
	// settings
	port int
}

func NewServer(port int) (*Server, error) {
	db, err := db.NewPostgresDB()
	if err != nil {
		return nil, err
	}

	ipResolverClient, err := ipresolver.NewClient()
	if err != nil {
		return nil, err
	}

	server := &Server{
		Server:           &http.Server{Addr: fmt.Sprintf(":%d", port)},
		ipResolverClient: ipResolverClient,
		db:               db,
		port:             port,
	}

	router := mux.NewRouter()

	// add logging middleware
	router.Use(loggingMiddleware)

	// bind routes
	router.HandleFunc("/signup", server.signupRoute).Methods("POST")
	router.HandleFunc("/users", server.getUsers).Methods("GET")

	server.Server.Handler = router

	return server, nil
}

// Start listening for incoming requests.
// This function should be run in a separate goroutine to
// prevent blocking the main execution thread.
// Returns an error if failed to listen on earlier configured port.
func (s *Server) Serve() error {
	log.Logger.Info("Listening on port %d", s.port)

	if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("Failed to listen")
	}

	return nil
}

// Gracefully shutdown the server,
// by closing database connection first.
// Returns an error if something goes wrong.
func (s *Server) Shutdown() error {
	// Close db connection
	defer s.db.Close()

	// Shutdown the http server
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
