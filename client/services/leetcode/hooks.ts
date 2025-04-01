import { useQuery } from '@tanstack/react-query';
import { fetchRecentSubmissions, fetchStreakData, fetchUserProblemProfile } from './queries';

// Hook for recent submissions
export const useRecentSubmissions = (username: string, limit = 15) => {
  return useQuery({
    queryKey: ['recentSubmissions', username, limit],
    queryFn: () => fetchRecentSubmissions(username, limit),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
};

// Hook for user problem profile
export const useUserProblemProfile = (username: string) => {
  return useQuery({
    queryKey: ['userProblemProfile', username],
    queryFn: () => fetchUserProblemProfile(username),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
};

// Hook for streak data
export const useStreakData = (username: string) => {
  return useQuery({
    queryKey: ['streakData', username],
    queryFn: () => fetchStreakData(username),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
};