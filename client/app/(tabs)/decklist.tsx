import React, { useState } from 'react';
import { View, Text, FlatList, TouchableOpacity, ActivityIndicator, Alert, Modal, TextInput } from 'react-native';
import { MenuProvider, Menu, MenuOptions, MenuOption, MenuTrigger } from 'react-native-popup-menu';
import { useQueryClient } from '@tanstack/react-query'; // Import useQueryClient
import { useDecks, useCreateDeck, useDeleteDeck, Deck } from '../../services/api/decks'; // Import useDeleteDeck
import { Ionicons } from '@expo/vector-icons';
import { router } from 'expo-router';
import Button from '../../components/ui/Button'; // Import the Button component
import Toast from 'react-native-toast-message'; // Import Toast
import DeckItem from '../../components/ui/DeckItem'; // Import the DeckItem component

export default function DecksScreen() {
  const { data: responseData = { public_decks: [], user_decks: [] }, isLoading, error } = useDecks();
  const publicDecks = responseData.public_decks || [];
  const userDecks = responseData.user_decks || [];
  const createDeck = useCreateDeck();
  const deleteDeckMutation = useDeleteDeck(); // Initialize the delete mutation
  const [showModal, setShowModal] = useState(false);
  const [deckName, setDeckName] = useState('');


  const handleDeleteDeck = (deckId: number) => {
    // Directly call the mutation without confirmation alert
    if (deleteDeckMutation.isPending) return; // Prevent multiple clicks

    deleteDeckMutation.mutate(deckId, {
      onSuccess: () => {
        Toast.show({
          type: 'success',
          text1: 'Deck Deleted',
          text2: 'The deck was successfully deleted.',
        });
        // List will refresh via query invalidation
      },
      onError: (error) => {
        Toast.show({
          type: 'error',
          text1: 'Deletion Failed',
          text2: error.message || 'Could not delete the deck.',
        });
      },
    });
  };

  // renderDeckItem function removed, logic moved to DeckItem component

  // Note: The 'if (isLoading)' check below was duplicated, removing the second instance.
  if (isLoading) {
    return (
      <View className="flex-1 justify-center items-center bg-[#131C24]">
        <ActivityIndicator size="large" color="#6366F1" />
      </View>
    );
  }

  // Prepare combined data for a single list
  type ListItem = { type: 'header'; title: string } | { type: 'deck'; data: Deck };
  const combinedData: ListItem[] = [];
  if (publicDecks.length > 0) {
    combinedData.push({ type: 'header', title: 'Public Decks' });
    publicDecks.forEach((deck: Deck) => combinedData.push({ type: 'deck', data: deck }));
  }
  if (userDecks.length > 0) {
    combinedData.push({ type: 'header', title: 'Your Decks' });
    userDecks.forEach((deck: Deck) => combinedData.push({ type: 'deck', data: deck }));
  }

  // Render item function for the combined list
  const renderCombinedItem = ({ item }: { item: ListItem }) => {
    if (item.type === 'header') {
      return (
        <View className="p-2 mb-2">
          <Text className="text-[#F8F9FB] text-lg font-bold">{item.title}</Text>
        </View>
      );
    } else { // item.type === 'deck'
      return <DeckItem deck={item.data} onDelete={handleDeleteDeck} />;
    }
  };

  // Key extractor for the combined list
  const keyExtractor = (item: ListItem, index: number) => {
    if (item.type === 'header') {
      return `header-${item.title}-${index}`;
    } else {
      return `deck-${item.data.id.toString()}`;
    }
  };
  
  return (
    <MenuProvider>
      <View className="flex-1 bg-[#131C24]">
      {/* Header */}
      <View className="flex-row items-center justify-center p-4">
        <Text className="text-[#F8F9FB] text-xl font-bold">Problem Decks</Text>
      </View>
      
      <FlatList
        data={combinedData}
        renderItem={renderCombinedItem}
        keyExtractor={keyExtractor}
        contentContainerStyle={{ paddingHorizontal: 16, paddingBottom: 80 }}
        ListEmptyComponent={
          <View className="flex-1 justify-center items-center p-4 mt-10">
            <Text className="text-[#8A9DC0] text-center">
              No decks found. Create your first deck!
            </Text>
          </View>
        }
      />

      {/* Floating Action Button remains the same */}
      <View className="absolute bottom-6 right-6">
        <TouchableOpacity
          className="bg-[#6366F1] p-4 rounded-full shadow-lg"
          onPress={() => setShowModal(true)}
          activeOpacity={0.8}
        >
          <Ionicons name="add" size={28} color="#F8F9FB" />
        </TouchableOpacity>
      </View>

      <Modal
        animationType="slide"
        transparent={true}
        visible={showModal}
        onRequestClose={() => setShowModal(false)}
      >
        {/* Semi-transparent background */}
        <View className="flex-1 justify-center items-center bg-black/60 px-4">
          {/* Modal content container with dark theme */}
          <View className="bg-[#1E2A3A] p-6 rounded-xl w-full max-w-sm shadow-xl">
            <Text className="text-[#F8F9FB] text-xl font-bold mb-5 text-center">Create New Deck</Text>
            <TextInput
              className="bg-[#29374C] border border-[#32415D] rounded-lg p-3 mb-5 text-[#F8F9FB] placeholder:text-[#8A9DC0]"
              placeholder="Deck name"
              placeholderTextColor="#8A9DC0" // Explicitly set placeholder color
              value={deckName}
              onChangeText={setDeckName}
              autoFocus={true} // Focus input on modal open
            />
            {/* Button container */}
            <View className="flex-row justify-end space-x-3">
              {/* Cancel Button - styled for dark theme */}
              <Button
                title="Cancel"
                onPress={() => setShowModal(false)}
                variant="ghost"
                className="px-4 py-2"
                textStyle={{ color: '#8A9DC0' }} // Lighter gray for dark background
              />
              {/* Create Button */}
              <Button
                title="Create"
                variant="primary"
                className="px-5 py-2 rounded-lg" // Slightly more padding
                isLoading={createDeck.isPending} // Show loading state
                onPress={() => {
                  if (deckName.trim()) {
                    createDeck.mutate({
                      name: deckName,
                      description: '',
                      is_public: false
                    }, {
                      onSuccess: () => {
                        setDeckName('');
                        setShowModal(false);
                      },
                      onError: (error: Error) => {
                        Alert.alert('Error', `Failed to create deck: ${error.message}`);
                      }
                    });
                  }
                }}
              />
            </View>
          </View>
        </View>
      </Modal>
      </View>
    </MenuProvider>
  );
}