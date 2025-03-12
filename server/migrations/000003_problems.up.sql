-- public.problems definition

-- Drop table

-- DROP TABLE public.problems;

CREATE TABLE problems (
	id int4 NOT NULL,
	frontend_id int4 NOT NULL,
	title varchar NOT NULL,
	title_slug varchar NOT NULL,
	difficulty varchar NOT NULL,
	is_paid_only bool NOT NULL,
	"content" text NOT NULL,
	topic_tags jsonb NULL,
	example_testcases text NULL,
	similar_questions jsonb NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT problems_pkey PRIMARY KEY (id),
	CONSTRAINT problems_title_slug_key UNIQUE (title_slug)
);