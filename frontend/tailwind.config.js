import { transform } from 'typescript';

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      spacing:{
        'less-nav': 'calc(100vh - 4rem)',
      },
      
      animation: {
        'spin-slow': 'spin 3s linear infinite',
        'appear-bot': 'fadeIn 500ms cubic-bezier(0.4, 0, 0.2, 1)',
      },
      keyframes: {
        fadeIn: {
          '0%': { 
            opacity: '0',
            transform: 'translateY(20px)',
           },
          '100%': { 
            opacity: '1',
            transform: 'translateY(0)',
           },
        },
      },
    },
  },
  plugins: [
    require('daisyui'),
  ],
  daisyui: {
    themes: ["light", "dark", 
      {
        corporate: {
          ...require("daisyui/src/theming/themes")["corporate"],
          primary: "blue",
          "primary-focus": "mediumblue",
        },
      }
    ],
  },
}

