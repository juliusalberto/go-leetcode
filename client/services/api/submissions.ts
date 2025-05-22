import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';
import { getApiUrl } from '../../utils/apiUrl'; // Import getApiUrl

// Submission API service
export interface SubmissionRequest {
  is_internal: boolean;
  leetcode_submission_id?: string;
  title: string;
  title_slug: string;
  submitted_at: string;
  rating?: number; // Add optional rating property
}

export interface SubmissionResponse {
  success: boolean;
  submission_id: string;
  next_review_at: string;
  days_until_review: number;
  is_due: boolean;
  title: string;
  title_slug: string;
}

// React hook version for use in components
export const useSubmissionsApi = () => {
  const { post, get } = useFetchWithAuth();
  
  return {
    // Create a submission with auth
    createSubmission: async (submission: SubmissionRequest): Promise<SubmissionResponse> => {
      console.log("Creating submission:", submission);
      const url = getApiUrl('/api/reviews/process-submission'); // Use getApiUrl with relative path
      return post<SubmissionResponse>(url, submission);
    },
    
    // Get submissions for a user with auth
    getSubmissionsByUser: async (userId: number): Promise<SubmissionResponse[]> => {
      console.log("Fetching submissions for user ID:", userId);
      // Construct the relative path including query parameters
      const relativePath = `/api/submissions?user_id=${userId}`;
      const url = getApiUrl(relativePath); // Use getApiUrl with relative path
      return get<SubmissionResponse[]>(url);
    }
  };
};

