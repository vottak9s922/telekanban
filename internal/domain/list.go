package domain

import "time"

type List struct {
	ID        string    `json:"id"`
	BoardID   string    `json:"board_id"`
	Name      string    `json:"name"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
