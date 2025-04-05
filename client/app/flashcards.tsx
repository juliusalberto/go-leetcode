import React, { useState, useEffect } from 'react';
// Import Alert from react-native
import { View, Text, TouchableOpacity, ScrollView, ActivityIndicator, Platform, Alert } from 'react-native';
import { useFlashcardReviews, useSubmitFlashcardRating } from '../services/api/flashcards';
import { useDecks } from '../services/api/decks'; // Keep for potential future use (e.g., deck selector)
import { useSolutionsApi } from '../services/api/solutions';
import { WebView } from 'react-native-webview';
import { Ionicons } from '@expo/vector-icons';
import { router, useLocalSearchParams } from 'expo-router';
import CodeHighlighter from '../components/CodeHighlighter';
import RatingButtons from '../components/RatingButtons';
import ScreenHeader from '../components/ui/ScreenHeader';
import Button from '../components/ui/Button'; // Use consistent Button component
import LanguageTabs from '../components/ui/LanguageTabs'; // Import the new component

export default function FlashcardsScreen() {
  const params = useLocalSearchParams<{ deckId?: string }>();
  const initialDeckIdFromQuery = params.deckId ? parseInt(params.deckId, 10) : undefined;

  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [selectedDeckId, setSelectedDeckId] = useState<number | undefined>(
    initialDeckIdFromQuery && !isNaN(initialDeckIdFromQuery) ? initialDeckIdFromQuery : undefined
  );
  const [showSolutionApproach, setShowSolutionApproach] = useState(false);
  const [showFullSolution, setShowFullSolution] = useState(false);
  const [reviewCompleted, setReviewCompleted] = useState(false);
  const [selectedLanguage, setSelectedLanguage] = useState('python');
  const [solutions, setSolutions] = useState<Record<string, string>>({});
  const [solutionsLoading, setSolutionsLoading] = useState(false);
  const [isSubmittingRating, setIsSubmittingRating] = useState(false); // <-- Keep loading state
  
  const solutionsApi = useSolutionsApi();
  // Get infinite query properties
  const {
    data,
    isLoading: isLoadingReviews, // Rename to avoid conflict if needed later
    refetch,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage
  } = useFlashcardReviews(selectedDeckId);
  const submitRating = useSubmitFlashcardRating();

  // Flatten the pages into a single array of flashcards
  const flashcards = React.useMemo(() => data?.pages.flatMap(page => page.reviews) || [], [data]);
  const totalDueCount = data?.pages[0]?.total ?? 0; // Get total from the first page if available
  const currentCard = flashcards[currentCardIndex];

  // Proactively fetch next page when user gets close to the end
  useEffect(() => {
    const threshold = 3; // Fetch when 3 cards or less remaining in current set
    const remainingInSet = flashcards.length - currentCardIndex;

    if (remainingInSet <= threshold && hasNextPage && !isFetchingNextPage) {
      console.log('Proactively fetching next page...'); // Optional: for debugging
      fetchNextPage();
    }
  }, [currentCardIndex, flashcards.length, hasNextPage, isFetchingNextPage, fetchNextPage]);


  // Fetch solutions when card changes
  useEffect(() => {
    const fetchSolutions = async () => {
      if (currentCard?.problem?.id) {
        setSolutionsLoading(true);
        setSolutions({});
        try {
          const solutionsData = await solutionsApi.fetchSolutionByID(currentCard.problem.id.toString());
          setSolutions(solutionsData);
          const availableLanguages = Object.keys(solutionsData);
          if (availableLanguages.length > 0 && !solutionsData[selectedLanguage]) {
            setSelectedLanguage(availableLanguages[0]); // Default to first available if current selection is invalid
          }
        } catch (error) {
          console.error('Error fetching solutions:', error);
          setSolutions({});
        } finally {
          setSolutionsLoading(false);
        }
      } else {
        setSolutions({});
        setSolutionsLoading(false);
      }
    };
    fetchSolutions();
  }, [currentCard]); // Re-fetch when currentCard changes

  // Handle rating submission
  const handleRating = async (rating: 1 | 2 | 3 | 4) => {
    // Prevent submission if no card or already submitting
    if (!currentCard || isSubmittingRating) return;
 
    setIsSubmittingRating(true); // <-- Set loading state
 
    try {
      // Submit the rating (invalidation is now removed from the hook)
      await submitRating.mutateAsync({ review_id: currentCard.id, rating });
 
      // --- Immediate UI Update ---
      // Check if it's the last card in the currently loaded list
      const isLastCardInLoadedSet = currentCardIndex === flashcards.length - 1;
 
      if (isLastCardInLoadedSet) {
        // If it was the last loaded card, check if more pages might exist
        // (based on hasNextPage flag from useInfiniteQuery)
        if (!hasNextPage) {
          // No more pages expected, review is complete for now
          setReviewCompleted(true);
        } else {
          // More pages might exist, but we don't advance index here.
          // The proactive fetch useEffect (lines 49-57) should handle fetching.
          // We might stay on the "last" card briefly until new data loads.
          // Alternatively, could show a specific "fetching next batch..." message.
          console.log("Rated last loaded card, waiting for proactive fetch...");
        }
      } else {
        // Not the last card in the set, advance index immediately
        setCurrentCardIndex(prev => prev + 1);
        setShowSolutionApproach(false); // Reset view state for next card
        setShowFullSolution(false);
      }
      // --- End Immediate UI Update ---
 
    } catch (error) {
      console.error('Error submitting rating:', error);
      Alert.alert("Error", "Failed to submit rating. Please try again.");
    } finally {
      setIsSubmittingRating(false); // <-- Reset loading state regardless of outcome
    }
  };
 
  // Reset state and refetch reviews
  const resetAndRefetch = () => {
    setReviewCompleted(false);
    setCurrentCardIndex(0);
    setShowSolutionApproach(false);
    setShowFullSolution(false);
    // Call the refetch function from useInfiniteQuery
    // It will refetch from the first page.
    refetch();
  };
 
  // --- Render Logic ---
 
  // Loading State (Use the renamed variable)
  if (isLoadingReviews && !data) { // Check if loading initial data
    return (
      <View className="flex-1 justify-center items-center bg-[#131C24]">
        <ActivityIndicator size="large" color="#6366F1" />
      </View>
    );
  }

  // Completion State
  if (reviewCompleted || flashcards.length === 0) {
    return (
      <View className="flex-1 justify-center items-center bg-[#131C24] p-6">
        <Ionicons name="checkmark-done-circle" size={64} color="#4CD137" />
        <Text className="text-[#F8F9FB] text-2xl font-semibold mt-5 text-center">
          All Done!
        </Text>
        <Text className="text-[#ADBAC7] text-base mt-2 text-center mb-8">
          You've completed all available reviews for now.
        </Text>
        <View className="flex-row space-x-3 w-full max-w-xs">
          <Button
            title="Browse Decks"
            onPress={() => router.push('/(tabs)/decklist')}
            variant="primary"
            className="flex-1"
          />
          <Button
            title="Review Again"
            onPress={resetAndRefetch}
            variant="secondary"
            className="flex-1"
          />
        </View>
      </View>
    );
  }

  // Main Flashcard View
  return (
    <View className="flex-1 bg-[#131C24]">
      <ScreenHeader
        title="Flashcard Review"
        rightElement={
          <Text className="text-[#8A9DC0] text-sm font-medium">
            {currentCardIndex + 1} / {flashcards.length}
          </Text>
        }
      />

      <ScrollView className="flex-1" contentContainerStyle={{ paddingBottom: 24 }}>
        <View className="px-4 pt-4">

          {/* Problem Title & Difficulty */}
          <View className="mb-5">
            <Text className="text-[#F8F9FB] text-xl font-semibold mb-2">
              {currentCard.problem.title}
            </Text>
            <View className="flex-row">
              <Text className={`text-xs font-medium px-2 py-1 rounded-full ${
                currentCard.problem.difficulty === 'Easy' ? 'bg-[#4CD137]/20 text-[#4CD137]' :
                currentCard.problem.difficulty === 'Medium' ? 'bg-[#F39C12]/20 text-[#F39C12]' :
                'bg-[#E74C3C]/20 text-[#E74C3C]'
              }`}>
                {currentCard.problem.difficulty}
              </Text>
            </View>
          </View>

          {/* Problem Description (Only show if solution approach is hidden) */}
          {!showSolutionApproach && (
            <View className="bg-[#1E2A3A] rounded-lg p-4 mb-5">
              {Platform.OS === 'web' ? (
                <div
                  style={{ color: '#D1D5DB', fontFamily: '-apple-system, sans-serif', fontSize: '15px', lineHeight: 1.6 }}
                  dangerouslySetInnerHTML={{
                    __html: `
                      <style>
                        body { margin: 0; padding: 0; }
                        p { margin-bottom: 1em; }
                        code { font-family: monospace; background-color: #29374C; padding: 0.2em 0.4em; border-radius: 4px; font-size: 0.9em; }
                        pre { background-color: #131C24; padding: 1em; border-radius: 6px; overflow-x: auto; font-family: monospace; font-size: 0.9em; }
                        strong { font-weight: 600; }
                        ul, ol { padding-left: 1.5em; margin-bottom: 1em; }
                        li { margin-bottom: 0.5em; }
                      </style>
                      ${currentCard?.problem?.content || ''}
                    `
                  }}
                />
              ) : (
                <WebView
                  originWhitelist={['*']}
                  source={{
                    html: `
                    <html><head><meta name="viewport" content="width=device-width, initial-scale=1.0">
                    <style>
                      body { font-family: -apple-system, sans-serif; margin: 0; padding: 0; color: #D1D5DB; background-color: #1E2A3A; font-size: 15px; line-height: 1.6; word-wrap: break-word; overflow-wrap: break-word; }
                      p { margin-bottom: 1em; }
                      code { font-family: monospace; background-color: #29374C; padding: 0.2em 0.4em; border-radius: 4px; font-size: 0.9em; }
                      pre { background-color: #131C24; padding: 1em; border-radius: 6px; overflow-x: auto; font-family: monospace; font-size: 0.9em; }
                      strong { font-weight: 600; }
                      ul, ol { padding-left: 1.5em; margin-bottom: 1em; }
                      li { margin-bottom: 0.5em; }
                    </style></head>
                    <body>${currentCard?.problem?.content || ''}</body></html>`
                  }}
                  style={{ backgroundColor: '#1E2A3A', minHeight: 150 }} // Set minHeight for native WebView
                  scalesPageToFit={false}
                />
              )}
            </View>
          )}

          {/* --- Solution Section --- */}

          {/* Reveal Solution Approach Button (Only show if solution approach is hidden) */}
          {!showSolutionApproach && (
            <Button
              title="Show Solution Approach"
              onPress={() => setShowSolutionApproach(true)}
              variant="primary"
              className="mb-6"
            />
          )}

          {/* Solution Revealed Content (Only show if solution approach is shown) */}
          {showSolutionApproach && (
            <>
              {/* Solution Approach Display */}
              <View className="bg-[#1E2A3A] p-4 rounded-lg mb-4">
                <Text className="text-[#ADBAC7] text-sm font-medium mb-1">
                  Solution Approach
                </Text>
                <Text className="text-[#F8F9FB] text-base">
                  {currentCard.problem.solution_approach || "Approach details not available."}
                </Text>
              </View>

              {/* Hide Solution Button */}
              <Button
                title="Hide Solution"
                onPress={() => {
                  setShowSolutionApproach(false);
                  setShowFullSolution(false); // Also hide full solution
                }}
                variant="secondary"
                className="mb-4"
              />

              {/* Use LanguageTabs component */}
              <LanguageTabs
                availableLanguages={Object.keys(solutions)}
                selectedLanguage={selectedLanguage}
                onSelectLanguage={setSelectedLanguage}
                loading={solutionsLoading}
              />

              {/* Full Solution Code Button/Content */}
              {!showFullSolution ? (
                 <Button
                  title="Show Full Solution Code"
                  onPress={() => setShowFullSolution(true)}
                  variant="secondary"
                  className="mb-6" // Margin before rating buttons
                />
              ) : solutionsLoading ? (
                <View className="bg-[#1E2A3A] rounded-lg p-4 mb-6 items-center justify-center min-h-[100px]">
                  <ActivityIndicator size="small" color="#6366F1" />
                  <Text className="text-[#8A9DC0] mt-2 text-sm">Loading solution...</Text>
                </View>
              ) : solutions[selectedLanguage] ? (
                <CodeHighlighter
                  language={selectedLanguage}
                  style={{ marginBottom: 24 }} // Keep margin bottom before rating buttons
                >
                  {solutions[selectedLanguage]}
                </CodeHighlighter>
              ) : (
                <View className="bg-[#1E2A3A] rounded-lg p-4 mb-6">
                  <Text className="text-[#ADBAC7] text-sm">
                    Solution code not available for {selectedLanguage === 'cpp' ? 'C++' : selectedLanguage}.
                  </Text>
                </View>
              )}

              {/* Rating Buttons - Pass disabled state */}
              <RatingButtons onRate={handleRating} disabled={isSubmittingRating} />
            </>
          )}
        </View>
      </ScrollView>
    </View>
  );
}