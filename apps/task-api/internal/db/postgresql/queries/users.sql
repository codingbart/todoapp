-- name: FindUserById :one
SELECT * FROM users WHERE id = $1;

-- name: FindUserByKeycloakId :one
SELECT * FROM users WHERE keycloak_id = $1;

-- name: SaveUser :one
INSERT INTO users (keycloak_id, name, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;
