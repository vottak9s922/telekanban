package postgres

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/telekanban/internal/application/ports"
	"github.com/telekanban/internal/domain"
)

type BoardRepo struct {
	pool *pgxpool.Pool
}

func NewBoardRepo(pool *pgxpool.Pool) ports.BoardRepository {
	return &BoardRepo{pool: pool}
}

func (r *BoardRepo) Create(ctx context.Context, board *domain.Board) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO boards (id, name, owner_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (name, owner_id) DO NOTHING`,
		board.ID, board.Name, board.OwnerID, board.CreatedAt, board.UpdatedAt,
	)

	return err
}

func (r *BoardRepo) FindByOwner(ctx context.Context, ownerID string) ([]*domain.Board, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT * FROM boards 
		WHERE owner_id = $1 
		ORDER BY created_at DESC`, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boards []*domain.Board
	err = pgxscan.ScanAll(&boards, rows)

	return boards, err
}
