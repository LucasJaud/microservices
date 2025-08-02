package domain

import "time";

type OrderItem struct {
	ProductCode string
	UnitPrice float32
	Quantity int32
}

type Order struct {
	ID int64
	CustomerID int64
	Status string
	OrderItems []OrderItem
	CreatedAt int64
}

func NewOrder(customerId int64, orderItems []OrderItem,) Order{
	return Order{
		CreatedAt: time.Now().Unix(),
		Status: "Pending",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}