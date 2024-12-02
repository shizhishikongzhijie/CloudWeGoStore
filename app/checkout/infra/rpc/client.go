package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/conf"
	checkoututils "github.com/cloudwego/biz-demo/gomall/app/checkout/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
)

func InitClient() {
	once.Do(func() {
		// registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		// serviceName = conf.GetConf().Kitex.Service
		// commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
		// 	CurrentServiceName: serviceName,
		// 	RegistryAddr:       registryAddr,
		// })
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initProductClient() {
	// 初始化客户端选项
	// 该代码段初始化了一个 client.Option 类型的切片，用于配置客户端的特定选项。
	// 主要目的是设置创建 gRPC 客户端所需的通用参数，如服务注册地址和服务名称。
	opts := []client.Option{
		// 使用 gRPC 客户端的通用配置套件
		// 包括设置服务注册地址和当前服务名称，这些对于服务发现和通信至关重要。
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoututils.MustHandleError(err)
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       RegistryAddr,
			CurrentServiceName: ServiceName,
		}),
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       RegistryAddr,
			CurrentServiceName: ServiceName,
		}),
	}
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoututils.MustHandleError(err)
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       RegistryAddr,
			CurrentServiceName: ServiceName,
		}),
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	checkoututils.MustHandleError(err)
}
