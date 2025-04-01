// app/services/api/auth.ts
import { getApiUrl } from '../../utils/apiUrl';

interface AuthStatusResponse {
    authenticated: boolean;
    profile_exists: boolean;
    user_id?: string; // Or number, depending on your backend ID type
    username?: string;
    leetcode_username?: string;
}

interface CompleteProfileRequest {
    username: string;
    leetcode_username: string;
}

export const fetchAuthStatus = async (token: string): Promise<AuthStatusResponse> => {
    const url = getApiUrl(`http://localhost:8080/api/auth/status`); // Your backend status check endpoint
    console.log("API: Fetching auth status from", url);

    try {
        console.log("API: Sending auth status request with token:", token.substring(0, 10) + "...");
        
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        });

        console.log("API: Raw response status:", response.status);
        const data = await response.json();
        console.log("API: Auth status response data:", data);

        if (!response.ok) {
            // Try to extract error message from backend response
            const message = data?.errors?.[0]?.message || `Auth status check failed (${response.status})`;
            throw new Error(message);
        }

        // Ensure the response has the expected fields in the data property
        if (typeof data?.data?.profile_exists !== 'boolean') {
            throw new Error("Invalid auth status response format from backend");
        }

        // Return the data property which contains the actual auth status
        return data.data as AuthStatusResponse;
    } catch (error) {
        console.error('API: Error fetching auth status:', error);
        throw error; // Re-throw the error to be caught by the caller
    }
};

export const completeUserProfile = async (token: string, profileData: CompleteProfileRequest): Promise<AuthStatusResponse> => {
    const url = getApiUrl(`http://localhost:8080/api/users/profile`); 
    console.log("API: Completing user profile at", url);

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(profileData),
        });

        const data = await response.json();
        console.log("API: Complete profile response:", response.status, data);

        if (!response.ok) {
            // Try to extract error message from backend response
            const message = data?.errors?.[0]?.message || `Failed to complete profile (${response.status})`;
            throw new Error(message);
        }
        
        // After successful profile creation, fetch the updated auth status
        return await fetchAuthStatus(token);
    } catch (error) {
        console.error('API: Error completing profile:', error);
        throw error; // Re-throw the error to be caught by the caller
    }
};