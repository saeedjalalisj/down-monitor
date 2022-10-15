package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/saeedjalalisj/down-monitor/internal/domain"
)

type serviceUsecase struct {
	serviceRepo    domain.ServiceRepository
	contextTimeout time.Duration
}

func NewServiceUsecase(s domain.ServiceRepository, timeout time.Duration) domain.ServiceUsecase {
	return &serviceUsecase{
		serviceRepo:    s,
		contextTimeout: timeout,
	}
}

func (s *serviceUsecase) Create(ctx context.Context, service *domain.CreateServiceDto) (id uuid.UUID, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	id, err = s.serviceRepo.Create(ctx, service)
	return id, err
}
