CREATE TABLE season_ratings (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    season_id INT NOT NULL REFERENCES seasons(id) ON DELETE CASCADE,

    rating SMALLINT NOT NULL
        CHECK (rating BETWEEN 1 AND 5),

    review TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(user_id, season_id)
);
