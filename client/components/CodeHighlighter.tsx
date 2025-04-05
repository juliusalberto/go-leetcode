import React from 'react';
import { Platform, StyleProp, TextStyle, ViewStyle, View } from 'react-native'; // Added View
import WebSyntaxHighlighter from 'react-syntax-highlighter'; // Use alias for web
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import RNSyntaxHighlighter from 'react-native-syntax-highlighter'; // Use alias for native

interface CodeHighlighterProps {
  language: string;
  children: string;
  style?: StyleProp<ViewStyle>; // Optional container style
  customTextStyle?: StyleProp<TextStyle>; // Optional text style for native
  fontSize?: number; // Optional font size for native
  lineHeight?: number; // Optional line height for native
}

const CodeHighlighter: React.FC<CodeHighlighterProps> = ({
  language,
  children,
  style,
  customTextStyle,
  fontSize = 14, // Default font size
  lineHeight = 24, // Default line height
}) => {
  const highlighterStyle = {
    ...atomOneDark, // Base style
    'hljs': { // Ensure background is transparent or matches container
      ...(atomOneDark.hljs || {}),
      backgroundColor: 'transparent', // Make highlighter background transparent
    },
  };

  const containerStyle: StyleProp<ViewStyle> = [ // Explicitly type containerStyle
    { backgroundColor: '#1E2A3A', borderRadius: 8, padding: 16, overflow: 'hidden' }, // Default container style
    style, // Apply custom container style if provided
  ];

  return (
    <View style={containerStyle}>
      {Platform.OS === 'web' ? (
        <WebSyntaxHighlighter
          language={language}
          style={highlighterStyle}
          customStyle={{ // Web customStyle applies to the <pre> tag
            margin: 0, // Remove default margins
            padding: 0, // Remove default padding, handled by container
            backgroundColor: 'transparent', // Ensure no background override
          }}
          codeTagProps={{ // Style the inner <code> tag if needed
            style: {
              fontFamily: 'monospace', // Consistent font
              fontSize: `${fontSize}px`,
              lineHeight: `${lineHeight}px`,
            }
          }}
        >
          {children}
        </WebSyntaxHighlighter>
      ) : (
        <RNSyntaxHighlighter
          language={language}
          fontSize={fontSize}
          style={highlighterStyle} // Pass the theme style
          customStyle={{ // Native customStyle applies to the Text component
            padding: 0, // Remove default padding
            margin: 0,
            backgroundColor: 'transparent', // Ensure no background override
            ...(customTextStyle as object), // Apply custom text styles
          }}
          highlighter='hljs' // Specify the highlighter engine
          lineHeight={lineHeight} // Apply line height
        >
          {children}
        </RNSyntaxHighlighter>
      )}
    </View>
  );
};

export default CodeHighlighter;