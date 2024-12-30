package data

import (
	"bizserver/internal/conf"
	"strings"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientV3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewHelloRepo, NewEtcdRegistrar)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
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
