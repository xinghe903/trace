package service

import (
	"context"

	v1 "bizserver/api/bizserver/v1"
	"bizserver/internal/biz"
)

// BizServerService is a greeter service.
type BizServerService struct {
	v1.UnimplementedBizServerServer

	uc *biz.BizServerUsercase
}

// NewBizServerService new a greeter service.
func NewBizServerService(uc *biz.BizServerUsercase) *BizServerService {
	return &BizServerService{uc: uc}
}

// SayHello implements bizserver.GreeterServer.
func (s *BizServerService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateHello(ctx, &biz.Hello{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
