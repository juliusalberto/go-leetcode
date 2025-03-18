export interface Submission {
  id: string;
  title: string;
  titleSlug: string;
  timestamp: number;
}

export interface StreakData {
  currentStreak: number;
  calendar: {
    activeYears?: number[];
    streak?: number;
    totalActiveDays?: number;
    submissionCalendar?: string;
  };
}

export interface StreakCounter {
  streakCount: number;
  daysSkipped: number;
  currentDayCompleted: boolean;
}

export interface UserCalendar {
  activeYears: number[];
  streak: number;
  totalActiveDays: number;
  submissionCalendar: string;
}

export interface TopicTag {
  name: string;
  slug: string;
}

export interface Problem {
  id: number;
  frontend_id: number;
  title: string;
  title_slug: string;
  difficulty: string;
  is_paid_only: boolean;
  user_rating?: number;
  content?: string;
  topic_tags: TopicTag[];
  example_testcases?: string;
  similar_questions?: any[];
  created_at: string;
  completed?: boolean;
}

export interface ProblemResponse {
  data: Problem[];
  meta: {
    pagination: {
      total: number;
      page: number;
      per_page: number;
    };
    timestamp: string;
  };
  errors: any[];
}