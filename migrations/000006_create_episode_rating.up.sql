CREATE TABLE episode_ratings (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    episode_id INT NOT NULL REFERENCES episodes(id) ON DELETE CASCADE,

    rating SMALLINT NOT NULL
        CHECK (rating BETWEEN 1 AND 5),

    review TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(user_id, episode_id)
);
