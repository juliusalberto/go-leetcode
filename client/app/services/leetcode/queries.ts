import { leetCodeGraphQL } from './api';
import { Submission, StreakData, StreakCounter, UserCalendar} from './types';

// Fetch recent successful submissions
export const fetchRecentSubmissions = async (username: string, limit = 15): Promise<Submission[]> => {
  const query = `
    query recentAcSubmissions($username: String!, $limit: Int!) {
      recentAcSubmissionList(username: $username, limit: $limit) {
        id
        title
        titleSlug
        timestamp
      }
    }
  `;
  
  const variables = { username, limit };
  const data = await leetCodeGraphQL(query, variables);
  return data.data?.recentAcSubmissionList || [];
};

export const fetchUserProblemProfile = async (username: string): Promise<Map<string, number>> => {
    const query = `
      query userProblemsSolved($username: String!) {
        matchedUser(username: $username) {
          submitStatsGlobal {
            acSubmissionNum {
              difficulty
              count
            }
          }
        }
      }
    `

    const variables = { username };
    const data = await leetCodeGraphQL(query, variables);

    interface ProblemCount {
      difficulty: string;
      count: number;
    }
    
    // Extract the array of problem counts
    const problemCounts: ProblemCount[] = data.data?.matchedUser?.submitStatsGlobal?.acSubmissionNum || [];

    // turn it into a map[key, int]

    const problemCountMap = new Map<string, number>();

    problemCounts.forEach(element => {
      problemCountMap.set(element.difficulty, element.count)
    });

    return problemCountMap;
}

// Fetch user's streak data
export const fetchStreakData = async (username: string): Promise<StreakData> => {
  // Get current streak
  const streakQuery = `
    query getStreakCounter {
      streakCounter {
        streakCount
        daysSkipped
        currentDayCompleted
      }
    }
  `;
  
  // Get calendar data for visualization
  const calendarQuery = `
    query userProfileCalendar($username: String!, $year: Int) {
      matchedUser(username: $username) {
        userCalendar(year: $year) {
          activeYears
          streak
          totalActiveDays
          submissionCalendar
        }
      }
    }
  `;
  
  const currentYear = new Date().getFullYear();
  const variables = { username, year: currentYear };
  
  // Run both queries
  const [streakData, calendarData] = await Promise.all([
    leetCodeGraphQL(streakQuery),
    leetCodeGraphQL(calendarQuery, variables)
  ]);
  
  return {
    currentStreak: streakData.data?.streakCounter?.streakCount || 0,
    calendar: calendarData.data?.matchedUser?.userCalendar || {}
  };
};

// Format timestamps for display
export const formatTimeAgo = (timestamp: number): string => {
  const now = new Date();
  const submissionDate = new Date(timestamp * 1000); // Convert to milliseconds
  const diffDays = Math.floor((now.getTime() - submissionDate.getTime()) / (1000 * 60 * 60 * 24));
  
  if (diffDays === 0) return 'Today';
  if (diffDays === 1) return 'Yesterday';
  if (diffDays < 7) return `${diffDays} days ago`;
  
  return submissionDate.toLocaleDateString();
};