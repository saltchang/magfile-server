-- name: CreatePostTag :one
INSERT INTO post_tag (
  post_id,
  tag_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAllTagFromPost :many
SELECT * FROM post_tag
WHERE post_id = $1
ORDER BY created_at;

-- name: GetAllPostFromTag :many
SELECT * FROM post_tag
WHERE tag_id = $1
ORDER BY created_at;

-- name: UpdatePostTag :one
UPDATE post_tag
SET (
  post_id,
  tag_id
) = (
  $2, $3
)
WHERE id = $1
RETURNING *;

-- name: DeletePostTag :exec
DELETE FROM post_tag
WHERE id = $1;
