package usecase

import (
	"context"
	"mango_crm/internal/orchard/domain/usecase"

	"mango_crm/internal/orchard/domain/entity"
	"mango_crm/internal/orchard/domain/repository"
	"mango_crm/pkg/logger"
)

// impl is a struct that encapsulates repository interaction and logging functionality for use case operations.
type impl struct {
	repo   repository.Repository
	logger *logger.Logger
}

// NewOrchardUseCase initializes and returns an implementation of the OrchardUseCase interface using the provided dependencies.
func NewOrchardUseCase(repo repository.Repository, logger *logger.Logger) usecase.OrchardUseCase {
	return &impl{
		repo:   repo,
		logger: logger,
	}
}

// Create persists a new orchard entity to the repository.
// It logs an informational message and passes the entity to the repository's Save method.
// Returns an error if the save operation fails.
func (i *impl) Create(ctx context.Context, o *entity.Orchard) error {
	i.logger.Info("Creating new orchard")
	return i.repo.Save(ctx, o)
}

// GetById retrieves an Orchard entity by its unique identifier from the repository.
// Returns the Orchard entity if found, or an error if retrieval fails.
func (i *impl) GetById(ctx context.Context, id string) (*entity.Orchard, error) {
	i.logger.Info("Getting orchard by ID")
	return i.repo.FindById(ctx, id)
}

// List retrieves a list of all orchard entities from the repository and logs the operation.
func (i *impl) List(ctx context.Context) ([]*entity.Orchard, error) {
	i.logger.Info("Listing all orchards")
	return i.repo.FindAll(ctx)
}

// Update updates an existing orchard in the repository and logs the operation. Returns an error if the update fails.
func (i *impl) Update(ctx context.Context, o *entity.Orchard) error {
	i.logger.Info("Updating orchard")
	return i.repo.Update(ctx, o)
}

// Delete removes an orchard record identified by the specified ID from the repository. Returns an error if the operation fails.
func (i *impl) Delete(ctx context.Context, id string) error {
	i.logger.Info("Deleting orchard")
	return i.repo.Delete(ctx, id)
}
