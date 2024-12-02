package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	err            error
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddress
	ServiceName    = frontendutils.ServiceName
)

func Init() {
	once.Do(func() {
		iniUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func iniUserClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	// frontendutils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}
func initProductClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	// frontendutils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}
func initCartClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	// frontendutils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}
func initCheckoutClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	// frontendutils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initOrderClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress)
	// frontendutils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}
