package adder

import (
	"context"
	"testR/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

