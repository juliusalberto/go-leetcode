import React, { useState } from 'react';
import { TouchableOpacity, View, TextInput, KeyboardAvoidingView, Platform, Text } from 'react-native';
import { useFonts, Roboto_400Regular, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { router } from 'expo-router';
import Button from '../components/ui/Button'; // Import the Button component

export default function UsernameInputScreen() {
  const [username, setUsername] = useState('');
  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_700Bold,
  });

  const handleStart = () => {
    // Validate username
    if (username.trim() === '') {
      // Show error or toast message
      return;
    }
    
    // Navigate to the main app or save the username
    console.log('Username submitted:', username);
    
    // Navigate to the next screen
    router.push('/(tabs)/dashboard');
  };

  if (!fontsLoaded) {
    return null;
  }

  return (
    <KeyboardAvoidingView 
      className="flex-1"
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
    >
      <View className="flex-1 bg-[#131C24]">
        {/* Header - empty space where the magnifying glass was */}
        <View className="h-14 px-4 pb-2 flex-row justify-end items-center" />
        
        {/* Title */}
        <Text 
          className="text-[#F8F9FB] text-[22px] font-bold text-center px-4 pt-5 pb-3"
          style={{ fontFamily: 'Roboto_700Bold', letterSpacing: -0.3 }}
        >
          Welcome to LeetCode Flashcards
        </Text>
        
        {/* Description */}
        <Text 
          className="text-[#F8F9FB] text-base text-center px-4 pt-1 pb-3"
          style={{ fontFamily: 'Roboto_400Regular', lineHeight: 24 }}
        >
          Enter your LeetCode username to start
        </Text>
        
        {/* Input Field */}
        <View className="px-4 py-3">
          <TextInput
            className="h-14 bg-[#29374C] rounded-xl px-4 text-[#F8F9FB]"
            placeholder="Username"
            placeholderTextColor="#8A9DC0"
            value={username}
            onChangeText={setUsername}
            autoCapitalize="none"
            autoCorrect={false}
            style={{ fontFamily: 'Roboto_400Regular', fontSize: 16 }}
          />
        </View>
        {/* Start Button */}
        <View className="px-4 py-3">
          <Button
            title="Start"
            onPress={handleStart}
            // Override default styles using className and textStyle
            className="bg-[#F4C753] h-12 rounded-xl" // Keep original background, height, radius
            textStyle={{
              color: '#141C24', // Original text color
              fontSize: 16, // Original text size (approximated from text-base)
              fontFamily: 'Roboto_700Bold',
              letterSpacing: 0.24
            }}
          />
        </View>
        
        {/* Bottom spacer */}
        {/* Bottom spacer */}
        <View className="h-5 bg-[#131C24]" />
      </View>
    </KeyboardAvoidingView>
  );
}