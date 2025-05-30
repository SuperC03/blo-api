-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS api_key (
  id          BIGSERIAL PRIMARY KEY,
  title       TEXT NOT NULL,
  description TEXT,
  key         TEXT UNIQUE NOT NULL,
  permissions TEXT[],
  quota       BIGINT NOT NULL,
  usage       BIGINT DEFAULT 0 NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE api_key;
-- +goose StatementEnd
