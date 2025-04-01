import React, { useState, useEffect } from 'react';
import { Platform, View, Text, TouchableOpacity, ScrollView, ActivityIndicator } from 'react-native';
import { useLocalSearchParams, router } from 'expo-router';
import { useFonts, Roboto_400Regular, Roboto_500Medium, Roboto_700Bold } from '@expo-google-fonts/roboto';
import { Ionicons } from '@expo/vector-icons';
import { WebView } from 'react-native-webview';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomOneDark } from 'react-syntax-highlighter/dist/esm/styles/hljs';
import RNHighlighter from 'react-native-syntax-highlighter';

// Import API hooks
import { useProblemsApi } from '../../services/api/problems';
import { useSolutionsApi } from '../../services/api/solutions';

export default function ProblemDetailScreen() {
  const { slug } = useLocalSearchParams();
  const [problem, setProblem] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [solutions, setSolutions] = useState<Record<string, string>>({});
  const [selectedLanguage, setSelectedLanguage] = useState('python');
  const [webViewHeight, setWebViewHeight] = useState(1);
  
  // Initialize API hooks
  const problemsApi = useProblemsApi();
  const solutionsApi = useSolutionsApi();
  
  const [fontsLoaded] = useFonts({
    Roboto_400Regular,
    Roboto_500Medium,
    Roboto_700Bold,
  });
  
  useEffect(() => {
    const loadProblem = async () => {
      try {
        if (typeof slug === 'string') {
          const problemData = await problemsApi.fetchProblemBySlug(slug);
          
          // Set problem data
          setProblem(problemData);
          
          // Fetch solutions for this problem
          if (problemData.id) {
            const solutionsData = await solutionsApi.fetchSolutionByID(problemData.id.toString());
            
            // Solutions data is already in the format of { language: code }
            setSolutions(solutionsData);
            
            // Set selected language to the first available one if current selection isn't available
            const availableLanguages = Object.keys(solutionsData);
            if (availableLanguages.length > 0 && !solutionsData[selectedLanguage]) {
              setSelectedLanguage(availableLanguages[0]);
            }
          }
        }
      } catch (error) {
        console.error('Error fetching problem details:', error);
      } finally {
        setLoading(false);
      }
    };
    
    loadProblem();
  }, []);
  
  
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
                  height: webViewHeight,
                  backgroundColor: '#131C24' 
                }}
                onMessage={(event) => {
                  setWebViewHeight(Number(event.nativeEvent.data));
                }}
                injectedJavaScript={`
                  setTimeout(function() {
                    window.ReactNativeWebView.postMessage(document.body.scrollHeight);
                  }, 500);
                  true;
                `}
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
            {Object.keys(solutions).length > 0 ? (
              Object.keys(solutions).map(lang => (
                <TouchableOpacity
                  key={lang}
                  onPress={() => setSelectedLanguage(lang)}
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
              ))
            ) : (
              // Default language tabs if no solutions available yet
              ['python', 'cpp'].map(lang => (
                <TouchableOpacity
                  key={lang}
                  onPress={() => setSelectedLanguage(lang)}
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
              ))
            )}
          </View>
          
          {/* Solution Code Block */}
          {solutions[selectedLanguage] ? (
            <View className="bg-[#1E2A3A] rounded-lg p-4 mb-8">
              {Platform.OS === 'web' ? (
                <SyntaxHighlighter
                  language={selectedLanguage}
                  style={atomOneDark}
                  customStyle={{
                    backgroundColor: '#1E2A3A',
                    padding: 0,
                    margin: 0,
                    fontSize: 14,
                    lineHeight: 24
                  }}
                >
                  {solutions[selectedLanguage]}
                </SyntaxHighlighter>
              ) : (
                <RNHighlighter
                  language={selectedLanguage}
                  fontSize={14}
                  style={atomOneDark}
                  customStyle={{ backgroundColor: '#1E2A3A' }}
                  lineHeight={24}
                >
                  {solutions[selectedLanguage]}
                </RNHighlighter>
              )}
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