package service

import (
	"context"
	"fmt"

	api "github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(r *api.Request) (resp *api.Response, err error) {
	// Finish your business logic.
	info := rpcinfo.GetRPCInfo(s.ctx)
	fmt.Printf("ServiceName: %s\n", info.From().ServiceName())
	return &api.Response{Message: r.Message}, nil
}
