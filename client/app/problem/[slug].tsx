import React, { useState, useEffect } from 'react';
import { Platform, View, Text, TouchableOpacity, ScrollView, ActivityIndicator } from 'react-native';
import { useLocalSearchParams, router } from 'expo-router';
import { useFonts, Roboto_400Regular, Roboto_500Medium, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { Ionicons } from '@expo/vector-icons';
import { WebView } from 'react-native-webview';

// Import API services
import { fetchProblemBySlug} from '../services/api/problems';

export default function ProblemDetailScreen() {
  const { slug } = useLocalSearchParams();
  const [problem, setProblem] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [selectedLanguage, setSelectedLanguage] = useState('python');
  
  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_500Medium,
    Roboto_700Bold,
  });
  
  useEffect(() => {
    const loadProblem = async () => {
      try {
        if (typeof slug === 'string') {
          const problemData = await fetchProblemBySlug(slug);
          
          // Extract example inputs and outputs from the HTML content
          if (!problemData.content) {
            problemData.content = "";
          }
          const examples = extractExamples(problemData.content);
          
          // Combine the data
          setProblem({
            ...problemData,
            examples,
            solutions: {} // Will be populated when a language is selected
          });
        }
      } catch (error) {
        console.error('Error fetching problem details:', error);
      } finally {
        setLoading(false);
      }
    };
    
    loadProblem();
  }, [slug]);
  
  // Function to extract examples from HTML content
  const extractExamples = (htmlContent: string): {input: string, output: string}[] => {
    // This is a simple regex-based extraction - a proper HTML parser would be better
    // but this works for our basic example format
    const examples: {input: string, output: string}[] = [];
    
    // Extract blocks between <pre> tags
    const preBlockRegex = /<pre>([\s\S]*?)<\/pre>/g;
    let match;
    
    while ((match = preBlockRegex.exec(htmlContent)) !== null) {
      const block = match[1];
      
      // Extract input and output from the block
      const inputMatch = block.match(/<strong>Input:<\/strong>([\s\S]*?)<strong>Output:<\/strong>/);
      const outputMatch = block.match(/<strong>Output:<\/strong>([\s\S]*?)(?:<strong>|$)/);
      
      if (inputMatch && outputMatch) {
        examples.push({
          input: inputMatch[1].trim(),
          output: outputMatch[1].trim().split('<strong>Explanation:')[0].trim()
        });
      }
    }
    
    return examples;
  };
  
  const handleBack = () => {
    router.back();
  };
  
  if (!fontsLoaded) {
    return null;
  }
  
  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Header with back button */}
      <View className="flex-row items-center p-4 pb-2">
        <TouchableOpacity
          className="p-2"
          onPress={handleBack}
        >
          <Ionicons name="chevron-back" size={24} color="#F8F9FB" />
        </TouchableOpacity>
      </View>
      
      {loading ? (
        <View className="flex-1 justify-center items-center">
          <ActivityIndicator size="large" color="#6366F1" />
        </View>
      ) : problem ? (
        <ScrollView className="flex-1 px-4">
          {/* Problem Title */}
          <Text 
            className="text-[#F8F9FB] text-3xl font-bold mb-6"
            style={{ fontFamily: 'Roboto_700Bold' }}
          >
            {problem.title}
          </Text>
          
        {/* Problem Description */}
        <View className="mb-6">
            {Platform.OS === 'web' ? (
                <div 
                    style={{ 
                        color: '#F8F9FB',
                        backgroundColor: '#131C24',
                        fontSize: '16px',
                        lineHeight: 1.5,
                        fontFamily: 'Roboto, sans-serif',
                    }}
                    dangerouslySetInnerHTML={{ 
                        __html: `
                        <style>
                            p { margin-bottom: 16px; }
                            code {
                            font-family: monospace;
                            background-color: #29374C;
                            padding: 2px 4px;
                            border-radius: 4px;
                            }
                            pre {
                            background-color: #1E2A3A;
                            padding: 16px;
                            border-radius: 8px;
                            overflow-x: auto;
                            font-family: monospace;
                            }
                            strong { font-weight: bold; }
                        </style>
                        ${problem.content || ''}
                        `
                    }} 
                />
            ) : (
                <WebView
                originWhitelist={['*']}
                source={{ 
                    html: `
                    <html>
                    <head>
                        <meta name="viewport" content="width=device-width, initial-scale=1.0">
                        <style>
                        body {
                            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
                            padding: 0;
                            margin: 0;
                            color: #F8F9FB;
                            background-color: #131C24;
                            font-size: 16px;
                            line-height: 1.5;
                        }
                        p {
                            margin-bottom: 16px;
                        }
                        code {
                            font-family: monospace;
                            background-color: #29374C;
                            padding: 2px 4px;
                            border-radius: 4px;
                        }
                        pre {
                            background-color: #1E2A3A;
                            padding: 16px;
                            border-radius: 8px;
                            overflow-x: auto;
                            font-family: monospace;
                        }
                        strong {
                            font-weight: bold;
                        }
                        </style>
                    </head>
                    <body>
                        ${problem.content || ''}
                    </body>
                    </html>
                    `
                }}
                style={{ 
                    height: 300,
                    backgroundColor: '#131C24' 
                }}
                />
            )}
        </View>
          
          {/* Solution Section */}
          <Text 
            className="text-[#F8F9FB] text-xl font-bold mb-4"
            style={{ fontFamily: 'Roboto_700Bold' }}
          >
            Solution:
          </Text>
          
          {/* Language Selection Tabs */}
          <View className="flex-row flex-wrap mb-4">
            {['python', 'cpp'].map(lang => (
              <TouchableOpacity
                key={lang}
                className={`mr-2 mb-2 px-4 py-2 rounded-full ${
                  selectedLanguage === lang ? 'bg-[#29374C]' : 'bg-[#1E2A3A]'
                }`}
              >
                <Text 
                  className="text-[#F8F9FB] text-base capitalize"
                  style={{ fontFamily: 'Roboto_500Medium' }}
                >
                  {lang === 'cpp' ? 'C++' : lang.charAt(0).toUpperCase() + lang.slice(1)}
                </Text>
              </TouchableOpacity>
            ))}
          </View>
          
          {/* Solution Code Block */}
          {problem.solutions && problem.solutions[selectedLanguage] ? (
            <View className="bg-[#1E2A3A] rounded-lg p-4 mb-8">
              <Text 
                className="text-[#F8F9FB] font-mono text-sm"
                style={{ lineHeight: 24 }}
              >
                {problem.solutions[selectedLanguage]}
              </Text>
            </View>
          ) : (
            <Text 
              className="text-[#8A9DC0] text-base mb-8"
              style={{ fontFamily: 'Roboto_400Regular' }}
            >
              Solution not available for this language
            </Text>
          )}
          
          {/* Bottom padding to ensure content is visible */}
          <View className="h-10" />
        </ScrollView>
      ) : (
        <View className="flex-1 justify-center items-center p-4">
          <Text 
            className="text-[#F8F9FB] text-base text-center"
            style={{ fontFamily: 'Roboto_500Medium' }}
          >
            Problem not found
          </Text>
        </View>
      )}
    </View>
  );
}