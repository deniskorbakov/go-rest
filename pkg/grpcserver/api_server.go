package grpcserver

import (
	"net"

	"github.com/go-list-templ/grpc/config"
	"google.golang.org/grpc"
)

type ApiServer struct {
	server *grpc.Server
	config *config.Server
	errors chan error
}

func NewApiServer(cfg *config.Server) *ApiServer {
	grpcServer := grpc.NewServer()

	return &ApiServer{
		server: grpcServer,
		config: cfg,
		errors: make(chan error, 1),
	}
}

func (s *ApiServer) Notify() <-chan error {
	return s.errors
}

func (s *ApiServer) Start() {
	go func() {
		lis, err := net.Listen("tcp", net.JoinHostPort("", s.config.GRPCPort))
		if err != nil {
			s.errors <- err
		}

		s.errors <- s.server.Serve(lis)
		close(s.errors)
	}()
}

func (s *ApiServer) Stop() {
	s.server.GracefulStop()
}
