/** @type {import('tailwindcss').Config} */

function withOpacity(variableName) {
  return ({ opacityValue }) => {
    if (opacityValue !== undefined)
      return `rgba(var(${variableName}), ${opacityValue})`
    else return `rgb(var(${variableName}))`
  }
}

export default {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {
      textColor: {
        widget: {
          primary: withOpacity('--synapse-widget-primary-color'),
          accent: withOpacity('--synapse-widget-accent-color'),
        },
      },
      backgroundColor: {
        widget: {
          primary: withOpacity('--synapse-widget-background-color'),
          surface: withOpacity('--synapse-widget-surface-color'),
        },
      },
      borderColor: {
        widget: {
          background: withOpacity('--synapse-widget-background-color'),
          separator: withOpacity('--synapse-widget-separator-color'),
        },
      },
    },
  },
  plugins: [],
}
