const defaultTheme = require('tailwindcss/defaultTheme')
const colors = require('tailwindcss/colors')

// make some colored shadows cause gradients
const makeShadow = (name, rgb) => {
  const obj = {}

  obj[name + '-xs'] = `0 0 0 1px rgba(${rgb}, 0.05)`
  obj[name + '-xs'] = `0 0 0 1px rgba(${rgb}, 0.05)`
  obj[name + '-sm'] = `0 1px 2px 0 rgba(${rgb}, 0.05)`
  obj[name] = `0 1px 3px 0 rgba(${rgb}, 0.1), 0 1px 2px 0 rgba(${rgb}, 0.06)`
  obj[
    name + '-md'
  ] = `0 4px 6px -1px rgba(${rgb}, 0.1), 0 2px 4px -1px rgba(${rgb}, 0.06)`
  obj[
    name + '-lg'
  ] = `0 10px 15px -3px rgba(${rgb}, 0.1), 0 4px 6px -2px rgba(${rgb}, 0.05)`
  obj[
    name + '-xl'
  ] = `0 20px 25px -5px rgba(${rgb}, 0.1), 0 10px 10px -5px rgba(${rgb}, 0.04)`
  obj[name + '-2xl'] = `0 25px 50px -12px rgba(${rgb}, 0.25)`
  obj[name + '-inner'] = `inset 0 2px 4px 0 rgba(${rgb}, 0.06)`
  return obj
}

module.exports = {
  darkMode: 'class',
  content: ['./**/*.{js,jsx,ts,tsx}'],
  content: [
    './pages/**/*.{js,ts,jsx,tsx}',
    './pages/*.{js,ts,jsx,tsx}',

    './components/**/*.{js,ts,jsx,tsx}',
    './styles/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    screens: {
      xs: '475px',
      mdl: '928px',
      ...defaultTheme.screens,
    },
    extend: {
      screens: {
        '3xl': '1600px',
      },
      minWidth: {
        0: '0',
        '1/5': '20%',
        '1/4': '25%',
        '1/3': '33.33%',
        '2/5': '40%',
        '1/2': '50%',
        '3/5': '60%',
        '2/3': '66.66%',
        '3/4': '75%',
        '4/5': '80%',
        full: '100%',
      },
      borderRadius: {
        xl: '3rem',
        lg: '1.0rem',
        md: '0.50rem',
        sm: '0.25rem',
      },
      spacing: {
        xl: '3rem',
        lg: '1.0rem',
        md: '0.5rem',
        sm: '0.25rem',
      },
      borderWidth: {
        3: '3px',
      },
      backgroundColor: {
        lightish: {
          0: 'rgba(255, 255, 255, 0)',
          10: 'rgba(255, 255, 255, 0.1)',
          20: 'rgba(255, 255, 255, 0.2)',
          30: 'rgba(255, 255, 255, 0.3)',
          40: 'rgba(255, 255, 255, 0.4)',
          50: 'rgba(255, 255, 255, 0.5)',
          60: 'rgba(255, 255, 255, 0.6)',
          70: 'rgba(255, 255, 255, 0.7)',
          80: 'rgba(255, 255, 255, 0.8)',
          90: 'rgba(255, 255, 255, 0.9)',
          100: 'rgba(255, 255, 255, 1.0)',
        },
        darkish: {
          0: 'rgba(0, 0, 0, 0)',
          10: 'rgba(0, 0, 0, 0.1)',
          20: 'rgba(0, 0, 0, 0.2)',
          30: 'rgba(0, 0, 0, 0.3)',
          40: 'rgba(0, 0, 0, 0.4)',
          50: 'rgba(0, 0, 0, 0.5)',
          60: 'rgba(0, 0, 0, 0.6)',
          70: 'rgba(0, 0, 0, 0.7)',
          80: 'rgba(0, 0, 0, 0.8)',
          90: 'rgba(0, 0, 0, 0.9)',
          100: 'rgba(0, 0, 0, 1.0)',
        },
      },
      colors: {
        // updated colors
        strong: '#FCFCFD',
        primary: '#EEEDEF',
        secondary: '#C0BCC2',
        separator: '#565058',
        surface: '#343036',
        tint: '#252226',
        background: '#151315',
        synapsePurple: '#D747FF',
        blueText: '#99E6FF',
        greenText: '#66e595',
        greenProgress: 'hsl(105deg 100% 60%)',
        // previous colors
        bgBase: `${colors.slate[400]}`,
        bgDarker: '#111111',
        bgLight: '#353038',
        bgLighter: '#443F47',
        bgLightest: '#58535B',
        primaryTextColor: '#EEEDEF',
        secondaryTextColor: '#cccad3',
        primaryHover: '#ff00ff',
        default: '#0e103c',
        light: '#c7d4ed',
        dark: '#41526A',
        sky: {
          50: '#f0f9ff',
          100: '#e0f2fe',
          200: '#bae6fd',
          300: '#7dd3fc',
          400: '#38bdf8',
          500: '#0ea5e9',
          600: '#0284c7',
          700: '#0369a1',
          800: '#075985',
          900: '#0c4a6e',
        },
        lightish: {
          0: 'rgba(255, 255, 255, 0)',
          10: 'rgba(255, 255, 255, 0.1)',
          20: 'rgba(255, 255, 255, 0.2)',
          30: 'rgba(255, 255, 255, 0.3)',
          40: 'rgba(255, 255, 255, 0.4)',
          50: 'rgba(255, 255, 255, 0.5)',
          60: 'rgba(255, 255, 255, 0.6)',
          70: 'rgba(255, 255, 255, 0.7)',
          80: 'rgba(255, 255, 255, 0.8)',
          90: 'rgba(255, 255, 255, 0.9)',
          100: 'rgba(255, 255, 255, 1.0)',
        },
        darkish: {
          0: 'rgba(0, 0, 0, 0)',
          10: 'rgba(0, 0, 0, 0.1)',
          20: 'rgba(0, 0, 0, 0.2)',
          30: 'rgba(0, 0, 0, 0.3)',
          40: 'rgba(0, 0, 0, 0.4)',
          50: 'rgba(0, 0, 0, 0.5)',
          60: 'rgba(0, 0, 0, 0.6)',
          70: 'rgba(0, 0, 0, 0.7)',
          80: 'rgba(0, 0, 0, 0.8)',
          90: 'rgba(0, 0, 0, 0.9)',
          100: 'rgba(0, 0, 0, 1.0)',
        },
        ...colors,
      },
      fontFamily: {
        sans: ['Matter', ...defaultTheme.fontFamily.sans],
        mono: ['Monospace', ...defaultTheme.fontFamily.mono],
      },
      boxShadow: {
        'custom-shadow': 'inset 0 3px 3px 0 rgba(0, 0, 0, 0.25)', // replace 'custom-shadow' with a more appropriate name
        ...makeShadow('cool-gray', '71, 85, 104'),
        ...makeShadow('gray', '75, 85, 98'),
        ...makeShadow('red', '223, 39, 44'),
        ...makeShadow('orange', '207, 57, 24'),
        ...makeShadow('yellow', '158, 88, 28'),
        ...makeShadow('green', '16, 122, 87'),
        ...makeShadow('teal', '13, 116, 128'),
        ...makeShadow('sky', '56, 189, 248'),
        ...makeShadow('blue', '29, 100, 236'),
        ...makeShadow('indigo', '87, 81, 230'),
        ...makeShadow('purple', '125, 59, 236'),
        ...makeShadow('pink', '213, 34, 105'),
      },
      linearBorderGradients: ({ theme }) => ({
        colors: {
          'magenta-melrose': ['#ff00ff', '#ac8fff'],
        },
        background: {
          bgDarker: '#111111',
          bgLight: '#353038',
          ...colors,
        },
      }),
      keyframes: {
        movement: {
          '0%, 100%': { transform: 'translate3d(0, -25px, 0) scale(1)' },
          '50%': { transform: 'translate3d(0, 25px, 0) scale(1.5)' },
        },
      },
      animation: {
        'movement': 'movement 2s ease-in-out infinite',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    require('@tailwindcss/aspect-ratio'),
    require('tailwind-scrollbar-hide'),
    require('tailwindcss-border-gradient-radius'),
  ],
  safelist: [...Array(300).keys()].map((i) => `h-[${i + 1}px]`),
}
