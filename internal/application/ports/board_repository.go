package ports

import (
	"context"

	"github.com/telekanban/internal/domain"
)

type BoardRepository interface {
	Create(ctx context.Context, board *domain.Board) error
	FindByOwner(ctx context.Context, ownerID string) ([]*domain.Board, error)
}
