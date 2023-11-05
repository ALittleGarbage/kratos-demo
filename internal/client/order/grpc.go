package order

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	userv1 "kratos-demo/api/user/v1"
	"kratos-demo/internal/conf"
)

type GRPCClient struct {
	Uc userv1.UserClient
}

func NewGRPCClient(c *conf.Server, reg *nacos.Registry, logger log.Logger) (*GRPCClient, func(), error) {
	/*	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint("discovery:///user"),
		kgrpc.WithDiscovery(reg),
	)*/
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		conn.Close()
	}

	userClient := userv1.NewUserClient(conn)

	return &GRPCClient{Uc: userClient}, cleanup, nil
}
