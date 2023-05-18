-- name: CreatePost :one
INSERT INTO posts (
    title,
    content,
    category,
    user_id
) VALUES (
     $1, $2, $3, $4
)RETURNING *;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET
    title = coalesce(sqlc.narg('title'), title),
    category = coalesce(sqlc.narg('category'), category),
    content = coalesce(sqlc.narg('content'), content)
WHERE id = sqlc.arg('id') RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;