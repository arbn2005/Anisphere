CREATE TYPE watch_status AS ENUM (
    'PLAN_TO_WATCH',
    'WATCHING',
    'ON_HOLD',
    'COMPLETED',
    'DROPPED'
);

CREATE TABLE watch_progress (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    anime_id INT NOT NULL REFERENCES anime(id) ON DELETE CASCADE,

    season_id INT REFERENCES seasons(id) ON DELETE CASCADE,

    episode_id INT REFERENCES episodes(id) ON DELETE CASCADE,

    status watch_status NOT NULL DEFAULT 'PLAN_TO_WATCH',

    episodes_watched INT DEFAULT 0,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(user_id, anime_id)
);
