CREATE TABLE problems (
    id int PRIMARY KEY,
    frontend_id int NOT NULL,
    title VARCHAR NOT NULL,
    title_slug VARCHAR NOT NULL UNIQUE,
    difficulty VARCHAR NOT NULL,
    is_paid_only BOOLEAN NOT NULL,
    content TEXT NOT NULL,
    topic_tags JSONB,
    example_testcases TEXT,
    similar_questions JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);