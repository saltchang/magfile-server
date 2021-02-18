-- name: CreateToken :one
INSERT INTO token (
  user_id,
  access_token,
  expired_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetToken :one
SELECT * FROM token
WHERE user_id = $1
LIMIT 1;

-- name: UpdateToken :one
UPDATE token
SET (
  access_token,
  expired_at
) = (
  $2, $3
)
WHERE user_id = $1
RETURNING *;

-- name: DeleteToken :exec
DELETE FROM token
WHERE user_id = $1;
