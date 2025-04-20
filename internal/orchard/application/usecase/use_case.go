package usecase

import (
	"context"
	"mango_crm/internal/orchard/domain/entity"
)

type OrchardUseCase interface {
	Create(ctx context.Context, o *entity.Orchard) error
	GetById(ctx context.Context, id string) (*entity.Orchard, error)
	List(ctx context.Context) ([]*entity.Orchard, error)
	Update(ctx context.Context, o *entity.Orchard) error
	Delete(ctx context.Context, id string) error
}
