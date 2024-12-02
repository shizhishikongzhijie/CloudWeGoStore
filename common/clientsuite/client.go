package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonGrpcClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonGrpcClientSuite) Options() []client.Option {
	// 创建Consul解析器
	// s.RegistryAddr 是服务注册地址
	r, err := consul.NewConsulResolver(s.RegistryAddr)
	// 如果创建过程中出现错误，则抛出panic
	if err != nil {
		panic(err)
	}

	// 初始化客户端选项
	opts := []client.Option{
		// 使用Consul作为服务发现和负载均衡的解析器
		client.WithResolver(r),
		// 使用transmeta.ClientHTTP2Handler作为元数据处理器
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		// 设置传输协议为GRPC
		client.WithTransportProtocol(transport.GRPC),
	}

	// 添加客户端基本信息和追踪套件配置到选项中
	opts = append(opts,
		// 设置服务名称为当前服务名称
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// 使用tracing.NewClientSuite()作为追踪套件
		client.WithSuite(tracing.NewClientSuite()),
	)

	// 返回配置好的客户端选项
	return opts
}
