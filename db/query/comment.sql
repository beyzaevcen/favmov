-- name: GetComments :many
SELECT comments.id, comments.content, comments.created_at, users.name, users.image
FROM comments 
INNER JOIN users ON users.id = comments.user_id
WHERE comments.movie_id = $1;

-- name: AddComment :one
INSERT INTO comments(user_id, movie_id, content) VALUES ($1, $2, $3) RETURNING *;