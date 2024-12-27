package data

import (
	"bizclient/internal/conf"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	ggrpc "google.golang.org/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBizServerRepo)

// Data .
type Data struct {
	BizServer *ggrpc.ClientConn
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func newBizServerClient() *ggrpc.ClientConn {
	ctx := context.Background()
	conn, err := grpc.Dial(ctx)
	if err != nil {
		panic("bizserver grpc dial failed" + err.Error())
	}
	return conn
}
