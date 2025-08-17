package ports

import (
	"context"

	"github.com/LucasJaud/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(ctx context.Context,order *domain.Order) error
}