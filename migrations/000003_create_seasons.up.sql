CREATE TABLE seasons (
    id SERIAL PRIMARY KEY,

    anime_id INT NOT NULL REFERENCES anime(id) ON DELETE CASCADE,

    season_number INT NOT NULL,

    title TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(anime_id, season_number)
);
