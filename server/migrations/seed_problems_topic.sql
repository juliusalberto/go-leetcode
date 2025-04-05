-- Sample data for problems_topic table
-- Adjust problem_id and topic_slug based on your seed_test_problems.sql

INSERT INTO public.problems_topic (problem_id, topic_slug) VALUES (1, 'array') ON CONFLICT DO NOTHING;
INSERT INTO public.problems_topic (problem_id, topic_slug) VALUES (1, 'hash-table') ON CONFLICT DO NOTHING;

INSERT INTO public.problems_topic (problem_id, topic_slug) VALUES (15, 'array') ON CONFLICT DO NOTHING;
INSERT INTO public.problems_topic (problem_id, topic_slug) VALUES (15, 'two-pointers') ON CONFLICT DO NOTHING;
INSERT INTO public.problems_topic (problem_id, topic_slug) VALUES (15, 'sorting') ON CONFLICT DO NOTHING;
