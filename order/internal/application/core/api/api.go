package api

import (
	"github.com/LucasJaud/microservices/order/internal/application/core/domain"
	"github.com/LucasJaud/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db: db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	
	Quantity := int32(0)
	for _,item := range order.OrderItems {
		Quantity += item.Quantity
	}
	if Quantity > 50 {
		return domain.Order{}, status.Error(codes.InvalidArgument, "Payment cannot be more than a thousand.")
	}

	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		order.Status = "Canceled"
		if saveErr := a.db.Save(&order); saveErr != nil {
			return domain.Order{}, saveErr
		}
		return order, paymentErr
	}
	order.Status = "Paid"
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

