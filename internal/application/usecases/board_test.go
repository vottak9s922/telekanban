package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/telekanban/internal/domain"
)

type mockBoardsRepo struct {
	createCalled bool
	lastBoard    *domain.Board
}

func (m *mockBoardsRepo) Create(ctx context.Context, board *domain.Board) error {
	m.createCalled = true
	m.lastBoard = board
	return nil
}

func (m *mockBoardsRepo) FindByOwner(ctx context.Context, ownerID string) ([]*domain.Board, error) {
	return nil, nil
}

func TestCreateBoard_Success(t *testing.T) {
	repo := &mockBoardsRepo{}
	uc := NewBoardUsecase(repo)

	board, err := uc.CreateBoard(context.Background(), "My Board", "test_owner")

	assert.NoError(t, err)
	assert.NotNil(t, board)
	assert.Equal(t, "My Board", board.Name)
	assert.True(t, repo.createCalled)
	assert.Equal(t, board, repo.lastBoard)
}

func TestCreateBoard_Invalid(t *testing.T) {
	repo := &mockBoardsRepo{}
	uc := NewBoardUsecase(repo)

	board, err := uc.CreateBoard(context.Background(), "", "test_owner")

	assert.Error(t, err)
	assert.Nil(t, board)
}
