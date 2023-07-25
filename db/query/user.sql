-- name: RegisterUser :one
INSERT INTO users (name, email) VALUES ($1, $2) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING id;

-- name: UpdateNameOfUser :one
UPDATE users SET "name" = $1 WHERE id = $2 RETURNING *;  