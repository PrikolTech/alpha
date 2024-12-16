CREATE TABLE IF NOT EXISTS project (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    is_archived BOOLEAN DEFAULT FALSE,
    code VARCHAR(255) UNIQUE NOT NULL
);