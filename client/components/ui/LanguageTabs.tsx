import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';

interface LanguageTabsProps {
  availableLanguages: string[];
  selectedLanguage: string;
  onSelectLanguage: (language: string) => void;
  loading?: boolean; // Optional loading state
}

const LanguageTabs: React.FC<LanguageTabsProps> = ({
  availableLanguages,
  selectedLanguage,
  onSelectLanguage,
  loading = false,
}) => {
  // Helper to format language names (e.g., cpp -> C++)
  const formatLanguageName = (lang: string): string => {
    if (lang === 'cpp') return 'C++';
    if (lang === 'csharp') return 'C#';
    return lang.charAt(0).toUpperCase() + lang.slice(1);
  };

  return (
    <View className="flex-row flex-wrap mb-4">
      {availableLanguages.length > 0 ? (
        availableLanguages.map(lang => (
          <TouchableOpacity
            key={lang}
            onPress={() => onSelectLanguage(lang)}
            disabled={loading} // Disable while loading solutions
            className={`
              mr-2 mb-2 px-3 py-1.5 rounded-md border
              ${loading ? 'opacity-50' : ''}
              ${selectedLanguage === lang
                ? 'bg-[#29374C] border-[#32415D]' // Active state
                : 'bg-transparent border-[#32415D]' // Inactive state
              }
            `}
          >
            <Text className={`
              text-sm capitalize
              ${selectedLanguage === lang ? 'text-[#F8F9FB]' : 'text-[#ADBAC7]'}
            `}>
              {formatLanguageName(lang)}
            </Text>
          </TouchableOpacity>
        ))
      ) : loading ? (
        <Text className="text-[#8A9DC0] text-sm italic py-1.5">Loading languages...</Text>
      ) : (
        <Text className="text-[#8A9DC0] text-sm italic py-1.5">No solutions found.</Text>
      )}
    </View>
  );
};

export default LanguageTabs;