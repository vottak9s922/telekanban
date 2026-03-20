package domain

import (
	"fmt"
	"time"
)

type Board struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBoard(id, name, ownerID string) (*Board, error) {
	if name == "" {
		return nil, fmt.Errorf("board cannot be empty")
	}
	return &Board{
		ID:        id,
		Name:      name,
		OwnerID:   ownerID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}

// domain behavior

func (b *Board) Rename(newName string) error {
	if newName == "" {
		return fmt.Errorf("board cannot be empty")
	}
	b.Name = newName
	b.UpdatedAt = time.Now().UTC()
	return nil
}
