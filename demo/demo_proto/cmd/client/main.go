package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"

	// "github.com/cloudwego/kitex/transport"

	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatal(err)
	}
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r),
		// client.WithTransportProtocol(transport.GRPC), 不使用 gRPC [由于对于windows有些问题]
		client.WithShortConnection(), // 使用短链接
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware),
	)
	if err != nil {
		klog.Fatal(err)
	}

	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	// WithPersistentValue 和 WithValue 区别
	// WithPersistentValue 能一直传送下去
	// WithValue 只能传递给下一个服务
	res, err := c.Echo(ctx, &pbapi.Request{Message: "error"})
	// res, err := c.Echo(context.TODO(), &pbapi.Request{Message: "hello"})
	var bizErr *kerrors.GRPCBizStatusError
	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Printf("%#v", bizErr)
		}
		klog.Fatal(err)
	}
	fmt.Printf("%v", res)
}
