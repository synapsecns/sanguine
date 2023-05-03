import { Chain } from '@types'

const CUSTOM_YELLOW = '[#ecae0b]'
const CUSTOM_YELLOW_DARK = '[#3c3c44]'
const ETH_BASE = '[#5170ad]'
const ETH_DARK = '[#3f4f8c]'
const ETH_EXTRA_DARK = '[#314367]'
const ETH_LIGHT = '[#78a5ff]'

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

export const getNetworkTextColorContrast = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `text-${CUSTOM_YELLOW}`
  }
  return 'text-white'
}

export const getNetworkTextColorContrastHover = (
  chainColor: string
): string => {
  if (chainColor === 'yellow') {
    return `group-hover:text-${CUSTOM_YELLOW}`
  }
  return 'group-hover:text-white'
}

export const getNetworkBgClassName = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return 'bg-stone-800'
  } else if (chainColor === 'eth') {
    return `bg-${ETH_BASE}`
  }
  return `bg-${chainColor}-500`
}

export const getNetworkBgClassNameLightDark = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `bg-${CUSTOM_YELLOW}`
  } else if (chainColor === 'eth') {
    return `bg-${ETH_BASE}`
  }
  return `bg-${chainColor}-500`
}

export const getNetworkShadow = (chainColor: string): string => {
  if (chainColor === 'eth') {
    return `!shadow-blue-xl hover:!shadow-blue-2xl`
  }
  return `!shadow-${chainColor}-xl hover:!shadow-${chainColor}-2xl`
}

export const getNetworkHover = (chainColor: string): string => {
  if (chainColor === 'yellow') {
    return `hover:!bg-${CUSTOM_YELLOW} hover:!bg-opacity-20`
  } else if (chainColor === 'eth') {
    return `hover:!bg-${ETH_BASE} hover:!bg-opacity-20`
  }
  return `hover:!bg-${chainColor}-500 hover:!bg-${chainColor}-20`
}
