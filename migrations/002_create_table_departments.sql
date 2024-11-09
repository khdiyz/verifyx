-- +goose Up
CREATE TABLE IF NOT EXISTS "departments"(
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS "departments";
