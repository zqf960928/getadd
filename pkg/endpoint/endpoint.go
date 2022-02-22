package endpoint

import (
	"context"
	"github.com/zqf960928/getadd/pkg/grpc/pb"
	"github.com/zqf960928/getadd/pkg/service"
	"github.com/go-kit/kit/endpoint"

	//endpoint "github.com/go-kit/kit/endpoint"
)

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	Res int   `json:"res"`
	Err error `json:"err"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.GetAddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.AddRequest)
		res, _ := s.Add(ctx, &req)
		return &pb.AddReply{
			Res: res.Res,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, a int, b int) (res int, err error) {
	request := AddRequest{
		A: a,
		B: b,
	}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).Res, response.(AddResponse).Err
}
