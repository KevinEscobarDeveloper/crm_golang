package impl

import (
	"context"
	"mango_crm/internal/orchard/application/usecase"
	"mango_crm/internal/orchard/domain/entity"
	"mango_crm/internal/orchard/domain/repository"
	"mango_crm/pkg/logger"
)

type impl struct {
	repo   repository.Repository
	logger *logger.Logger
}

func NewOrchardUseCase(repo repository.Repository, logger *logger.Logger) usecase.OrchardUseCase {
	return &impl{
		repo:   repo,
		logger: logger,
	}
}

func (i *impl) Create(ctx context.Context, o *entity.Orchard) error {
	i.logger.Info("Creating new orchard")
	return i.repo.Save(ctx, o)
}

func (i *impl) GetById(ctx context.Context, id string) (*entity.Orchard, error) {
	i.logger.Info("Getting orchard by ID")
	return i.repo.FindById(ctx, id)
}

func (i *impl) List(ctx context.Context) ([]*entity.Orchard, error) {
	i.logger.Info("Listing all orchards")
	return i.repo.FindAll(ctx)
}

func (i *impl) Update(ctx context.Context, o *entity.Orchard) error {
	i.logger.Info("Updating orchard")
	return i.repo.Update(ctx, o)
}

func (i *impl) Delete(ctx context.Context, id string) error {
	i.logger.Info("Deleting orchard")
	return i.repo.Delete(ctx, id)
}
