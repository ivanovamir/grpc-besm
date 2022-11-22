package service

import (
	"context"
	"github.com/ivanovamir/grpc-besm/api/internal/repository"
	pb "github.com/ivanovamir/grpc-besm/gen/proto"
)

type Service struct {
	Product
}

type Product interface {
	Update(ctx context.Context, req *pb.UpdateRequest) *pb.UpdateResponse
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Product: NewProductService(repo),
	}
}
