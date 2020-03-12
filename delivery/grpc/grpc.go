package grpc

import (
	micro "github.com/micro/go-micro/v2"
	pbOrder "github.com/phanletrunghieu/demo-go-micro-order-srv/proto/order"
	pb "github.com/phanletrunghieu/demo-go-micro-payment-srv/proto/payment"
)

func New() micro.Service {
	srv := micro.NewService(
		micro.Name("srv.grpc.payment"),
	)

	srv.Init()

	orderClient := pbOrder.NewOrderService("srv.grpc.order", srv.Client())
	pb.RegisterPaymentServiceHandler(srv.Server(), NewHandler(orderClient))

	return srv
}
