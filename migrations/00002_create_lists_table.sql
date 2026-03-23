-- +goose Up
CREATE TABLE IF NOT EXISTS lists (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  board_id UUID NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  position INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

  UNIQUE (board_id, name)
);

CREATE INDEX idx_lists_board_id ON lists(board_id);

-- +goose Down
DROP TABLE lists;
