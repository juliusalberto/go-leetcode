import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';
import { getApiUrl } from '../../utils/apiUrl'; // Import getApiUrl

// React hook version for use in components
export const useSolutionsApi = () => {
  const { get } = useFetchWithAuth();
  
  return {
    // Function to fetch solutions for a problem ID with auth
    fetchSolutionByID: async (id: string) => {
      console.log("Fetching solutions for problem ID:", id);
      const relativePath = `/api/solutions?id=${encodeURIComponent(id)}`;
      const url = getApiUrl(relativePath); // Use getApiUrl with relative path
      return get(url);
    },
    
    // Function to fetch a specific solution by problem ID and language with auth
    fetchSolutionByLanguage: async (id: string, language: string) => {
      console.log(`Fetching ${language} solution for problem ID:`, id);
      const relativePath = `/api/solutions?id=${encodeURIComponent(id)}&language=${encodeURIComponent(language)}`;
      const url = getApiUrl(relativePath); // Use getApiUrl with relative path
      return get(url);
    }
  };
};