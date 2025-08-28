package ports

import (
	"context"
	"github.com/LucasJaud/microservices/shipping/internal/application/core/domain"
)

type APIPort interface {
	Process(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error)
}





