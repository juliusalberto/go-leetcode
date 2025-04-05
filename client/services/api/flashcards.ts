import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';
// Import useInfiniteQuery
import { useInfiniteQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { getApiUrl } from '../../utils/apiUrl';

// Define the expected shape of a single page response from the API
interface FlashcardPage {
  reviews: FlashcardReview[];
  total: number;
}

export interface FlashcardReview {
  id: number;
  problem_id: number;
  user_id: string;
  deck_id: number;
  fsrs_card: {
    Due: string;
    Stability: number;
    Difficulty: number;
    ElapsedDays: number;
    ScheduledDays: number;
    Reps: number;
    Lapses: number;
    State: number;
    LastReview: string;
  };
  problem: {
    id: number;
    frontend_id: number;
    title: string;
    title_slug: string;
    difficulty: string;
    is_paid_only: boolean;
    content: string;
    solution_approach: string;
  };
}

export interface FlashcardRating {
  review_id: number;
  rating: 1 | 2 | 3 | 4; // 1=Very Hard, 2=Hard, 3=Good, 4=Easy
}

const FLASHCARD_PAGE_LIMIT = 20; // Fetch 20 cards per page

// Get due flashcard reviews using infinite query, optionally filtered by deck
export const useFlashcardReviews = (deckId?: number) => {
  const { get } = useFetchWithAuth();

  return useInfiniteQuery<FlashcardPage, Error>({ // Specify types for data and error
    // Query key identifies the data source, including the deck filter
    queryKey: ['flashcardReviews', deckId],
    
    // queryFn fetches a single page
    queryFn: async ({ pageParam = 0 }) => { // pageParam will be the offset
      const offset = pageParam as number;
      const url = getApiUrl(
        `/api/flashcards/reviews?${deckId ? `deck_id=${deckId}&` : ''}limit=${FLASHCARD_PAGE_LIMIT}&offset=${offset}` // Use relative path
      );
      const fullResponse = await get(url);
      // Return the data part which should match FlashcardPage interface
      return fullResponse?.data || { reviews: [], total: 0 };
    },

    // initialPageParam defines the offset for the first page
    initialPageParam: 0,

    // getNextPageParam calculates the offset for the next page
    getNextPageParam: (lastPage, allPages) => {
      // Calculate how many reviews have been fetched across all pages
      const fetchedCount = allPages.flatMap(page => page.reviews).length;
      
      // If the number fetched is less than the total reported by the last page,
      // calculate the next offset. Otherwise, return undefined (no more pages).
      if (fetchedCount < lastPage.total) {
        return fetchedCount; // Next offset is the current total count
      }
      return undefined;
    },
  });
};

// Submit flashcard rating
export const useSubmitFlashcardRating = () => {
  const { post } = useFetchWithAuth();
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: async (rating: FlashcardRating) => {
      const url = getApiUrl(`/api/flashcards/reviews`); // Use relative path
      return await post(url, {
        review_id: rating.review_id,
        rating: rating.rating
      });
    }
    // REMOVED: onSuccess invalidation to prevent refetching after every rating.
    // We will rely on fetchNextPage triggered by the component when needed.
  });
};