package repository

import (
	"context"
	pb "github.com/ivanovamir/grpc-besm/gen/proto"
	"gorm.io/gorm"
)

type Repository struct {
	Product
}

type Product interface {
	Update(ctx context.Context, req *pb.UpdateRequest) *pb.UpdateResponse
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Product: NewProductRepository(db),
	}
}
