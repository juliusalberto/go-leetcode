import React, { useState, useEffect } from 'react';
import {
  TouchableOpacity,
  View,
  ImageBackground,
  Text,
  Platform,
  Alert,
  ActivityIndicator
} from 'react-native';
import { useFonts, Roboto_400Regular, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { router } from 'expo-router';
import { Ionicons } from '@expo/vector-icons';
import { supabase } from '../lib/supabase';
import * as WebBrowser from 'expo-web-browser';
import { useAuth } from '../contexts/AuthContext';
import { GoogleSignin } from '@react-native-google-signin/google-signin';

// Enable WebBrowser.maybeCompleteAuthSession for the OAuth redirect flow
WebBrowser.maybeCompleteAuthSession();

export default function IndexScreen() {
  const { setSession } = useAuth();
  const [isLoading, setIsLoading] = useState(false);
  const [fontTimeout, setFontTimeout] = useState(false);
  const [fontsLoaded, fontError] = useFonts({
    Roboto_400Regular,
    Roboto_700Bold,
  });

  // Log any font loading errors
  useEffect(() => {
    if (fontError) {
      console.log('Error loading fonts:', fontError);
    }
  }, [fontError]);

  // Add a timeout for font loading to prevent indefinite black screen
  useEffect(() => {
    const timer = setTimeout(() => {
      if (!fontsLoaded) {
        console.log('Font loading timeout reached - proceeding anyway');
        setFontTimeout(true);
      }
    }, 3000); // 3 second timeout
    
    return () => clearTimeout(timer);
  }, [fontsLoaded]);

  // Initialize Google Sign-In configuration
  useEffect(() => {
    console.log('Initializing Google Sign-In');
    if (Platform.OS !== 'web') {
      try {
        GoogleSignin.configure({
          webClientId: process.env.EXPO_PUBLIC_GOOGLE_WEB_CLIENT_ID,
          offlineAccess: true, // If you need a refresh token
          scopes: ['profile', 'email'],
        });
        console.log('Google Sign-In initialized successfully');
      } catch (error) {
        console.error('Failed to initialize Google Sign-In:', error);
      }
    }
  }, []);
  
  const handleGoogleSignIn = async () => {
    setIsLoading(true);
    try {
      if (Platform.OS === 'web') {
        // Web platform uses standard OAuth flow with a callback
        console.log('Using standard web OAuth flow');
        
        // For web, we'll use the standard OAuth flow
        const { error } = await supabase.auth.signInWithOAuth({
          provider: 'google',
          options: {
            redirectTo: window.location.origin, // Redirect to the same origin
            queryParams: {
              access_type: 'offline',
              prompt: 'consent'
            }
            // Don't skip browser redirect on web - let Supabase handle it
          }
        });
        
        if (error) throw error;
        // No need for further handling on web as the page will redirect to Google
        
      } else {
        // Native platforms use Google Sign-In native flow
        console.log('Using native Google Sign-In');
        
        try {
          // Check if Play Services are available (Android)
          await GoogleSignin.hasPlayServices();
          
          // Sign in with Google
          const userInfo = await GoogleSignin.signIn();
          
          // Using type assertion to accommodate the library's type structure
          // This is safer than relying on the specific structure that might change
          const userInfoAny = userInfo as any;
          const idToken = userInfoAny.idToken ||
                         (userInfoAny.user && userInfoAny.user.idToken) ||
                         (userInfoAny.data && userInfoAny.data.idToken);
          
          if (idToken) {
            console.log('Got ID token from Google, signing into Supabase');
            
            // Use the ID token to sign in with Supabase
            const { data, error } = await supabase.auth.signInWithIdToken({
              provider: 'google',
              token: idToken,
            });
            
            if (error) throw error;
            console.log('Successfully signed in with ID token:', data);
            
          } else {
            throw new Error('No ID token present in Google Sign-In response');
          }
        } catch (nativeError: any) {
          console.error('Native Google Sign-In error:', nativeError);
          
          // Handle specific Google Sign-In errors
          if (nativeError.code === 'SIGN_IN_CANCELLED') {
            console.log('User cancelled the sign-in flow');
          } else if (nativeError.code === 'IN_PROGRESS') {
            console.log('Sign-in already in progress');
          } 
        }
      }
    } catch (error: any) {
      console.error('Error during Google Sign-In:', error);
      Alert.alert('Sign In Error', error.message || 'An unexpected error occurred during sign-in.');
      setIsLoading(false);
    }
  };

  const handleExplore = () => {
    router.push('/(tabs)/dashboard');
  };

  // Show a loading indicator instead of returning null
  if (!fontsLoaded && !fontTimeout) {
    return (
      <View style={{ flex: 1, backgroundColor: '#121212', justifyContent: 'center', alignItems: 'center' }}>
        <ActivityIndicator size="large" color="#ffffff" />
        <Text style={{ color: '#ffffff', marginTop: 16, fontWeight: '500' }}>
          Loading SpaceCode...
        </Text>
      </View>
    );
  }

  return (
    <View className="flex-1 bg-dark-bg">
      {/* Header */}
      <View className="flex-row items-center justify-center p-4 pb-2 bg-dark-bg">
        <Text 
          className="text-light-text text-lg font-bold text-center px-12"
          style={{ fontFamily: 'Roboto_700Bold', letterSpacing: -0.24 }}
        >
          SpaceCode
        </Text>
      </View>

      {/* Hero Image */}
      <View className="w-full px-4 py-3">
        <ImageBackground
          source={{ uri: 'https://img.icons8.com/fluency/240/rocket.png' }} 
          className="w-full h-80 rounded-xl overflow-hidden bg-accent"
          resizeMode="contain"
        />
      </View>

      {/* Content */}
      <Text 
        className="text-light-text text-2xl font-bold text-center px-4 pt-5 pb-3"
        style={{ fontFamily: 'Roboto_700Bold', lineHeight: 34 }}
      >
        Master coding with spaced repetition
      </Text>
      
      <Text 
        className="text-light-text text-base text-center px-4 pt-1 pb-3"
        style={{ fontFamily: 'Roboto_400Regular', lineHeight: 24 }}
      >
        Improve your coding skills by practicing with SpaceCode's spaced repetition system.
      </Text>

      {/* Primary Button */}
      <View className="flex-row justify-center px-4 py-3 gap-3">
        <TouchableOpacity
          onPress={handleGoogleSignIn}
          className='flex-row items-center justify-center w-full max-w-xs p-3 rounded-lg shadow-md bg-white dark:bg-gray-700'
        > 
            <Ionicons name="logo-google" size={24} color={Platform.OS === 'web' ? '#4285F4' : '#FFFFFF'} className="mr-3" />
          
          <Text className="text-lg font-medium text-black dark:text-white">
            Sign in with Google
          </Text>
        </TouchableOpacity>
      </View>
      
      {/* Secondary Button
      <View className="flex-row justify-between px-4 gap-3">
        <TouchableOpacity 
          className="flex-1 h-12 rounded-xl justify-center items-center px-4"
          onPress={handleExplore}
        >
          <Text 
            className="text-light-text text-base font-bold"
            style={{ fontFamily: 'Roboto_700Bold' }}
          >
            Explore Problems
          </Text>
        </TouchableOpacity>
      </View> */}
      
      {/* Bottom spacer */}
      <View className="h-5 bg-dark-bg" />
    </View>
  );
}