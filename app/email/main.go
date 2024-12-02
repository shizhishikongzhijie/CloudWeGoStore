package main

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/cloudwego/biz-demo/gomall/app/email/biz/consumer"
	"github.com/cloudwego/biz-demo/gomall/app/email/conf"
	"github.com/cloudwego/biz-demo/gomall/app/email/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"github.com/cloudwego/biz-demo/gomall/common/serversuite"
	commonutils "github.com/cloudwego/biz-demo/gomall/common/utils"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email/emailservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	mq.Init()
	consumer.Init()
	MetricsAddress := conf.GetConf().Kitex.MetricsPort
	// MetricsAddress := conf.GetConf().Kitex.MetricsPort
	// 初始化MetricsAddress，获取配置文件中的Metrics端口信息
	// 如果MetricsAddress以":"开头，表明需要动态获取本地IP地址
	if strings.HasPrefix(MetricsAddress, ":") {
		// 使用commonutils.MustGetLocalIPv4()函数获取本地IPv4地址
		localIp := commonutils.MustGetLocalIPv4()
		// 将本地IP地址与MetricsAddress拼接，以构造完整的Metrics地址
		MetricsAddress = localIp + MetricsAddress
	}
	// 调用mtl.InitMetric函数初始化Metrics
	// 参数分别为服务名称，Metrics地址和注册地址
	mtl.InitMetric(ServiceName, MetricsAddress, RegistryAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	opts := kitexInit()

	svr := emailservice.NewServer(new(EmailServiceImpl), opts...)

	err := svr.Run()
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
