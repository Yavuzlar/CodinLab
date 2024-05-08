-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_log_type (
    id TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_log_type;
-- +goose StatementEnd
