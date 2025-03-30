import { useInfiniteQuery, InfiniteData } from '@tanstack/react-query';

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

const fetchReviews = async ({ pageParam = 1, limit = 10 }: FetchReviewsParams): Promise<Review[]> => {
  const params = new URLSearchParams({
    status: "due",
    per_page: limit.toString(),
    page: pageParam.toString(),
  });

  const response = await fetch(`http://localhost:8080/api/reviews?${params.toString()}`, {
    headers: { 'Content-Type': 'application/json' }
  });

  console.log(response)

  if (!response.ok) {
    throw new Error('Failed to fetch reviews');
  }

  const data = await response.json();
  // Add null check to handle cases where data.data might be null or undefined
  return data?.data ? data.data.slice(0, limit) : [];
};

export const useReviews = () => {
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
    }
  });
};