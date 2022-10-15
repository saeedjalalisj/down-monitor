package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Url       string    `json:"url" validate:"required"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ServiceUsecase interface {
	Create(ctx context.Context, csd *CreateServiceDto) (uuid.UUID, error)
	// GenerateToken(ctx context.Context, sID uuid.UUID) string
	// Update()
	// Delete()
	// GetByIDs()
}

type ServiceRepository interface {
	Create(ctx context.Context, csd *CreateServiceDto) (uuid.UUID, error)
	// GenerateToken(ctx context.Context, sID uuid.UUID) string
	// Update()
	// Delete()
	// GetByIDs()
}

type CreateServiceDto struct {
	Name string
	Url  string
}
