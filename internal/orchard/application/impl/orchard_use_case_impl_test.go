package impl

import (
	"context"
	"errors"
	"testing"

	"mango_crm/internal/orchard/domain/entity"
	"mango_crm/pkg/logger"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) Save(ctx context.Context, o *entity.Orchard) error {
	args := m.Called(ctx, o)
	return args.Error(0)
}

func (m *mockRepository) FindById(ctx context.Context, id string) (*entity.Orchard, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Orchard), args.Error(1)
}

func (m *mockRepository) FindAll(ctx context.Context) ([]*entity.Orchard, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Orchard), args.Error(1)
}

func (m *mockRepository) Update(ctx context.Context, o *entity.Orchard) error {
	args := m.Called(ctx, o)
	return args.Error(0)
}

func (m *mockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestImpl_Create(t *testing.T) {
	mockRepo := new(mockRepository)
	logger := &logger.Logger{}
	useCase := NewOrchardUseCase(mockRepo, logger)

	tests := []struct {
		name    string
		input   *entity.Orchard
		mockErr error
		wantErr bool
	}{
		{"success", &entity.Orchard{ID: "1", Name: "TestOrchard"}, nil, false},
		{"repo error", &entity.Orchard{ID: "2", Name: "ErrorOrchard"}, errors.New("save error"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Save", mock.Anything, tt.input).Return(tt.mockErr).Once()

			err := useCase.Create(context.TODO(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestImpl_GetById(t *testing.T) {
	mockRepo := new(mockRepository)
	logger := &logger.Logger{}
	useCase := NewOrchardUseCase(mockRepo, logger)

	tests := []struct {
		name    string
		id      string
		mockRes *entity.Orchard
		mockErr error
		wantErr bool
	}{
		{"found", "1", &entity.Orchard{ID: "1", Name: "TestOrchard"}, nil, false},
		{"not found", "2", nil, errors.New("not found"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("FindById", mock.Anything, tt.id).Return(tt.mockRes, tt.mockErr).Once()

			res, err := useCase.GetById(context.TODO(), tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockRes, res)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestImpl_List(t *testing.T) {
	mockRepo := new(mockRepository)
	logger := &logger.Logger{}
	useCase := NewOrchardUseCase(mockRepo, logger)

	tests := []struct {
		name    string
		mockRes []*entity.Orchard
		mockErr error
		wantErr bool
	}{
		{"success", []*entity.Orchard{{ID: "1", Name: "TestOrchard"}}, nil, false},
		{"repo error", nil, errors.New("list error"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("FindAll", mock.Anything).Return(tt.mockRes, tt.mockErr).Once()

			res, err := useCase.List(context.TODO())

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockRes, res)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestImpl_Update(t *testing.T) {
	mockRepo := new(mockRepository)
	logger := &logger.Logger{}
	useCase := NewOrchardUseCase(mockRepo, logger)

	tests := []struct {
		name    string
		input   *entity.Orchard
		mockErr error
		wantErr bool
	}{
		{"success", &entity.Orchard{ID: "1", Name: "UpdatedOrchard"}, nil, false},
		{"repo error", &entity.Orchard{ID: "2", Name: "ErrorOrchard"}, errors.New("update error"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Update", mock.Anything, tt.input).Return(tt.mockErr).Once()

			err := useCase.Update(context.TODO(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestImpl_Delete(t *testing.T) {
	mockRepo := new(mockRepository)
	logger := &logger.Logger{}
	useCase := NewOrchardUseCase(mockRepo, logger)

	tests := []struct {
		name    string
		id      string
		mockErr error
		wantErr bool
	}{
		{"success", "1", nil, false},
		{"repo error", "2", errors.New("delete error"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Delete", mock.Anything, tt.id).Return(tt.mockErr).Once()

			err := useCase.Delete(context.TODO(), tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
