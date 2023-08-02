
CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "login" VARCHAR(24) NOT NULL,
    "password" VARCHAR(24) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
