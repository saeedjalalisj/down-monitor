package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	jwtService "github.com/saeedjalalisj/down-monitor/infra/jwt"
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
	token, err := s.GenerateToken(ctx, service.Name)

	if err != nil {
		return id, err
	}

	if service.IsActive {
		err := s.Monitor(ctx, service.Url)
		if err != nil {
			fmt.Println("err", err)
		}
	}

	id, err = s.serviceRepo.Create(ctx, &domain.Service{
		Name:     service.Name,
		Url:      service.Url,
		Token:    token,
		IsActive: service.IsActive,
	})

	if err != nil {
		return id, err
	}
	
	return id, nil
}

func (s *serviceUsecase) GenerateToken(ctx context.Context, serviceName string) (token string, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	token, err = jwtService.GenerateJWT(serviceName)
	return token, err
}

func (s *serviceUsecase) Monitor(ctx context.Context, url string) (err error) {
	return nil
}

func (s *serviceUsecase) callApi(ctx context.Context, url string) {
	fmt.Println("called")
}
