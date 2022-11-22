package handler

import (
	"context"
	pb "github.com/ivanovamir/grpc-besm/gen/proto"
)

func (h *Handler) DBUpdate(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return h.service.Update(ctx, req), nil
}
