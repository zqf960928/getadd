package grpc

import (
	"context"
	endpoint "github.com/zqf960928/getadd/pkg/endpoint"
	pb "github.com/zqf960928/getadd/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
)

// makeAddHandler creates the handler logic
func makeAddHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddEndpoint, decodeAddRequest, encodeAddResponse, options...)
}

// decodeAddResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Add request.
// TODO implement the decoder
func decodeAddRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.AddRequest)
	return pb.AddRequest{A: req.A, B: req.B}, nil

	//return nil, errors.New("'GetAdd' Decoder is not impelemented")
}

// encodeAddResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.AddReply)
	return resp, nil
	//return nil, errors.New("'GetAdd' Encoder is not impelemented")
}
func (g *grpcServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	_, rep, err := g.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddReply), nil
}
