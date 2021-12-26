-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1 LIMIT 1;

-- name: InsertUser :one
INSERT INTO users (email, password)
VALUES (@email, @password) RETURNING *;

