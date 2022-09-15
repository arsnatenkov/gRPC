package adder

import (
	"context"
	"gRPC/pkg/api/pkg/api"
)

type GRPCServer struct {
}

func (s *GRPCServer) Add(ctx context.Context, request *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: request.GetX() + request.GetY()}, nil
}
