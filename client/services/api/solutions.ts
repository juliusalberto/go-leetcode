import { useFetchWithAuth } from '@/hooks/useFetchWithAuth';

// React hook version for use in components
export const useSolutionsApi = () => {
  const { get } = useFetchWithAuth();
  
  return {
    // Function to fetch solutions for a problem ID with auth
    fetchSolutionByID: async (id: string) => {
      console.log("Fetching solutions for problem ID:", id);
      const url = `http://localhost:8080/api/solutions?id=${encodeURIComponent(id)}`;
      return get(url);
    },
    
    // Function to fetch a specific solution by problem ID and language with auth
    fetchSolutionByLanguage: async (id: string, language: string) => {
      console.log(`Fetching ${language} solution for problem ID:`, id);
      const url = `http://localhost:8080/api/solutions?id=${encodeURIComponent(id)}&language=${encodeURIComponent(language)}`;
      return get(url);
    }
  };
};