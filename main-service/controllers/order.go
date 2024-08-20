package controllers

import "order-service/repositories/interfaces"

type OrderControllers struct {
	repo interfaces.OrderInterfaces
}

func NewOrderControllers(repo interfaces.OrderInterfaces) *OrderControllers {
	return &OrderControllers{
		repo: repo,
	}
}
