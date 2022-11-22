package service

import (
	"context"
	"github.com/ivanovamir/grpc-besm/api/internal/repository"
	pb "github.com/ivanovamir/grpc-besm/gen/proto"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}
func (s *ProductService) Update(ctx context.Context, req *pb.UpdateRequest) *pb.UpdateResponse {
	return s.repo.Update(ctx, req)
}
