import React, { useState, useEffect } from 'react';
import { Platform, View, Text, TouchableOpacity, ScrollView, ActivityIndicator } from 'react-native';
import { useLocalSearchParams, router } from 'expo-router';
// Removed useFonts import as custom fonts are not explicitly applied via style prop anymore
import { Ionicons } from '@expo/vector-icons';
import { WebView } from 'react-native-webview';
import CodeHighlighter from '../../components/CodeHighlighter';
import ScreenHeader from '../../components/ui/ScreenHeader'; // Import the new header
import LanguageTabs from '../../components/ui/LanguageTabs'; // Import LanguageTabs

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
  
  // Removed fontsLoaded state and useFonts hook

  const languageMap: Record<string, string> = {
    "C++": "cpp",
    "JavaScript": "javascript",
    "Python": "python",
    "Java": "java",
    "C#": "csharp",
  };
  
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
  
  // Removed !fontsLoaded check

  return (
    <View className="flex-1 bg-[#131C24]">
      {/* Use the reusable ScreenHeader component */}
      <ScreenHeader />
      
      {loading ? (
        <View className="flex-1 justify-center items-center">
          <ActivityIndicator size="large" color="#6366F1" />
        </View>
      ) : problem ? (
        <ScrollView className="flex-1" contentContainerStyle={{ paddingBottom: 24, paddingTop: 16, paddingHorizontal: 16 }}>
          {/* Problem Title */}
          <Text
            className="text-[#F8F9FB] text-2xl font-semibold mb-4" // Adjusted size/weight/margin
            // style={{ fontFamily: 'Roboto_700Bold' }} // Remove font family if using default nativewind/system fonts
          >
            {problem.title}
          </Text>

          {/* Problem Description */}
          <View className="bg-[#1E2A3A] rounded-lg p-4 mb-6"> {/* Added background, padding, rounding */}
            {Platform.OS === 'web' ? (
              <div
                style={{ color: '#D1D5DB', fontFamily: '-apple-system, sans-serif', fontSize: '15px', lineHeight: 1.6 }} // Consistent styles
                dangerouslySetInnerHTML={{
                  __html: `
                    <style>
                      body { margin: 0; padding: 0; }
                      p { margin-bottom: 1em; }
                      code { font-family: monospace; background-color: #29374C; padding: 0.2em 0.4em; border-radius: 4px; font-size: 0.9em; }
                      pre { background-color: #131C24; padding: 1em; border-radius: 6px; overflow-x: auto; font-family: monospace; font-size: 0.9em; }
                      strong { font-weight: 600; }
                      ul, ol { padding-left: 1.5em; margin-bottom: 1em; }
                      li { margin-bottom: 0.5em; }
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
                  <html><head><meta name="viewport" content="width=device-width, initial-scale=1.0">
                  <style>
                    body { font-family: -apple-system, sans-serif; margin: 0; padding: 0; color: #D1D5DB; background-color: #1E2A3A; font-size: 15px; line-height: 1.6; word-wrap: break-word; overflow-wrap: break-word; }
                    p { margin-bottom: 1em; }
                    code { font-family: monospace; background-color: #29374C; padding: 0.2em 0.4em; border-radius: 4px; font-size: 0.9em; }
                    pre { background-color: #131C24; padding: 1em; border-radius: 6px; overflow-x: auto; font-family: monospace; font-size: 0.9em; }
                    strong { font-weight: 600; }
                    ul, ol { padding-left: 1.5em; margin-bottom: 1em; }
                    li { margin-bottom: 0.5em; }
                  </style></head>
                  <body>${problem.content || ''}</body></html>`
                }}
                style={{
                  height: webViewHeight, // Keep dynamic height
                  backgroundColor: '#1E2A3A', // Match container background
                  opacity: 0.99 // Minor hack sometimes needed for WebView sizing
                }}
                onMessage={(event) => {
                  // Consider adding a maximum height check if needed
                  setWebViewHeight(Number(event.nativeEvent.data));
                }}
                injectedJavaScript={`
                  const height = document.body.scrollHeight;
                  window.ReactNativeWebView.postMessage(height);
                  true; // note: this is required, or you'll sometimes get silent failures
                `}
                androidLayerType="hardware" // Potential performance improvement on Android
                scalesPageToFit={false}
              />
            )}
          </View>
          
          {/* Solution Section Title */}
          <Text
            className="text-[#F8F9FB] text-lg font-semibold mb-3" // Adjusted size/weight/margin
            // style={{ fontFamily: 'Roboto_700Bold' }} // Remove font family
          >
            Solution
          </Text>

          {/* Use LanguageTabs component */}
          <LanguageTabs
            availableLanguages={Object.keys(solutions)}
            selectedLanguage={selectedLanguage}
            onSelectLanguage={setSelectedLanguage}
            // Pass loading state if applicable, otherwise omit or pass false
            // loading={solutionsLoading} // Example if loading state exists here
          />

          {/* Solution Code Block */}
          {solutions[selectedLanguage] ? (
            <CodeHighlighter
              language={languageMap[selectedLanguage] || selectedLanguage.toLowerCase()}
              style={{ marginBottom: 24 }} // Adjusted margin
            >
              {solutions[selectedLanguage]}
            </CodeHighlighter>
          ) : (
            // Container for "Not Available" message for consistent spacing
            <View className="bg-[#1E2A3A] rounded-lg p-4 mb-6">
              <Text className="text-[#ADBAC7] text-sm"> {/* Adjusted style */}
                Solution code not available for {languageMap[selectedLanguage] || selectedLanguage}.
              </Text>
            </View>
          )}
          
          {/* Removed explicit bottom padding view, handled by ScrollView contentContainerStyle */}
        </ScrollView>
      ) : (
        // Styled "Not Found" state
        <View className="flex-1 justify-center items-center p-6 bg-[#131C24]">
          <Ionicons name="alert-circle-outline" size={48} color="#8A9DC0" />
          <Text className="text-[#ADBAC7] text-lg mt-4 text-center">
            Problem Not Found
          </Text>
          <Text className="text-[#8A9DC0] text-sm mt-1 text-center">
            The requested problem could not be loaded.
          </Text>
        </View>
      )}
    </View>
  );
}