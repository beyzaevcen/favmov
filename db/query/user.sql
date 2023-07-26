-- name: RegisterUser :one
INSERT INTO users (name, image, password_hash, email ) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING id;

-- name: UpdateNameOfUser :one
UPDATE users SET "name" = $1 WHERE id = $2 RETURNING *;

-- name: GetImageAndNameOfUser :one
SELECT name, image FROM users WHERE id = $1 LIMIT 1;