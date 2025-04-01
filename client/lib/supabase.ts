import 'react-native-url-polyfill/auto'; // Keep this
import AsyncStorage from '@react-native-async-storage/async-storage';
import { createClient, Session } from '@supabase/supabase-js';
import { Platform } from 'react-native'; // Import Platform

// --- Define a Noop Storage for the Server ---
const createNoopStorage = () => {
  return {
    getItem: (_key: string): Promise<string | null> => {
      // console.log("SSR Storage getItem:", key); // Optional: log server access
      return Promise.resolve(null);
    },
    setItem: (_key: string, _value: string): Promise<void> => {
      // console.log("SSR Storage setItem:", key); // Optional: log server access
      return Promise.resolve();
    },
    removeItem: (_key: string): Promise<void> => {
      // console.log("SSR Storage removeItem:", key); // Optional: log server access
      return Promise.resolve();
    },
  };
};

// Determine the correct storage based on the environment
const storage = Platform.OS === 'web' && typeof window === 'undefined'
  ? createNoopStorage() // Use noop storage on the server (web platform, but no window object)
  : AsyncStorage;      // Use AsyncStorage on native or on the web client-side

// --- Supabase Client Initialization ---
let supabaseUrl = process.env.EXPO_PUBLIC_SUPABASE_URL;
const supabaseAnonKey = process.env.EXPO_PUBLIC_SUPABASE_ANON_KEY;

if (!supabaseUrl || !supabaseAnonKey) {
  throw new Error("Supabase URL or Anon Key is missing. Check your environment variables.");
}

// Handle different platforms for Supabase URL
if (supabaseUrl) {
  if (Platform.OS === 'android') {
    // For Android emulators, 10.0.2.2 is the special IP to reach host machine's localhost
    if (supabaseUrl.includes('127.0.0.1') || supabaseUrl.includes('localhost')) {
      console.log('Running on Android - replacing localhost with 10.0.2.2');
      supabaseUrl = supabaseUrl.replace('127.0.0.1', '10.0.2.2').replace('localhost', '10.0.2.2');
    }
    
    // Log the URL we're using for more visibility
    console.log(`Android: Using Supabase URL: ${supabaseUrl}`);
  } else if (Platform.OS === 'ios') {
    // For iOS simulators, localhost works but you can use host.docker.internal if needed
    console.log('Running on iOS - using original URL');
    console.log(`iOS: Using Supabase URL: ${supabaseUrl}`);
  } else {
    console.log('Running on web - using original URL');
    console.log(`Web: Using Supabase URL: ${supabaseUrl}`);
  }
}

export const supabase = createClient(supabaseUrl, supabaseAnonKey, {
  auth: {
    storage: storage, // Use the conditional storage object
    autoRefreshToken: true,
    persistSession: true,
    detectSessionInUrl: Platform.OS === 'web', // Enable URL session detection on web, disable on native
  },
});

export type { Session }; // Keep type export