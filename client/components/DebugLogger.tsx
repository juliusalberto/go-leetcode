import React, { useEffect, useState } from 'react';
import { View, Text, ScrollView, StyleSheet, Platform } from 'react-native';

// Global log collection
const logs: string[] = [];
const MAX_LOGS = 50;

// Replace console methods to capture logs
const originalConsoleLog = console.log;
const originalConsoleError = console.error;
const originalConsoleWarn = console.warn;

console.log = (...args) => {
  const log = args.map(arg => 
    typeof arg === 'object' ? JSON.stringify(arg) : String(arg)
  ).join(' ');
  logs.unshift(`[LOG] ${log}`);
  if (logs.length > MAX_LOGS) logs.pop();
  originalConsoleLog(...args);
};

console.error = (...args) => {
  const log = args.map(arg => 
    typeof arg === 'object' ? JSON.stringify(arg) : String(arg)
  ).join(' ');
  logs.unshift(`[ERROR] ${log}`);
  if (logs.length > MAX_LOGS) logs.pop();
  originalConsoleError(...args);
};

console.warn = (...args) => {
  const log = args.map(arg => 
    typeof arg === 'object' ? JSON.stringify(arg) : String(arg)
  ).join(' ');
  logs.unshift(`[WARN] ${log}`);
  if (logs.length > MAX_LOGS) logs.pop();
  originalConsoleWarn(...args);
};

// Add an explicit log for debugging
export const debugLog = (message: string) => {
  console.log(`[DEBUG] ${message}`);
};

interface DebugLoggerProps {
  visible?: boolean;
}

const DebugLogger: React.FC<DebugLoggerProps> = ({ visible = __DEV__ }) => {
  const [logState, setLogState] = useState<string[]>([]);

  useEffect(() => {
    // Update log state every second
    const timer = setInterval(() => {
      setLogState([...logs]);
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  if (!visible) return null;

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Debug Logs (Platform: {Platform.OS})</Text>
      <ScrollView style={styles.scrollView}>
        {logState.map((log, index) => (
          <Text key={index} style={[
            styles.logText,
            log.includes('[ERROR]') && styles.errorText,
            log.includes('[WARN]') && styles.warningText,
            log.includes('[DEBUG]') && styles.debugText
          ]}>
            {log}
          </Text>
        ))}
      </ScrollView>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    position: 'absolute',
    bottom: 0,
    left: 0,
    right: 0,
    backgroundColor: 'rgba(0, 0, 0, 0.8)',
    maxHeight: 300,
    zIndex: 9999,
  },
  title: {
    color: 'white',
    fontWeight: 'bold',
    padding: 5,
    backgroundColor: '#333',
    textAlign: 'center',
  },
  scrollView: {
    padding: 10,
  },
  logText: {
    color: 'white',
    fontSize: 10,
    fontFamily: Platform.OS === 'ios' ? 'Menlo' : 'monospace',
    marginBottom: 2,
  },
  errorText: {
    color: '#ff6666',
  },
  warningText: {
    color: '#ffcc00',
  },
  debugText: {
    color: '#66ff66',
  },
});

export default DebugLogger;