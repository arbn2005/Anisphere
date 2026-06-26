CREATE TABLE anime (
    id SERIAL PRIMARY KEY,

    title TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,

    description TEXT,

    poster_url TEXT,
    banner_url TEXT,

    status TEXT,
    release_year INT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
