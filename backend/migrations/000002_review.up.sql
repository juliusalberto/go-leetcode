CREATE TABLE review_schedules (
    id SERIAL PRIMARY KEY,
    submission_id VARCHAR REFERENCES submissions(id),
    next_review_at TIMESTAMP NOT NULL,
    interval_days INTEGER NOT NULL,  -- current interval (1,3,7,14,30)
    times_reviewed INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
