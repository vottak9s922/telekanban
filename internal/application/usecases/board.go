package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/telekanban/internal/domain"
)

type BoardUsecase struct {
	BoardRepo domain.BoardRepository
}

func NewBoardUsecase(r domain.BoardRepository) *BoardUsecase {
	return &BoardUsecase{BoardRepo: r}
}

func (u *BoardUsecase) CreateBoard(ctx context.Context, name, ownerID string) (*domain.Board, error) {
	if name == "" || ownerID == "" {
		return nil, fmt.Errorf("board cannot be empty")
	}

	board := &domain.Board{
		ID:        uuid.NewString(),
		Name:      name,
		OwnerID:   ownerID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := u.BoardRepo.Create(ctx, board)
	if err != nil {
		return nil, err
	}
	return board, nil
}
