import React, { useState, useCallback } from 'react';
import { View, Text, TextInput, TouchableOpacity, FlatList, ActivityIndicator, Modal } from 'react-native';
import { useFonts, Roboto_400Regular, Roboto_500Medium, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { Ionicons } from '@expo/vector-icons';
import ProblemCard from '../../components/ui/ProblemCard';
import { router } from 'expo-router';
import { MenuProvider, Menu, MenuTrigger, MenuOptions, MenuOption } from 'react-native-popup-menu';
import Toast from 'react-native-toast-message';
import debounce from 'lodash/debounce';
import { useInfiniteQuery, useQueryClient, QueryKey, InfiniteData } from '@tanstack/react-query'; // Import QueryKey and InfiniteData types
import { TOPIC_DISPLAY_NAMES, TOPIC_OPTIONS } from '../../constants/topics';
import { useAuth } from '../../contexts/AuthContext'; // <-- Import useAuth
 
// Import types and API hooks
import { Problem, ProblemWithStatus, ProblemWithStatusResponse } from '../../services/leetcode/types'; // Assuming ProblemWithStatusResponse is defined here, e.g., { data: ProblemWithStatus[], total?: number }
import { useProblemsApi } from '../../services/api/problems';
import { useSubmissionsApi } from '../../services/api/submissions';
import { useDecks, useAddProblemToDeck } from '../../services/api/decks';
import DropdownFilter from "../../components/ui/DropdownFilter";
import ScreenHeader from '../../components/ui/ScreenHeader'; // Import the new header

// Problem difficulty colors
const difficultyColors: Record<string, string> = {
  Easy: '#1CBABA', // Green
  Medium: '#FFB700', // Orange
  Hard: '#F63737'  // Red
};

export default function ProblemsScreen() {
  console.log("ProblemsScreen rendering");
  const queryClient = useQueryClient();
  const problemsApi = useProblemsApi();
  const submissionsApi = useSubmissionsApi();
  const { session } = useAuth(); // <-- Get session
  
  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_500Medium,
    Roboto_700Bold,
  });
  console.log("Fonts loaded:", fontsLoaded);

  // States
  const [searchInput, setSearchInput] = useState('');
  const [searchQuery, setSearchQuery] = useState('');

  // Filter states
  const [difficulty, setDifficulty] = useState<string | null>(null);
  const [tags, setTags] = useState<string | null>(null);

  const structuredTagOptions = TOPIC_DISPLAY_NAMES.map(displayName => ({
    label: displayName,
    value: TOPIC_OPTIONS[displayName as keyof typeof TOPIC_OPTIONS]
  }));

  const structuredDiffOptions = ["Easy", "Medium", "Hard"].map(displayName => ({
    label: displayName,
    value: displayName
  }))

  // Define the query key (ensure it matches the one in useInfiniteQuery)
  const queryKey: QueryKey = ['problems', { difficulty, tags, searchQuery }];

  const {
    data,
    fetchNextPage,
    hasNextPage,
    isFetching,
    isFetchingNextPage,
     // Type useInfiniteQuery with the *page* response type, error type, and the full InfiniteData structure for selection/mutation
  } = useInfiniteQuery<ProblemWithStatusResponse, Error, InfiniteData<ProblemWithStatusResponse>, QueryKey, number>({
    queryKey: queryKey,
    queryFn: ({ pageParam = 0 }) => // pageParam is correctly inferred as number here
      problemsApi.fetchProblemsWithStatus({
        difficulty: difficulty || undefined,
        tags: tags || undefined,
        search: searchQuery,
        offset: pageParam, // Use pageParam directly
        limit: 20,
      }),
    getNextPageParam: (lastPage, allPages) => {
       // lastPage is now correctly typed as ProblemWithStatusResponse
       if (lastPage?.data?.length === 20) {
         // Calculate the next offset based on the total number of items fetched across all pages
         const totalFetched = allPages.reduce((acc, page) => acc + (page?.data?.length || 0), 0);
         return totalFetched; // This becomes the next pageParam
       }
       return undefined; // No next page
    },
    initialPageParam: 0, // The initial pageParam (offset)
    enabled: !!session,
  });

  // Submission state
  const [submittingProblemId, setSubmittingProblemId] = useState<number | null>(null);
  
  // Deck modal state
  const [showDeckModal, setShowDeckModal] = useState(false);
  const [selectedProblem, setSelectedProblem] = useState<ProblemWithStatus | null>(null);
  const { data: decksData } = useDecks();
  const addProblemToDeck = useAddProblemToDeck();

  // Create debounced search function using lodash
  const debouncedSearch = useCallback(
    debounce((value: string) => {
      setSearchQuery(value);
    }, 800),
    []
  );

  // Handle adding a submission for a problem
  // Modified handleAddSubmission with Optimistic Update
  const handleAddSubmission = async (problemWithStatus: ProblemWithStatus) => {
    const problem = problemWithStatus.problem;
    setSubmittingProblemId(problem.id); // Show loading indicator immediately

    // 1. Cancel any outgoing refetches so they don't overwrite our optimistic update
    await queryClient.cancelQueries({ queryKey: queryKey });

    // 2. Snapshot the previous value
    // 2. Snapshot the previous value using InfiniteData<PageResponseType>
    const previousProblemsData = queryClient.getQueryData<InfiniteData<ProblemWithStatusResponse>>(queryKey);

    // 3. Optimistically update to the new value
    // 3. Optimistically update to the new value
    if (previousProblemsData) {
      const newData: InfiniteData<ProblemWithStatusResponse> = {
        ...previousProblemsData,
        // Map over the pages array within the InfiniteData structure
        pages: previousProblemsData.pages.map(page => ({
          ...page,
          // Map over the data array within each page
          data: page.data.map((item: ProblemWithStatus) => // item is ProblemWithStatus here
            item.problem.id === problem.id
              ? { ...item, completed: true } // Mark the specific problem as completed
              : item
          ),
        })),
      };
      // Set the updated data structure back into the cache
      queryClient.setQueryData<InfiniteData<ProblemWithStatusResponse>>(queryKey, newData);
    }

    try {
      // 4. Make the API call
      await submissionsApi.createSubmission({
        is_internal: true,
        title: problem.title,
        title_slug: problem.title_slug,
        submitted_at: new Date().toISOString(),
      });

      // 5. On success: Optionally invalidate to refetch in the background and ensure consistency
      // queryClient.invalidateQueries({ queryKey }); // You might keep this commented out if optimistic update is enough

      Toast.show({
        type: 'success',
        text1: 'Submission Added',
        text2: `Successfully recorded submission for "${problem.title}"`,
      });

    } catch (error) {
      console.error('Error adding submission:', error);

      // 6. On error: Rollback to the previous data
      // 6. On error: Rollback to the previous data
      if (previousProblemsData) { // Rollback using the snapshot
        queryClient.setQueryData<InfiniteData<ProblemWithStatusResponse>>(queryKey, previousProblemsData);
      }

      Toast.show({
        type: 'error',
        text1: 'Error',
        text2: 'Failed to add submission. Reverted changes.', // Inform user about rollback
      });
    } finally {
      // 7. Always clear the loading indicator
      setSubmittingProblemId(null);
    }
  };

  const handleProblemPress = (problemWithStatus: ProblemWithStatus) => {
    router.push(`/problem/${problemWithStatus.problem.title_slug}`);
  };

  if (!fontsLoaded) {
    return null;
  }
  // Access data correctly using the InfiniteData structure
  console.log('Problems data length:', data?.pages.flatMap(page => page.data)?.length ?? 0, 'hasNextPage:', hasNextPage);

  return (
    <MenuProvider>
      <View className="flex-1 bg-[#131C24]">
        {/* Use the reusable ScreenHeader component */}
        <ScreenHeader
          title="Problem Library"
          showBackButton={false}
          centerTitle={true}
        />

        {/* Search Bar */}
        <View className="px-4 py-3">
          <View className="flex flex-col min-w-40 h-12 w-full">
            <View className="flex w-full flex-1 items-stretch rounded-xl h-full flex-row">
              <View className="text-[#8A9DC0] flex border-none bg-[#29374C] items-center justify-center pl-4 rounded-l-xl border-r-0">
                <Ionicons name="search" size={24} color="#8A9DC0" />
              </View>
              <TextInput
                placeholder="Find a problem"
                className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-xl text-[#F8F9FB] focus:outline-0 focus:ring-0 border-none bg-[#29374C] focus:border-none h-full placeholder:text-[#8A9DC0] px-4 rounded-l-none border-l-0 pl-2 text-base font-normal leading-normal"
                placeholderTextColor="#8A9DC0"
                value={searchInput}
                onChangeText={(text) => {
                  setSearchInput(text);
                  debouncedSearch(text);
                }}
                style={{ fontFamily: 'Roboto_400Regular' }}
              />
            </View>
          </View>
        </View>

        {/* Filter Buttons */}
        <View className='flex flex-row flex-wrap px-4 py-2 gap-2'>
          <DropdownFilter
              label="Difficulty"
              selectedValue={difficulty}
              options={structuredDiffOptions}
              onSelect={(value) => setDifficulty(value)}
          />

          <DropdownFilter
            label="Tag"
            selectedValue={tags}
            options={structuredTagOptions}
            onSelect={(selectedSlug: string | null) => {
              setTags(selectedSlug); 
            }}
          />
        </View>

        {/* Problems List */}
        <FlatList
          // Flatten pages data and filter out potential null/undefined items if necessary
          data={data?.pages.flatMap((page) => page.data).filter((item): item is ProblemWithStatus => !!item) || []}
          maintainVisibleContentPosition={{ minIndexForVisible: 0 }}
          keyExtractor={(item, index) => `${item?.problem.id}-${item?.problem.frontend_id}-${index}`}
          renderItem={({ item }) => {
            return (
            item ? (
            <ProblemCard
              title={`${item?.problem.frontend_id}. ${item?.problem.title}`}
              subtitle={item?.problem.difficulty}
              subtitleColor={difficultyColors[item?.problem.difficulty] || '#8A9DC0'}
              iconName="checkmark-circle-outline"
              completed={item?.completed}
              onPress={() => handleProblemPress(item)}
              rightElement={
                <Menu>
                  <MenuTrigger>
                    <View className="p-2 flex items-center justify-center">
                      {submittingProblemId === item.problem.id ? (
                        <ActivityIndicator size="small" color="#6366F1" />
                      ) : (
                        <Ionicons name="ellipsis-vertical" size={20} color="#8A9DC0" />
                      )}
                    </View>
                  </MenuTrigger>
                  <MenuOptions>
                    <MenuOption style={{ borderRadius: 100 }} onSelect={() => handleAddSubmission(item)}>
                      <Text className={'flex items-center justify-center p-2 text-black'}>Add Submission</Text>
                    </MenuOption>
                    <MenuOption style={{ borderRadius: 100 }} onSelect={() => {
                      setSelectedProblem(item);
                      setShowDeckModal(true);
                    }}>
                      <Text className={'flex items-center justify-center p-2 text-black'}>Add to Deck</Text>
                    </MenuOption>
                  </MenuOptions>
                </Menu>
              }
            />) : null
          )}}
          onEndReached={() => hasNextPage && fetchNextPage()}
          onEndReachedThreshold={0.5}
          ListFooterComponent={() =>
            isFetchingNextPage ? (
              <View className="py-4">
                <ActivityIndicator size="small" color="#6366F1" />
              </View>
            ) : null
          }
          ListEmptyComponent={() => (
            <View className="p-4 items-center">
              <Text className="text-[#8A9DC0] text-base text-center">
                No problems found. Try adjusting your search or filters.
              </Text>
            </View>
          )}
        />
        
        {/* Deck Selection Modal */}
        <Modal
          animationType="slide"
          transparent={true}
          visible={showDeckModal}
          onRequestClose={() => setShowDeckModal(false)}
        >
          <View className="flex-1 justify-end bg-black/50">
            <View className="bg-[#1E2A3A] rounded-t-xl p-6 max-h-[70%]">
              <View className="flex-row justify-between items-center mb-4">
                <Text
                  className="text-[#F8F9FB] text-xl font-bold"
                  style={{ fontFamily: 'Roboto_700Bold' }}
                >
                  Add to Deck
                </Text>
                <TouchableOpacity onPress={() => setShowDeckModal(false)}>
                  <Ionicons name="close" size={24} color="#8A9DC0" />
                </TouchableOpacity>
              </View>
              
              <Text
                className="text-[#8A9DC0] mb-4"
                style={{ fontFamily: 'Roboto_400Regular' }}
              >
                Select a deck to add "{selectedProblem?.problem.title}"
              </Text>
              
              <FlatList
                data={decksData?.user_decks || []}
                keyExtractor={(item) => item.id.toString()}
                renderItem={({ item: deck }) => (
                  <TouchableOpacity
                    className="bg-[#29374C] p-4 rounded-lg mb-2 flex-row justify-between items-center"
                    onPress={() => {
                      if (selectedProblem) {
                        addProblemToDeck.mutate({
                          deckId: deck.id,
                          problemId: selectedProblem.problem.id
                        }, {
                          onSuccess: () => {
                            Toast.show({
                              type: 'success',
                              text1: 'Success',
                              text2: `Added to "${deck.name}" deck`,
                            });
                            setShowDeckModal(false);
                          },
                          onError: () => {
                            Toast.show({
                              type: 'error',
                              text1: 'Error',
                              text2: 'Failed to add problem to deck',
                            });
                          }
                        });
                      }
                    }}
                  >
                    <View>
                      <Text
                        className="text-[#F8F9FB] text-lg font-medium"
                        style={{ fontFamily: 'Roboto_500Medium' }}
                      >
                        {deck.name}
                      </Text>
                      <Text
                        className="text-[#8A9DC0]"
                        style={{ fontFamily: 'Roboto_400Regular' }}
                      >
                        {deck.problem_count || 0} problems
                      </Text>
                    </View>
                    <Ionicons name="chevron-forward" size={20} color="#8A9DC0" />
                  </TouchableOpacity>
                )}
                ListEmptyComponent={
                  <View className="items-center p-4">
                    <Text
                      className="text-[#8A9DC0] text-center"
                      style={{ fontFamily: 'Roboto_400Regular' }}
                    >
                      You don't have any decks yet. Create one first.
                    </Text>
                  </View>
                }
              />
            </View>
          </View>
        </Modal>
      </View>
    </MenuProvider>
  );
}