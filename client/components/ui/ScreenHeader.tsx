import React from 'react';
import { View, Text, TouchableOpacity, StyleProp, ViewStyle, TextStyle } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';
import { useFonts, Roboto_700Bold } from '@expo-google-fonts/roboto'; // Import font

interface ScreenHeaderProps {
  title?: string;
  showBackButton?: boolean;
  onBackPress?: () => void;
  rightElement?: React.ReactNode;
  centerTitle?: boolean;
  containerStyle?: StyleProp<ViewStyle>;
  titleStyle?: StyleProp<TextStyle>;
}

const ScreenHeader: React.FC<ScreenHeaderProps> = ({
  title,
  showBackButton = true, // Default to true
  onBackPress = () => router.back(), // Default to router.back()
  rightElement,
  centerTitle = false,
  containerStyle,
  titleStyle,
}) => {
  const [fontsLoaded] = useFonts({ Roboto_700Bold });

  if (!fontsLoaded) {
    // Render a placeholder or null while fonts load to prevent layout shifts
    return <View style={[{ minHeight: 56 }, containerStyle]} />; // Adjust minHeight as needed
  }

  const defaultContainerStyle: ViewStyle = {
    flexDirection: 'row',
    alignItems: 'center',
    paddingHorizontal: 16,
    paddingVertical: 8, // Adjusted padding
    backgroundColor: '#131C24',
    minHeight: 56, // Ensure consistent height
  };

  const defaultTitleStyle: TextStyle = {
    color: '#F8F9FB',
    fontSize: 18,
    fontFamily: 'Roboto_700Bold',
    textAlign: centerTitle ? 'center' : 'left',
    flex: 1, // Allow title to take available space
    marginHorizontal: showBackButton || rightElement ? 8 : 0, // Add margin if buttons/elements exist
  };

  return (
    <View style={[defaultContainerStyle, containerStyle]}>
      {showBackButton ? (
        <TouchableOpacity onPress={onBackPress} style={{ padding: 8, marginLeft: -8 }}>
          <Ionicons name="chevron-back" size={24} color="#F8F9FB" />
        </TouchableOpacity>
      ) : (
        // Optional: Render a spacer if no back button but right element exists for alignment
        rightElement ? <View style={{ width: 40 }} /> : null
      )}

      {title && (
        <Text style={[defaultTitleStyle, titleStyle]} numberOfLines={1}>
          {title}
        </Text>
      )}

      {rightElement ? (
        <View style={{ padding: 8, marginRight: -8 }}>{rightElement}</View>
      ) : (
         // Optional: Render a spacer if back button exists but no right element for alignment
         showBackButton && title ? <View style={{ width: 40 }} /> : null
      )}
    </View>
  );
};

export default ScreenHeader;