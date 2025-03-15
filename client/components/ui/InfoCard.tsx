import React from 'react';
import { View, Text } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { Roboto_400Regular, Roboto_500Medium } from '@expo-google-fonts/roboto';

interface InfoCardProps {
  icon: keyof typeof Ionicons.glyphMap;
  title: string;
  subtitle: string;
}

export default function InfoCard({ icon, title, subtitle }: InfoCardProps) {
  return (
    <View className="flex-row items-center gap-4 bg-[#131C24] px-4 min-h-[72px] py-2">
      <View className="items-center justify-center rounded-lg bg-[#29374C] shrink-0 w-12 h-12">
        <Ionicons name={icon} size={24} color="#F8F9FB" />
      </View>
      <View className="flex-col justify-center">
        <Text 
          className="text-[#F8F9FB] text-base leading-normal"
          style={{ fontFamily: 'Roboto_500Medium' }}
        >
          {title}
        </Text>
        <Text 
          className="text-[#8A9DC0] text-sm leading-normal"
          style={{ fontFamily: 'Roboto_400Regular' }}
        >
          {subtitle}
        </Text>
      </View>
    </View>
  );
}