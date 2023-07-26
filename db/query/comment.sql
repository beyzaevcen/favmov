-- name: GetComments :many
SELECT comments.id, comments.content, comments.created_at, users.name, users.image
FROM comments 
INNER JOIN users ON users.id = comments.user_id
WHERE comments.movie_id = $1;

-- name: AddComment :one
INSERT INTO comments(user_id, movie_id, content) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteComment :one
DELETE FROM comments 
WHERE comments.id = $1 
AND user_id = $2 RETURNING id;

-- name: GetMyComments :many
SELECT comments.id, comments.content, comments.created_at
FROM comments 
INNER JOIN movies ON movies.id = comments.movie_id
WHERE comments.movie_id = $1 AND comments.user_id = $2;
