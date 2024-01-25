CREATE TABLE IF NOT EXISTS "project"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(64) NOT NULL,
    "describe" VARCHAR(255),
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "task"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(64) NOT NULL,
    "describe" VARCHAR(255),
    "is_completed" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deadline" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "project_id" INTEGER NOT NULL,

    FOREIGN KEY("project_id") REFERENCES "project"("id")
);