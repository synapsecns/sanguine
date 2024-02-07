/** @type {import('tailwindcss').Config} */

export default {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  plugins: [],
  corePlugins: {
    preflight: false,
  },
  theme: {
    extend: {
      keyframes: {
        'slide-down': {
          '0%': { transform: 'translateY(-4px) scale(1, .95)', opacity: 0.67 },
          '100%': { transform: 'translateY(0) scale(1, 1)', opacity: 1 },
        },
      },
      animation: {
        'slide-down': 'slide-down .3s cubic-bezier(0, .5, .2, 1)',
      },
    },
  },
}
