package main

import (
	"log"
	"net"
	"time"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	
	opts := kitexInit()

	svr := echoservice.NewServer(new(EchoServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	// middleware 启用中间件
	opts = append(opts, server.WithServiceAddr(addr),server.WithMiddleware(middleware.Middleware))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// 创建一个新的Consul注册器，用于服务发现和健康检查
	// 参数 conf.GetConf().Kitex.Service 是Consul代理的地址, 而不是直接写死
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	// 如果在创建Consul注册器时发生错误，则记录错误并退出程序
	if err != nil {
		log.Fatal(err)
	}
	// 将Consul注册器添加到服务器的配置选项中，以便服务器可以使用Consul进行服务注册和发现
	opts = append(opts, server.WithRegistry(r))

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
