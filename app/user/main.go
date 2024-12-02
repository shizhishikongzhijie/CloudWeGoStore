package main

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/app/user/conf"
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"github.com/cloudwego/biz-demo/gomall/common/serversuite"
	commonutils "github.com/cloudwego/biz-demo/gomall/common/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	MetricsAddress := conf.GetConf().Kitex.MetricsPort
	if strings.HasPrefix(MetricsAddress, ":") {
		localIp := commonutils.MustGetLocalIPv4()
		MetricsAddress = localIp + MetricsAddress
	}
	mtl.InitMetric(ServiceName, MetricsAddress, RegistryAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := commonutils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	// opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
	// 	CurrentServiceName: ServiceName,
	// 	RegistryAddr:       RegistryAddr,
	// }))
	// // 创建一个新的Consul注册器，用于服务发现和健康检查
	// // 参数 conf.GetConf().Kitex.Service 是Consul代理的地址, 而不是直接写死
	// r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	// // 如果在创建Consul注册器时发生错误，则记录错误并退出程序
	// if err != nil {
	// 	klog.Fatal(err)
	// }
	// // 将Consul注册器添加到服务器的配置选项中，以便服务器可以使用Consul进行服务注册和发现
	// opts = append(opts, server.WithRegistry(r))
	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
