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

export const getButtonStyleForCoin = (tokenColor: ColorOptions) => {
  switch (tokenColor) {
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

export const getMenuItemStyleForCoin = (tokenColor: string): string =>
  `hover:bg-${tokenColor}-50 focus:bg-${tokenColor}-50 active:bg-${tokenColor}-50`

export const getMenuItemStyleForCoinDark = (tokenColor: string): string =>
  `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-${tokenColor}-500 dark:focus:bg-${tokenColor}-500 dark:active:bg-${tokenColor}-500`

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

// export const getBorderStyleForCoinHover = (tokenColor: string): string =>
//   `hover:!border-${tokenColor}-300`

// export const getBorderStyleForCoinHover = (tokenColor: string) => {
//   switch (tokenColor) {
//     case ColorOptions.gray:
//       return 'hover:!border-gray-300'
//     case ColorOptions.yellow:
//       return 'hover:!border-yellow-300'
//     case ColorOptions.green:
//       return 'hover:!border-green-300'
//     case ColorOptions.lime:
//       return 'hover:!border-lime-300'
//     case ColorOptions.sky:
//       return 'hover:!border-sky-300'
//     case ColorOptions.blue:
//       return 'hover:!border-blue-300'
//     case ColorOptions.orange:
//       return 'hover:!border-orange-300'
//     case ColorOptions.purple:
//       return 'hover:!border-purple-300'
//     case ColorOptions.indigo:
//       return 'hover:!border-indigo-300'
//     case ColorOptions.cyan:
//       return 'hover:!border-cyan-300'
//     case ColorOptions.red:
//       return 'hover:!border-red-300'
//     default:
//       return ''
//   }
// }

/* eslint-disable */
/**
 * @param {Token} coin
 */
export function getBorderStyleForCoinHover(color: string) {
  switch (color) {
    case 'yellow':
      return 'hover:!border-yellow-300'
    case 'green':
      return 'hover:!border-green-300'
    case 'lime':
      return 'hover:!border-lime-300'
    case 'sky':
      return 'hover:!border-sky-300'
    case 'blue':
      return 'hover:!border-blue-300'
    case 'orange':
      return 'hover:!border-orange-300'
    case 'purple':
      return 'hover:!border-purple-300'
    case 'indigo':
      return 'hover:!border-indigo-300'
    case 'cyan':
      return 'hover:!border-cyan-300'
    case 'red':
      return 'hover:!border-red-300'
    default:
      return 'hover:!border-gray-200'
  }
}
