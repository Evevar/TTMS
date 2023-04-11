package main

import (
	order "TTMS/kitex_gen/order"
	"context"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// AddOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) AddOrder(ctx context.Context, req *order.AddOrderRequest) (resp *order.AddOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, req *order.DeleteOrderRequest) (resp *order.DeleteOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderRequest) (resp *order.UpdateOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetAllOrder(ctx context.Context, req *order.GetAllOrderRequest) (resp *order.GetAllOrderResponse, err error) {
	// TODO: Your code here...
	return
}

// GetOrderAnalysis implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderAnalysis(ctx context.Context, req *order.GetOrderAnalysisRequest) (resp *order.GetOrderAnalysisResponse, err error) {
	// TODO: Your code here...
	return
}
