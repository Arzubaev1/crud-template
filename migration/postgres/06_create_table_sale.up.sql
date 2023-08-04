CREATE TABLE "sales"(
    id UUID PRIMARY KEY,
    user_id UUID,
    total INT,
    count INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
);