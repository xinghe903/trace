package data

import (
	"context"

	"bizserver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type helloRepo struct {
	data *Data
	log  *log.Helper
}

// NewhelloRepo .
func NewHelloRepo(data *Data, logger log.Logger) biz.HelloRepo {
	return &helloRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *helloRepo) Save(ctx context.Context, g *biz.Hello) (*biz.Hello, error) {
	return g, nil
}

func (r *helloRepo) Update(ctx context.Context, g *biz.Hello) (*biz.Hello, error) {
	return g, nil
}

func (r *helloRepo) FindByID(context.Context, int64) (*biz.Hello, error) {
	return nil, nil
}

func (r *helloRepo) ListByHello(context.Context, string) ([]*biz.Hello, error) {
	return nil, nil
}

func (r *helloRepo) ListAll(context.Context) ([]*biz.Hello, error) {
	return nil, nil
}
