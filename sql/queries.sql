-- name: CreateTask :exec
INSERT INTO task (title, describe, is_completed, created_at, deadline, project_id) VALUES ($1,$2,$3,$4,$5,$6);

-- name: DeleteTask :exec
DELETE FROM task WHERE id = $1 AND project_id = $2;

-- name: GetTask :one
SELECT * FROM task WHERE id = $1;

-- name: GetTasks :many
SELECT * FROM task;

-- name: GetTasksByProject :many
SELECT * FROM task WHERE project_id = $1 ORDER BY is_completed;

-- name: MarkTaskAsCompleted :exec
UPDATE task SET is_completed = true WHERE id = $1 AND project_id = $2;

-- name: CreateProject :exec
INSERT INTO project (title, describe, created_at) VALUES ($1, $2, $3);

-- name: DeleteProject :exec
DELETE FROM project WHERE id = $1;

-- name: GetProject :one
SELECT * FROM project WHERE id = $1;

-- name: GetProjects :many
SELECT * FROM project;