-- name: FindTaskById :one
SELECT * FROM tasks WHERE id = $1;

-- name: FindAllTasksByUserId :many
SELECT * FROM tasks WHERE user_id = $1 ORDER BY created_at DESC;

-- name: CreateTask :one
INSERT INTO tasks (user_id, title, description, status, priority, due_date)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks SET title = $2, description = $3, status = $4, priority = $5, due_date = $6, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: FindAllTasksByUserIdAndStatus :many
SELECT * FROM tasks WHERE user_id = $1 AND status = $2 ORDER BY created_at DESC;

-- name: FindAllTasksByUserIdAndPriority :many
SELECT * FROM tasks WHERE user_id = $1 AND priority = $2 ORDER BY created_at DESC;

-- name: DeleteTaskById :exec
DELETE FROM tasks WHERE id = $1;

-- name: CountTasksByUserIdGroupedByStatus :many
SELECT status, COUNT(*) as count FROM tasks WHERE user_id = $1 GROUP BY status;
