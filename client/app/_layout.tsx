import { DarkTheme, DefaultTheme, ThemeProvider } from '@react-navigation/native';
import { useFonts } from 'expo-font';
import { Stack } from 'expo-router';
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from 'expo-status-bar';
import { useEffect, useState } from 'react';
import { View, Text, ActivityIndicator, Platform } from 'react-native';
import 'react-native-reanimated';
import '../global.css';
import { AppQueryClientProvider } from '../contexts/QueryClientProvider';
import Toast from 'react-native-toast-message';
import { AuthProvider } from '../contexts/AuthContext';
import ClientOnly from '../components/ClientOnly';
import DebugLogger, { debugLog } from '../components/DebugLogger';

import { useColorScheme } from '@/hooks/useColorScheme';

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync().catch(() => {
  // Ignore errors related to splash screen
  console.log('Failed to prevent auto hide of splash screen');
});

export default function RootLayout() {
  const colorScheme = useColorScheme();
  const [loaded, error] = useFonts({
    SpaceMono: require('../assets/fonts/SpaceMono-Regular.ttf'),
  });
  const [fontTimeout, setFontTimeout] = useState(false);

  // Set up global error handler for debugging
  useEffect(() => {
    // Log environment and platform information
    debugLog(`App initializing on ${Platform.OS}`);
    debugLog(`Supabase URL: ${process.env.EXPO_PUBLIC_SUPABASE_URL}`);
    debugLog(`Google Web Client ID: ${process.env.EXPO_PUBLIC_GOOGLE_WEB_CLIENT_ID ? 'Set' : 'Not Set'}`);
    debugLog(`Google Android Client ID: ${process.env.EXPO_PUBLIC_GOOGLE_ANDROID_CLIENT_ID ? 'Set' : 'Not Set'}`);
    
    if (Platform.OS === 'android') {
      // @ts-ignore - ErrorUtils is a global object in React Native
      const originalHandler = global.ErrorUtils.getGlobalHandler();
      
      // @ts-ignore
      global.ErrorUtils.setGlobalHandler((error, isFatal) => {
        console.log(`Global error caught: ${error.message}`);
        console.log(error.stack);
        
        // Still call original handler
        originalHandler(error, isFatal);
      });
    }
    
    return () => {
      // Reset global handler if needed
    };
  }, []);

  // Handle splash screen and add a timeout
  useEffect(() => {
    const hideSplashScreen = async () => {
      try {
        await SplashScreen.hideAsync();
      } catch (e) {
        console.log('Error hiding splash screen:', e);
      }
    };

    if (loaded) {
      hideSplashScreen();
    } else {
      // Add a safety timeout to ensure splash screen is hidden even if fonts fail to load
      const timer = setTimeout(() => {
        console.log('Font loading timeout - hiding splash screen anyway');
        setFontTimeout(true);
        hideSplashScreen();
      }, 5000); // 5 second timeout
      
      return () => clearTimeout(timer);
    }
  }, [loaded]);

  // This will display a loading UI instead of returning null
  if (!loaded && !fontTimeout) {
    return (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center', backgroundColor: '#000' }}>
        <ActivityIndicator size="large" color="#fff" />
        <Text style={{ color: '#fff', marginTop: 10 }}>Loading app...</Text>
      </View>
    );
  }

  return (
    <AuthProvider>
      <AppQueryClientProvider>
        <ThemeProvider value={colorScheme === 'dark' ? DarkTheme : DefaultTheme}>
          <Stack>
            <Stack.Screen name="index" options={{ headerShown: false }} />
            <Stack.Screen name="username" options={{ headerShown: false }} />
            <Stack.Screen name="signin" options={{ headerShown: false }} />
            <Stack.Screen name="(tabs)" options={{ headerShown: false }} />
            <Stack.Screen name="problem/[slug]" options={{ headerShown: false }} />
            <Stack.Screen name="complete-profile" options={{ headerShown: false }} />
            <Stack.Screen name="+not-found" />
          </Stack>
          <StatusBar style="auto" />
        </ThemeProvider>
        <ClientOnly>
          <Toast />
        </ClientOnly>
      </AppQueryClientProvider>
    </AuthProvider>
  );
}
