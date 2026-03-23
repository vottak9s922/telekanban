package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/telekanban/internal/domain"
)

type mockListsRepo struct {
	createCalled bool
	lastList     *domain.List
}

func (m *mockListsRepo) Create(ctx context.Context, list *domain.List) error {
	m.createCalled = true
	m.lastList = list
	return nil
}

func (m *mockListsRepo) FindByBoard(ctx context.Context, boardID string) ([]*domain.List, error) {
	return nil, nil
}

func TestCreateList_Success(t *testing.T) {
	repo := &mockListsRepo{}
	uc := NewListUsecase(repo)

	list, err := uc.CreateList(context.Background(), "TODO", "105a520a-2961-408d-9b00-34b19911647f", 0)

	assert.NoError(t, err)
	assert.NotNil(t, list)
	assert.Equal(t, "TODO", list.Name)
	assert.Equal(t, "105a520a-2961-408d-9b00-34b19911647f", list.BoardID)
	assert.Equal(t, 0, list.Position)
	assert.True(t, repo.createCalled)
	assert.Equal(t, list, repo.lastList)
}

func TestCreateList_Invalid(t *testing.T) {
	repo := &mockListsRepo{}
	uc := NewListUsecase(repo)

	list, err := uc.CreateList(context.Background(), "", "", 0)

	assert.Error(t, err)
	assert.Nil(t, list)
}
