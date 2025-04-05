import { useInfiniteQuery, InfiniteData } from '@tanstack/react-query';
import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';
import { getApiUrl } from '../../utils/apiUrl';

export interface Review {
  id: string;
  title: string;
  due_date: string;
  status: 'due' | 'upcoming';
  next_review_at?: string;
  created_at: string;
  title_slug: string;
}

interface FetchReviewsParams {
  pageParam?: number;
  limit?: number;
}

interface ReviewsResponse {
  data: Review[];
}

// Define options type for the hook
interface UseReviewsOptions {
  enabled?: boolean;
}

export const useReviews = (options?: UseReviewsOptions) => { // Accept options
  const { get } = useFetchWithAuth();
  const enabled = options?.enabled ?? true; // Default to true if not provided

  const fetchReviews = async ({ pageParam = 1, limit = 10 }: FetchReviewsParams): Promise<Review[]> => {
    const params = new URLSearchParams({
      status: "due",
      per_page: limit.toString(),
      page: pageParam.toString(),
    });
    
    try {
      const url = getApiUrl(`/api/reviews?${params.toString()}`); // Use relative path
      const response = await get<ReviewsResponse>(url);
      
      // Add null check to handle cases where data.data might be null or undefined
      return response?.data ? response.data.slice(0, limit) : [];
    } catch (error) {
      console.error('Error fetching reviews:', error);
      throw error;
    }
  };

  return useInfiniteQuery<Review[], Error, InfiniteData<Review[]>, [string], number>({
    queryKey: ['recentReviews'],
    queryFn: ({ pageParam }) => fetchReviews({
      pageParam,
      limit: 10
    }),
    initialPageParam: 1,
    getNextPageParam: (lastPage, allPages) => {
      // Only fetch next page if we received a full page of results
      // This ensures we don't keep trying to fetch when there are no more results
      return lastPage && lastPage.length >= 10 ? allPages.length + 1 : undefined;
    },
    enabled: enabled, // <-- Pass the enabled option here
  });
};