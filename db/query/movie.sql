-- name: GetMovies :many
SELECT * FROM movies;

-- name: AddMovie :one
INSERT INTO movies (title, description, score, image, release_date) VALUES ($1,$2,$3,$4,$5) RETURNING *;