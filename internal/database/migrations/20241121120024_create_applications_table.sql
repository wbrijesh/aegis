-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS applications (
    id TEXT PRIMARY KEY,
    developer_id TEXT NOT NULL,
    name TEXT NOT NULL,
    rp_id TEXT NOT NULL UNIQUE,
    rp_display_name TEXT,
    rp_origins TEXT[] NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (developer_id) REFERENCES developers(id),
    CONSTRAINT unique_application_name UNIQUE(developer_id, name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS applications;
-- +goose StatementEnd
