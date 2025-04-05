import React from 'react';
import { View, Text, FlatList, TouchableOpacity, ActivityIndicator, Alert } from 'react-native'; // Added Alert
import { useDeckProblems, useStartPracticePublicDeck } from '../../services/api/decks'; // Import useStartPracticePublicDeck
import { Ionicons } from '@expo/vector-icons';
import { router, useLocalSearchParams } from 'expo-router';
import { Problem } from '../../services/leetcode/types';
import ProblemCard from '../../components/ui/ProblemCard'; // Import ProblemCard

// Function to get difficulty color (re-added)
const getDifficultyColor = (difficulty: string) => {
  switch (difficulty) {
    case 'Easy': return '#4CD137';
    case 'Medium': return '#F39C12';
    case 'Hard': return '#E74C3C';
    default: return '#8A9DC0'; // Default color
  }
};

export default function DeckDetailScreen() {
  // Read route parameters (id) and query parameters (name, is_public)
  const { id, name: deckNameParam, is_public: isPublicParam } = useLocalSearchParams<{ id: string; name?: string; is_public?: string }>();
  const deckId = parseInt(id as string);
  const deckName = deckNameParam || 'Deck Details'; // Default name if not passed
  const isPublic = isPublicParam === 'true'; // Convert string back to boolean

  const {
    data, // This is now an object { pages: Problem[][], pageParams: any[] }
    fetchNextPage,
    hasNextPage,
    isLoading: isLoadingProblems, // Initial load state
    isFetchingNextPage, // State for loading more pages
  } = useDeckProblems(deckId);
  const startPracticeMutation = useStartPracticePublicDeck(); // Initialize the new mutation hook

  // Flatten the pages data into a single array for the FlatList
  const problems = React.useMemo(() => data?.pages.flatMap(page => page) ?? [], [data]);

  // Updated render function using ProblemCard
  const renderProblemItem = ({ item }: {item: Problem}) => {
    return (
      <View className="mb-2 rounded-lg overflow-hidden"> {/* Add rounded corners and overflow hidden */}
        <ProblemCard
          title={item.title}
          subtitle={item.difficulty} // Use difficulty as subtitle
          subtitleColor={getDifficultyColor(item.difficulty)} // Use dynamic color based on difficulty
          backgroundColor="#1E2A3A" // Explicitly set the lighter background color
          onPress={() => router.push(`/problem/${item.title_slug}`)}
          // iconName can be customized if needed, e.g., based on completion status if available
          // completed={item.isCompleted} // Example if completion status exists
        />
      </View>
    );
  };
  
  if (isLoadingProblems) { // Use the renamed variable
    return (
      <View className="flex-1 justify-center items-center bg-[#131C24]">
        <ActivityIndicator size="large" color="#6366F1" />
      </View>
    );
  }

  // Define the handler function *before* the return statement
  const handlePracticePress = async () => {
    if (isPublic) {
      try {
        // Call the mutation to ensure flashcards exist for this public deck
        await startPracticeMutation.mutateAsync(deckId);
        // Navigate only after successful preparation
        router.push(`/flashcards?deckId=${deckId}`);
      } catch (err) { // Use 'err' to avoid conflict with Error type
        console.error("Failed to prepare public deck for practice:", err);
        Alert.alert("Error", "Could not prepare the deck for practice. Please try again.");
      }
    } else {
      // For private decks, navigate directly
      router.push(`/flashcards?deckId=${deckId}`);
    }
  };
  
  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Header */}
      <View className="flex-row items-center p-4">
        <TouchableOpacity onPress={() => router.back()}>
          <Ionicons name="chevron-back" size={24} color="#F8F9FB" />
        </TouchableOpacity>
        {/* Use the passed deck name */}
        <Text className="text-[#F8F9FB] text-xl font-bold ml-4 flex-1">
          {deckName}
        </Text>
      </View>
      
      <FlatList
        data={problems}
        renderItem={renderProblemItem}
        keyExtractor={(item) => item.id.toString()}
        contentContainerStyle={{ padding: 16 }}
        ListHeaderComponent={
          // Corrected structure: Single row, justify-between, items-center
          <View className="mb-4 flex-row justify-between items-center">
            {/* Left side: Title and count */}
            <View className="pl-4"> {/* Added pl-4 for left padding */}
              {/* Use the passed deck name here as well */}
              <Text className="text-[#F8F9FB] text-2xl font-bold mb-1">
                {deckName}
              </Text>
              {/* Display fetched problem count. Note: This won't show the *total* count until all pages are loaded. */}
              {/* Consider fetching total count separately if needed immediately. */}
              <Text className="text-[#8A9DC0] text-sm">
                {problems.length} {problems.length === 1 ? 'problem' : 'problems'} loaded
              </Text>
            </View>
            {/* Right side: Button */}
            <TouchableOpacity
              className={`px-4 py-2 rounded-lg ${startPracticeMutation.isPending ? 'bg-gray-500' : 'bg-[#6366F1]'}`} // Style change on loading
              onPress={handlePracticePress} // Use the new handler
              disabled={startPracticeMutation.isPending} // Disable button while mutating
            >
              {startPracticeMutation.isPending ? (
                <ActivityIndicator size="small" color="#FFFFFF" />
              ) : (
                <Text className="text-white font-medium">Practice</Text>
              )}
            </TouchableOpacity>
          </View>
        }
        ListEmptyComponent={
          <View className="flex-1 justify-center items-center p-4">
            <Text className="text-[#8A9DC0] text-center">
              No problems in this deck yet.
            </Text>
          </View>
        }
      />
    </View>
  );
}