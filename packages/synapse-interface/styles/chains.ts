/* eslint-disable */
const CustomColors = {
  CUSTOM_YELLOW: '[#ecae0b]',
  CUSTOM_YELLOW_DARK: '[#3c3c44]',
  ETH_BASE: '[#5170ad]',
  ETH_DARK: '[#3f4f8c]',
  ETH_EXTRA_DARK: '[#314367]',
  ETH_LIGHT: '[#78a5ff]',
}

const ColorOptions = {
  GRAY: 'gray',
  YELLOW: 'yellow',
  GREEN: 'green',
  LIME: 'lime',
  SKY: 'sky',
  BLUE: 'blue',
  ORANGE: 'orange',
  PURPLE: 'purple',
  INDIGO: 'indigo',
  CYAN: 'cyan',
  RED: 'red',
  ETH: 'eth',
  TEAL: 'teal',
}

// TODO obviously this can be cleaned up
// 1. define custom yellow/eth with tailwind
// 2. remove all the if/else
// or
// combine all of this in one function with a switch for the different desired classes.

// Revisit if any of these can be deleted.

// export const getNetworkCurrencyColor = (chain: Chain): string => {
//   const { color, nativeCurrency } = chain
//   switch (color) {
//     case ColorOptions.YELLOW:
//       return `text-[#ecae0b] dark:text-[#ecae0b]`
//     case ColorOptions.ETH:
//       return `text-[#5170ad] dark:text-[#78a5ff]`
//     case ColorOptions.GRAY:
//       return `text-gray-500 dark:text-gray-500`
//     case ColorOptions.GREEN:
//       return `text-green-500 dark:text-green-500`
//     case ColorOptions.LIME:
//       return `text-lime-500 dark:text-lime-500`
//     case ColorOptions.SKY:
//       return `text-sky-500 dark:text-sky-500`
//     case ColorOptions.BLUE:
//       return `text-blue-500 dark:text-blue-500`
//     case ColorOptions.ORANGE:
//       return `text-orange-500 dark:text-orange-500`
//     case ColorOptions.PURPLE:
//       return `text-purple-500 dark:text-purple-500`
//     case ColorOptions.INDIGO:
//       return `text-indigo-500 dark:text-indigo-500`
//     case ColorOptions.CYAN:
//       return `text-cyan-500 dark:text-cyan-500`
//     case ColorOptions.RED:
//       return `text-red-500 dark:text-red-500`
//     default:
//       return `text-gray-500 dark:text-gray-500`
//   }
// }

export const getNetworkButtonBgClassName = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `bg-[#ecae0b] hover:bg-[#ecae0b] active:bg-[#ecae0b]`
    case ColorOptions.ETH:
      return `bg-[#5170ad] hover:bg-[#3f4f8c] active:bg-[#314367]`
    case ColorOptions.GRAY:
      return `bg-gray-500 hover:bg-gray-600 active:bg-gray-700`
    case ColorOptions.GREEN:
      return `bg-green-500 hover:bg-green-600 active:bg-green-700`
    case ColorOptions.LIME:
      return `bg-lime-500 hover:bg-lime-600 active:bg-lime-700`
    case ColorOptions.SKY:
      return `bg-sky-500 hover:bg-sky-600 active:bg-sky-700`
    case ColorOptions.BLUE:
      return `bg-blue-500 hover:bg-blue-600 active:bg-blue-700`
    case ColorOptions.ORANGE:
      return `bg-orange-500 hover:bg-orange-600 active:bg-orange-700`
    case ColorOptions.PURPLE:
      return `bg-purple-500 hover:bg-purple-600 active:bg-purple-700`
    case ColorOptions.INDIGO:
      return `bg-indigo-500 hover:bg-indigo-600 active:bg-indigo-700`
    case ColorOptions.CYAN:
      return `bg-cyan-500 hover:bg-cyan-600 active:bg-cyan-700`
    case ColorOptions.RED:
      return `bg-red-500 hover:bg-red-600 active:bg-red-700`
    case ColorOptions.TEAL:
      return `bg-teal-500 hover:bg-teal-600 active:bg-teal-700`
    default:
      return `bg-gray-500 hover:bg-gray-600 active:bg-gray-700`
  }
}

export const getNetworkButtonBgClassNameActive = (
  chainColor: string
): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `dark:active:bg-[#ecae0b] active:bg-[#ecae0b] `
    case ColorOptions.ETH:
      return `dark:active:bg-[#314367] active:bg-[#314367] `
    case ColorOptions.GRAY:
      return `dark:active:bg-gray-500 active:bg-gray-500 `
    case ColorOptions.GREEN:
      return `dark:active:bg-green-500 active:bg-green-500 `
    case ColorOptions.LIME:
      return `dark:active:bg-lime-500 active:bg-lime-500 `
    case ColorOptions.SKY:
      return `dark:active:bg-sky-500 active:bg-sky-500 `
    case ColorOptions.BLUE:
      return `dark:active:bg-blue-500 active:bg-blue-500 `
    case ColorOptions.ORANGE:
      return `dark:active:bg-orange-500 active:bg-orange-500 `
    case ColorOptions.PURPLE:
      return `dark:active:bg-purple-500 active:bg-purple-500 `
    case ColorOptions.INDIGO:
      return `dark:active:bg-indigo-500 active:bg-indigo-500 `
    case ColorOptions.CYAN:
      return `dark:active:bg-cyan-500 active:bg-cyan-500 `
    case ColorOptions.RED:
      return `dark:active:bg-red-500 active:bg-red-500 `
    case ColorOptions.TEAL:
      return `dark:active:bg-teal-500 active:bg-teal-500 `
    default:
      return `dark:active:bg-gray-500 active:bg-gray-500 `
  }
}

export const getNetworkButtonBorderHover = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `hover:border-[#ecae0b]`
    case ColorOptions.ETH:
      return `hover:border-[#5170ad]`
    case ColorOptions.GRAY:
      return `hover:border-gray-300`
    case ColorOptions.GREEN:
      return `hover:border-green-300`
    case ColorOptions.LIME:
      return `hover:border-lime-300`
    case ColorOptions.SKY:
      return `hover:border-sky-300`
    case ColorOptions.BLUE:
      return `hover:border-blue-300`
    case ColorOptions.ORANGE:
      return `hover:border-orange-300`
    case ColorOptions.PURPLE:
      return `hover:border-purple-300`
    case ColorOptions.INDIGO:
      return `hover:border-indigo-300`
    case ColorOptions.CYAN:
      return `hover:border-cyan-300`
    case ColorOptions.RED:
      return `hover:border-red-300`
    case ColorOptions.TEAL:
      return `hover:border-teal-300`
    default:
      return `hover:border-gray-300`
  }
}

export const getNetworkButtonBorderActive = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `active:border-[#ecae0b]`
    case ColorOptions.ETH:
      return `active:border-[#5170ad]`
    case ColorOptions.GRAY:
      return `active:border-gray-500`
    case ColorOptions.GREEN:
      return `active:border-green-500`
    case ColorOptions.LIME:
      return `active:border-lime-500`
    case ColorOptions.SKY:
      return `active:border-sky-500`
    case ColorOptions.BLUE:
      return `active:border-blue-500`
    case ColorOptions.ORANGE:
      return `active:border-orange-500`
    case ColorOptions.PURPLE:
      return `active:border-purple-500`
    case ColorOptions.INDIGO:
      return `active:border-indigo-500`
    case ColorOptions.CYAN:
      return `active:border-cyan-500`
    case ColorOptions.RED:
      return `active:border-red-500`
    case ColorOptions.RED:
      return `active:border-teal-500`
    default:
      return `active:border-gray-500`
  }
}

export const getNetworkButtonBorder = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `border-[#ecae0b] dark:border-[#ecae0b]`
    case ColorOptions.ETH:
      return `border-[#5170ad] dark:border-[#5170ad]`
    case ColorOptions.GRAY:
      return `border-gray-500 dark:border-gray-500`
    case ColorOptions.GREEN:
      return `border-green-500 dark:border-green-500`
    case ColorOptions.LIME:
      return `border-lime-500 dark:border-lime-500`
    case ColorOptions.SKY:
      return `border-sky-500 dark:border-sky-500`
    case ColorOptions.BLUE:
      return `border-blue-500 dark:border-blue-500`
    case ColorOptions.ORANGE:
      return `border-orange-500 dark:border-orange-500`
    case ColorOptions.PURPLE:
      return `border-purple-500 dark:border-purple-500`
    case ColorOptions.INDIGO:
      return `border-indigo-500 dark:border-indigo-500`
    case ColorOptions.CYAN:
      return `border-cyan-500 dark:border-cyan-500`
    case ColorOptions.RED:
      return `border-red-500 dark:border-red-500`
    case ColorOptions.RED:
      return `border-teal-500 dark:border-teal-500`
    default:
      return `border-gray-500 dark:border-gray-500`
  }
}

export const getNetworkButtonBorderImportant = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `border-[#ecae0b] dark:border-[#ecae0b]`
    case ColorOptions.ETH:
      return `border-[#5170ad] dark:border-[#5170ad]`
    case ColorOptions.GRAY:
      return `border-gray-500 dark:border-gray-500`
    case ColorOptions.GREEN:
      return `border-green-500 dark:border-green-500`
    case ColorOptions.LIME:
      return `border-lime-500 dark:border-lime-500`
    case ColorOptions.SKY:
      return `border-sky-500 dark:border-sky-500`
    case ColorOptions.BLUE:
      return `border-blue-500 dark:border-blue-500`
    case ColorOptions.ORANGE:
      return `border-orange-500 dark:border-orange-500`
    case ColorOptions.PURPLE:
      return `border-purple-500 dark:border-purple-500`
    case ColorOptions.INDIGO:
      return `border-indigo-500 dark:border-indigo-500`
    case ColorOptions.CYAN:
      return `border-cyan-500 dark:border-cyan-500`
    case ColorOptions.RED:
      return `border-red-500 dark:border-red-500`
    case ColorOptions.RED:
      return `border-teal-500 dark:border-teal-500`
    default:
      return `border-gray-500 dark:border-gray-500`
  }
}

export const getNetworkTextColor = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `text-[#ecae0b] dark:text-[#ecae0b]`
    case ColorOptions.ETH:
      return `text-[#5170ad] dark:text-[#78a5ff]`
    case ColorOptions.GRAY:
      return `text-gray-500 dark:text-gray-500`
    case ColorOptions.GREEN:
      return `text-green-500 dark:text-green-500`
    case ColorOptions.LIME:
      return `text-lime-500 dark:text-lime-500`
    case ColorOptions.SKY:
      return `text-sky-500 dark:text-sky-500`
    case ColorOptions.BLUE:
      return `text-blue-500 dark:text-blue-500`
    case ColorOptions.ORANGE:
      return `text-orange-500 dark:text-orange-500`
    case ColorOptions.PURPLE:
      return `text-purple-500 dark:text-purple-500`
    case ColorOptions.INDIGO:
      return `text-indigo-500 dark:text-indigo-500`
    case ColorOptions.CYAN:
      return `text-cyan-500 dark:text-cyan-500`
    case ColorOptions.RED:
      return `text-red-500 dark:text-red-500`
    case ColorOptions.TEAL:
      return `text-teal-500 dark:text-teal-500`
    default:
      return `text-gray-500 dark:text-gray-500`
  }
}

export const getNetworkLinkTextColor = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `text-gray-800 hover:text-[#ecae0b] dark:text-[#ecae0b] dark:hover:text-[#ecae0b]`
    case ColorOptions.ETH:
      return `text-[#5170ad] hover:text-[#78a5ff]`
    case ColorOptions.GRAY:
      return `text-gray-500 hover:text-gray-600 dark:hover:text-gray-500`
    case ColorOptions.GREEN:
      return `text-green-500 hover:text-green-600 dark:hover:text-green-500`
    case ColorOptions.LIME:
      return `text-lime-500 hover:text-lime-600 dark:hover:text-lime-500`
    case ColorOptions.SKY:
      return `text-sky-500 hover:text-sky-600 dark:hover:text-sky-500`
    case ColorOptions.BLUE:
      return `text-blue-500 hover:text-blue-600 dark:hover:text-blue-500`
    case ColorOptions.ORANGE:
      return `text-orange-500 hover:text-orange-600 dark:hover:text-orange-500`
    case ColorOptions.PURPLE:
      return `text-purple-500 hover:text-purple-600 dark:hover:text-purple-500`
    case ColorOptions.INDIGO:
      return `text-indigo-500 hover:text-indigo-600 dark:hover:text-indigo-500`
    case ColorOptions.CYAN:
      return `text-cyan-500 hover:text-cyan-600 dark:hover:text-cyan-500`
    case ColorOptions.RED:
      return `text-red-500 hover:text-red-600 dark:hover:text-red-500`
    case ColorOptions.TEAL:
      return `text-teal-500 hover:text-teal-600 dark:hover:text-teal-500`
    default:
      return `text-gray-500 hover:text-gray-600 dark:hover:text-gray-500`
  }
}

export const getMenuItemStyleForChain = (color: string): string => {
  switch (color) {
    case ColorOptions.ETH:
      return `hover:bg-opacity-20 hover:bg-[#3f4f8c] focus:bg-[#3f4f8c] active:bg-[#314367] active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.GRAY:
      return `hover:bg-opacity-20 hover:bg-gray-500 focus:bg-gray-500 active:bg-gray-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.YELLOW:
      return `hover:bg-opacity-20 hover:bg-yellow-500 focus:bg-yellow-500 active:bg-yellow-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.GREEN:
      return `hover:bg-opacity-20 hover:bg-green-500 focus:bg-green-500 active:bg-green-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.LIME:
      return `hover:bg-opacity-20 hover:bg-lime-500 focus:bg-lime-500 active:bg-lime-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.SKY:
      return `hover:bg-opacity-20 hover:bg-sky-500 focus:bg-sky-500 active:bg-sky-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.BLUE:
      return `hover:bg-opacity-20 hover:bg-blue-500 focus:bg-blue-500 active:bg-blue-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.ORANGE:
      return `hover:bg-opacity-20 hover:bg-orange-500 focus:bg-orange-500 active:bg-orange-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.PURPLE:
      return `hover:bg-opacity-20 hover:bg-purple-500 focus:bg-purple-500 active:bg-purple-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.INDIGO:
      return `hover:bg-opacity-20 hover:bg-indigo-500 focus:bg-indigo-500 active:bg-indigo-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.CYAN:
      return `hover:bg-opacity-20 hover:bg-cyan-500 focus:bg-cyan-500 active:bg-cyan-500 active:bg-opacity-20 focus:bg-opacity-20`
    case ColorOptions.RED:
      return `hover:bg-opacity-20 hover:bg-red-500 focus:bg-red-500 active:bg-red-500 active:bg-opacity-20 focus:bg-opacity-20`
    default:
      return ''
  }
}

export const getNetworkTextColorContrast = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `text-[#ecae0b]`
    case ColorOptions.ETH:
      return `text-[#5170ad]`
    case ColorOptions.GRAY:
    case ColorOptions.GREEN:
    case ColorOptions.LIME:
    case ColorOptions.SKY:
    case ColorOptions.BLUE:
    case ColorOptions.ORANGE:
    case ColorOptions.PURPLE:
    case ColorOptions.INDIGO:
    case ColorOptions.CYAN:
    case ColorOptions.RED:
    case ColorOptions.TEAL:
      return 'text-white'
    default:
      return 'text-white'
  }
}

export const getNetworkTextColorContrastHover = (
  chainColor: string
): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `group-hover:text-[#ecae0b]`
    case ColorOptions.ETH:
      return `group-hover:text-[#5170ad]`
    case ColorOptions.GRAY:
    case ColorOptions.GREEN:
    case ColorOptions.LIME:
    case ColorOptions.SKY:
    case ColorOptions.BLUE:
    case ColorOptions.ORANGE:
    case ColorOptions.PURPLE:
    case ColorOptions.INDIGO:
    case ColorOptions.CYAN:
    case ColorOptions.RED:
    case ColorOptions.TEAL:
      return 'group-hover:text-white'
    default:
      return 'group-hover:text-white'
  }
}

export const getNetworkBgClassName = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return 'bg-stone-800'
    case ColorOptions.ETH:
      return `bg-[#5170ad]`
    case ColorOptions.GRAY:
      return `bg-gray-500`
    case ColorOptions.GREEN:
      return `bg-green-500`
    case ColorOptions.LIME:
      return `bg-lime-500`
    case ColorOptions.SKY:
      return `bg-sky-500`
    case ColorOptions.BLUE:
      return `bg-blue-500`
    case ColorOptions.ORANGE:
      return `bg-orange-500`
    case ColorOptions.PURPLE:
      return `bg-purple-500`
    case ColorOptions.INDIGO:
      return `bg-indigo-500`
    case ColorOptions.CYAN:
      return `bg-cyan-500`
    case ColorOptions.RED:
      return `bg-red-500`
    case ColorOptions.TEAL:
      return `bg-teal-500`
    default:
      return `bg-gray-500`
  }
}

export const getNetworkBgClassNameLightDark = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `bg-[#ecae0b]`
    case ColorOptions.ETH:
      return `bg-[#5170ad]`
    case ColorOptions.GRAY:
      return `bg-gray-500`
    case ColorOptions.GREEN:
      return `bg-green-500`
    case ColorOptions.LIME:
      return `bg-lime-500`
    case ColorOptions.SKY:
      return `bg-sky-500`
    case ColorOptions.BLUE:
      return `bg-blue-500`
    case ColorOptions.ORANGE:
      return `bg-orange-500`
    case ColorOptions.PURPLE:
      return `bg-purple-500`
    case ColorOptions.INDIGO:
      return `bg-indigo-500`
    case ColorOptions.CYAN:
      return `bg-cyan-500`
    case ColorOptions.RED:
      return `bg-red-500`
    case ColorOptions.TEAL:
      return `bg-teal-500`
    default:
      return `bg-gray-500`
  }
}

export const getNetworkShadow = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.ETH:
      return `shadow-blue-xl hover:shadow-blue-2xl`
    case ColorOptions.YELLOW:
      return `shadow-[#ecae0b] hover:shadow-[#ecae0b]`
    case ColorOptions.GRAY:
      return `shadow-gray-xl hover:shadow-gray-2xl`
    case ColorOptions.GREEN:
      return `shadow-green-xl hover:shadow-green-2xl`
    case ColorOptions.LIME:
      return `shadow-lime-xl hover:shadow-lime-2xl`
    case ColorOptions.SKY:
      return `shadow-sky-xl hover:shadow-sky-2xl`
    case ColorOptions.BLUE:
      return `shadow-blue-xl hover:shadow-blue-2xl`
    case ColorOptions.ORANGE:
      return `shadow-orange-xl hover:shadow-orange-2xl`
    case ColorOptions.PURPLE:
      return `shadow-purple-xl hover:shadow-purple-2xl`
    case ColorOptions.INDIGO:
      return `shadow-indigo-xl hover:shadow-indigo-2xl`
    case ColorOptions.CYAN:
      return `shadow-cyan-xl hover:shadow-cyan-2xl`
    case ColorOptions.RED:
      return `shadow-red-xl hover:shadow-red-2xl`
    case ColorOptions.TEAL:
      return `shadow-teal-xl hover:shadow-teal-2xl`
    default:
      return `shadow-gray-xl hover:shadow-gray-2xl`
  }
}

export const getNetworkHover = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.ETH:
      return `hover:bg-[#5170ad] hover:bg-opacity-30`
    case ColorOptions.YELLOW:
      return `hover:bg-[#ecae0b] hover:bg-opacity-30`
    case ColorOptions.GRAY:
      return `hover:bg-gray-500 hover:bg-opacity-30`
    case ColorOptions.GREEN:
      return `hover:bg-green-500 hover:bg-opacity-30`
    case ColorOptions.LIME:
      return `hover:bg-lime-500 hover:bg-opacity-30`
    case ColorOptions.SKY:
      return `hover:bg-sky-500 hover:bg-opacity-30`
    case ColorOptions.BLUE:
      return `hover:bg-blue-500 hover:bg-opacity-30`
    case ColorOptions.ORANGE:
      return `hover:bg-orange-500 hover:bg-opacity-30`
    case ColorOptions.PURPLE:
      return `hover:bg-purple-500 hover:bg-opacity-30`
    case ColorOptions.INDIGO:
      return `hover:bg-indigo-500 hover:bg-opacity-30`
    case ColorOptions.CYAN:
      return `hover:bg-cyan-500 hover:bg-opacity-30`
    case ColorOptions.RED:
      return `hover:bg-red-500 hover:bg-opacity-30`
    case ColorOptions.TEAL:
      return `hover:bg-teal-500 hover:bg-opacity-30`
    default:
      return `hover:bg-gray-500 hover:bg-opacity-30`
  }
}
