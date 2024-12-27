package biz

import (
	"context"

	v1 "bizserver/api/bizserver/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Hello struct {
	Hello string
}

type HelloRepo interface {
	Save(context.Context, *Hello) (*Hello, error)
	Update(context.Context, *Hello) (*Hello, error)
	FindByID(context.Context, int64) (*Hello, error)
	ListByHello(context.Context, string) ([]*Hello, error)
	ListAll(context.Context) ([]*Hello, error)
}

type BizServerUsercase struct {
	repo HelloRepo
	log  *log.Helper
}

func NewBizServerUsercase(repo HelloRepo, logger log.Logger) *BizServerUsercase {
	return &BizServerUsercase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BizServerUsercase) CreateHello(ctx context.Context, g *Hello) (*Hello, error) {
	uc.log.WithContext(ctx).Infof("CreateHello: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
