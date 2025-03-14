import React from 'react';
import { TouchableOpacity, View, ImageBackground, Text } from 'react-native';
import { useFonts, Roboto_400Regular, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { router } from 'expo-router';

export default function IndexScreen() {
  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_700Bold,
  });

  const handleContinueWithGoogle = () => {
    router.push('/signin');
  };

  const handleExplore = () => {
    router.push('/(tabs)/dashboard');
  };

  if (!fontsLoaded) {
    return null; // You can replace this with a loading indicator if you like
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
      <View className="flex-row justify-between px-4 py-3 gap-3">
        <TouchableOpacity 
          className="flex-1 h-12 bg-primary rounded-xl justify-center items-center px-4"
          onPress={handleContinueWithGoogle}
        >
          <Text 
            className="text-light-text text-base font-bold"
            style={{ fontFamily: 'Roboto_700Bold', letterSpacing: 0.24 }}
          >
            Continue with Google
          </Text>
        </TouchableOpacity>
      </View>
      
      {/* Secondary Button */}
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
      </View>
      
      {/* Bottom spacer */}
      <View className="h-5 bg-dark-bg" />
    </View>
  );
}