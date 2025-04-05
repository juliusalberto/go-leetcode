import React from 'react';
import { TouchableOpacity, Text, StyleSheet, StyleProp, ViewStyle, TextStyle, ActivityIndicator } from 'react-native';

interface ButtonProps {
  title: string;
  onPress: () => void;
  variant?: 'primary' | 'secondary' | 'ghost'; // Example variants
  disabled?: boolean;
  isLoading?: boolean;
  style?: StyleProp<ViewStyle>;
  textStyle?: StyleProp<TextStyle>;
  className?: string; // Allow passing Tailwind classes
}

const Button: React.FC<ButtonProps> = ({
  title,
  onPress,
  variant = 'primary', // Default variant
  disabled = false,
  isLoading = false,
  style,
  textStyle,
  className,
}) => {
  const getVariantStyles = () => {
    switch (variant) {
      case 'secondary':
        return {
          button: 'bg-[#29374C]', // Example secondary background
          text: 'text-[#F8F9FB]',   // Example secondary text
        };
      case 'ghost':
         return {
          button: 'bg-transparent',
          text: 'text-[#6366F1]', // Example ghost text (link-like)
        };
      case 'primary':
      default:
        return {
          button: 'bg-[#6366F1]', // Default primary background
          text: 'text-white',     // Default primary text
        };
    }
  };

  const variantStyles = getVariantStyles();
  const opacity = disabled || isLoading ? 'opacity-50' : '';

  // Combine base, variant, passed className, and disabled opacity
  const buttonClasses = `py-3 px-6 rounded-lg items-center justify-center ${variantStyles.button} ${opacity} ${className || ''}`;
  const textClasses = `font-medium text-center ${variantStyles.text}`;

  return (
    <TouchableOpacity
      onPress={onPress}
      disabled={disabled || isLoading}
      className={buttonClasses}
      style={style}
    >
      {isLoading ? (
        <ActivityIndicator size="small" color={variant === 'primary' ? '#FFFFFF' : '#6366F1'} />
      ) : (
        <Text className={textClasses} style={textStyle}>
          {title}
        </Text>
      )}
    </TouchableOpacity>
  );
};

export default Button;