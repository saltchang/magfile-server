-- name: CreateSeries :one
INSERT INTO series (
  semantic_id,
  author_id,
  title,
  abstract,
  is_archived,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetSeries :one
SELECT * FROM series
WHERE id = $1
LIMIT 1;

-- name: GetSeriesBySemanticID :one
SELECT * FROM series
WHERE author_id = $1
AND semantic_id = $2
LIMIT 1;

-- name: GetAllSeriesFromAuthor :many
SELECT * FROM series
WHERE author_id = $1
ORDER BY created_at;

-- name: UpdateSeries :one
UPDATE series
SET (
  semantic_id,
  author_id,
  title,
  abstract,
  is_archived,
  updated_at
) = (
  $2, $3, $4, $5, $6, $7
)
WHERE id = $1
RETURNING *;

-- name: DeleteSeries :exec
DELETE FROM series
WHERE id = $1;
