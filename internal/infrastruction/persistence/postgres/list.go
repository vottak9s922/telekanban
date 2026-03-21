package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/telekanban/internal/domain"
)

type ListRepo struct {
	pool *pgxpool.Pool
}

func NewListRepo(pool *pgxpool.Pool) domain.ListRepository {
	return &ListRepo{pool: pool}
}
