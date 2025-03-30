-- Create review_logs table for tracking review history
CREATE TABLE review_logs (
	id serial4 NOT NULL,
	review_schedule_id int4 NOT NULL,
	rating int2 NOT NULL,
	review_date timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	elapsed_days int4 NOT NULL,
	scheduled_days int4 NOT NULL,
	state int2 NOT NULL,
	CONSTRAINT review_logs_pkey PRIMARY KEY (id),
	CONSTRAINT review_logs_rating_check CHECK ((rating = ANY (ARRAY[1, 2, 3, 4]))),
	CONSTRAINT review_logs_state_check CHECK ((state = ANY (ARRAY[0, 1, 2, 3])))
);

-- Add foreign key constraints
ALTER TABLE review_logs ADD CONSTRAINT fk_review_schedule FOREIGN KEY (review_schedule_id) REFERENCES review_schedules(id) ON DELETE CASCADE;