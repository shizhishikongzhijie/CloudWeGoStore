package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func InitMetric(serviceName string, metricsPort string, registryAddr string) (registry.Registry, *registry.Info) {
	// 初始化Prometheus监控的Registry对象
	Registry = prometheus.NewRegistry()

	// 注册Go语言运行时的指标收集器
	Registry.MustRegister(collectors.NewGoCollector())

	// 注册进程相关的指标收集器
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	// 创建Consul注册器实例
	r, _ := consul.NewConsulRegister(registryAddr)

	// 解析metrics服务监听的TCP地址
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)

	// 构建注册信息对象，用于在Consul中注册Prometheus metrics服务
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}

	// 在Consul中注册Prometheus metrics服务
	_ = r.Register(registryInfo)

	// 注册服务关闭钩子，用于在服务关闭时注销Consul中的服务注册信息
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo) //nolint:errcheck
	})

	// 设置Prometheus metrics HTTP处理函数
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	// 启动HTTP服务器监听指定端口，提供metrics服务
	go http.ListenAndServe(metricsPort, nil) //nolint:errcheck

	// 返回Consul注册器实例和服务注册信息
	return r, registryInfo
}
