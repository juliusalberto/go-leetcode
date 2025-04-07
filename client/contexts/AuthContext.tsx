// contexts/AuthContext.tsx (create this file)
import { createContext, useState, useEffect, useContext, ReactNode } from 'react';
import { Platform } from 'react-native';
import { supabase, Session } from '../lib/supabase'; // Adjust path if needed
import { fetchAuthStatus } from '../services/api/auth';
import { router, usePathname } from 'expo-router';
import Toast from 'react-native-toast-message';

interface AuthContextType {
  session: Session | null;
  isLoading: boolean; // Indicates if initial session/profile check is loading
  profileExists: boolean | null; // Tracks if backend profile exists
  leetcodeUsername: string | null; // Stores the leetcode username from backend
  setSession: (session: Session | null) => void; // Allow manual setting if needed
  updateProfileStatus: (exists: boolean) => void; // Manually update profile status
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface AuthProviderProps {
  children: ReactNode;
}

// ... (AuthContextType and AuthContext definition) ...

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [session, setSession] = useState<Session | null>(null);
  const [profileExists, setProfileExists] = useState<boolean | null>(null);
  const [leetcodeUsername, setLeetcodeUsername] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  // Flag to track if we're currently on the complete profile page
  const [isOnCompletePage, setIsOnCompletePage] = useState(false);

  // Use pathname to detect when we're on the complete-profile page
  const pathname = usePathname();
  
  useEffect(() => {
    // Update flag when pathname changes
    setIsOnCompletePage(pathname === '/complete-profile');
  }, [pathname]);
  
  // --- Function to check profile and navigate ---
  const checkProfileAndNavigate = async (currentSession: Session) => {
    console.log("AuthProvider: Checking profile status for user:", currentSession.user.id);

    // If we already know the profile status, use it directly for navigation checks
    if (profileExists !== null) {
        console.log("AuthProvider: Profile status already known:", profileExists, "Skipping fetch.");
        setIsLoading(false); // Ensure loading is false if we skip fetch

        // --- Navigation Logic (using known status) ---
        if (profileExists) {
            // Check if effectively on root page (handles web hash issue)
            const isEffectivelyRoot = Platform.OS === 'web'
                ? window.location.pathname === '/'
                : (pathname === '/' || pathname === '');

            if (isEffectivelyRoot) {
                console.log("AuthProvider: Profile exists (known) and user is effectively on root page, navigating to dashboard.");
                router.replace('/(tabs)/dashboard');
            } else {
                console.log(`AuthProvider: Profile exists (known) but user is not on root page (pathname: ${pathname}, web path: ${Platform.OS === 'web' ? window.location.pathname : 'N/A'}), not navigating.`);
            }
        } else { // Profile does not exist (known)
             if (isOnCompletePage) {
                console.log("AuthProvider: Profile missing (known) but already on complete-profile page, not navigating.");
             } else {
                console.log("AuthProvider: Profile missing (known), navigating to complete-profile.");
                router.replace('/complete-profile');
             }
        }
        return; // Exit early since we didn't need to fetch
    }

    // --- If profile status is unknown (null), proceed to fetch ---
    console.log("AuthProvider: Profile status unknown, fetching from API...");
    setIsLoading(true); // Set loading true only when fetching
    let fetchedProfileExists: boolean | null = null; // Variable to hold the result
    try {
      const status = await fetchAuthStatus(currentSession.access_token); // Fetch status
      console.log("AuthProvider: Full profile status response:", status);
      fetchedProfileExists = status.profile_exists; // Store fetched status
      setProfileExists(fetchedProfileExists); // Update state
      setLeetcodeUsername(status.leetcode_username || null);

      // --- Navigation Logic ---
      if (fetchedProfileExists) {
        // Only redirect to dashboard if profile exists AND we are on the root index page
        // Check if effectively on root page (handles web hash issue)
        const isEffectivelyRoot = Platform.OS === 'web'
            ? window.location.pathname === '/'
            : (pathname === '/' || pathname === '');

        if (isEffectivelyRoot) {
          console.log("AuthProvider: Profile exists (fetched) and user is effectively on root page, navigating to dashboard.");
          router.replace('/(tabs)/dashboard');
        } else {
          console.log(`AuthProvider: Profile exists (fetched) but user is not on root page (pathname: ${pathname}, web path: ${Platform.OS === 'web' ? window.location.pathname : 'N/A'}), not navigating.`);
        }
      } else {
        // Skip navigation if we're already on the complete-profile page
        // This prevents the loop by using our React state rather than checking window.location
        if (isOnCompletePage) {
          console.log("AuthProvider: Already on complete-profile page, not navigating.");
        } else {
          console.log("AuthProvider: Profile missing (fetched), navigating to complete-profile.");
          router.replace('/complete-profile');
        }
      }
      console.log("AuthProvider: Navigation completed after fetch");
    } catch (error: any) {
       // ... error handling ...
       console.error('AuthProvider: Failed to fetch auth status:', error);
       Toast.show({type: 'error', text1: 'Auth Error', text2: 'Could not verify profile status.'});
       setProfileExists(null); // Reset on error
       try { await supabase.auth.signOut(); } catch {}
       // Avoid complex back logic if possible, just go to root on critical fetch error
       router.replace('/');
    } finally {
      setIsLoading(false);
    }
  };

  // Method to manually set profile status - can be called from complete-profile page
  const updateProfileStatus = (exists: boolean) => {
    console.log("AuthProvider: Manually updating profile status to:", exists);
    setProfileExists(exists);
  };
  
  useEffect(() => {
    // 1. Get Initial Session & Check Profile only on first mount
    const getInitialSession = async () => {
      try {
        const { data: { session: initialSession } } = await supabase.auth.getSession();
        console.log('AuthProvider: Initial session:', initialSession ? 'Found' : 'Not Found');
        setSession(initialSession);
        
        // Skip profile check if we're already on the complete-profile page
        if (initialSession && !isOnCompletePage) {
          await checkProfileAndNavigate(initialSession);
        } else {
          setIsLoading(false);
        }
      } catch (err) {
        console.error("AuthProvider: Error fetching initial session:", err);
        setIsLoading(false);
      }
    };
    
    getInitialSession();
    
    // 2. Listen for Auth State Changes
    const { data: authListener } = supabase.auth.onAuthStateChange(
      async (_event, newSession) => {
        console.log(`AuthProvider: Auth state changed (${_event}):`, newSession ? 'Got Session' : 'No Session');
        
        // Only update session state if it has actually changed
        if (JSON.stringify(newSession) !== JSON.stringify(session)) {
          setSession(newSession);
        }

        if (_event === 'SIGNED_IN' && newSession) {
           // Use the existing isOnCompletePage state value (can't use hooks inside callbacks)
           if (!isOnCompletePage) {
             await checkProfileAndNavigate(newSession);
           } else {
             console.log("AuthProvider: On complete-profile page during sign-in, skipping navigation check");
             setIsLoading(false);
           }

        } else if (_event === 'SIGNED_OUT') {
          setProfileExists(null);
          setIsLoading(false);
          console.log("AuthProvider: Signed out, navigating to /");
          router.replace('/'); // Navigate to root on sign out

        } else if (_event === 'TOKEN_REFRESHED' && newSession){
          console.log("AuthProvider: Token refreshed, NOT re-checking profile.");
          // Skip re-checking profile on token refresh to avoid excessive refreshes
          // The supabase client auto-refreshes tokens frequently, so this is a major source of loops
          setIsLoading(false);

        } else if (!newSession) {
          // Handles cases like USER_DELETED or other events resulting in no session
          setProfileExists(null);
          setIsLoading(false);
          router.replace('/');
        }
      }
    );

    return () => {
      authListener?.subscription.unsubscribe();
    };
  // Remove session and isLoading from the dependency array to prevent the rerun loop
  }, [isOnCompletePage]);

  const value = {
    session,
    isLoading,
    profileExists,
    leetcodeUsername,
    setSession,
    updateProfileStatus
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

// Custom hook to use the auth context
export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};