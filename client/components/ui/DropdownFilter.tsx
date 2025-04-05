import React, { useState } from 'react';
import { View, TouchableOpacity, Text, Modal, TouchableWithoutFeedback, FlatList } from 'react-native';
import { Ionicons } from '@expo/vector-icons';

// Define the structure for each option
interface OptionType {
  label: string; // The display name (e.g., "Hash Table")
  value: string;  // The internal value/slug (e.g., "hash-table")
}

interface DropdownFilterProps {
  label: string;
  selectedValue: string | null; // This still holds the slug/value
  options: OptionType[];       // Now expects an array of OptionType objects
  onSelect: (value: string | null) => void; // Still called with the slug/value or null
}

export default function DropdownFilter({ label, selectedValue, options, onSelect }: DropdownFilterProps) {
  const [open, setOpen] = useState(false);

  // Find the label corresponding to the currently selected value (slug)
  const getSelectedLabel = () => {
    if (selectedValue === null) {
      return label; // Show default label if nothing is selected
    }
    // Find the option object whose 'value' matches the selectedValue slug
    const selectedOption = options.find(option => option.value === selectedValue);
    // Return the found label, or the default label as a fallback
    return selectedOption ? selectedOption.label : label;
  };

  // This function now receives the VALUE (slug) of the selected option
  const handleSelect = (optionValue: string) => {
    // Compare the current selectedValue (slug) with the clicked option's value (slug)
    onSelect(selectedValue === optionValue ? null : optionValue);
    setOpen(false);
  };

  return (
    <>
      <TouchableOpacity
        // Style based on whether selectedValue (slug) is present
        className={`flex-row items-center px-3 py-2 rounded-xl ${selectedValue ? 'bg-[#6366F1]' : 'bg-[#29374C]'}`}
        onPress={() => setOpen(true)}
      >
        <Text className={`${selectedValue ? 'text-white' : 'text-[#8A9DC0]'} text-sm font-medium mr-1`}>
          {/* Display the calculated selected label */}
          {getSelectedLabel()}
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
          <View className="flex-1 bg-black/50 justify-center items-center px-4">
              {/* Added max-h to prevent overly tall dropdowns */}
            <View className="bg-[#29374C] rounded-xl shadow-lg p-2 w-1/2 max-w-s max-h-[70vh]">
              <FlatList
                data={options} // Pass the array of {label, value} objects
                // Use the unique value (slug) as the key
                keyExtractor={(item) => item.value}
                renderItem={({ item }) => ( // item is now { label: string, value: string }
                  <TouchableOpacity
                    // Compare selectedValue (slug) with item.value (slug) for highlighting
                    className={`px-4 py-3 rounded-md ${selectedValue === item.value ? 'bg-[#6366F1]' : ''}`}
                    // Pass the item's value (slug) to handleSelect
                    onPress={() => handleSelect(item.value)}
                  >
                    {/* Display the item's label */}
                    <Text className="text-[#F8F9FB] text-center font-medium">{item.label}</Text>
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