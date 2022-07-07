-- name: ListMovies :many
SELECT * FROM favorite_movies;

-- name: AddMovie :one
-- INSERT INTO favorite_movies VALUES ($1, $2);