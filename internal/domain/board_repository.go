package domain

import "context"

type BoardRepository interface {
	Create(ctx context.Context, board *Board) error
	FindByOwner(ctx context.Context, ownerID string) ([]*Board, error)
}
