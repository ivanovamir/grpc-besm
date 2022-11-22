package server

import (
	"fmt"
	"github.com/ivanovamir/grpc-besm/api/internal/handler"
	pb "github.com/ivanovamir/grpc-besm/gen/proto"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	GrpcHandler *handler.Handler
	srv         *grpc.Server
}

func (s *GrpcServer) NewServer(handler *handler.Handler) *GrpcServer {
	return &GrpcServer{
		srv:         grpc.NewServer(),
		GrpcHandler: handler,
	}
}

func (s *GrpcServer) ListenAndServe(connType, port string) error {
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen(connType, addr)
	if err != nil {
		return err
	}
	pb.RegisterDBServer(s.srv, s.GrpcHandler)

	if err := s.srv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *GrpcServer) Stop() {
	s.srv.GracefulStop()
}
