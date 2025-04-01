import { Platform } from 'react-native';

/**
 * Converts localhost URLs to the appropriate address for Android emulators
 * On Android emulators, 'localhost' refers to the emulator itself, not the host machine.
 * To reach the host machine, we need to use 10.0.2.2 (for standard Android emulator)
 * 
 * @param url The original URL that might contain 'localhost'
 * @returns The converted URL with the appropriate host for the current platform
 */
export function getApiUrl(url: string): string {
  if (Platform.OS === 'android') {
    return url.replace('localhost', '10.0.2.2');
  }
  return url;
}

/**
 * Returns the base API URL with platform-specific adjustments
 */
export function getBaseApiUrl(): string {
  return getApiUrl('http://localhost:8080');
}