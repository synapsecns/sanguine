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

export const getButtonStyleForCoin = (tokenColor: string) => {
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

export const getMenuItemStyleForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `hover:bg-opacity-20 hover:!bg-gray-500 focus:bg-gray-500 active:bg-gray-500`
    case ColorOptions.yellow:
      return `hover:bg-opacity-20 hover:bg-yellow-500 focus:bg-yellow-500 active:bg-yellow-500`
    case ColorOptions.green:
      return `hover:bg-opacity-20 hover:!bg-green-500 focus:bg-green-500 active:bg-green-500`
    case ColorOptions.lime:
      return `hover:bg-opacity-20 hover:!bg-lime-500 focus:bg-lime-500 active:bg-lime-500`
    case ColorOptions.sky:
      return `hover:bg-opacity-20 hover:bg-sky-500 focus:bg-sky-500 active:bg-sky-500`
    case ColorOptions.blue:
      return `hover:bg-opacity-20 hover:bg-blue-500 focus:bg-blue-500 active:bg-blue-500`
    case ColorOptions.orange:
      return `hover:bg-opacity-20 hover:bg-orange-500 focus:bg-orange-500 active:bg-orange-500`
    case ColorOptions.purple:
      return `hover:bg-opacity-20 hover:bg-purple-500 focus:bg-purple-500 active:bg-purple-500`
    case ColorOptions.indigo:
      return `hover:bg-opacity-20 hover:bg-indigo-500 focus:bg-indigo-500 active:bg-indigo-500`
    case ColorOptions.cyan:
      return `hover:bg-opacity-20 hover:bg-cyan-500 focus:bg-cyan-500 active:bg-cyan-500`
    case ColorOptions.red:
      return `hover:bg-opacity-20 hover:bg-red-500 focus:bg-red-500 active:bg-red-500`
    default:
      return ''
  }
}

export const getMenuItemStyleForCoinDark = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-gray-500 dark:focus:bg-gray-500 dark:active:bg-gray-500`
    case ColorOptions.yellow:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-yellow-500 dark:focus:bg-yellow-500 dark:active:bg-yellow-500`
    case ColorOptions.green:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-green-500 dark:focus:bg-green-500 dark:active:bg-green-500`
    case ColorOptions.lime:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-lime-500 dark:focus:bg-lime-500 dark:active:bg-lime-500`
    case ColorOptions.sky:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-sky-500 dark:focus:bg-sky-500 dark:active:bg-sky-500`
    case ColorOptions.blue:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-blue-500 dark:focus:bg-blue-500 dark:active:bg-blue-500`
    case ColorOptions.orange:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-orange-500 dark:focus:bg-orange-500 dark:active:bg-orange-500`
    case ColorOptions.purple:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-purple-500 dark:focus:bg-purple-500 dark:active:bg-purple-500`
    case ColorOptions.indigo:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-indigo-500 dark:focus:bg-indigo-500 dark:active:bg-indigo-500`
    case ColorOptions.cyan:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-cyan-500 dark:focus:bg-cyan-500 dark:active:bg-cyan-500`
    case ColorOptions.red:
      return `dark:hover:bg-opacity-20 dark:focus:bg-opacity-20 dark:active:bg-opacity-20 dark:hover:bg-red-500 dark:focus:bg-red-500 dark:active:bg-red-500`
    default:
      return ''
  }
}

export const getMenuItemStyleForCoinCombined = (tokenColor: string): string => {
  return `${getMenuItemStyleForCoin(tokenColor)}
  ${getMenuItemStyleForCoinDark(tokenColor)}`
}

export const getSwapHoverStyleForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `hover:shadow-gray-xl border-gray-100`
    case ColorOptions.yellow:
      return `hover:shadow-yellow-xl border-yellow-100`
    case ColorOptions.green:
      return `hover:shadow-green-xl border-green-100`
    case ColorOptions.lime:
      return `hover:shadow-lime-xl border-lime-100`
    case ColorOptions.sky:
      return `hover:shadow-sky-xl border-sky-100`
    case ColorOptions.blue:
      return `hover:shadow-blue-xl border-blue-100`
    case ColorOptions.orange:
      return `hover:shadow-orange-xl border-orange-100`
    case ColorOptions.purple:
      return `hover:shadow-purple-xl border-purple-100`
    case ColorOptions.indigo:
      return `hover:shadow-indigo-xl border-indigo-100`
    case ColorOptions.cyan:
      return `hover:shadow-cyan-xl border-cyan-100`
    case ColorOptions.red:
      return `hover:shadow-red-xl border-red-100`
    default:
      return ''
  }
}

export const getMenuItemBgForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `bg-gray-50 dark:bg-opacity-20 dark:bg-gray-700`
    case ColorOptions.yellow:
      return `bg-yellow-50 dark:bg-opacity-20 dark:bg-yellow-700`
    case ColorOptions.green:
      return `bg-green-50 dark:bg-opacity-20 dark:bg-green-700`
    case ColorOptions.lime:
      return `bg-lime-50 dark:bg-opacity-20 dark:bg-lime-700`
    case ColorOptions.sky:
      return `bg-sky-50 dark:bg-opacity-20 dark:bg-sky-700`
    case ColorOptions.blue:
      return `bg-blue-50 dark:bg-opacity-20 dark:bg-blue-700`
    case ColorOptions.orange:
      return `bg-orange-50 dark:bg-opacity-20 dark:bg-orange-700`
    case ColorOptions.purple:
      return `bg-purple-50 dark:bg-opacity-20 dark:bg-purple-700`
    case ColorOptions.indigo:
      return `bg-indigo-50 dark:bg-opacity-20 dark:bg-indigo-700`
    case ColorOptions.cyan:
      return `bg-cyan-50 dark:bg-opacity-20 dark:bg-cyan-700`
    case ColorOptions.red:
      return `bg-red-50 dark:bg-opacity-20 dark:bg-red-700`
    default:
      return ''
  }
}

export const getMenuItemHoverBgForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `hover:bg-opacity-20 hover:bg-gray-500 dark:hover:bg-opacity-20 dark:hover:bg-gray-700`
    case ColorOptions.yellow:
      return `hover:bg-opacity-20 hover:bg-yellow-500 dark:hover:bg-opacity-20 dark:hover:bg-yellow-700`
    case ColorOptions.green:
      return `hover:bg-opacity-20 hover:bg-green-500 dark:hover:bg-opacity-20 dark:hover:bg-green-700`
    case ColorOptions.lime:
      return `hover:bg-opacity-20 hover:bg-lime-500 dark:hover:bg-opacity-20 dark:hover:bg-lime-700`
    case ColorOptions.sky:
      return `hover:bg-opacity-20 hover:bg-sky-500 dark:hover:bg-opacity-20 dark:hover:bg-sky-700`
    case ColorOptions.blue:
      return `hover:bg-opacity-20 hover:bg-blue-500 dark:hover:bg-opacity-20 dark:hover:bg-blue-700`
    case ColorOptions.orange:
      return `hover:bg-opacity-20 hover:bg-orange-500 dark:hover:bg-opacity-20 dark:hover:bg-orange-700`
    case ColorOptions.purple:
      return `hover:bg-opacity-20 hover:bg-purple-500 dark:hover:bg-opacity-20 dark:hover:bg-purple-700`
    case ColorOptions.indigo:
      return `hover:bg-opacity-20 hover:bg-indigo-500 dark:hover:bg-opacity-20 dark:hover:bg-indigo-700`
    case ColorOptions.cyan:
      return `hover:bg-opacity-20 hover:bg-cyan-500 dark:hover:bg-opacity-20 dark:hover:bg-cyan-700`
    case ColorOptions.red:
      return `hover:bg-opacity-20 hover:bg-red-500 dark:hover:bg-opacity-20 dark:hover:bg-red-700`
    default:
      return ''
  }
}

export const getCoinTextColor = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `text-gray-500 group-hover:text-gray-400`
    case ColorOptions.yellow:
      return `text-yellow-500 group-hover:text-yellow-400`
    case ColorOptions.green:
      return `text-green-500 group-hover:text-green-400`
    case ColorOptions.lime:
      return `text-lime-500 group-hover:text-lime-400`
    case ColorOptions.sky:
      return `text-sky-500 group-hover:text-sky-400`
    case ColorOptions.blue:
      return `text-blue-500 group-hover:text-blue-400`
    case ColorOptions.orange:
      return `text-orange-500 group-hover:text-orange-400`
    case ColorOptions.purple:
      return `text-purple-500 group-hover:text-purple-400`
    case ColorOptions.indigo:
      return `text-indigo-500 group-hover:text-indigo-400`
    case ColorOptions.cyan:
      return `text-cyan-500 group-hover:text-cyan-400`
    case ColorOptions.red:
      return `text-red-500 group-hover:text-red-400`
    default:
      return ''
  }
}

export const getCoinTextColorDark = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `dark:text-gray-500 dark:group-hover:text-gray-400`
    case ColorOptions.yellow:
      return `dark:text-yellow-500 dark:group-hover:text-yellow-400`
    case ColorOptions.green:
      return `dark:text-green-500 dark:group-hover:text-green-400`
    case ColorOptions.lime:
      return `dark:text-lime-500 dark:group-hover:text-lime-400`
    case ColorOptions.sky:
      return `dark:text-sky-500 dark:group-hover:text-sky-400`
    case ColorOptions.blue:
      return `dark:text-blue-500 dark:group-hover:text-blue-400`
    case ColorOptions.orange:
      return `dark:text-orange-500 dark:group-hover:text-orange-400`
    case ColorOptions.purple:
      return `dark:text-purple-500 dark:group-hover:text-purple-400`
    case ColorOptions.indigo:
      return `dark:text-indigo-500 dark:group-hover:text-indigo-400`
    case ColorOptions.cyan:
      return `dark:text-cyan-500 dark:group-hover:text-cyan-400`
    case ColorOptions.red:
      return `dark:text-red-500 dark:group-hover:text-red-400`
    default:
      return ''
  }
}

export const getCoinTextColorAlt = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `dark:text-gray-500 dark:group-hover:text-gray-400`
    case ColorOptions.yellow:
      return `dark:text-yellow-500 dark:group-hover:text-yellow-400`
    case ColorOptions.green:
      return `dark:text-green-500 dark:group-hover:text-green-400`
    case ColorOptions.lime:
      return `dark:text-lime-500 dark:group-hover:text-lime-400`
    case ColorOptions.sky:
      return `dark:text-sky-500 dark:group-hover:text-sky-400`
    case ColorOptions.blue:
      return `dark:text-blue-500 dark:group-hover:text-blue-400`
    case ColorOptions.orange:
      return `dark:text-orange-500 dark:group-hover:text-orange-400`
    case ColorOptions.purple:
      return `dark:text-purple-500 dark:group-hover:text-purple-400`
    case ColorOptions.indigo:
      return `dark:text-indigo-500 dark:group-hover:text-indigo-400`
    case ColorOptions.cyan:
      return `dark:text-cyan-500 dark:group-hover:text-cyan-400`
    case ColorOptions.red:
      return `dark:text-red-500 dark:group-hover:text-red-400`
    default:
      return ''
  }
}

export const getCoinTextColorCombined = (tokenColor: string): string => {
  return `${getCoinTextColor(tokenColor)} ${getCoinTextColorDark(tokenColor)}`
}

export const getCardStyleByPool = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `shadow-gray-lg hover:shadow-gray-2xl`
    case ColorOptions.yellow:
      return `shadow-yellow-lg hover:shadow-yellow-2xl`
    case ColorOptions.green:
      return `shadow-green-lg hover:shadow-green-2xl`
    case ColorOptions.lime:
      return `shadow-lime-lg hover:shadow-lime-2xl`
    case ColorOptions.sky:
      return `shadow-sky-lg hover:shadow-sky-2xl`
    case ColorOptions.blue:
      return `shadow-blue-lg hover:shadow-blue-2xl`
    case ColorOptions.orange:
      return `shadow-orange-lg hover:shadow-orange-2xl`
    case ColorOptions.purple:
      return `shadow-purple-lg hover:shadow-purple-2xl`
    case ColorOptions.indigo:
      return `shadow-indigo-lg hover:shadow-indigo-2xl`
    case ColorOptions.cyan:
      return `shadow-cyan-lg hover:shadow-cyan-2xl`
    case ColorOptions.red:
      return `shadow-red-lg hover:shadow-red-2xl`
    default:
      return ''
  }
}

export const getInputBorderFocusStyleForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `focus-within:border-gray-200 dark:focus-within:border-gray-500`
    case ColorOptions.yellow:
      return `focus-within:border-yellow-200 dark:focus-within:border-yellow-500`
    case ColorOptions.green:
      return `focus-within:border-green-200 dark:focus-within:border-green-500`
    case ColorOptions.lime:
      return `focus-within:border-lime-200 dark:focus-within:border-lime-500`
    case ColorOptions.sky:
      return `focus-within:border-sky-200 dark:focus-within:border-sky-500`
    case ColorOptions.blue:
      return `focus-within:border-blue-200 dark:focus-within:border-blue-500`
    case ColorOptions.orange:
      return `focus-within:border-orange-200 dark:focus-within:border-orange-500`
    case ColorOptions.purple:
      return `focus-within:border-purple-200 dark:focus-within:border-purple-500`
    case ColorOptions.indigo:
      return `focus-within:border-indigo-200 dark:focus-within:border-indigo-500`
    case ColorOptions.cyan:
      return `focus-within:border-cyan-200 dark:focus-within:border-cyan-500`
    case ColorOptions.red:
      return `focus-within:border-red-200 dark:focus-within:border-red-500`
    default:
      return ''
  }
}

export const getSwapBorderStyleForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `border-gray-50 dark:border-opacity-20 dark:border-gray-700`
    case ColorOptions.yellow:
      return `border-yellow-50 dark:border-opacity-20 dark:border-yellow-700`
    case ColorOptions.green:
      return `border-green-50 dark:border-opacity-20 dark:border-green-700`
    case ColorOptions.lime:
      return `border-lime-50 dark:border-opacity-20 dark:border-lime-700`
    case ColorOptions.sky:
      return `border-sky-50 dark:border-opacity-20 dark:border-sky-700`
    case ColorOptions.blue:
      return `border-blue-50 dark:border-opacity-20 dark:border-blue-700`
    case ColorOptions.orange:
      return `border-orange-50 dark:border-opacity-20 dark:border-orange-700`
    case ColorOptions.purple:
      return `border-purple-50 dark:border-opacity-20 dark:border-purple-700`
    case ColorOptions.indigo:
      return `border-indigo-50 dark:border-opacity-20 dark:border-indigo-700`
    case ColorOptions.cyan:
      return `border-cyan-50 dark:border-opacity-20 dark:border-cyan-700`
    case ColorOptions.red:
      return `border-red-50 dark:border-opacity-20 dark:border-red-700`
    default:
      return ''
  }
}

export const getSwapBorderHoverStyleForCoin = (tokenColor: string): string => {
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

export const getBorderStyleForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `border-gray-300`
    case ColorOptions.yellow:
      return `border-yellow-300`
    case ColorOptions.green:
      return `border-green-300`
    case ColorOptions.lime:
      return `border-lime-300`
    case ColorOptions.sky:
      return `border-sky-300`
    case ColorOptions.blue:
      return `border-blue-300`
    case ColorOptions.orange:
      return `border-orange-300`
    case ColorOptions.purple:
      return `border-purple-300`
    case ColorOptions.indigo:
      return `border-indigo-300`
    case ColorOptions.cyan:
      return `border-cyan-300`
    case ColorOptions.red:
      return `border-red-300`
    default:
      return ''
  }
}

export const getFromStyleForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `from-gray-300`
    case ColorOptions.yellow:
      return `from-yellow-300`
    case ColorOptions.green:
      return `from-green-300`
    case ColorOptions.lime:
      return `from-lime-300`
    case ColorOptions.sky:
      return `from-sky-300`
    case ColorOptions.blue:
      return `from-blue-300`
    case ColorOptions.orange:
      return `from-orange-300`
    case ColorOptions.purple:
      return `from-purple-300`
    case ColorOptions.indigo:
      return `from-indigo-300`
    case ColorOptions.cyan:
      return `from-cyan-300`
    case ColorOptions.red:
      return `from-red-300`
    default:
      return ''
  }
}

export const getToStyleForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `to-gray-300`
    case ColorOptions.yellow:
      return `to-yellow-300`
    case ColorOptions.green:
      return `to-green-300`
    case ColorOptions.lime:
      return `to-lime-300`
    case ColorOptions.sky:
      return `to-sky-300`
    case ColorOptions.blue:
      return `to-blue-300`
    case ColorOptions.orange:
      return `to-orange-300`
    case ColorOptions.purple:
      return `to-purple-300`
    case ColorOptions.indigo:
      return `to-indigo-300`
    case ColorOptions.cyan:
      return `to-cyan-300`
    case ColorOptions.red:
      return `to-red-300`
    default:
      return ''
  }
}

export const getSwapCardShadowStyleForCoin = (tokenColor: string): string => {
  switch (tokenColor) {
    case ColorOptions.gray:
      return `shadow-gray-xl hover:shadow-gray-2xl`
    case ColorOptions.yellow:
      return `shadow-yellow-xl hover:shadow-yellow-2xl`
    case ColorOptions.green:
      return `shadow-green-xl hover:shadow-green-2xl`
    case ColorOptions.lime:
      return `shadow-lime-xl hover:shadow-lime-2xl`
    case ColorOptions.sky:
      return `shadow-sky-xl hover:shadow-sky-2xl`
    case ColorOptions.blue:
      return `shadow-blue-xl hover:shadow-blue-2xl`
    case ColorOptions.orange:
      return `shadow-orange-xl hover:shadow-orange-2xl`
    case ColorOptions.purple:
      return `shadow-purple-xl hover:shadow-purple-2xl`
    case ColorOptions.indigo:
      return `shadow-indigo-xl hover:shadow-indigo-2xl`
    case ColorOptions.cyan:
      return `shadow-cyan-xl hover:shadow-cyan-2xl`
    case ColorOptions.red:
      return `shadow-red-xl hover:shadow-red-2xl`
    default:
      return ''
  }
}

export const getBorderStyleForCoinHover = (tokenColor: string) => {
  switch (tokenColor) {
    case ColorOptions.yellow:
      return 'hover:border-yellow-300'
    case ColorOptions.green:
      return 'hover:border-green-300'
    case ColorOptions.lime:
      return 'hover:border-lime-300'
    case ColorOptions.sky:
      return 'hover:border-sky-300'
    case ColorOptions.blue:
      return 'hover:border-blue-300'
    case ColorOptions.orange:
      return 'hover:border-orange-300'
    case ColorOptions.purple:
      return 'hover:border-purple-300'
    case ColorOptions.indigo:
      return 'hover:border-indigo-300'
    case ColorOptions.cyan:
      return 'hover:border-cyan-300'
    case ColorOptions.red:
      return 'hover:border-red-300'
    default:
      return 'hover:border-gray-200'
  }
}
