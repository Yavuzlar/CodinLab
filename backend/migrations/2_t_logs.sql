-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_logs (
    id TEXT PRIMARY KEY NOT NULL,
    user_id TEXT NOT NULL,
    language_id INTEGER NOT NULL,
    type TEXT NOT NULL,
    lab_path_id INTEGER,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES t_users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_logs;
-- +goose StatementEnd
