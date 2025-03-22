-- Create table for storing problem solutions
CREATE TABLE problem_solutions (
    id SERIAL PRIMARY KEY,
    problem_id int4 NOT NULL REFERENCES problems(id),
    language varchar NOT NULL,
    solution_code text NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_problem_language UNIQUE (problem_id, language)
);

-- Create index for faster lookups by problem_id
CREATE INDEX idx_problem_solutions_problem_id ON problem_solutions(problem_id);