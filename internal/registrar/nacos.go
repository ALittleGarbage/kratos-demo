package registrar

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"kratos-demo/internal/conf"
)

func NewNacos(reg *conf.Registrar, logger log.Logger) (*nacos.Registry, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(reg.GetNacos().GetAddr(), reg.GetNacos().GetPort()),
	}

	cc := constant.ClientConfig{
		NamespaceId: reg.GetNacos().GetNamespaceId(),
		TimeoutMs:   5000,
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)

	if err != nil {
		log.NewHelper(logger).Errorf("nacos初始化失败,原因:%v", err)
		return nil, err
	}

	r := nacos.New(client)
	return r, nil
}
