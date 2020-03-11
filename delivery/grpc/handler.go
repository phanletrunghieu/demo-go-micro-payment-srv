package grpc

import (
	"context"
	"errors"

	pbOrder "github.com/phanletrunghieu/demo-go-micro-order-srv/proto/order"
	pb "github.com/phanletrunghieu/demo-go-micro-payment-srv/proto/payment"
)

// mock
var transactions map[string]*pb.Transaction

func init() {
	transactions = make(map[string]*pb.Transaction)
}

type Handler struct {
	// db *gorm.DB
	orderClient pbOrder.OrderService
}

func NewHandler(orderClient pbOrder.OrderService) *Handler {
	return &Handler{
		orderClient: orderClient,
	}
}

func (h *Handler) CreateTransaction(ctx context.Context, input *pb.Transaction, output *pb.Response) error {
	if input == nil {
		return errors.New("order cannot nil")
	}

	transactions[input.Id] = input
	updateOrder := &pbOrder.Order{
		Id:     input.OrderId,
		Status: pbOrder.OrderStatus_PAID,
	}

	_, err := h.orderClient.UpdateStatus(ctx, updateOrder)
	if err != nil {
		return err
	}

	output.Success = true
	output.Transaction = input
	return nil
}
