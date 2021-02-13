package adder

import (
	context "context"
	"grpc/pkg/api"
)

type AdderServer struct {
	api.UnimplementedAdderServer
	x string
}

func (a *AdderServer) Add(ctx context.Context, request *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: request.GetX() + request.GetY()}, nil
}
