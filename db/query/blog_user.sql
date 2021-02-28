-- name: CreateBlogUser :one
INSERT INTO blog_user (
  username,
  email,
  full_name,
  gender,
  current_location,
  password_hash,
  logined_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetBlogUser :one
SELECT * FROM blog_user
WHERE id = $1 LIMIT 1;

-- name: GetBlogUserByUsername :one
SELECT * FROM blog_user
WHERE username = $1 LIMIT 1;

-- name: GetBlogUserByEmail :one
SELECT * FROM blog_user
WHERE email = $1 LIMIT 1;

-- name: UpdateBlogUser :one
UPDATE blog_user
SET (
  username,
  email,
  full_name,
  gender,
  current_location,
  password_hash,
  logined_at
) = (
  $2, $3, $4, $5, $6, $7, $8
)
WHERE id = $1
RETURNING *;

-- name: DeleteBlogUser :exec
DELETE FROM blog_user
WHERE id = $1;
