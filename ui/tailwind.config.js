/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'twitter-blue': '#1DA1F2',
        'fresh-green': '#23d5ab',
        'soft-pink': '#e73c7e',
      },
      fontFamily: {
        sans: ['Inter', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'Roboto', 'sans-serif'],
      },
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
    require('daisyui')
  ],
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/theming/themes")["light"],
          "primary": "#1DA1F2",
          "secondary": "#23d5ab",
          "accent": "#e73c7e",
          "base-100": "#ffffff",
          "base-200": "#f7f9fa",
          "base-300": "#e1e8ed",
        },
      },
    ],
    base: true,
    styled: true,
    utils: true,
  },
}
