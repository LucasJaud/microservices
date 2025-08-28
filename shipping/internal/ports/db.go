package ports

import (
	"github.com/LucasJaud/microservices/shipping/internal/application/core/domain"
	"context"
)
type DBPort interface {
	Save(ctx context.Context ,shipping *domain.Shipping) error
}