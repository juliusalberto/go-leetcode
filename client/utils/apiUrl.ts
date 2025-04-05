import { Platform } from 'react-native';

// Ensure environment variable is read correctly. Provide a default for safety, though it should be set.
const BASE_API_URL = process.env.EXPO_PUBLIC_API_URL || 'http://localhost:8080';

/**
 * Constructs the full API URL by combining the base URL (from environment variables)
 * with the provided relative path, adjusting for Android emulators.
 *
 * @param path The relative path for the API endpoint (e.g., '/api/decks')
 * @returns The full, platform-adjusted API URL
 */
export function getApiUrl(path: string): string {
  let baseUrl = BASE_API_URL;

  // Adjust for Android emulator if the base URL is localhost
  if (Platform.OS === 'android' && baseUrl.includes('localhost')) {
    baseUrl = baseUrl.replace('localhost', '10.0.2.2');
  } else if (Platform.OS === 'android' && baseUrl.includes('127.0.0.1')) {
    // Also handle 127.0.0.1 for Android
    baseUrl = baseUrl.replace('127.0.0.1', '10.0.2.2');
  }

  // Ensure path starts with a slash if it doesn't already
  const formattedPath = path.startsWith('/') ? path : `/${path}`;

  // Remove trailing slash from baseUrl if present, before concatenating
  const cleanBaseUrl = baseUrl.endsWith('/') ? baseUrl.slice(0, -1) : baseUrl;

  return `${cleanBaseUrl}${formattedPath}`;
}

// getBaseApiUrl function is no longer needed as getApiUrl now handles base URL internally.