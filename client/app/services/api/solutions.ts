// Function to fetch solutions for a problem ID
export const fetchSolutionByID = async (id: string) => {
    try {
      console.log("Fetching solutions for problem ID:", id);
      
      const url = `http://localhost:8080/api/solutions?id=${encodeURIComponent(id)}`;
      console.log("Fetching from URL:", url);
      
      // Make the API request
      const response = await fetch(url);
      
      console.log("Response status:", response.status);
      
      if (!response.ok) {
        throw new Error(`API request failed with status ${response.status}`);
      }
      
      const data = await response.json();
      console.log("Received solution data:", data);
      
      return data;
    } catch (error) {
      console.error('Error fetching solutions:', error);
      throw error;
    }
  };

// Function to fetch a specific solution by problem ID and language
export const fetchSolutionByLanguage = async (id: string, language: string) => {
  try {
    console.log(`Fetching ${language} solution for problem ID:`, id);
    
    const url = `http://localhost:8080/api/solutions?id=${encodeURIComponent(id)}&language=${encodeURIComponent(language)}`;
    console.log("Fetching from URL:", url);
    
    // Make the API request
    const response = await fetch(url);
    
    if (!response.ok) {
      throw new Error(`API request failed with status ${response.status}`);
    }
    
    const data = await response.json();
    return data;
  } catch (error) {
    console.error(`Error fetching ${language} solution:`, error);
    throw error;
  }
};