-- name: CreateTag :one
INSERT INTO tag (
  author_id,
  name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetTag :one
SELECT * FROM tag
WHERE id = $1
LIMIT 1;

-- name: GetTagByAuthorAndName :one
SELECT * FROM tag
WHERE author_id = $1
AND name = $2
LIMIT 1;

-- name: GetAllTagFromAuthor :many
SELECT * FROM tag
WHERE author_id = $1
ORDER BY created_at;

-- name: UpdateTag :one
UPDATE tag
SET (
  author_id,
  name
) = (
  $2, $3
)
WHERE id = $1
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tag
WHERE id = $1;
