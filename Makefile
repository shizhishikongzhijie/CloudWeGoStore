.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC -module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC -module github.com/cloudwego/biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY:gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP -module github.com/cloudwego/biz-demo/gomall/app/frontend --service frontend --idl ../../idl/frontend/order_page.proto

.PHONY:gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user -module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC -module github.com/cloudwego/biz-demo/gomall/app/user --service user --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/user.proto

.PHONY:gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product -module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC -module github.com/cloudwego/biz-demo/gomall/app/product --service product --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/product.proto

.PHONY:gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart -module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC -module github.com/cloudwego/biz-demo/gomall/app/cart --service cart --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/cart.proto


.PHONY:gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment -module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC -module github.com/cloudwego/biz-demo/gomall/app/payment --service payment --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/payment.proto

.PHONY:gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout -module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC -module github.com/cloudwego/biz-demo/gomall/app/checkout --service checkout --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/checkout.proto

.PHONY:gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order -module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC -module github.com/cloudwego/biz-demo/gomall/app/order --service order --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/order.proto

.PHONY:gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email -module github.com/cloudwego/biz-demo/gomall/rpc_gen --I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC -module github.com/cloudwego/biz-demo/gomall/app/email --service email --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" --I ../../idl  --idl ../../idl/email.proto
