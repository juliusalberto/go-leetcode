import React from 'react';
import { View, Text } from 'react-native';
import { Roboto_500Medium, Roboto_700Bold } from '@expo-google-fonts/roboto';

interface StatCardProps {
  title: string;
  value: string | number;
}

export default function StatCard({ title, value }: StatCardProps) {
  return (
    <View className="flex-1 min-w-[158px] flex-col gap-2 rounded-xl p-6 border border-[#32415D]">
      <Text 
        className="text-[#F8F9FB] text-base leading-normal"
        style={{ fontFamily: 'Roboto_500Medium' }}
      >
        {title}
      </Text>
      <Text 
        className="text-[#F8F9FB] text-2xl font-bold leading-tight"
        style={{ fontFamily: 'Roboto_700Bold' }}
      >
        {value}
      </Text>
    </View>
  );
}