import { getApiUrl } from '../../utils/apiUrl';

// Base GraphQL client function for LeetCode
export const leetCodeGraphQL = async (query: string, variables = {}) => {
    try {
      // Use our proxy endpoint instead of directly calling LeetCode API
      const apiUrl = getApiUrl('http://localhost:8080/api/proxy/leetcode');
      
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          query,
          variables,
        }),
      });
      
      if (!response.ok) {
        throw new Error(`API request failed with status ${response.status}`);
      }
      
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error fetching from LeetCode GraphQL:', error);
      throw error;
    }
  };