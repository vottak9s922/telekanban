package domain

import "context"

type ListRepository interface {
	Create(ctx context.Context, list *List) error
	FindByBoard(ctx context.Context, boardID string) ([]*List, error)
}
