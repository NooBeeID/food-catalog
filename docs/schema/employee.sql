
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY NOT NULL,
    nip VARCHAR,
    name VARCHAR,
    address VARCHAR,
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);