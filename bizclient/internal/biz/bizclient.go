package biz

import (
	"context"

	v1 "bizclient/api/bizclient/v1"
	bizserverv1 "bizclient/api/bizserver/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type BizServerRepo interface {
	SayHello(context.Context, *bizserverv1.HelloRequest) (*bizserverv1.HelloReply, error)
}

type BizClientUsecase struct {
	repo BizServerRepo
	log  *log.Helper
}

func NewBizClientUsecase(repo BizServerRepo, logger log.Logger) *BizClientUsecase {
	return &BizClientUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BizClientUsecase) HelloServer(ctx context.Context, name string) (string, error) {
	if name == "failed" {
		return "", ErrUserNotFound
	}
	rsp, err := uc.repo.SayHello(ctx, &bizserverv1.HelloRequest{Name: name})
	if err != nil {
		uc.log.Error("say hello error", "error", err)
		return "", err
	}
	return rsp.Message, nil
}
