const SHARED_OPACITY_OVERRIDE = `
  dark:hover:bg-opacity-20
  dark:focus:bg-opacity-20
  dark:active:bg-opacity-20
  `

const ColorOptions = {
  gray: 'gray',
  yellow: 'yellow',
  green: 'green',
  lime: 'lime',
  sky: 'sky',
  blue: 'blue',
  orange: 'orange',
  purple: 'purple',
  indigo: 'indigo',
  cyan: 'cyan',
  red: 'red',
}

// export const getButtonStyleForCoin = (tokenColor: string): string => {
//   `shadow-${tokenColor}-xl border-${tokenColor}-100 dark:border-opacity-50 dark:border-${tokenColor}-700`

// }

export const getButtonStyleForCoin = (color: string) => {
  switch (color) {
    case ColorOptions.gray:
      return `hover:border-gray-100 dark:hover:border-opacity-50 dark:hover:!border-gray-700`
    case ColorOptions.yellow:
      return `hover:border-yellow-100 dark:hover:border-opacity-50 dark:hover:!border-yellow-700`
    case ColorOptions.green:
      return `hover:border-green-100 dark:hover:border-opacity-50 dark:hover:!border-green-700`
    case ColorOptions.lime:
      return `hover:border-lime-100 dark:hover:border-opacity-50 dark:hover:!border-lime-700`
    case ColorOptions.sky:
      return `hover:border-sky-100 dark:hover:border-opacity-50 dark:hover:!border-sky-700`
    case ColorOptions.blue:
      return `hover:border-blue-100 dark:hover:border-opacity-50 dark:hover:!border-blue-700`
    case ColorOptions.orange:
      return `hover:border-orange-100 dark:hover:border-opacity-50 dark:hover:!border-orange-700`
    case ColorOptions.purple:
      return `hover:border-purple-100 dark:hover:border-opacity-50 dark:hover:!border-purple-700`
    case ColorOptions.indigo:
      return `hover:border-indigo-100 dark:hover:border-opacity-50 dark:hover:!border-indigo-700`
    case ColorOptions.cyan:
      return `hover:border-cyan-100 dark:hover:border-opacity-50 dark:hover:!border-cyan-700`
    case ColorOptions.red:
      return `hover:border-red-100 dark:hover:border-opacity-50 dark:hover:!border-red-700`
    default:
      return ''
  }
}

// export const getMenuItemStyleForCoin = (tokenColor: string): string =>
//   `hover:bg-${tokenColor}-50 focus:bg-${tokenColor}-50 active:bg-${tokenColor}-50`

export const getMenuItemStyleForCoin = (color: string): string => {
  switch (color) {
    case ColorOptions.gray:
      return `hover:bg-gray-50 focus:bg-gray-50 active:bg-gray-50`
    case ColorOptions.yellow:
      return `hover:bg-yellow-50 focus:bg-yellow-50 active:bg-yellow-50`
    case ColorOptions.green:
      return `hover:bg-green-50 focus:bg-green-50 active:bg-green-50`
    case ColorOptions.lime:
      return `hover:bg-lime-50 focus:bg-lime-50 active:bg-lime-50`
    case ColorOptions.sky:
      return `hover:bg-sky-50 focus:bg-sky-50 active:bg-sky-50`
    case ColorOptions.blue:
      return `hover:bg-blue-50 focus:bg-blue-50 active:bg-blue-50`
    case ColorOptions.orange:
      return `hover:bg-orange-50 focus:bg-orange-50 active:bg-orange-50`
    case ColorOptions.purple:
      return `hover:bg-purple-50 focus:bg-purple-50 active:bg-purple-50`
    case ColorOptions.indigo:
      return `hover:bg-indigo-50 focus:bg-indigo-50 active:bg-indigo-50`
    case ColorOptions.cyan:
      return `hover:bg-cyan-50 focus:bg-cyan-50 active:bg-cyan-50`
    case ColorOptions.red:
      return `hover:bg-red-50 focus:bg-red-50 active:bg-red-50`
    default:
      return ''
  }
}

// export const getMenuItemStyleForCoinDark = (tokenColor: string): string =>
//   `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-${tokenColor}-500 dark:focus:bg-${tokenColor}-500 dark:active:bg-${tokenColor}-500`

export const getMenuItemStyleForCoinDark = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-gray-500 dark:focus:bg-gray-500 dark:active:bg-gray-500`
    case ColorOptions.yellow:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-yellow-500 dark:focus:bg-yellow-500 dark:active:bg-yellow-500`
    case ColorOptions.green:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-green-500 dark:focus:bg-green-500 dark:active:bg-green-500`
    case ColorOptions.lime:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-lime-500 dark:focus:bg-lime-500 dark:active:bg-lime-500`
    case ColorOptions.sky:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-sky-500 dark:focus:bg-sky-500 dark:active:bg-sky-500`
    case ColorOptions.blue:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-blue-500 dark:focus:bg-blue-500 dark:active:bg-blue-500`
    case ColorOptions.orange:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-orange-500 dark:focus:bg-orange-500 dark:active:bg-orange-500`
    case ColorOptions.purple:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-purple-500 dark:focus:bg-purple-500 dark:active:bg-purple-500`
    case ColorOptions.indigo:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-indigo-500 dark:focus:bg-indigo-500 dark:active:bg-indigo-500`
    case ColorOptions.cyan:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-cyan-500 dark:focus:bg-cyan-500 dark:active:bg-cyan-500`
    case ColorOptions.red:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-red-500 dark:focus:bg-red-500 dark:active:bg-red-500`
    default:
      return ''
  }
}

export const getMenuItemStyleForCoinCombined = (tokenColor: string): string => {
  return `${getMenuItemStyleForCoin(tokenColor)}
  ${getMenuItemStyleForCoinDark(tokenColor)}`
}

export const getSwapHoverStyleForCoin = (tokenColor: string): string =>
  `hover:shadow-${tokenColor}-xl border-${tokenColor}-100`

export const getMenuItemBgForCoin = (tokenColor: string): string =>
  `bg-${tokenColor}-50 dark:bg-opacity-20 dark:bg-${tokenColor}-700`

export const getMenuItemHoverBgForCoin = (tokenColor: string): string =>
  ` hover:bg-${tokenColor}-100 dark:hover:bg-opacity-20 dark:hover:bg-${tokenColor}-700`

export const getCoinTextColor = (tokenColor: string): string =>
  `text-${tokenColor}-500 group-hover:text-${tokenColor}-400`

export const getCoinTextColorDark = (tokenColor: string): string =>
  `dark:text-${tokenColor}-500 dark:group-hover:text-${tokenColor}-400`

export const getCoinTextColorAlt = (tokenColor: string): string =>
  `dark:text-${tokenColor}-500 dark:group-hover:text-${tokenColor}-400`

export const getCoinTextColorCombined = (tokenColor: string): string => {
  return `${getCoinTextColor(tokenColor)} ${getCoinTextColorDark(tokenColor)}`
}

export const getCardStyleByPool = (tokenColor: string): string =>
  `shadow-${tokenColor}-lg hover:shadow-${tokenColor}-2xl`

export const getInputBorderFocusStyleForCoin = (tokenColor: string): string =>
  `focus-within:border-${tokenColor}-200 dark:focus-within:border-${tokenColor}-500`

export const getSwapBorderStyleForCoin = (tokenColor: string): string =>
  `border-${tokenColor}-50 dark:border-opacity-20 dark:border-${tokenColor}-700`

export const getSwapBorderHoverStyleForCoin = (tokenColor: string): string =>
  `hover:border-${tokenColor}-100 dark:hover:border-opacity-50 dark:hover:!border-${tokenColor}-700`

/**
 * @param {Token} coin
 */
export const getBorderStyleForCoin = (tokenColor: string): string =>
  `border-${tokenColor}-300`

export const getFromStyleForCoin = (tokenColor: string): string =>
  `from-${tokenColor}-300`

export const getToStyleForCoin = (tokenColor: string): string =>
  `to-${tokenColor}-300`

export const getSwapCardShadowStyleForCoin = (tokenColor: string): string =>
  `shadow-${tokenColor}-xl hover:shadow-${tokenColor}-2xl`

export const getBorderStyleForCoinHover = (color: string) => {
  switch (color) {
    case ColorOptions.yellow:
      return 'hover:!border-yellow-300'
    case ColorOptions.green:
      return 'hover:!border-green-300'
    case ColorOptions.lime:
      return 'hover:!border-lime-300'
    case ColorOptions.sky:
      return 'hover:!border-sky-300'
    case ColorOptions.blue:
      return 'hover:!border-blue-300'
    case ColorOptions.orange:
      return 'hover:!border-orange-300'
    case ColorOptions.purple:
      return 'hover:!border-purple-300'
    case ColorOptions.indigo:
      return 'hover:!border-indigo-300'
    case ColorOptions.cyan:
      return 'hover:!border-cyan-300'
    case ColorOptions.red:
      return 'hover:!border-red-300'
    default:
      return 'hover:border-gray-200'
  }
}
