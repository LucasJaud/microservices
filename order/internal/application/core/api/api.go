package api

import (
"github.com/LucasJaud/microservices/order/internal/ports"
"github.com/LucasJaud/microservices/order/internal/application/core/domain"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

