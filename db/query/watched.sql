-- name: GetWatchedMovies :many
SELECT movies.* 
FROM watched_movie
INNER JOIN movies ON movies.id = watched_movie.movie_id
WHERE watched_movie.user_id = $1;

-- name: AddToWatchedMovies :one
INSERT INTO watched_movie(user_id, movie_id) VALUES ($1, $2) RETURNING *;

-- name: DeleteFromWatchedMovies :one
DELETE FROM watched_movie 
WHERE movie_id = $1 
AND user_id = $2 RETURNING *;
