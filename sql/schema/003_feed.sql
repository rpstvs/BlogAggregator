-- +goose Up
CREATE TABLE feed(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID,
    url text UNIQUE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose Down
DROP TABLE feed;