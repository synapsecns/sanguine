import { Chain } from '@types'

const CUSTOM_YELLOW = '[#ecae0b]'
const CUSTOM_YELLOW_DARK = '[#3c3c44]'
const ETH_BASE = '[#5170ad]'
const ETH_DARK = '[#3f4f8c]'
const ETH_EXTRA_DARK = '[#314367]'
const ETH_LIGHT = '[#78a5ff]'

const CustomClasses = {
  CUSTOM_YELLOW_BG: `bg-${CUSTOM_YELLOW}`,
  HOVER_CUSTOM_YELLOW_BG: `hover:!bg-${CUSTOM_YELLOW}`,
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
}

// TODO obviously this can be cleaned up
// 1. define custom yellow/eth with tailwind
// 2. remove all the if/else
// or
// combine all of this in one function with a switch for the different desired classes.

// Revisit if any of these can be deleted.

export const getNetworkCurrencyColor = (chain: Chain): string => {
  if (chain.color === 'yellow') {
    return `text-${CUSTOM_YELLOW} dark:text-${CUSTOM_YELLOW}`
  } else if (chain.nativeCurrency.symbol === 'ETH') {
    return `text-${ETH_BASE} dark:text-${ETH_LIGHT}`
  }
  return `text-${chain.color}-500 dark:text-${chain.color}-500`
}

export const getNetworkButtonBgClassName = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `bg-stone-800 hover:bg-stone-900 active:bg-${CUSTOM_YELLOW_DARK}`
  } else if (chainColor === 'eth') {
    return `bg-${ETH_BASE} hover:bg-${ETH_DARK} active:bg-${ETH_EXTRA_DARK}`
  }
  return `bg-${chainColor}-500 hover:bg-${chainColor}-600 active:bg-${chainColor}-700`
}
export const getNetworkButtonBgClassNameActive = (
  chainColor: string
): string => {
  if (chainColor === 'yellow') {
    return `dark:active:!bg-${CUSTOM_YELLOW_DARK}`
  } else if (chainColor === 'eth') {
    return `dark:active:!bg-${ETH_EXTRA_DARK}`
  }
  return `dark:active:!bg-${chainColor}-500`
}
export const getNetworkButtonBorderHover = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `hover:!border-${CUSTOM_YELLOW}`
  } else if (chainColor === 'eth') {
    return `hover:!border-${ETH_BASE}`
  }
  return `hover:!border-${chainColor}-500'`
}
export const getNetworkButtonBorderActive = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `focus:!border-${CUSTOM_YELLOW} active:!border-${CUSTOM_YELLOW}`
  } else if (chainColor === 'eth') {
    return `focus:!border-${ETH_BASE} active:!border-${ETH_BASE}`
  }
  return `focus:!border-${chainColor}-500 active:!border-${chainColor}-500`
}

export const getNetworkButtonBorder = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `border-${CUSTOM_YELLOW} dark:border-${CUSTOM_YELLOW}`
  } else if (chainColor === 'eth') {
    return `border-${ETH_BASE} dark:border-${ETH_BASE}`
  }
  return `border-${chainColor}-500 dark:border-${chainColor}-500`
}

export const getNetworkButtonBorderImportant = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `!border-${CUSTOM_YELLOW} dark:!border-${CUSTOM_YELLOW}`
  } else if (chainColor === 'eth') {
    return `!border-${ETH_BASE} dark:!border-${ETH_BASE}`
  }
  return `!border-${chainColor}-500 dark:!border-${chainColor}-500`
}

export const getNetworkTextColor = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `text-${CUSTOM_YELLOW} dark:text-${CUSTOM_YELLOW}`
  } else if (chainColor === 'eth') {
    return `text-${ETH_BASE} dark:text-${ETH_LIGHT}`
  }
  return `text-${chainColor}-500 dark:text-${chainColor}-500`
}

export const getNetworkLinkTextColor = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `!text-gray-800 hover:!text-${CUSTOM_YELLOW} dark:!text-${CUSTOM_YELLOW} dark:hover:!text-${CUSTOM_YELLOW}`
  } else if (chainColor === 'eth') {
    return `!text-${ETH_BASE} hover:!text-${ETH_LIGHT} dark:hover:!text-${ETH_LIGHT}`
  }
  return `!text-${chainColor}-500 hover:!text-${chainColor}-600 dark:hover:!text-${chainColor}-500`
}

// export const getNetworkTextColorContrast = (chainColor: string): string => {
//   if (chainColor === 'yellow') {
//     return `text-${CUSTOM_YELLOW}`
//   }
//   return 'text-white'
// }

export const getNetworkTextColorContrast = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `text-${CUSTOM_YELLOW}`
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
      return 'text-white'
    default:
      return 'text-white'
  }
}

// export const getNetworkTextColorContrastHover = (
//   chainColor: string
// ): string => {
//   if (chainColor === 'yellow') {
//     return `group-hover:text-${CUSTOM_YELLOW}`
//   }
//   return 'group-hover:text-white'
// }

export const getNetworkTextColorContrastHover = (
  chainColor: string
): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `group-hover:text-${CUSTOM_YELLOW}`
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
      return `bg-${ETH_BASE}`
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
    default:
      return `bg-gray-500`
  }
}

// export const getNetworkBgClassName = (chainColor: string): string => {
//   if (chainColor === 'yellow') {
//     return 'bg-stone-800'
//   } else if (chainColor === 'eth') {
//     return `bg-${ETH_BASE}`
//   }
//   return `bg-${chainColor}-500`
// }

export const getNetworkBgClassNameLightDark = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.YELLOW:
      return `${CustomClasses.CUSTOM_YELLOW_BG}`
    case ColorOptions.ETH:
      return `bg-${ETH_BASE}`
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
    default:
      return `bg-gray-500`
  }
}

// export const getNetworkBgClassNameLightDark = (chainColor: string): string => {
//   if (chainColor === 'yellow') {
//     return `bg-${CUSTOM_YELLOW}`
//   } else if (chainColor === 'eth') {
//     return `bg-${ETH_BASE}`
//   }
//   return `bg-${chainColor}-500`
// }

export const getNetworkShadow = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.ETH:
      return `!shadow-blue-xl hover:!shadow-blue-2xl`
    case ColorOptions.YELLOW:
      return `!shadow-[#ecae0b] hover:!shadow-[#ecae0b]`
    case ColorOptions.GRAY:
      return `!shadow-gray-xl hover:!shadow-gray-2xl`
    case ColorOptions.GREEN:
      return `!shadow-green-xl hover:!shadow-green-2xl`
    case ColorOptions.LIME:
      return `!shadow-lime-xl hover:!shadow-lime-2xl`
    case ColorOptions.SKY:
      return `!shadow-sky-xl hover:!shadow-sky-2xl`
    case ColorOptions.BLUE:
      return `!shadow-blue-xl hover:!shadow-blue-2xl`
    case ColorOptions.ORANGE:
      return `!shadow-orange-xl hover:!shadow-orange-2xl`
    case ColorOptions.PURPLE:
      return `!shadow-purple-xl hover:!shadow-purple-2xl`
    case ColorOptions.INDIGO:
      return `!shadow-indigo-xl hover:!shadow-indigo-2xl`
    case ColorOptions.CYAN:
      return `!shadow-cyan-xl hover:!shadow-cyan-2xl`
    case ColorOptions.RED:
      return `!shadow-red-xl hover:!shadow-red-2xl`
    default:
      return `!shadow-gray-xl hover:!shadow-gray-2xl`
  }
}

// export const getNetworkShadow = (chainColor: string): string => {
//   if (chainColor === 'eth') {
//     return `!shadow-blue-xl hover:!shadow-blue-2xl`
//   }
//   return `!shadow-${chainColor}-xl hover:!shadow-${chainColor}-2xl`
// }

// export const getNetworkHover = (chainColor: string): string => {
//   if (chainColor === 'yellow') {
//     return `hover:!bg-${CUSTOM_YELLOW} hover:!bg-opacity-20`
//   } else if (chainColor === 'eth') {
//     return `hover:!bg-${ETH_BASE} hover:!bg-opacity-20`
//   }
//   return `hover:!bg-${chainColor}-500 hover:!bg-${chainColor}-20`
// }

export const getNetworkHover = (chainColor: string): string => {
  switch (chainColor) {
    case ColorOptions.ETH:
      return `hover:!bg-[#5170ad] hover:bg-opacity-20`
    case ColorOptions.YELLOW:
      return `${CustomClasses.HOVER_CUSTOM_YELLOW_BG} hover:bg-opacity-20`
    case ColorOptions.GRAY:
      return `hover:!bg-gray-500 hover:bg-opacity-20`
    case ColorOptions.GREEN:
      return `hover:!bg-green-500 hover:bg-opacity-20`
    case ColorOptions.LIME:
      return `hover:!bg-lime-500 hover:bg-opacity-20`
    case ColorOptions.SKY:
      return `hover:!bg-sky-500 hover:bg-opacity-20`
    case ColorOptions.BLUE:
      return `hover:!bg-blue-500 hover:bg-opacity-20`
    case ColorOptions.ORANGE:
      return `hover:!bg-orange-500 hover:bg-opacity-20`
    case ColorOptions.PURPLE:
      return `hover:!bg-purple-500 hover:bg-opacity-20`
    case ColorOptions.INDIGO:
      return `hover:!bg-indigo-500 hover:bg-opacity-20`
    case ColorOptions.CYAN:
      return `hover:!bg-cyan-500 hover:bg-opacity-20`
    case ColorOptions.RED:
      return `hover:!bg-red-500 hover:bg-opacity-20`
    default:
      return `hover:!bg-gray-500 hover:bg-opacity-20`
  }
}
