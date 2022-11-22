package handler

import (
	"github.com/ivanovamir/grpc-besm/api/internal/service"
	pb "github.com/ivanovamir/grpc-besm/gen/proto"
)

//type HHandler interface {
//	DBUpdate(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error)
//}

type Handler struct {
	pb.UnimplementedDBServer
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		UnimplementedDBServer: pb.UnimplementedDBServer{},
		service:               service,
	}
}
