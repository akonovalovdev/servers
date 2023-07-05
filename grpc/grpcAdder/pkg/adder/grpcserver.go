package adder

import (
	"context"
	api "github.com/akonovalovdev/servers/grpc/grpcAdder/pkg/api/genproto/adder"
)

// GRPCServer ...
type GRPCServer struct {
	api.UnimplementedAdderServer
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {

	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
