import { ProblemResponse, Problem, ProblemWithStatus, ProblemWithStatusResponse } from '../leetcode/types';

// Function to fetch problems with filters
export const fetchProblems = async ({
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
  try {
    console.log("fetchProblems called with:", { limit, offset, difficulty, order_by, order_dir, search, tags, paid_only });
    
    // Build query parameters
    const queryParams = new URLSearchParams();
    
    if (limit) queryParams.append('limit', limit.toString());
    if (offset) queryParams.append('offset', offset.toString());
    if (difficulty) queryParams.append('difficulty', difficulty);
    if (order_by) queryParams.append('order_by', order_by);
    if (order_dir) queryParams.append('order_dir', order_dir);
    if (search) queryParams.append('search', search);
    if (tags) queryParams.append('tags', tags);
    if (paid_only !== undefined) queryParams.append('paid_only', paid_only.toString());
    
    const url = `http://localhost:8080/api/problems/list?${queryParams.toString()}`;
    console.log("Fetching from URL:", url);
    
    // Make the API request
    const response = await fetch(url);
    
    console.log("Response status:", response.status);
    
    if (!response.ok) {
      throw new Error(`API request failed with status ${response.status}`);
    }
    
    const data = await response.json();
    console.log("Received data:", data);
    return data;
  } catch (error) {
    console.error('Error fetching problems:', error);
    throw error;
  }
};

// Function to fetch problems with completion status
export const fetchProblemsWithStatus = async ({
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
  try {
    console.log("fetchProblemsWithStatus called with:", { limit, offset, difficulty, order_by, order_dir, search, tags, paid_only });
    
    // Build query parameters
    const queryParams = new URLSearchParams();
    
    if (limit) queryParams.append('limit', limit.toString());
    if (offset) queryParams.append('offset', offset.toString());
    if (difficulty) queryParams.append('difficulty', difficulty);
    if (order_by) queryParams.append('order_by', order_by);
    if (order_dir) queryParams.append('order_dir', order_dir);
    if (search) queryParams.append('search', search);
    if (tags) queryParams.append('tags', tags);
    if (paid_only !== undefined) queryParams.append('paid_only', paid_only.toString());
    
    const url = `http://localhost:8080/api/problems/with-status?${queryParams.toString()}`;
    console.log("Fetching from URL:", url);
    
    // Make the API request
    const response = await fetch(url);
    
    console.log("Response status:", response.status);
    
    if (!response.ok) {
      throw new Error(`API request failed with status ${response.status}`);
    }
    
    const data = await response.json();
    console.log("Received data with status:", data);
    return data;
  } catch (error) {
    console.error('Error fetching problems with status:', error);
    throw error;
  }
};

// Function to fetch a specific problem by slug
export const fetchProblemBySlug = async (slug: string): Promise<Problem> => {
  try {
    console.log("Fetching problem details for slug:", slug);
    
    const url = `http://localhost:8080/api/problems/by-slug?slug=${encodeURIComponent(slug)}`;
    console.log("Fetching from URL:", url);
    
    // Make the API request
    const response = await fetch(url);
    
    console.log("Response status:", response.status);
    
    if (!response.ok) {
      throw new Error(`API request failed with status ${response.status}`);
    }
    
    const data = await response.json();
    console.log("Received problem data:", data);
    return data.data;
  } catch (error) {
    console.error('Error fetching problem details:', error);
    throw error;
  }
};
