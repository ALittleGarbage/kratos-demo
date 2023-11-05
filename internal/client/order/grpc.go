package order

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	userv1 "kratos-demo/api/user/v1"
	"kratos-demo/internal/conf"
)

type GRPCClient struct {
	Uc userv1.UserClient
}

func NewGRPCClient(c *conf.Server, reg *nacos.Registry, logger log.Logger) (*GRPCClient, func(), error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user.grpc"),
		grpc.WithDiscovery(reg),
	)
	if err != nil {
		log.NewHelper(logger).Errorf("创建连接时发生错误,原因:%v", err)
		return nil, nil, err
	}
	cleanup := func() {
		conn.Close()
	}

	userClient := userv1.NewUserClient(conn)

	return &GRPCClient{Uc: userClient}, cleanup, nil
}
