-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    application_id TEXT NOT NULL,
    name TEXT NOT NULL,
    display_name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (application_id) REFERENCES applications(id),
    CONSTRAINT unique_user_name UNIQUE(application_id, name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
