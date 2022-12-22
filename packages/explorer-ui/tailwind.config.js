module.exports = {
  darkMode: 'class',
  content: ['./**/*.{js,jsx,ts,tsx}'],
  content: [
    './pages/**/*.{js,ts,jsx,tsx}',
    './components/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      backgroundImage: {
        'wavylines-stretch': "url('wavylinesstretch.svg')",
      },
    },
  },
  plugins: [],
  safelist: [...Array(300).keys()].map((i) => `h-[${i + 1}px]`),
}
