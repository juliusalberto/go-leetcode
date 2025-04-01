import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, ActivityIndicator } from 'react-native';
import { router } from 'expo-router';
import { useFonts, Roboto_400Regular, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { Ionicons } from '@expo/vector-icons';
import { StatusBar } from 'expo-status-bar';
import { useAuth } from '../contexts/AuthContext';
import { completeUserProfile } from '../services/api/auth';

export default function CompleteProfile() {
  const { session, updateProfileStatus } = useAuth();
  const [username, setUsername] = useState('');
  const [leetcodeUsername, setLeetcodeUsername] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_700Bold,
  });

  const handleBack = () => {
    router.back();
  };

  const handleContinue = async () => {
    if (!username.trim()) {
      setError('Username is required');
      return;
    }

    if (!leetcodeUsername.trim()) {
      setError('LeetCode username is required');
      return;
    }

    if (!session?.access_token) {
      setError('Not authenticated');
      return;
    }

    setIsLoading(true);
    setError(null);

    try {
      // Send profile data to your custom backend API and get updated status
      const updatedStatus = await completeUserProfile(session.access_token, {
        username: username.trim(),
        leetcode_username: leetcodeUsername.trim(),
      });
      
      console.log("Profile completed successfully, status:", updatedStatus);
      
      // Tell the auth context that the profile now exists
      // This is crucial for preventing the navigation loop
      updateProfileStatus(true);
      
      // Navigate to dashboard on success
      router.replace('/(tabs)/dashboard');
    } catch (err: any) {
      console.error('Error creating profile:', err);
      setError(err.message || 'Failed to create profile');
    } finally {
      setIsLoading(false);
    }
  };

  if (!fontsLoaded) {
    return null;
  }

  return (
    <View className="flex-1 bg-[#131C24]">
      <StatusBar style="light" />
      
      {/* Header with back button */}
      <View className="flex-row items-center bg-[#131C24] p-4 pb-2 justify-between">
        <TouchableOpacity onPress={handleBack} className="text-[#F8F9FB] flex size-12 shrink-0 items-center justify-center">
          <Ionicons name="arrow-back" size={24} color="#F8F9FB" />
        </TouchableOpacity>
      </View>
      
      {/* Welcome Text */}
      <Text 
        className="text-[#F8F9FB] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 text-left pb-3 pt-5"
        style={{ fontFamily: 'Roboto_700Bold' }}
      >
        Welcome back!
      </Text>
      
      {/* Username Input */}
      <View className="flex flex-wrap gap-4 px-4 py-3">
        <View className="flex-col min-w-40 flex-1">
          <TextInput
            placeholder="username"
            className="w-full resize-none overflow-hidden rounded-xl text-[#F8F9FB] bg-[#29374C] h-14 placeholder:text-[#8A9DC0] p-4 text-base font-normal leading-normal"
            placeholderTextColor="#8A9DC0"
            value={username}
            onChangeText={setUsername}
          />
        </View>
      </View>
      
      {/* LeetCode Username Input */}
      <View className="flex  flex-wrap  gap-4 px-4 py-3">
        <View className="flex-col min-w-40 flex-1">
          <TextInput
            placeholder="LeetCode username"
            className="w-full resize-none overflow-hidden rounded-xl text-[#F8F9FB] bg-[#29374C] h-14 placeholder:text-[#8A9DC0] p-4 text-base font-normal leading-normal"
            placeholderTextColor="#8A9DC0"
            value={leetcodeUsername}
            onChangeText={setLeetcodeUsername}
          />
        </View>
      </View>
      
      {/* Error Message */}
      {error ? (
        <Text className="text-red-500 px-4 py-1 text-sm">{error}</Text>
      ) : null}
      
      {/* Continue Button */}
      <View className="flex justify-center">
        <View className="flex flex-1 gap-3 max-w-[480px] flex-col items-stretch px-4 py-3">
          <TouchableOpacity
            onPress={handleContinue}
            disabled={isLoading}
            className="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-xl h-10 px-4 bg-[#F4C753]"
          >
            {isLoading ? (
              <ActivityIndicator size="small" color="#141C24" />
            ) : (
              <Text 
                className="text-[#141C24] text-sm font-bold leading-normal tracking-[0.015em]"
                style={{ fontFamily: 'Roboto_700Bold' }}
              >
                Continue
              </Text>
            )}
          </TouchableOpacity>
        </View>
      </View>
      
      {/* Bottom Spacer */}
      <View className="h-5 bg-[#131C24]" />
    </View>
  );
}