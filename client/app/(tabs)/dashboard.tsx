import React from 'react';
import { View, Text, ScrollView, TouchableOpacity, Dimensions, ActivityIndicator } from 'react-native';
import { useFonts, Roboto_400Regular, Roboto_500Medium, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { router } from 'expo-router';
import { LineChart } from 'react-native-chart-kit';
import { formatTimeAgo } from '../../services/leetcode/queries';
import { format, subDays } from 'date-fns';
import InfoCard from '../../components/ui/InfoCard';
import StatCard from '../../components/ui/StatCard';
import { useRecentSubmissions, useStreakData, useUserProblemProfile } from '../../services/leetcode/hooks';
import { useAuth } from '../../contexts/AuthContext';
import { useReviews } from '../../services/api/reviews';

export default function DashboardScreen() {
  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_500Medium,
    Roboto_700Bold,
  });

  const { session, leetcodeUsername } = useAuth();
  const username = leetcodeUsername || 'elhazen'; // Fallback to example if not available
  const screenWidth = Dimensions.get('window').width - 40;
  
  // Use Tanstack Query hooks
  const {
    data: recentSubmissions = [],
    isLoading: submissionsLoading
  } = useRecentSubmissions(username);
  
  const {
    data: streakDataResponse,
    isLoading: streakLoading
  } = useStreakData(username);
  
  const {
    data: userProblemProfile = new Map<string, number>(),
    isLoading: profileLoading
  } = useUserProblemProfile(username);
  
  // Fetch due reviews count - this uses Tanstack Query caching
  const {
    data: reviewsData,
    isLoading: reviewsLoading
  } = useReviews();
  
  // Calculate number of due reviews from the cached data
  const dueReviewsCount = reviewsData?.pages?.[0]?.length || 0;

  // Helper function to calculate streak
  const calculateStreak = (dateToCountMap: Map<string, number> | undefined) => {
    let streak = 0;
    let currentDate = new Date();
    currentDate.setHours(0, 0, 0, 0);
    
    // If dateToCountMap is undefined, return 0
    if (!dateToCountMap) {
      return streak;
    }
    
    // Start with today and go backwards
    while (true) {
      const dateKey = format(subDays(currentDate, 1), 'yyyy-MM-dd');
      
      // Check if this date has submissions
      if (dateToCountMap.has(dateKey) && dateToCountMap.get(dateKey)! > 0) {
        streak++;
        // Move to previous day
        currentDate = subDays(currentDate, 1);
      } else {
        // Break the streak when we find a day with no submissions
        break;
      }
    }
    
    return streak;
  };

  // Process streak data
  const streakHistory = [0, 0, 0, 0, 0, 0, 0]; // Default empty data
  let currentStreak = 0;
  
  if (streakDataResponse?.calendar?.submissionCalendar) {
    const calendar = JSON.parse(streakDataResponse.calendar.submissionCalendar);
    
    // Get last 7 days
    const dateToCountMap = new Map();
    for (const timestamp in calendar) {
      const date = new Date(parseInt(timestamp) * 1000)
      const dateKey = format(date, "yyyy-MM-dd");
      dateToCountMap.set(dateKey, calendar[timestamp])
    }

    const today = new Date();
    today.setHours(0, 0, 0, 0);

    // Fill streak history
    for (let i = 6; i >= 0; i--) {
      const date = subDays(today, i);
      const dateKey = format(date, "yyyy-MM-dd");
      const count = dateToCountMap.get(dateKey) || 0
      streakHistory[6-i] = count;
    }

    // Calculate current streak
    currentStreak = calculateStreak(dateToCountMap);
  }
  
  // Combine loading states
  const loading = submissionsLoading || streakLoading || profileLoading || reviewsLoading || !fontsLoaded;
  
  if (!fontsLoaded) {
    return null;
  }

  const handleProblemPress = (problem: { slug: string }) => {
    // Navigate to problem details
    router.push(`/problem/${problem.slug}`);
  };
  
  const handleReviewsPress = () => {
    // Navigate to reviews tab
    router.push('/(tabs)/reviews');
  };

  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Header */}
      <View className="flex-row items-center bg-[#131C24] p-4 pb-2 justify-between">
        <Text 
          className="text-[#F8F9FB] text-lg font-bold leading-tight flex-1 text-center px-12"
          style={{ fontFamily: 'Roboto_700Bold', letterSpacing: -0.24 }}
        >
          Dashboard
        </Text>
      </View>

      <ScrollView>
        {/* Current Streak */}
        <InfoCard
          icon="flame-outline"
          title="Current Streak"
          subtitle={`${currentStreak} days`}
        />

        {/* Upcoming Reviews - Clickable with actual count */}
        <TouchableOpacity onPress={handleReviewsPress}>
          <InfoCard
            icon="calendar-outline"
            title="Upcoming Reviews"
            subtitle={`${dueReviewsCount} reviews`}
          />
        </TouchableOpacity>

        {/* Recently Attempted */}
        <Text 
          className="text-[#F8F9FB] text-[22px] font-bold leading-tight px-4 pb-3 pt-5"
          style={{ fontFamily: 'Roboto_700Bold', letterSpacing: -0.24 }}
        >
          Recently Attempted
        </Text>

        {loading ? (
          <View className="flex-1 justify-center items-center py-4">
            <ActivityIndicator size="small" color="#6366F1" />
          </View>
        ): recentSubmissions.length > 0? ( 
          recentSubmissions.slice(0, 4).map((submission) => (
            <TouchableOpacity 
              key={submission.id}
              className="flex-row items-center gap-4 bg-[#131C24] px-4 min-h-14 justify-between"
              onPress={() => handleProblemPress({ slug: submission.titleSlug })}
            >
              <Text 
                className="text-[#F8F9FB] text-base leading-normal flex-1 truncate"
                style={{ fontFamily: 'Roboto_400Regular' }}
              >
                {submission.title}
              </Text>
              <View className="shrink-0">
                <Text 
                  className="text-[#8A9DC0] text-sm leading-normal"
                  style={{ fontFamily: 'Roboto_400Regular' }}
                >
                  {formatTimeAgo(submission.timestamp)}
                </Text>
              </View>
            </TouchableOpacity>
        ))
        ): (
          <Text className='text-[#8A9DC0] text-base px-4 py-2' style={{ fontFamily: 'Roboto_400Regular'}}>
            No recent submissions found
          </Text>
        )}

        {/* Stats Cards */}
        <View className="flex-row flex-wrap gap-4 p-4">
          <StatCard 
            title="Total Problems Solved" 
            value={userProblemProfile.get("All") || 0} 
          />
          <StatCard 
            title="Problems Solved This Week" 
            value={streakHistory.reduce((acc, curr) => acc + curr, 0)} 
          />
        </View>
        
        {/* Streak Progress Chart - Simple Placeholder */}
        <View className="mx-4 mb-6 p-4 rounded-xl border border-[#32415D]">
          <Text 
            className="text-[#F8F9FB] text-base leading-normal mb-4"
            style={{ fontFamily: 'Roboto_500Medium' }}
          >
            Streak Progress (Last 7 Days)
          </Text>
          
          
          <View className="h-[200px] items-center justify-center">
            <LineChart
                  data={{
                    labels: ['', '', '', '', '', ''],
                    datasets: [
                      {
                        data: streakHistory
                      }
                    ]
                  }}
                  width={screenWidth}
                  height={180}
                  yAxisInterval={2}
                  chartConfig={{
                    backgroundColor: 'transparent',
                    backgroundGradientFrom: 'transparent',
                    backgroundGradientTo: 'transparent',
                    decimalPlaces: 0,
                    color: (opacity = 1) => `rgba(99, 102, 241, ${opacity})`,
                    labelColor: (opacity = 1) => `rgba(138, 157, 192, ${opacity})`,
                    style: {
                      borderRadius: 16
                    },
                    propsForDots: {
                      r: '5',
                      strokeWidth: '2',
                      stroke: '#6366F1'
                    },
                    
                  }}
                  bezier
                  style={{
                    marginVertical: 8,
                    borderRadius: 16
                  }}
              />
          </View>
          
          <View className="flex-row justify-between ">
            <Text 
              className="text-[#8A9DC0] text-xs"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              7 days ago
            </Text>
            <Text 
              className="text-[#8A9DC0] text-xs"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              Today
            </Text>
          </View>
        </View>
      </ScrollView>
    </View>
  );
}