package postgres

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/telekanban/internal/domain"
)

type ListRepo struct {
	pool *pgxpool.Pool
}

func NewListRepo(pool *pgxpool.Pool) domain.ListRepository {
	return &ListRepo{pool: pool}
}

func (r *ListRepo) Create(ctx context.Context, list *domain.List) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO lists (id, board_id, name, position, created_at, updated_at)
		 VALUES $1, $2, $3, $4, $5, $6
		 ON CONFLICT (name) DO NOTHING`,
		list.ID, list.BoardID, list.Name, list.Position, list.CreatedAt, list.UpdatedAt,
	)

	return err
}

func (r *ListRepo) FindByBoard(ctx context.Context, boardID string) ([]*domain.List, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT * FROM lists 
		WHERE board_id = $1 
		ORDER BY created_at DESC`, boardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []*domain.List
	err = pgxscan.ScanAll(&lists, rows)

	return lists, err
}
