package data

import (
	"context"

	bizserverv1 "bizclient/api/bizserver/v1"
	"bizclient/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type bizServerRepo struct {
	data *Data
	log  *log.Helper
}

// NewbizServerRepo .
func NewBizServerRepo(data *Data, logger log.Logger) biz.BizServerRepo {
	return &bizServerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (b *bizServerRepo) SayHello(ctx context.Context, req *bizserverv1.HelloRequest) (*bizserverv1.HelloReply, error) {
	return b.data.BizServer.SayHello(ctx, req)
}
