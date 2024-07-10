-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserInfoWithCredentials :one
SELECT users.user_id as user_id, users.email as email, user_credentials.password_hash as pwd FROM users
INNER JOIN user_credentials ON users.user_id = user_credentials.user_id
WHERE users.email = $1 LIMIT 1;

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