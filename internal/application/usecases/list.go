package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/telekanban/internal/domain"
)

type ListUsecase struct {
	ListRepo domain.ListRepository
}

func NewListUsecase(r domain.ListRepository) *ListUsecase {
	return &ListUsecase{ListRepo: r}
}

func (u *ListUsecase) Create(ctx context.Context, name, board_id string, position int) (*domain.List, error) {
	if name == "" {
		return nil, fmt.Errorf("List should have a name")
	}

	if board_id == "" {
		return nil, fmt.Errorf("Board should be specified")
	}

	list := &domain.List{
		ID:        uuid.NewString(),
		BoardID:   board_id,
		Name:      name,
		Position:  position,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := u.ListRepo.Create(ctx, list)
	if err != nil {
		return nil, err
	}
	return list, err
}
