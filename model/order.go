package model

type Order struct {
	CustomerName string    `json:"customerName"`
	ListProducts []Product `json:"listProducts"`
}
