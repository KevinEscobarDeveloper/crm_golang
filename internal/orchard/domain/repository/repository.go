package repository

import (
	"context"
	"mango_crm/internal/orchard/domain/entity"
)

type Repository interface {
	Save(ctx context.Context, o *entity.Orchard) error
	FindById(ctx context.Context, id string) (*entity.Orchard, error)
	FindAll(ctx context.Context) ([]*entity.Orchard, error)
	Update(ctx context.Context, o *entity.Orchard) error
	Delete(ctx context.Context, id string) error
}
