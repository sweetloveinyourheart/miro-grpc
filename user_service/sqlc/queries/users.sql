-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, first_name, last_name)
VALUES ($1, $2, $3)
RETURNING user_id;

-- name: CreateUserCredential :exec
INSERT INTO user_credentials (user_id, password_hash)
VALUES ($1, $2);