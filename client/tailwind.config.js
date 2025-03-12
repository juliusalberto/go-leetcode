/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [    // Include all screens/components in your `app` folder:
    "./app/**/*.{js,jsx,ts,tsx}",

    // Include all components in your `components` folder:
    "./components/**/*.{js,jsx,ts,tsx}",
    
    // Optionally include other folders like hooks, etc.:
    "./hooks/**/*.{js,jsx,ts,tsx}",],
  theme: {
    extend: {
      colors: {
        "dark-bg": "#121212",
        "light-text": "#f5f5f5",
        primary: "#4f46e5",
        accent: "#f59e0b",
      }
    },
  },
  plugins: [],
  presets: [require("nativewind/preset")],
}

