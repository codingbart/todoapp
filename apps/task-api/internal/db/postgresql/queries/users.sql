-- name: FindUserById :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (keycloak_id)
VALUES ($1)
RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;
