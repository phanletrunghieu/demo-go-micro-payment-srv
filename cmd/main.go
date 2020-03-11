package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	micro "github.com/micro/go-micro/v2"
	pbOrder "github.com/phanletrunghieu/demo-go-micro-order-srv/proto/order"
	"github.com/phanletrunghieu/demo-go-micro-payment-srv/delivery/grpc"
	pb "github.com/phanletrunghieu/demo-go-micro-payment-srv/proto/payment"
)

func main() {
	srv := micro.NewService(
		micro.Name("service.payment"),
	)

	srv.Init()

	orderClient := pbOrder.NewOrderService("service.order", srv.Client())
	pb.RegisterPaymentServiceHandler(srv.Server(), grpc.NewHandler(orderClient))

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- srv.Run()
	}()

	log.Println("exit", <-errs)
}
