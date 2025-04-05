import React, { useCallback, useRef, useState } from 'react';
import { View, Text, FlatList, ActivityIndicator, TouchableOpacity, Animated, Alert } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { useAuth } from '../../contexts/AuthContext'; // <-- Import useAuth
import { useReviews, Review } from '../../services/api/reviews';
import ProblemCard from '../../components/ui/ProblemCard';
import { useQueryClient } from '@tanstack/react-query';
import { useSubmissionsApi, SubmissionRequest } from '../../services/api/submissions';
import { MenuProvider, Menu, MenuTrigger, MenuOptions, MenuOption } from 'react-native-popup-menu';
import Toast from 'react-native-toast-message';

export default function ReviewsScreen() {
  const queryClient = useQueryClient();
  const submissionsApi = useSubmissionsApi();
  const [removingIds, setRemovingIds] = useState<Set<string>>(new Set());
  const { session } = useAuth(); // <-- Get session
  
  // Fetch reviews using infinite query, enable only when session exists
  const {
    data,
    isLoading,
    error,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage
  } = useReviews({ enabled: !!session }); // <-- Pass enabled option

  // Debug the data and error
  console.log("Reviews data:", data);
  if (error) console.error("Reviews error:", error);

  // Animation values for removing reviews
  const animatedValues = useRef<Map<string, Animated.Value>>(new Map());
  
  // Get animation value for a review
  const getAnimatedValue = (reviewId: string) => {
    if (!animatedValues.current.has(reviewId)) {
      animatedValues.current.set(reviewId, new Animated.Value(1));
    }
    return animatedValues.current.get(reviewId)!;
  };

  // Mark review as completed
  const markAsCompleted = useCallback(async (review: Review) => {
    try {
      setRemovingIds(prev => new Set(prev).add(review.id));
      
      // Create animation
      const animValue = getAnimatedValue(review.id);
      
      // Start fade out animation
      Animated.timing(animValue, {
        toValue: 0,
        duration: 500,
        useNativeDriver: true
      }).start();
      
      // TODO: change the title_slug: review.title to review.title_slug (probably have to return the title slug as well)
      const submission: SubmissionRequest = {
        is_internal: true,
        title: review.title,
        title_slug: review.title_slug,
        submitted_at: new Date().toISOString()
      }
      
      // Process submission
      await submissionsApi.createSubmission(submission);
      
      // Show success toast
      Toast.show({
        type: 'success',
        text1: 'Review Completed',
        text2: `Successfully marked "${review.title}" as reviewed`,
      });
      
      // After animation completes, invalidate query
      setTimeout(() => {
        // Invalidate queries to refresh data
      queryClient.invalidateQueries({ queryKey: ['recentReviews'] });
      
      // Remove from removing set after animation
      setRemovingIds(prev => {
        const newSet = new Set(prev);
        newSet.delete(review.id);
        return newSet;
      });
      }, 500);
    } catch (error) {
      console.error('Error marking review as completed:', error);
      
      Toast.show({
        type: 'error',
        text1: 'Error',
        text2: 'Failed to mark review as completed. Please try again.',
      });
      
      setRemovingIds(prev => {
        const newSet = new Set(prev);
        newSet.delete(review.id);
        return newSet;
      });
    }
  }, [queryClient]);
  
  // Flatten pages data for FlatList and ensure we never have null/undefined pages
  const reviews: Review[] = data?.pages?.flatMap(page => page || []) || [];
  
  // Log the flattened reviews for debugging
  console.log("Flattened reviews:", reviews?.length);
  
  // Handle end reached - load more data
  const handleLoadMore = () => {
    if (hasNextPage && !isFetchingNextPage) {
      fetchNextPage();
    }
  };
  
  // Render a review card with menu options
  const renderReviewCard = ({ item: review }: { item: Review }) => {
    const isRemoving = removingIds.has(review.id);
    const animValue = getAnimatedValue(review.id);
    
    return (
      <Animated.View style={{ opacity: animValue, transform: [{ scale: animValue }] }}>
        <ProblemCard
          key={review.id}
          title={review.title}
          subtitle="Next Review: Due Now"
          iconName="calendar-outline"
          completed={false}
          rightElement={
            <Menu>
              <MenuTrigger>
                <View className="p-2 flex items-center justify-center">
                  {isRemoving ? (
                    <ActivityIndicator size="small" color="#6366F1" />
                  ) : (
                    <Ionicons name="ellipsis-vertical" size={20} color="#8A9DC0" />
                  )}
                </View>
              </MenuTrigger>
              <MenuOptions>
                <MenuOption
                  style={{ borderRadius: 100 }}
                  onSelect={() => markAsCompleted(review)}
                  disabled={isRemoving}
                >
                  <Text className={'flex items-center justify-center p-2 text-black'}>
                    Mark Reviewed
                  </Text>
                </MenuOption>
              </MenuOptions>
            </Menu>
          }
        />
      </Animated.View>
    );
  };
  
  return (
    <MenuProvider>
      <View className="flex-1 bg-[#131C24] dark justify-between">
        <View className="flex-1">
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
              <Text className="text-[#8A9DC0] text-xs text-center mt-2">
                {error.message}
              </Text>
            </View>
          ) : (
            <FlatList
              data={reviews}
              renderItem={renderReviewCard}
              keyExtractor={(item) => item.id}
              contentContainerStyle={{ paddingBottom: 20 }}
              onEndReached={handleLoadMore}
              onEndReachedThreshold={0.5}
              ListFooterComponent={
                isFetchingNextPage ? (
                  <View className="py-4">
                    <ActivityIndicator size="small" color="#6366F1" />
                  </View>
                ) : null
              }
              ListEmptyComponent={
                !isLoading && (
                  <View className="p-4 items-center justify-center">
                    <Text className="text-[#8A9DC0] text-base text-center">
                      No reviews due. Great job!
                    </Text>
                  </View>
                )
              }
            />
          )}
        </View>
      </View>
      <Toast />
    </MenuProvider>
  );
}