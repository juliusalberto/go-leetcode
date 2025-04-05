import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';

interface RatingButtonProps {
  onRate: (rating: 1 | 2 | 3 | 4) => void;
}

const RatingButtons: React.FC<RatingButtonProps> = ({ onRate }) => {
  const ratings = [
    { label: 'Again', value: 1, textColor: 'text-red-500', borderColor: 'border-red-500', hoverBg: 'hover:bg-red-500/10', activeBg: 'active:bg-red-500/20' },
    { label: 'Hard', value: 2, textColor: 'text-orange-500', borderColor: 'border-orange-500', hoverBg: 'hover:bg-orange-500/10', activeBg: 'active:bg-orange-500/20' },
    { label: 'Good', value: 3, textColor: 'text-blue-500', borderColor: 'border-blue-500', hoverBg: 'hover:bg-blue-500/10', activeBg: 'active:bg-blue-500/20' },
    { label: 'Easy', value: 4, textColor: 'text-green-500', borderColor: 'border-green-500', hoverBg: 'hover:bg-green-500/10', activeBg: 'active:bg-green-500/20' },
  ] as const; // Use 'as const' for stricter typing on value

  return (
    <View className="mb-6 px-4"> {/* Added horizontal padding */}
      <Text className="text-[#ADBAC7] text-base font-normal mb-4 text-center"> {/* Adjusted text style */}
        How well did you recall this?
      </Text>
      <View className="flex-row justify-between space-x-2"> {/* Use space-x for spacing */}
        {ratings.map((rating) => (
          <TouchableOpacity
            key={rating.value}
            // Sleeker style: border, colored text, subtle background on interaction
            className={`
              border ${rating.borderColor} 
              rounded-lg 
              py-3 px-2 {/* Adjusted padding */}
              flex-1 
              items-center justify-center 
              transition-colors duration-150 ease-in-out 
              ${rating.hoverBg} ${rating.activeBg}
            `}
            onPress={() => onRate(rating.value)}
            activeOpacity={0.8} // Control opacity on press
          >
            <Text className={`${rating.textColor} text-center font-medium text-sm`}> {/* Adjusted text size */}
              {rating.label}
            </Text>
          </TouchableOpacity>
        ))}
      </View>
    </View>
  );
};

export default RatingButtons;