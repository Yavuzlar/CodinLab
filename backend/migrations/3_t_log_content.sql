-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_log_content (
    id TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    log_type_id TEXT NOT NULL,
    FOREIGN KEY (log_type_id) REFERENCES t_log_type(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_log_content;
-- +goose StatementEnd
