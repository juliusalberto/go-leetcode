import { useAuth } from '@/contexts/AuthContext';
import { getApiUrl } from '../utils/apiUrl';

interface FetchWithAuthOptions extends RequestInit {
  skipAuth?: boolean;
}

export function useFetchWithAuth() {
  const { session } = useAuth();
  
  /**
   * Enhanced fetch function that automatically includes bearer token
   * @param url The URL to fetch
   * @param options Regular fetch options plus skipAuth to bypass authentication
   * @returns Promise with the fetch response
   */
  const fetchWithAuth = async <T = any>(
    url: string,
    options: FetchWithAuthOptions = {}
  ): Promise<T> => {
    const { skipAuth = false, headers = {}, ...restOptions } = options;
    
    // Create a proper headers object with type safety
    const requestHeaders: Record<string, string> = {
      ...(headers as Record<string, string>)
    };
    
    // Add authorization header if session exists and skipAuth is false
    if (session?.access_token && !skipAuth) {
      requestHeaders['Authorization'] = `Bearer ${session.access_token}`;
    }
    
    // Add content-type if not present
    if (!requestHeaders['Content-Type']) {
      requestHeaders['Content-Type'] = 'application/json';
    }
    
    // Apply platform-specific URL conversion
    const platformUrl = getApiUrl(url);
    
    const response = await fetch(platformUrl, {
      ...restOptions,
      headers: requestHeaders,
    });
    
    // Handle non-2xx status codes
    if (!response.ok) {
      try {
        // Try to parse error response
        const errorData = await response.json();
        throw new Error(
          errorData?.errors?.[0]?.message || 
          `Request failed with status ${response.status}`
        );
      } catch (parseError) {
        // If JSON parsing fails, throw with status code
        throw new Error(`Request failed with status ${response.status}`);
      }
    }

    // Handle 204 No Content specifically
    if (response.status === 204) {
      // Return undefined or null, cast to T to satisfy the return type
      return undefined as T;
    }
    
    // Parse JSON response for other successful responses
    const data = await response.json();
    return data;
  };
  
  /**
   * Helper method for GET requests
   */
  const get = <T = any>(url: string, options: FetchWithAuthOptions = {}): Promise<T> => {
    return fetchWithAuth<T>(url, { 
      ...options,
      method: 'GET' 
    });
  };
  
  /**
   * Helper method for POST requests
   */
  const post = <T = any>(url: string, body: any, options: FetchWithAuthOptions = {}): Promise<T> => {
    return fetchWithAuth<T>(url, {
      ...options,
      method: 'POST',
      body: JSON.stringify(body)
    });
  };
  
  /**
   * Helper method for PUT requests
   */
  const put = <T = any>(url: string, body: any, options: FetchWithAuthOptions = {}): Promise<T> => {
    return fetchWithAuth<T>(url, {
      ...options,
      method: 'PUT',
      body: JSON.stringify(body)
    });
  };
  
  /**
   * Helper method for DELETE requests
   */
  const del = <T = any>(url: string, options: FetchWithAuthOptions = {}): Promise<T> => {
    return fetchWithAuth<T>(url, {
      ...options,
      method: 'DELETE'
    });
  };
  
  return {
    fetchWithAuth,
    get,
    post,
    put,
    delete: del  // 'delete' is a reserved word in JavaScript
  };
}