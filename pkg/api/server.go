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
	router.HandleFunc("/login", server.loginRoute).Methods("POST")

	return server, nil
}

func (s *Server) Serve() error {
	log.Logger.Info("Listening on port %d", s.port)

	if err := http.ListenAndServe(":3030", nil); err != http.ErrServerClosed {
		return fmt.Errorf("Failed to listen")
	}

	return nil
}

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
