CREATE TYPE integration_state as ENUM ();

CREATE TABLE IF NOT EXISTS integration(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES project(id),
    service_id UUID NOT NULL REFERENCES service(id),
    is_active BOOLEAN DEFAULT FALSE,
    state integration_state NOT NULL
);