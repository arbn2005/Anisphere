CREATE TABLE episodes (
    id SERIAL PRIMARY KEY,

    season_id INT NOT NULL REFERENCES seasons(id) ON DELETE CASCADE,

    episode_number INT NOT NULL,

    title TEXT NOT NULL,

    synopsis TEXT,

    duration INT,

    air_date DATE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(season_id, episode_number)
);
