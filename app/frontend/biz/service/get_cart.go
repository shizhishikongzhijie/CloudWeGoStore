package service

import (
	"context"
	"strconv"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	rpccart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var items []map[string]string
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: uint32(frontendutils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}
	var total float64
	for _, v := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.ProductId})
		if err != nil {
			continue
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Picture":     p.Picture,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Qty":         strconv.Itoa(int(v.Quantity)),
		})
		total += float64(v.Quantity) * float64(p.Price)
	}

	return utils.H{
		"Title": "Cart",
		"Item":  items,
		"Total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
