package ports

import "github.com/LucasJaud/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(Order domain.Order) error
}