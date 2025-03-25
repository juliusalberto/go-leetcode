import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { Ionicons } from '@expo/vector-icons';

// Define props for ProblemCard
interface ProblemCardProps {
  title: string;
  subtitle: string;
  subtitleColor?: string;
  iconName?: keyof typeof Ionicons.glyphMap;
  completed?: boolean;
  onPress?: () => void;
  rightElement?: React.ReactNode;
}

/**
 * A reusable card component for displaying problems and reviews
 */
export default function ProblemCard({
  title,
  subtitle,
  subtitleColor = "#8A9DC0",
  iconName = "checkmark-circle-outline",
  completed = false,
  onPress,
  rightElement
}: ProblemCardProps) {
  return (
    <View className="flex items-center bg-[#131C24] px-4 min-h-[80px] py-2">
      <View className="flex-row items-center w-full h-full justify-between">
        <TouchableOpacity
          className="flex-row items-center flex-1 gap-4"
          onPress={onPress}
          disabled={!onPress}
        >
          <View className="text-[#F8F9FB] flex items-center justify-center rounded-lg bg-[#29374C] shrink-0 size-12">
            {iconName === "checkmark-circle-outline" ? (
              <Ionicons
                name="checkmark-circle-outline"
                size={24}
                color={completed ? "#4CD137" : "#FFFFFF"}
              />
            ) : (
              <Ionicons name={iconName} size={24} color="#FFFFFF" />
            )}
          </View>
          <View className="flex flex-col justify-center flex-1">
            <Text className="text-[#F8F9FB] text-base font-medium leading-normal line-clamp-1">
              {title}
            </Text>
            <Text
              className="text-sm font-normal leading-normal line-clamp-2"
              style={{ color: subtitleColor }}
            >
              {subtitle}
            </Text>
          </View>
        </TouchableOpacity>
        
        {rightElement}
      </View>
    </View>
  );
}