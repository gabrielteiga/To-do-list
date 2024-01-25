-- name: CreateTask :exec
INSERT INTO task (title, describe, is_completed, created_at, deadline, project_id) VALUES ($1,$2,$3,$4,$5,$6);

-- name: GetTask :one
SELECT * FROM task WHERE id = $1;

-- name: GetTasks :many
SELECT * FROM task;

-- name: CreateProject :exec
INSERT INTO project (title, describe, created_at) VALUES ($1, $2, $3);

-- name: GetProject :one
SELECT * FROM project WHERE id = $1;

-- name: GetProjects :many
SELECT * FROM project;