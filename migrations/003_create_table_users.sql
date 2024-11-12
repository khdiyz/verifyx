-- +goose Up
CREATE TABLE IF NOT EXISTS "users"(
    "id" UUID PRIMARY KEY,
    "first_name" VARCHAR(64) NOT NULL,
    "last_name" VARCHAR(64) NOT NULL,
    "department_id" UUID NOT NULL,
    "phone_number" VARCHAR(64) NOT NULL,
    "profile_image" VARCHAR(64),
    "face_embedding" JSONB,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    FOREIGN KEY (department_id) REFERENCES departments(id)
);

-- +goose Down
DROP TABLE IF EXISTS "users";
