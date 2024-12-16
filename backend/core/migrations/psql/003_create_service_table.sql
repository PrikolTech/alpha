CREATE TYPE service_type as ENUM ();

CREATE TABLE IF NOT EXISTS service(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    type service_type NOT NULL,
    settings JSONB,
    is_active BOOLEAN DEFAULT FALSE,
    code VARCHAR(255) UNIQUE NOT NULL
);