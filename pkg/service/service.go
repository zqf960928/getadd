package service

import (
	"context"
	"fmt"
	"get_add/pkg/grpc/pb"
)

// GetAddService describes the service.
type GetAddService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Add(ctx context.Context, req *pb.AddRequest) (res *pb.AddReply, err error)
}

type basicGetAddService struct{}

func (ba *basicGetAddService) Add(ctx context.Context, req *pb.AddRequest) ( *pb.AddReply,  error) {
	// TODO implement the business logic of Add
	res := new(pb.AddReply)
	c := req.A + req.B
	fmt.Println("===c===",c)
	res.Res = c

	fmt.Println("===success===")
	return res, nil
}

// NewBasicGetAddService returns a naive, stateless implementation of GetAddService.
func NewBasicGetAddService() GetAddService {
	return &basicGetAddService{}
}

// New returns a GetAddService with all of the expected middleware wired in.
func New(middleware []Middleware) GetAddService {
	var svc GetAddService = NewBasicGetAddService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
