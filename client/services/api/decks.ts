import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';
import { useQuery, useMutation, useQueryClient, useInfiniteQuery } from '@tanstack/react-query';
import { getApiUrl } from '../../utils/apiUrl';
import { useAuth } from '@/contexts/AuthContext';
import { Problem } from '../leetcode/types'; // Import the Problem type

export interface Deck {
  id: number;
  name: string;
  description: string;
  is_public: boolean;
  created_at: string;
  user_id?: string;
  problem_count?: number;
}

export interface DeckProblem {
  id: number;
  title: string;
  difficulty: string;
  title_slug: string;
}

// Get all accessible decks (public + user's own)
export const useDecks = () => {
  const { get } = useFetchWithAuth();
  const { session, isLoading: isAuthLoading } = useAuth();
  
  return useQuery({
    queryKey: ['decks'],
    queryFn: async () => {
      const url = getApiUrl(`/api/decks`); // Use relative path
      const response = await get(url);
      console.log('Raw API response: ', response);
      
      const responseData = response?.data || {public_decks: [], user_decks: []}
      console.log('Response data extracted: ', responseData);
      
      return responseData;
    },
    enabled: !isAuthLoading && !!session?.access_token,
  });
};

// Get problems in a deck using infinite scrolling
export const useDeckProblems = (deckId: number) => {
  const { get } = useFetchWithAuth();
  const limit = 20; // Define how many problems to fetch per page

  return useInfiniteQuery({
    queryKey: ['deckProblems', deckId], // Keep deckId in the queryKey
    queryFn: async ({ pageParam = 0 }) => {
      // Ensure pageParam is treated as offset
      const offset = typeof pageParam === 'number' ? pageParam : 0;
      const url = getApiUrl(`/api/decks/${deckId}/problems?limit=${limit}&offset=${offset}`); // Use relative path
      const response = await get(url);
      // Assuming the API returns the array of problems directly in response.data
      // If it's nested (e.g., response.data.problems), adjust accordingly
      return response?.data || [];
    },
    initialPageParam: 0, // Start with offset 0
    getNextPageParam: (lastPage, allPages) => {
      // lastPage is the array of problems from the last fetch
      // allPages is an array of arrays (problems from all fetched pages)
      if (lastPage && lastPage.length === limit) {
        // If the last page was full, there might be more data
        // Calculate the next offset based on the total number of pages fetched so far
        const currentOffset = allPages.length * limit;
        return currentOffset;
      }
      // If the last page had fewer items than the limit, we've reached the end
      return undefined;
    },
    enabled: !!deckId // Only run the query if deckId is valid
  });
};

// Delete a deck
export const useDeleteDeck = () => {
  const { fetchWithAuth } = useFetchWithAuth(); // Use the generic fetch method
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (deckId: number) => {
      const url = getApiUrl(`/api/decks/${deckId}`); // Use relative path
      // Use fetchWithAuth and specify the DELETE method
      // fetchWithAuth now handles non-ok statuses by throwing and returns undefined for 204.
      // No need to check response.ok here. If it didn't throw, it was successful.
      await fetchWithAuth(url, { method: 'DELETE' });
      // Return nothing as the operation succeeded if no error was thrown.
      return;
    },
    onSuccess: () => {
      // Invalidate the decks query to refresh the list after deletion
      queryClient.invalidateQueries({ queryKey: ['decks'] });
    },
    // Optional: Add onError for error handling
    // onError: (error: Error) => {
    //   console.error("Failed to delete deck:", error);
    //   // Show error toast or alert
    // }
  });
};

// Create a new deck
export const useCreateDeck = () => {
  const { post } = useFetchWithAuth();
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: async (deck: Omit<Deck, 'id' | 'created_at'>) => {
      const url = getApiUrl(`/api/decks`); // Use relative path
      return await post(url, deck);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['decks'] });
    }
  });
};

// Add deck to flashcards
export const useAddDeckToFlashcards = () => {
  const { post } = useFetchWithAuth();
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: async (deckId: number) => {
      const url = getApiUrl(`/api/flashcards/deck/${deckId}`); // Use relative path
      return await post(url, {});
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['flashcardReviews'] });
      queryClient.invalidateQueries({ queryKey: ['flashcardStats'] });
    }
  });
};

// Add problem to a deck
export const useAddProblemToDeck = () => {
  const { post } = useFetchWithAuth();
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: async ({ deckId, problemId }: { deckId: number; problemId: number }) => {
      const url = getApiUrl(`/api/decks/${deckId}/problems`); // Use relative path
      return await post(url, { problem_id: problemId });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['decks'] });
      queryClient.invalidateQueries({ queryKey: ['deckProblems'] });
    }
  });
};

// Prepare a public deck for practice (adds flashcards if needed)
export const useStartPracticePublicDeck = () => {
  const { post } = useFetchWithAuth();
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (deckId: number) => {
      const url = getApiUrl(`/api/decks/${deckId}/start-practice`); // Use relative path
      // The body is empty as the necessary info (deckId, userId) is handled by the backend
      return await post(url, {});
    },
    onSuccess: () => {
      // Invalidate queries that might be affected by new flashcards being available
      queryClient.invalidateQueries({ queryKey: ['flashcardReviews'] });
      queryClient.invalidateQueries({ queryKey: ['flashcardStats'] });
      // Potentially invalidate deck list if problem counts change, though less likely needed here
      // queryClient.invalidateQueries({ queryKey: ['decks'] });
    }
  });
};