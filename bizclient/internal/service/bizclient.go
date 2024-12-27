package service

import (
	"context"

	v1 "bizclient/api/bizclient/v1"
	"bizclient/internal/biz"
)

type BizClientService struct {
	v1.UnimplementedBizClientServer

	uc *biz.BizClientUsecase
}

func NewBizClientService(uc *biz.BizClientUsecase) *BizClientService {
	return &BizClientService{uc: uc}
}

func (s *BizClientService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	_, err := s.uc.HelloServer(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello "}, nil
}
