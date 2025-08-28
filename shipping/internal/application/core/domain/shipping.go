package domain

import "time"

type Shipping struct {
	ID int64
	OrderID int64
	OrderItems []OrderItem
	DeliveryDays int32 
	CreatedAt int64
}

type OrderItem struct {
	ProductCode string
	Quantity int32
}

func NewShipping(orderID int64, orderItems []OrderItem) Shipping{
	deliveryDays := calcDeliveryDays(orderItems)
	return Shipping{
		OrderID: orderID,
		OrderItems: orderItems,
		DeliveryDays: deliveryDays,
		CreatedAt: time.Now().Unix(),
	}
}

func calcDeliveryDays(orderItems []OrderItem) int32 {
	totalQuant := int32(0)

	for _, item := range orderItems {
		totalQuant += item.Quantity
	}

	deliveryDays := int32(1)
	if totalQuant > 0 {
		adicionalDays := (totalQuant -1 ) / 5
		deliveryDays += adicionalDays
	}
	return deliveryDays
}