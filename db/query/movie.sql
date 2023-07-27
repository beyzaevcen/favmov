-- name: GetMovies :many
SELECT * FROM movies;

-- name: AddMovie :one
INSERT INTO movies (title, description, score, image, release_date) VALUES ($1,$2,$3,$4,$5) RETURNING *;

-- name: DeleteMovie :one
DELETE FROM movies 
WHERE id = $1 RETURNING id;

-- name: EditMovie :one
UPDATE movies SET "description" = $1, "score" = $2, "image" = $3  
WHERE id = $4 RETURNING title, description, score, image, release_date;  

-- name: GetMovie :one
SELECT title, description, score, image, release_date 
FROM movies 
WHERE id =$1;