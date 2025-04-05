import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { Menu, MenuOptions, MenuOption, MenuTrigger } from 'react-native-popup-menu';
import { Ionicons } from '@expo/vector-icons';
import { Deck } from '../../services/api/decks'; // Assuming Deck type is here
import { router } from 'expo-router';

interface DeckItemProps {
  deck: Deck;
  onDelete: (deckId: number) => void; // Callback for delete action
}

const DeckItem: React.FC<DeckItemProps> = ({ deck, onDelete }) => {
  const handlePress = () => {
    // Use object syntax with type assertion for typed routes
    router.push({
      pathname: `/deck/${deck.id}` as '/deck/[id]',
      params: {
        id: deck.id, // Include the dynamic segment parameter 'id'
        name: deck.name,
        is_public: deck.is_public.toString() // Convert boolean to string for params
      },
    });
  };

  return (
    <TouchableOpacity
      onPress={handlePress}
      activeOpacity={0.8}
    >
      {/* Use border instead of background, increase padding */}
      <View className="border border-[#32415D] rounded-lg p-4 mb-4">
        {/* Top row: Title and Menu */}
        <View className="flex-row justify-between items-start mb-2"> {/* Use items-start for better alignment */}
          <Text className="text-[#F8F9FB] text-lg font-semibold flex-1 mr-2">{deck.name}</Text> {/* Smaller title */}
          {!deck.is_public && (
            <Menu>
              <MenuTrigger>
                {/* Adjust padding/margin for trigger */}
                <View className="p-1 -mr-1">
                  <Ionicons name="ellipsis-vertical" size={20} color="#8A9DC0" />
                </View>
              </MenuTrigger>
              <MenuOptions>
                <MenuOption
                  // style={{ borderRadius: 100 }} // Optional: Style via options container if needed
                  onSelect={() => onDelete(deck.id)}
                >
                  <Text className="p-2 text-red-600">Delete Deck</Text> {/* Destructive text color */}
                </MenuOption>
              </MenuOptions>
            </Menu>
          )}
        </View>
        {/* Description */}
        {deck.description && ( // Only show if description exists
          <Text className="text-[#ADBAC7] text-sm mb-3">{deck.description}</Text>
        )}
        {/* Bottom row: Stats */}
        <View className="flex-row justify-end items-center"> {/* Changed justify-between to justify-end */}
          {/* Removed problem count display */}
          <Text className="text-[#8A9DC0] text-sm">
            {deck.is_public ? 'Public Deck' : 'Your Deck'}
          </Text>
        </View>
      </View>
    </TouchableOpacity>
  );
};

export default React.memo(DeckItem);