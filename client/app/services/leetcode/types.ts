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
