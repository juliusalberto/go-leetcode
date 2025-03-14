import React from 'react';
import { View, Text, ScrollView, TouchableOpacity, Dimensions, Platform } from 'react-native';
import { useFonts, Roboto_400Regular, Roboto_500Medium, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';
import { LineChart } from 'react-native-chart-kit';

export default function DashboardScreen() {
  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_500Medium,
    Roboto_700Bold,
  });

  // Sample streak data for the past 30 days (you would replace this with real data)
  const screenWidth = Dimensions.get('window').width - 40;
  const streakData = [2, 3, 2, 4, 3, 5, 4]
  
  if (!fontsLoaded) {
    return null;
  }

  const handleProblemPress = (problem) => {
    // Navigate to problem details
    router.push(`/problem/${problem.slug}`);
  };

  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Header */}
      <View className="flex-row items-center bg-[#131C24] p-4 pb-2 justify-between">
        <Text 
          className="text-[#F8F9FB] text-lg font-bold leading-tight flex-1 text-center px-12"
          style={{ fontFamily: 'Roboto_700Bold', letterSpacing: -0.24 }}
        >
          Dashboard
        </Text>
      </View>

      <ScrollView>
        {/* Current Streak */}
        <View className="flex-row items-center gap-4 bg-[#131C24] px-4 min-h-[72px] py-2">
          <View className="items-center justify-center rounded-lg bg-[#29374C] shrink-0 w-12 h-12">
            <Ionicons name="flame-outline" size={24} color="#F8F9FB" />
          </View>
          <View className="flex-col justify-center">
            <Text 
              className="text-[#F8F9FB] text-base leading-normal"
              style={{ fontFamily: 'Roboto_500Medium' }}
            >
              Current Streak
            </Text>
            <Text 
              className="text-[#8A9DC0] text-sm leading-normal"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              4 days
            </Text>
          </View>
        </View>

        {/* Upcoming Reviews */}
        <View className="flex-row items-center gap-4 bg-[#131C24] px-4 min-h-[72px] py-2">
          <View className="items-center justify-center rounded-lg bg-[#29374C] shrink-0 w-12 h-12">
            <Ionicons name="calendar-outline" size={24} color="#F8F9FB" />
          </View>
          <View className="flex-col justify-center">
            <Text 
              className="text-[#F8F9FB] text-base leading-normal"
              style={{ fontFamily: 'Roboto_500Medium' }}
            >
              Upcoming Reviews
            </Text>
            <Text 
              className="text-[#8A9DC0] text-sm leading-normal"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              5 reviews
            </Text>
          </View>
        </View>

        {/* Recently Attempted */}
        <Text 
          className="text-[#F8F9FB] text-[22px] font-bold leading-tight px-4 pb-3 pt-5"
          style={{ fontFamily: 'Roboto_700Bold', letterSpacing: -0.24 }}
        >
          Recently Attempted
        </Text>

        <TouchableOpacity 
          className="flex-row items-center gap-4 bg-[#131C24] px-4 min-h-14 justify-between"
          onPress={() => handleProblemPress({ slug: 'two-sum' })}
        >
          <Text 
            className="text-[#F8F9FB] text-base leading-normal flex-1 truncate"
            style={{ fontFamily: 'Roboto_400Regular' }}
          >
            Two Sum
          </Text>
          <View className="shrink-0">
            <Text 
              className="text-[#8A9DC0] text-sm leading-normal"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              3 days ago
            </Text>
          </View>
        </TouchableOpacity>

        <TouchableOpacity 
          className="flex-row items-center gap-4 bg-[#131C24] px-4 min-h-14 justify-between"
          onPress={() => handleProblemPress({ slug: 'add-two-numbers' })}
        >
          <Text 
            className="text-[#F8F9FB] text-base leading-normal flex-1 truncate"
            style={{ fontFamily: 'Roboto_400Regular' }}
          >
            Add Two Numbers
          </Text>
          <View className="shrink-0">
            <Text 
              className="text-[#8A9DC0] text-sm leading-normal"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              3 days ago
            </Text>
          </View>
        </TouchableOpacity>

        <TouchableOpacity 
          className="flex-row items-center gap-4 bg-[#131C24] px-4 min-h-14 justify-between"
          onPress={() => handleProblemPress({ slug: 'number-of-islands' })}
        >
          <Text 
            className="text-[#F8F9FB] text-base leading-normal flex-1 truncate"
            style={{ fontFamily: 'Roboto_400Regular' }}
          >
            Number of Islands
          </Text>
          <View className="shrink-0">
            <Text 
              className="text-[#8A9DC0] text-sm leading-normal"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              2 days ago
            </Text>
          </View>
        </TouchableOpacity>

        <TouchableOpacity 
          className="flex-row items-center gap-4 bg-[#131C24] px-4 min-h-14 justify-between"
          onPress={() => handleProblemPress({ slug: 'longest-substring-without-repeating-characters' })}
        >
          <Text 
            className="text-[#F8F9FB] text-base leading-normal flex-1 truncate"
            style={{ fontFamily: 'Roboto_400Regular' }}
          >
            Longest Substring Without Repeating Characters
          </Text>
          <View className="shrink-0">
            <Text 
              className="text-[#8A9DC0] text-sm leading-normal"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              1 day ago
            </Text>
          </View>
        </TouchableOpacity>

        {/* Stats Cards */}
        <View className="flex-row flex-wrap gap-4 p-4">
          <View className="flex-1 min-w-[158px] flex-col gap-2 rounded-xl p-6 border border-[#32415D]">
            <Text 
              className="text-[#F8F9FB] text-base leading-normal"
              style={{ fontFamily: 'Roboto_500Medium' }}
            >
              Total Problems Solved
            </Text>
            <Text 
              className="text-[#F8F9FB] text-2xl font-bold leading-tight"
              style={{ fontFamily: 'Roboto_700Bold' }}
            >
              234
            </Text>
          </View>

          <View className="flex-1 min-w-[158px] flex-col gap-2 rounded-xl p-6 border border-[#32415D]">
            <Text 
              className="text-[#F8F9FB] text-base leading-normal"
              style={{ fontFamily: 'Roboto_500Medium' }}
            >
              Problems Solved This Week
            </Text>
            <Text 
              className="text-[#F8F9FB] text-2xl font-bold leading-tight"
              style={{ fontFamily: 'Roboto_700Bold' }}
            >
              12
            </Text>
          </View>
        </View>
        
        {/* Streak Progress Chart - Simple Placeholder */}
        <View className="mx-4 mb-6 p-6 rounded-xl border border-[#32415D]">
          <Text 
            className="text-[#F8F9FB] text-base leading-normal mb-2"
            style={{ fontFamily: 'Roboto_500Medium' }}
          >
            Streak Progress (Last 7 Days)
          </Text>
          
          
          <View className="h-[150px] items-center justify-center">
            <LineChart
                  data={{
                    labels: ['', '', '', '', '', ''],
                    datasets: [
                      {
                        data: streakData
                      }
                    ]
                  }}
                  width={screenWidth}
                  height={180}
                  chartConfig={{
                    backgroundColor: 'transparent',
                    backgroundGradientFrom: 'transparent',
                    backgroundGradientTo: 'transparent',
                    decimalPlaces: 0,
                    color: (opacity = 1) => `rgba(99, 102, 241, ${opacity})`,
                    labelColor: (opacity = 1) => `rgba(138, 157, 192, ${opacity})`,
                    style: {
                      borderRadius: 16
                    },
                    propsForDots: {
                      r: '5',
                      strokeWidth: '2',
                      stroke: '#6366F1'
                    }
                  }}
                  bezier
                  style={{
                    marginVertical: 8,
                    borderRadius: 16
                  }}
              />
          </View>
          
          <View className="flex-row justify-between mt-2">
            <Text 
              className="text-[#8A9DC0] text-xs"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              30 days ago
            </Text>
            <Text 
              className="text-[#8A9DC0] text-xs"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              Today
            </Text>
          </View>
        </View>
      </ScrollView>
    </View>
  );
}