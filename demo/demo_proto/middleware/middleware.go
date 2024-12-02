package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func Middleware(next endpoint.Endpoint) endpoint.Endpoint{
	return func(ctx context.Context, request, response interface{}) (err error) {
		// do something before
		begin := time.Now()
		err = next(ctx, request, response)
		fmt.Println("cost:", time.Since(begin))
		return err
	}
}