-- name: CreateUser :one
INSERT INTO feed (id, created_at, updated_at, url, user_id)
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
    )
RETURNING *;