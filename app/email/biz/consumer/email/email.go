package email

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/email/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/app/email/infra/notify"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
)

func ConsumerInit() {
	// Connect to a server

	tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("email", func(m *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			klog.Error(err)
		}

		// consumer otel
		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
		// consumer otel
		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(m.Header))
		_, span := tracer.Start(ctx, "shop-email-consumer")
		defer span.End()

	})
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe() //nolint:errcheck
		mq.Nc.Close()
	})
}
