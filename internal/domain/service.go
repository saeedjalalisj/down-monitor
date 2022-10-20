package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name" validate:"required"`
	Url       string    `json:"url" validate:"required"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ServiceUsecase interface {
	Create(ctx context.Context, csd *CreateServiceDto) (uuid.UUID, error)
	GenerateToken(ctx context.Context, serviceName string) (token string, err error)
	// Update()
	// Delete()
	// GetByIDs()
}

type ServiceRepository interface {
	Create(ctx context.Context, s *Service) (uuid.UUID, error)
	// Update()
	// Delete()
	// GetByIDs()
}

type CreateServiceDto struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required"`
}
