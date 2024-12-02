package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/app/email/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return
	}
	if cartResult == nil || len(cartResult.Items) == 0 {
		err = errors.New("cart is empty")
		return
	}
	var (
		oi    []*order.OrderItem
		total float32
	)
	for _, cartItem := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}
	// create order

	var orderId string

	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)

	if orderResult != nil && orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}
	klog.Info(paymentResult)

	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)

	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You just created an order in CloudWeGo shop",
		Content:     "You just created an order in CloudWeGo shop",
	})
	// msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}

	// // otel inject
	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)
	klog.Info(orderResult)
	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
