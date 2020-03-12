package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/phanletrunghieu/demo-go-micro-payment-srv/delivery/grpc"
)

func main() {
	grpcSrv := grpc.New()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- grpcSrv.Run()
	}()

	log.Println("exit", <-errs)
}
