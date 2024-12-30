package data

import (
	v1 "bizclient/api/bizserver/v1"
	"bizclient/internal/conf"
	"context"
	"strings"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	clientV3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBizServerRepo, NewEtcdRegistrar, NewEtcdDiscovery)

// Data .
type Data struct {
	BizServer v1.BizServerClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, r registry.Discovery) (*Data, func(), error) {
	bizServerClient, fn := newBizServerClient(c.GetBizserver(), r)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		fn()
	}
	return &Data{
		BizServer: bizServerClient,
	}, cleanup, nil
}

func newBizServerClient(addr string, r registry.Discovery) (v1.BizServerClient, func()) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+addr),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic("bizserver grpc dial failed" + err.Error())
	}
	c := v1.NewBizServerClient(conn)
	return c, func() { conn.Close() }
}

func NewEtcdRegistrar(reg *conf.Data) registry.Registrar {
	addr := strings.Split(reg.GetEtcd().GetAddr(), ",")
	cli, err := clientV3.New(clientV3.Config{
		DialTimeout: reg.GetEtcd().GetTimeout().AsDuration(),
		Endpoints:   addr,
		Username:    reg.GetEtcd().GetUsername(),
		Password:    reg.GetEtcd().GetPassword(),
	})
	if err != nil {
		panic("registry etcd connect error: " + err.Error())
	}
	r := etcd.New(cli)
	return r
}

func NewEtcdDiscovery(reg *conf.Data) registry.Discovery {
	addr := strings.Split(reg.GetEtcd().GetAddr(), ",")
	cli, err := clientV3.New(clientV3.Config{
		DialTimeout: reg.GetEtcd().GetTimeout().AsDuration(),
		Endpoints:   addr,
		Username:    reg.GetEtcd().GetUsername(),
		Password:    reg.GetEtcd().GetPassword(),
	})
	if err != nil {
		panic("registry etcd connect error: " + err.Error())
	}
	r := etcd.New(cli)
	return r
}
