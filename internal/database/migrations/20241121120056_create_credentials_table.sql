-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS credentials (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    public_key BYTEA NOT NULL,
    credential_id BYTEA NOT NULL,
    sign_count INTEGER NOT NULL,
    aaguid BYTEA,
    clone_warning BOOLEAN NOT NULL DEFAULT FALSE,
    attachment TEXT,
    backup_eligible BOOLEAN NOT NULL DEFAULT FALSE,
    backup_state BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT unique_credential_id UNIQUE(user_id, credential_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS credentials;
-- +goose StatementEnd
