package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/saeedjalalisj/down-monitor/internal/domain"
)

type postgresServiceRepository struct {
	Conn *sql.DB
}

func NewPostgresServiceRepository(conn *sql.DB) domain.ServiceRepository {
	return &postgresServiceRepository{conn}
}

func (p *postgresServiceRepository) Create(ctx context.Context, s *domain.Service) (id uuid.UUID, err error) {
	query := `INSERT INTO "service" (name, url, token) VALUES ($1, $2, $3) RETURNING id;`
	err = p.Conn.QueryRowContext(ctx, query, s.Name, s.Name, s.Token).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
