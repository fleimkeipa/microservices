package models

type OrderRequest struct {
	OrderID string `json:"order_id"`
	SendBy  string `json:"send_by"`
}
