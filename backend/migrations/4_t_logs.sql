-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_logs (
    id TEXT PRIMARY KEY NOT NULL,
    title TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id TEXT NOT NULL,
    log_content_id TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES t_users(id) ON DELETE CASCADE,
    FOREIGN KEY (log_content_id) REFERENCES t_log_content(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_logs;
-- +goose StatementEnd
