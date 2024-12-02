package serversuite

import (
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonServerSuite) Options() []server.Option {
	// 初始化服务器配置选项
	opts := []server.Option{
		// 使用传输元数据的ServerHTTP2处理器
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		// 设置服务器基本信息，包括当前服务名称
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: s.CurrentServiceName}),
		// 配置Prometheus服务器追踪器，并禁用服务器端指标收集
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
		// 设置服务追踪器
		server.WithSuite(tracing.NewServerSuite()),
	}

	// 创建Consul注册器实例
	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		// 如果创建注册器失败，则记录错误并退出
		klog.Fatal(err)
	}

	// 将Consul注册器添加到服务器配置选项中
	opts = append(opts, server.WithRegistry(r))
	// 返回配置选项供服务器使用
	return opts
}