

CREATE TABLE IF NOT EXISTS menus (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR,
    category varchar,
    description varchar,
    price int,
    image_url varchar,
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);