package http

import (
	"context"
	"net/http"
	"project/internal/service"
)

type Server struct {
	server      *http.Server
	userService service.UserServiceI
}

func NewServer(addr string, userService service.UserServiceI) *Server {
	srv := &Server{
		userService: userService,
	}

	mux := http.NewServeMux()
	mux.Handle("/users/", srv.HandleUserByID())

	srv.server = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return srv
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
