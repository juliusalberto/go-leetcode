import React, { useState, useEffect } from 'react';
import { View, Text, TextInput, TouchableOpacity, FlatList, ActivityIndicator, ScrollView } from 'react-native';
import { useFonts, Roboto_400Regular, Roboto_500Medium, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';

// Import types and API function
import { Problem } from '../services/leetcode/types';
import { fetchProblems } from '../services/api/problems';
import DropdownFilter from "../../components/ui/DropdownFilter"

// Problem difficulty colors
const difficultyColors: Record<string, string> = {
  Easy: '#1CBABA', // Green
  Medium: '#FFB700', // Orange
  Hard: '#F63737'  // Red
};

export default function ProblemsScreen() {
  console.log("ProblemsScreen rendering");
  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_500Medium,
    Roboto_700Bold,
  });
  console.log("Fonts loaded:", fontsLoaded);

  // States
  const [problems, setProblems] = useState<Problem[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [offset, setOffset] = useState(1);
  const [hasMore, setHasMore] = useState(true);

  // Filter states
  const [difficulty, setDifficulty] = useState<string | null>(null);
  const [tags, setTags] = useState<string | null>(null);
  

  const loadProblems = async (requestedOffset = 0, append = false) => {
    console.log("loadProblems called with offset:", requestedOffset, "append:", append);
    
    if ((!hasMore && append)) {
      console.log("Early return from loadProblems hasMore:", hasMore, "append:", append);
      return;
    }
  
    try {
        console.log("Making API call to fetch problems...");
        const response = await fetchProblems({
        difficulty: difficulty || undefined,
        tags: tags || undefined,
        search: searchQuery || undefined,
        offset: requestedOffset,
        limit: 20, // can adjust this as needed
        });

        console.log("API response received:", response);
        
        if (append) {
          console.log("Appending problems to existing list");
          setProblems((prev) => [...prev, ...response.data]);
        } else {
          console.log("Setting new problems list");
          setProblems(response.data);
        }

        setHasMore(response.data.length === 20); // assumes more data exists if exactly 20 items returned
    } catch (error) {
        console.error('Error loading problems:', error);
        if (!append) setProblems([]);
    } 
  };
  
    useEffect(() => {
        console.log("First useEffect triggered. fontsLoaded:", fontsLoaded);
        if (!fontsLoaded) return;
        
        console.log("Calling loadProblems from first useEffect");
        loadProblems();
    }, [fontsLoaded, difficulty, tags, searchQuery]);

    useEffect(() => {
        console.log("Second useEffect triggered. difficulty:", difficulty, "tags:", tags, "search:", searchQuery);
        setOffset(0);
        setHasMore(true);
        console.log("Calling loadProblems from second useEffect");
        loadProblems(0, false);
    }, [difficulty, tags, searchQuery]);

    const handleProblemPress = (problem: Problem) => {
        router.push(`/problem/${problem.title_slug}`);
    };

  if (!fontsLoaded) {
    return null;
  }

  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Header */}
      <View className="flex items-center bg-[#131C24] p-4 pb-2 justify-between flex-row">
        <Text 
          className="text-[#F8F9FB] text-lg font-bold leading-tight flex-1 text-center"
          style={{ fontFamily: 'Roboto_700Bold' }}
        >
          Problem Library
        </Text>
      </View>

      {/* Search Bar */}
      <View className="px-4 py-3">
        <View className="flex flex-col min-w-40 h-12 w-full">
          <View className="flex w-full flex-1 items-stretch rounded-xl h-full flex-row">
            <View className="text-[#8A9DC0] flex border-none bg-[#29374C] items-center justify-center pl-4 rounded-l-xl border-r-0">
              <Ionicons name="search" size={24} color="#8A9DC0" />
            </View>
            <TextInput
              placeholder="Find a problem"
              className="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-xl text-[#F8F9FB] focus:outline-0 focus:ring-0 border-none bg-[#29374C] focus:border-none h-full placeholder:text-[#8A9DC0] px-4 rounded-l-none border-l-0 pl-2 text-base font-normal leading-normal"
              placeholderTextColor="#8A9DC0"
              value={searchQuery}
              onChangeText={setSearchQuery}
              style={{ fontFamily: 'Roboto_400Regular' }}
            />
          </View>
        </View>
      </View>

      {/* Filter Buttons */}
      <View className='flex flex-row flex-wrap px-4 py-2 gap-2'>
        <DropdownFilter
            label="Difficulty"
            selectedValue={difficulty}
            options={['Easy', 'Medium', 'Hard']}
            onSelect={(value) => setDifficulty(value)}
        />

        <DropdownFilter
            label="Tag"
            selectedValue={tags}
            options={['Array', 'String', 'Hash Table', 'Dynamic Programming', 'Math']}
            onSelect={(value) => setTags(value)}
        />
      </View>

      {/* Problems List */}
      {
        <FlatList
            data={problems}
            maintainVisibleContentPosition={{ minIndexForVisible: 0 }}
            keyExtractor={(item, index) => `${item.id}-${item.frontend_id}-${index}`}
            renderItem={({ item }) => (
                <TouchableOpacity 
                className="flex gap-4 bg-[#131C24] px-4 py-3"
                onPress={() => handleProblemPress(item)}
                >
                    <View className="flex flex-row gap-4">
                        <View className="text-[#F8F9FB] flex items-center justify-center rounded-lg bg-[#29374C] shrink-0 size-12">
                            {item.completed ? (
                                <Ionicons name="checkmark-circle-outline" size={24} color="#4CD137" />
                            ) : (
                                <Ionicons name="checkmark-circle-outline" size={24} color="#FFFFFF" />
                            )}
                        </View>
                        <View className="flex flex-1 flex-col justify-center">
                            <Text 
                                className="text-[#F8F9FB] text-base font-medium leading-normal"
                                style={{ fontFamily: 'Roboto_500Medium' }}
                            >
                                {item.title}
                            </Text>
                            <Text 
                                style={{ 
                                fontFamily: 'Roboto_400Regular',
                                fontSize: 14,
                                lineHeight: 20,
                                color: difficultyColors[item.difficulty] || '#8A9DC0'
                                }}
                            >
                                {item.difficulty}
                            </Text>
                        </View>
                    </View>
                </TouchableOpacity>
            )}
            onEndReached={() => {
                if (hasMore) {
                const newOffset = offset + 20;
                loadProblems(newOffset, true);
                setOffset(newOffset);
                }
            }}
            onEndReachedThreshold={0.5}
            ListFooterComponent={() =>
                offset > 0 ? (
                <View className="py-4">
                    <ActivityIndicator size="small" color="#6366F1" />
                </View>
                ) : null
            }
            ListEmptyComponent={() => (
                <View className="p-4 items-center">
                <Text className="text-[#8A9DC0] text-base text-center">
                    No problems found. Try adjusting your search or filters.
                </Text>
                </View>
            )}
        />
      }
    </View>
  );
}

// Filter button component
function FilterButton({ label, active, onPress }: { label: string; active: boolean; onPress: () => void }) {
  return (
    <TouchableOpacity 
      className={`flex-row items-center px-3 py-2 mr-3 rounded-xl ${active ? 'bg-[#6366F1]' : 'bg-[#29374C]'}`}
      onPress={onPress}
    >
      <Text 
        className={`${active ? 'text-white' : 'text-[#8A9DC0]'} text-sm font-medium mr-1`}
        style={{ fontFamily: 'Roboto_500Medium' }}
      >
        {label}
      </Text>
      <Ionicons 
        name={active ? "chevron-up" : "chevron-down"} 
        size={16} 
        color={active ? "#FFFFFF" : "#8A9DC0"} 
      />
    </TouchableOpacity>
  );
}