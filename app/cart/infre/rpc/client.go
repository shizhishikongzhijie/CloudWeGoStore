package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	cartutils "github.com/cloudwego/biz-demo/gomall/app/cart/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       RegistryAddr,
			CurrentServiceName: ServiceName,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
