-- +goose Up
CREATE TABLE IF NOT EXISTS boards (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  owner_id TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WIth TIME ZONE DEFAULT CURRENT_TIMESTAMP,

  UNIQUE (name, owner_id)
);

-- +goose Down
DROP TABLE boards;
