CREATE TABLE review_schedules (
	id serial4 NOT NULL,
	submission_id varchar NULL,
	next_review_at timestamp NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	stability float8 DEFAULT 0.0 NULL,
	difficulty float8 DEFAULT 0.0 NULL,
	elapsed_days int4 DEFAULT 0 NULL,
	scheduled_days int4 DEFAULT 0 NULL,
	reps int4 DEFAULT 0 NULL,
	lapses int4 DEFAULT 0 NULL,
	state int2 DEFAULT 0 NULL,
	last_review timestamp NULL,
	CONSTRAINT review_schedules_pkey PRIMARY KEY (id)
);


ALTER TABLE review_schedules ADD CONSTRAINT review_schedules_submission_id_fkey FOREIGN KEY (submission_id) REFERENCES submissions(id);