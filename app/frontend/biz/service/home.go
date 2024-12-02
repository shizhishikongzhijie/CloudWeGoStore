package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

//	func (h *HomeService) Run(req *home.Empty) (resp *home.Empty, err error) {
//		//defer func() {
//		// hlog.CtxInfof(h.Context, "req = %+v", req)
//		// hlog.CtxInfof(h.Context, "resp = %+v", resp)
//		//}()
//		// todo edit your code
//		return
//	}
func (h *HomeService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp = make(map[string]any) //初始化
	ctx := h.Context
	p, err := rpc.ProductClient.ListProducts(ctx, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	// item := []map[string]any{
	// 	{"Name": "wei-1", "Price": "100", "Picture": "https://e-esine.cn/img/时之世-img-1731068318824.png!/format/webp"},
	// 	{"Name": "wei-2", "Price": "100", "Picture": "https://e-esine.cn/img/时之世-img-1731068318824.png!/format/webp"},
	// 	{"Name": "wei-3", "Price": "100", "Picture": "https://e-esine.cn/img/时之世-img-1731068318824.png!/format/webp"},
	// 	{"Name": "wei-4", "Price": "100", "Picture": "https://e-esine.cn/img/时之世-img-1731068318824.png!/format/webp"},
	// 	{"Name": "wei-5", "Price": "100", "Picture": "https://e-esine.cn/img/时之世-img-1731068318824.png!/format/webp"},
	// 	{"Name": "wei-6", "Price": "100", "Picture": "https://e-esine.cn/img/时之世-img-1731068318824.png!/format/webp"},
	// 	{"Name": "wei-7", "Price": "100", "Picture": "https://e-esine.cn/img/时之世-img-1731068318824.png!/format/webp"},
	// }
	resp["Title"] = "Home"
	resp["Item"] = p.Products
	return resp, nil
}
