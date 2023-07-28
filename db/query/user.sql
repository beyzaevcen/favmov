-- name: RegisterUser :one
INSERT INTO users 
(name, firebase_uid, image, email )
VALUES ($1, $2, $3, $4) 
RETURNING id, name, image, email, created_at;

-- name: GetUsers :many
SELECT * FROM users;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1 RETURNING id;

-- name: UpdateNameOfUser :one
UPDATE users SET "name" = $1 
WHERE id = $2 RETURNING *;

-- name: GetImageAndNameOfUser :one
SELECT name, image FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = sqlc.arg(email) ::text LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = sqlc.arg(user_id) ::bigint LIMIT 1;

-- name: GetUserIDByUID :one
SELECT id FROM users WHERE firebase_uid = $1 LIMIT 1;