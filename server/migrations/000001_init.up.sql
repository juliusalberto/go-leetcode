CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	email varchar,
	username varchar NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	leetcode_username varchar NOT NULL,
	CONSTRAINT unique_leetcode_username UNIQUE (leetcode_username),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_username_key UNIQUE (username)
);

CREATE TABLE submissions (
	id varchar NOT NULL,
	title varchar NOT NULL,
	title_slug varchar NOT NULL,
	submitted_at timestamp NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	user_id uuid NOT NULL,
	CONSTRAINT submissions_pkey PRIMARY KEY (id)
);

ALTER TABLE submissions ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id);

CREATE INDEX idx_submissions_user_time ON submissions(user_id, submitted_at);