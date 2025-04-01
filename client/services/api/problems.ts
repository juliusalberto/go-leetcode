import { ProblemResponse, Problem, ProblemWithStatus, ProblemWithStatusResponse } from '../leetcode/types';
import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';
import { getApiUrl } from '../../utils/apiUrl';

// Create API clients with and without auth
// For React components, you'll use the hook directly
// For non-React contexts (like utility functions), we provide standalone functions

// Create standalone function for use outside of React components
const createQueryParams = (params: Record<string, any>): string => {
  const queryParams = new URLSearchParams();
  
  Object.entries(params).forEach(([key, value]) => {
    if (value !== undefined && value !== null) {
      queryParams.append(key, value.toString());
    }
  });
  
  return queryParams.toString();
};

// For use in React components
export const useProblemsApi = () => {
  const { get } = useFetchWithAuth();
  
  return {
    // Function to fetch problems with filters
    fetchProblems: async ({
      limit = 20,
      offset = 0,
      difficulty,
      order_by = 'frontend_id',
      order_dir = 'asc',
      search,
      tags,
      paid_only
    }: {
      limit?: number;
      offset?: number;
      difficulty?: string;
      order_by?: string;
      order_dir?: string;
      search?: string;
      tags?: string;
      paid_only?: boolean;
    }): Promise<ProblemResponse> => {
      console.log("fetchProblems called with:", { limit, offset, difficulty, order_by, order_dir, search, tags, paid_only });
      
      const params = {
        limit,
        offset,
        difficulty,
        order_by,
        order_dir,
        search,
        tags,
        paid_only
      };
      
      const queryString = createQueryParams(params);
      const url = `http://localhost:8080/api/problems/list?${queryString}`;
      console.log("Fetching from URL:", url);
      
      // Public endpoint, specify skipAuth
      return get<ProblemResponse>(url, { skipAuth: true });
    },
    
    // Function to fetch problems with completion status
    fetchProblemsWithStatus: async ({
      limit = 20,
      offset = 0,
      difficulty,
      order_by = 'frontend_id',
      order_dir = 'asc',
      search,
      tags,
      paid_only
    }: {
      limit?: number;
      offset?: number;
      difficulty?: string;
      order_by?: string;
      order_dir?: string;
      search?: string;
      tags?: string;
      paid_only?: boolean;
    }): Promise<ProblemWithStatusResponse> => {
      console.log("fetchProblemsWithStatus called with:", { limit, offset, difficulty, order_by, order_dir, search, tags, paid_only });
      
      const params = {
        limit,
        offset,
        difficulty,
        order_by,
        order_dir,
        search,
        tags,
        paid_only
      };
      
      const queryString = createQueryParams(params);
      const url = getApiUrl(`http://localhost:8080/api/problems/with-status?${queryString}`);
      console.log("Fetching from URL:", url);
      
      // This is an authenticated endpoint, token will be added automatically
      return get<ProblemWithStatusResponse>(url);
    },
    
    // Function to fetch a specific problem by slug
    fetchProblemBySlug: async (slug: string): Promise<Problem> => {
      console.log("Fetching problem details for slug:", slug);
      
      const url = getApiUrl(`http://localhost:8080/api/problems/by-slug?slug=${encodeURIComponent(slug)}`);
      console.log("Fetching from URL:", url);
      
      // Public endpoint, specify skipAuth
      const response = await get<{ data: Problem }>(url, { skipAuth: true });
      return response.data;
    }
  };
};