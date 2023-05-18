module.exports = {
  darkMode: 'class',
  content: ['./**/*.{js,jsx,ts,tsx}'],
  // @ts-expect-error TS(1117): An object literal cannot have multiple properties ... Remove this comment to see the full error message
  content: [
    './pages/**/*.{js,ts,jsx,tsx}',
    './components/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      backgroundImage: {
        'wavylines-stretch': "url('wavylinesstretch.svg')",
        'synapse-radial':
          'radial-gradient(249.66% 351.84% at -119.61% 97.91%, rgba(255, 0, 255, 0.33) 0%, rgba(172, 143, 255, 0.33) 100%)',
        'synapse-logo': "url('../assets/icons/synapselogo.svg')",
      },
      colors: {
        synapse: '#100C13',
      },
    },
  },
  plugins: [],
  // @ts-expect-error TS(2802): Type 'IterableIterator<number>' can only be iterat... Remove this comment to see the full error message
  safelist: [...Array(300).keys()].map((i) => `h-[${i + 1}px]`),
}
