import React from 'react';
import { View, Text, ScrollView, ActivityIndicator } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { useReviews } from '../services/api/reviews';
import ProblemCard from '../../components/ui/ProblemCard';

export default function ReviewsScreen() {
  // Fetch the 5 latest reviews using the useReviews hook
  const { data: reviews, isLoading, error } = useReviews(1); // Assuming user_id is 1
  
  // Use actual data if available, otherwise fallback to sample data
  const displayReviews = reviews;
  
  return (
    <View className="flex-1 bg-[#131C24] dark justify-between">
      <View>
        {/* Header */}
        <View className="flex items-center bg-[#131C24] p-4 pb-2 mb-4 justify-between">
          <Text className="text-[#F8F9FB] text-lg font-bold leading-tight tracking-[-0.015em] flex-1 text-center pl-12 pr-12">
            Upcoming Reviews
          </Text>
        </View>
        
        {/* Upcoming Reviews Section */}
        
        {isLoading ? (
          <View className="flex items-center justify-center py-10">
            <ActivityIndicator size="large" color="#6366F1" />
          </View>
        ) : error ? (
          <View className="p-4 items-center">
            <Text className="text-[#8A9DC0] text-base text-center">
              Error loading reviews. Please try again.
            </Text>
          </View>
        ) : displayReviews? (
          <ScrollView contentContainerStyle={{ paddingBottom: 20 }}>
            {displayReviews.map((review) => (
              <ProblemCard
                key={review.id}
                title={review.title}
                subtitle="Not Reviewed"
                iconName="calendar-outline"
                completed={false}
              />
            ))}
          </ScrollView>
        ): null}
      </View>
    </View>
  );
}