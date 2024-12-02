package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	cli, err := echo.NewClient("demo_thrift", client.WithHostPorts("127.0.0.1:8888"),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_thrift_client",
		}),
	)
	if err != nil {
		panic(err)
	}

	res, err := cli.Echo(context.Background(), &api.Request{Message: "hello"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", res)
}
