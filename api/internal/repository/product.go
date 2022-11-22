package repository

import (
	"context"
	pb "github.com/ivanovamir/grpc-besm/gen/proto"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Update(ctx context.Context, req *pb.UpdateRequest) *pb.UpdateResponse {
	if req.Link == "" {
		return &pb.UpdateResponse{Error: true}
	} else {
		return &pb.UpdateResponse{Error: false}
	}
}
