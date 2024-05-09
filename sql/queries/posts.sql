-- name: CreatePost :one
INSERT INTO posts (
        id,
        created_at,
        updated_at,
        title,
        description,
        published_at,
        feed_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7
    )
RETURNING *;
--
-- name: GetPostsForUser :many
SELECT posts.*
FROM posts
    JOIN feedfollow ON feedfollow.feed_id = posts.feed_id
WHERE feedfollow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;