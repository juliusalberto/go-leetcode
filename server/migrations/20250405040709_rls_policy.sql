-- Enable RLS for tables that store user-specific data or need controlled access
ALTER TABLE public.users ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.submissions ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.flashcard_reviews ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.flashcard_review_logs ENABLE ROW LEVEL SECURITY;

-- Optional: Enable RLS for generally public data for consistency/safety
-- If enabled, requires SELECT policies below. If disabled, access is unrestricted.
-- Uncomment the lines below if you want RLS on these public tables
-- ALTER TABLE public.problems ENABLE ROW LEVEL SECURITY;
-- ALTER TABLE public.problem_solutions ENABLE ROW LEVEL SECURITY;
-- ALTER TABLE public.problems_topic ENABLE ROW LEVEL SECURITY;

-- Note: RLS is already enabled for 'decks' and 'deck_problems' in your 20250405 migration.
-- Note: Consider if 'review_schedules' and 'review_logs' are still used. If so, enable RLS:
-- ALTER TABLE public.review_schedules ENABLE ROW LEVEL SECURITY;
-- ALTER TABLE public.review_logs ENABLE ROW LEVEL SECURITY;


-- Policies for 'users' table
CREATE POLICY "Allow users to view their own data" ON public.users
    FOR SELECT USING (auth.uid() = id);
CREATE POLICY "Allow users to update their own data" ON public.users
    FOR UPDATE USING (auth.uid() = id) WITH CHECK (auth.uid() = id);
-- Consider if users should be able to delete their own accounts via API
-- CREATE POLICY "Allow users to delete their own account" ON public.users
--     FOR DELETE USING (auth.uid() = id);


-- Policies for 'submissions' table
CREATE POLICY "Allow users to view their own submissions" ON public.submissions
    FOR SELECT USING (auth.uid() = user_id);
CREATE POLICY "Allow users to insert their own submissions" ON public.submissions
    FOR INSERT WITH CHECK (auth.uid() = user_id);
-- Add UPDATE/DELETE if needed, e.g.:
-- CREATE POLICY "Allow users to delete their own submissions" ON public.submissions
--     FOR DELETE USING (auth.uid() = user_id);


-- Policies for 'decks' table (RLS already enabled)
CREATE POLICY "Allow users to view own and public decks" ON public.decks
    FOR SELECT USING (auth.uid() = user_id OR is_public = true);
CREATE POLICY "Allow users to insert their own decks" ON public.decks
    FOR INSERT WITH CHECK (auth.uid() = user_id);
CREATE POLICY "Allow users to update their own decks" ON public.decks
    FOR UPDATE USING (auth.uid() = user_id) WITH CHECK (auth.uid() = user_id);
CREATE POLICY "Allow users to delete their own decks" ON public.decks
    FOR DELETE USING (auth.uid() = user_id);


-- Policies for 'deck_problems' table (RLS already enabled)
-- Allow viewing problems if the user can view the parent deck
CREATE POLICY "Allow users to view problems in accessible decks" ON public.deck_problems
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM public.decks d
            WHERE d.id = deck_problems.deck_id AND (d.user_id = auth.uid() OR d.is_public = true)
        )
    );
-- Allow inserting/deleting problems only if the user owns the parent deck
CREATE POLICY "Allow users to insert problems into own decks" ON public.deck_problems
    FOR INSERT WITH CHECK (
        EXISTS (
            SELECT 1 FROM public.decks d
            WHERE d.id = deck_problems.deck_id AND d.user_id = auth.uid()
        )
    );
CREATE POLICY "Allow users to delete problems from own decks" ON public.deck_problems
    FOR DELETE USING (
        EXISTS (
            SELECT 1 FROM public.decks d
            WHERE d.id = deck_problems.deck_id AND d.user_id = auth.uid()
        )
    );


-- Policies for 'flashcard_reviews' table
CREATE POLICY "Allow users to view their own flashcard reviews" ON public.flashcard_reviews
    FOR SELECT USING (auth.uid() = user_id);
CREATE POLICY "Allow users to insert their own flashcard reviews" ON public.flashcard_reviews
    FOR INSERT WITH CHECK (auth.uid() = user_id);
CREATE POLICY "Allow users to update their own flashcard reviews" ON public.flashcard_reviews
    FOR UPDATE USING (auth.uid() = user_id) WITH CHECK (auth.uid() = user_id);
-- Add DELETE if needed


-- Policies for 'flashcard_review_logs' table
-- Allow viewing logs if the user owns the parent review
CREATE POLICY "Allow users to view logs for their own reviews" ON public.flashcard_review_logs
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM public.flashcard_reviews fr
            WHERE fr.id = flashcard_review_logs.flashcard_review_id AND fr.user_id = auth.uid()
        )
    );
-- Allow inserting logs if the user owns the parent review (server usually handles this)
CREATE POLICY "Allow users to insert logs for their own reviews" ON public.flashcard_review_logs
    FOR INSERT WITH CHECK (
        EXISTS (
            SELECT 1 FROM public.flashcard_reviews fr
            WHERE fr.id = flashcard_review_logs.flashcard_review_id AND fr.user_id = auth.uid()
        )
    );
-- Logs are typically immutable, so UPDATE/DELETE might not be needed.


-- Policies for public data tables (only needed if RLS was enabled above)
-- Uncomment the policies below if you uncommented the ALTER TABLE statements above
-- CREATE POLICY "Allow authenticated users to read problems" ON public.problems
--     FOR SELECT USING (auth.role() = 'authenticated');

-- CREATE POLICY "Allow authenticated users to read problem solutions" ON public.problem_solutions
--     FOR SELECT USING (auth.role() = 'authenticated');

-- CREATE POLICY "Allow authenticated users to read problem topics" ON public.problems_topic
--     FOR SELECT USING (auth.role() = 'authenticated');


-- Policies for older review tables (if RLS enabled and still used)
-- Assuming review_schedules links to submissions which links to users
CREATE POLICY "Allow users to view own review schedules" ON public.review_schedules
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM public.submissions s
            WHERE s.id = review_schedules.submission_id AND s.user_id = auth.uid()
        )
    );

CREATE POLICY "Allow users to update own review schedules" ON public.review_schedules
    FOR UPDATE USING (
        EXISTS (
            SELECT 1 FROM public.submissions s
            WHERE s.id = review_schedules.submission_id AND s.user_id = auth.uid()
        )
    );

CREATE POLICY "Allow users to insert own review schedules" ON public.review_schedules
    FOR INSERT WITH CHECK (
        EXISTS (
            SELECT 1 FROM public.submissions s
            WHERE s.id = review_schedules.submission_id AND s.user_id = auth.uid()
        )
    );

CREATE POLICY "Allow users to delete own review schedules" ON public.review_schedules
    FOR DELETE USING (
        EXISTS (
            SELECT 1 FROM public.submissions s
            WHERE s.id = review_schedules.submission_id AND s.user_id = auth.uid()
        )
    );


-- Add INSERT/UPDATE/DELETE as needed, checking ownership via submissions table.

-- Assuming review_logs links to review_schedules
CREATE POLICY "Allow users to view own review logs" ON public.review_logs
    FOR SELECT USING (
        EXISTS (
            SELECT 1 FROM public.review_schedules rs
            JOIN public.submissions s ON rs.submission_id = s.id
            WHERE rs.id = review_logs.review_schedule_id AND s.user_id = auth.uid()
        )
    );