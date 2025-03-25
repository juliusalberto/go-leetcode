// Submission API service
export interface SubmissionRequest {
  is_internal: boolean;
  leetcode_submission_id?: string;
  user_id: number;
  title: string;
  title_slug: string;
  submitted_at: string;
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

export const createSubmission = async (submission: SubmissionRequest): Promise<SubmissionResponse> => {
  try {
    console.log("Creating submission:", submission);
    
    const url = 'http://localhost:8080/api/reviews/process-submission';
    console.log("Posting to URL:", url);
    
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(submission),
    });
    
    if (!response.ok) {
      throw new Error(`API request failed with status ${response.status}`);
    }
    
    const data = await response.json();
    console.log("Submission created successfully:", data);
    return data;
  } catch (error) {
    console.error('Error creating submission:', error);
    throw error;
  }
};

export const getSubmissionsByUser = async (userId: number): Promise<SubmissionResponse[]> => {
  try {
    console.log("Fetching submissions for user ID:", userId);
    
    const url = `http://localhost:8080/api/submissions?user_id=${userId}`;
    console.log("Fetching from URL:", url);
    
    const response = await fetch(url);
    
    if (!response.ok) {
      throw new Error(`API request failed with status ${response.status}`);
    }
    
    const data = await response.json();
    console.log("Received submissions data:", data);
    return data;
  } catch (error) {
    console.error('Error fetching user submissions:', error);
    throw error;
  }
};