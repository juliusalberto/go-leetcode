import React, { useState } from 'react';
import { View, TouchableOpacity, Text, Modal, TouchableWithoutFeedback, FlatList } from 'react-native';
import { Ionicons } from '@expo/vector-icons';

interface DropdownFilterProps {
  label: string;
  selectedValue: string | null;
  options: string[];
  onSelect: (value: string | null) => void;
}

export default function DropdownFilter({ label, selectedValue, options, onSelect }: DropdownFilterProps) {
  const [open, setOpen] = useState(false);

  const handleSelect = (option: string) => {
    onSelect(selectedValue === option ? null : option);
    setOpen(false);
  };

  return (
    <>
      <TouchableOpacity
        className={`flex-row items-center px-3 py-2 rounded-xl ${selectedValue ? 'bg-[#6366F1]' : 'bg-[#29374C]'}`}
        onPress={() => setOpen(true)}
      >
        <Text className={`${selectedValue ? 'text-white' : 'text-[#8A9DC0]'} text-sm font-medium mr-1`}>
          {selectedValue || label}
        </Text>
        <Ionicons
          name={open ? 'chevron-up' : 'chevron-down'}
          size={16}
          color={selectedValue ? '#FFFFFF' : '#8A9DC0'}
        />
      </TouchableOpacity>

      <Modal
        visible={open}
        transparent
        animationType="fade"
        onRequestClose={() => setOpen(false)}
      >
        <TouchableWithoutFeedback onPress={() => setOpen(false)}>
          <View className="flex-1 bg-black/50 justify-center items-center">
            <View className="bg-[#29374C] rounded-xl shadow-lg p-2 w-1/2 max-w-xs">
              <FlatList
                data={options}
                keyExtractor={(item) => item}
                renderItem={({ item }) => (
                  <TouchableOpacity
                    className={`px-4 py-3 rounded-md ${selectedValue === item ? 'bg-[#6366F1]' : ''}`}
                    onPress={() => handleSelect(item)}
                  >
                    <Text className="text-[#F8F9FB] text-center font-medium">{item}</Text>
                  </TouchableOpacity>
                )}
              />
            </View>
          </View>
        </TouchableWithoutFeedback>
      </Modal>
    </>
  );
}
