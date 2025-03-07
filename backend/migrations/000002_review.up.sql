CREATE TABLE review_schedules (
    id SERIAL PRIMARY KEY,
    submission_id VARCHAR REFERENCES submissions(id),
    next_review_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- FSRS specific fields
    stability FLOAT NOT NULL DEFAULT 0.0,
    difficulty FLOAT NOT NULL DEFAULT 0.0,
    elapsed_days BIGINT NOT NULL DEFAULT 0,
    scheduled_days BIGINT NOT NULL DEFAULT 0,
    reps BIGINT NOT NULL DEFAULT 0,
    lapses BIGINT NOT NULL DEFAULT 0,
    state SMALLINT NOT NULL DEFAULT 0,
    last_review TIMESTAMP DEFAULT NULL
);