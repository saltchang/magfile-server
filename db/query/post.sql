-- name: CreatePost :one
INSERT INTO post (
  semantic_id,
  author_id,
  series_id,
  order_in_series,
  title,
  abstract,
  content,
  views,
  is_archived,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: GetPost :one
SELECT * FROM post
WHERE id = $1
LIMIT 1;

-- name: GetPostBySemanticID :one
SELECT * FROM post
WHERE author_id = $1
AND semantic_id = $2
LIMIT 1;

-- name: GetAllPostFromAuthor :many
SELECT * FROM post
WHERE author_id = $1
ORDER BY created_at;

-- name: UpdatePost :one
UPDATE post
SET (
  semantic_id,
  author_id,
  series_id,
  order_in_series,
  title,
  abstract,
  content,
  views,
  is_archived,
  updated_at
) = (
  $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM post
WHERE id = $1;
