import { useQuery } from '@tanstack/react-query';

export interface Review {
  id: string;
  title: string;
  due_date: string;
  status: 'due' | 'upcoming';
  next_review_at: string;
}

const fetchReviews = async (userId: number, limit: number = 5): Promise<Review[]> => {
  const params = new URLSearchParams({
    user_id: userId.toString(),
    status: "due",
    per_page: limit.toString(),
    page: '1',
  });

  const response = await fetch(`http://localhost:8080/api/reviews?${params.toString()}`, {
    headers: { 'Content-Type': 'application/json' }
  });

  console.log(response)

  if (!response.ok) {
    throw new Error('Failed to fetch reviews');
  }

  const data = await response.json();
  return data.data.slice(0, limit);
};

export const useReviews = (userId: number) => {
  // useQuery provides loading state via isLoading, error, etc.
  return useQuery<Review[], Error>({
    queryKey: ['recentReviews', userId],
    queryFn: () => fetchReviews(userId, 10)
  });
};