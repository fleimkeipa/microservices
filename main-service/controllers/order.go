package controllers

import "order-service/repositories/interfaces"

type CarControllers struct {
	repo interfaces.OrderInterfaces
}

func NewCarControllers(repo interfaces.OrderInterfaces) *CarControllers {
	return &CarControllers{
		repo: repo,
	}
}
