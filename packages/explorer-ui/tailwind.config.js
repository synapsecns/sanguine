module.exports = {
  darkMode: 'class',
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
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
