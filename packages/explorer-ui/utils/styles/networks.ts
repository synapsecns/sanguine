import { CHAINS } from 'synapse-constants/dist'

const ChainId = CHAINS.ChainId

const ETH_CURRENCY_TEXT_CLASSNAME = 'text-[#5170ad] dark:text-[#78a5ff]'

const COLOR_NETWORK_MAP = {
  eth: [ChainId.ETH],
  gray: [ChainId.ARBITRUM],
  yellow: [ChainId.BSC],
  green: [],
  sky: [],
  teal: [ChainId.MOONBEAM, ChainId.METIS],
  blue: [ChainId.FANTOM, ChainId.CRONOS, ChainId.TERRA],
  orange: [ChainId.KLAYTN],
  purple: [ChainId.POLYGON, ChainId.MOONRIVER],
  indigo: [],
  cyan: [ChainId.HARMONY],
  lime: [ChainId.BOBA, ChainId.AURORA, ChainId.DFK],
  red: [ChainId.AVALANCHE, ChainId.OPTIMISM],
}

const NETWORK_COLORS = {}

for (const [colorName, chainIdArr] of Object.entries(COLOR_NETWORK_MAP)) {
  for (const someChainId of chainIdArr) {
    NETWORK_COLORS[someChainId] = colorName
  }
}

export { NETWORK_COLORS }

export function getNetworkCurrencyColor(chainId) {
  switch (parseInt(chainId)) {
    case ChainId.BSC:
      return 'text-[#ecae0b] dark:text-[#ecae0b]'
    case ChainId.ETH:
      return ETH_CURRENCY_TEXT_CLASSNAME
    case ChainId.POLYGON:
      return 'text-purple-500 dark:text-purple-500'
    case ChainId.FANTOM:
      return 'text-blue-500 dark:text-blue-500'
    case ChainId.CRONOS:
      return 'text-blue-500 dark:text-blue-500'
    case ChainId.BOBA:
      return ETH_CURRENCY_TEXT_CLASSNAME
    case ChainId.ARBITRUM:
      return ETH_CURRENCY_TEXT_CLASSNAME
    case ChainId.OPTIMISM:
      return ETH_CURRENCY_TEXT_CLASSNAME
    case ChainId.AURORA:
      return ETH_CURRENCY_TEXT_CLASSNAME
    case ChainId.AVALANCHE:
      return 'text-red-500 dark:text-red-500'
    case ChainId.DFK:
      return 'text-lime-500 dark:text-lime-500'
    case ChainId.HARMONY:
      return 'text-cyan-500 dark:text-cyan-500'
    case ChainId.MOONBEAM:
      return 'text-teal-500 dark:text-teal-500'
    case ChainId.METIS:
      return 'text-teal-500 dark:text-teal-500'
    case ChainId.KLAYTN:
      return 'text-orange-500 dark:text-orange-500'
    default:
      return ''
  }
}

export function getNetworkButtonBgClassName(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'bg-stone-800 hover:bg-stone-900 active:bg-[#3c3c44]'
    case 'eth':
      return 'bg-[#5170ad] hover:bg-[#3f4f8c] active:bg-[#314367]'
    case 'purple':
      return 'bg-purple-500 hover:bg-purple-600 active:bg-purple-700'
    case 'blue':
      return 'bg-blue-500 hover:bg-blue-600 active:bg-blue-700'
    case 'lime':
      return 'bg-lime-500 hover:bg-lime-600 active:bg-lime-700'
    case 'gray':
      return 'bg-gray-500 hover:bg-gray-600 active:bg-gray-700'
    case 'red':
      return 'bg-red-500 hover:bg-red-600 active:bg-red-700'
    case 'cyan':
      return 'bg-cyan-500 hover:bg-cyan-600 active:bg-cyan-700'
    case 'teal':
      return 'bg-teal-500 hover:bg-teal-600 active:bg-teal-700'
    case 'orange':
      return 'bg-orange-500 hover:bg-orange-600 active:bg-orange-700'
    default:
      return ''
  }
}

export function getNetworkButtonBgClassNameActive(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'dark:active:!bg-[#3c3c44] '
    case 'eth':
      return 'dark:active:!bg-[#314367] '
    case 'purple':
      return 'dark:active:!bg-purple-700 '
    case 'blue':
      return 'dark:active:!bg-blue-700 '
    case 'lime':
      return 'dark:active:!bg-lime-700 '
    case 'gray':
      return 'dark:active:!bg-gray-700 '
    case 'red':
      return 'dark:active:!bg-red-700 '
    case 'cyan':
      return 'dark:active:!bg-cyan-700 '
    case 'teal':
      return 'dark:active:!bg-teal-700 '
    case 'orange':
      return 'dark:active:!bg-orange-700 '
    default:
      return ''
  }
}

export function getNetworkButtonBorderHover(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'hover:!border-[#ecae0b]'
    case 'eth':
      return 'hover:!border-[#5170ad]'
    case 'purple':
      return 'hover:!border-purple-500'
    case 'blue':
      return 'hover:!border-blue-500'
    case 'lime':
      return 'hover:!border-lime-500'
    case 'gray':
      return 'hover:!border-gray-500'
    case 'red':
      return 'hover:!border-red-500'
    case 'cyan':
      return 'hover:!border-cyan-500'
    case 'teal':
      return 'hover:!border-teal-500'
    case 'orange':
      return 'hover:!border-orange-500'
    default:
      return ''
  }
}

export function getNetworkButtonBorderActive(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'focus:!border-[#ecae0b] active:!border-[#ecae0b]'
    case 'eth':
      return 'focus:!border-[#5170ad] active:!border-[#5170ad]'
    case 'purple':
      return 'focus:!border-purple-500 active:!border-purple-500'
    case 'blue':
      return 'focus:!border-blue-500 active:!border-blue-500'
    case 'lime':
      return 'focus:!border-lime-500 active:!border-lime-500'
    case 'gray':
      return 'focus:!border-gray-500 active:!border-gray-500'
    case 'red':
      return 'focus:!border-red-500 active:!border-red-500'
    case 'cyan':
      return 'focus:!border-cyan-500 active:!border-cyan-500'
    case 'teal':
      return 'focus:!border-teal-500 active:!border-teal-500'
    case 'orange':
      return 'focus:!border-orange-500 active:!border-orange-500'
    default:
      return ''
  }
}

export function getNetworkButtonBorder(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'border-[#ecae0b] dark:border-[#ecae0b]'
    case 'eth':
      return 'border-[#5170ad] dark:border-[#5170ad]'
    case 'purple':
      return 'border-purple-500 dark:border-purple-500'
    case 'blue':
      return 'border-blue-500 dark:border-blue-500'
    case 'lime':
      return 'border-lime-500 dark:border-lime-500'
    case 'gray':
      return 'border-gray-500 dark:border-gray-500'
    case 'red':
      return 'border-red-500 dark:border-red-500'
    case 'cyan':
      return 'border-cyan-500 dark:border-cyan-500'
    case 'teal':
      return 'border-teal-500 dark:border-teal-500'
    case 'orange':
      return 'border-orange-500 dark:border-orange-500'
    default:
      return ''
  }
}

export function getNetworkButtonBorderImportant(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return '!border-[#ecae0b] dark:!border-[#ecae0b]'
    case 'eth':
      return '!border-[#5170ad] dark:!border-[#5170ad]'
    case 'purple':
      return '!border-purple-500 dark:!border-purple-500'
    case 'blue':
      return '!border-blue-500 dark:!border-blue-500'
    case 'lime':
      return '!border-lime-500 dark:!border-lime-500'
    case 'gray':
      return '!border-gray-500 dark:!border-gray-500'
    case 'red':
      return '!border-red-500 dark:!border-red-500'
    case 'cyan':
      return '!border-cyan-500 dark:!border-cyan-500'
    case 'teal':
      return '!border-teal-500 dark:!border-teal-500'
    case 'orange':
      return '!border-orange-500 dark:!border-orange-500'
    default:
      return ''
  }
}

export function getNetworkTextColor(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'text-[#ecae0b] dark:text-[#ecae0b]'
    case 'eth':
      return 'text-[#5170ad] dark:text-[#78a5ff]'
    case 'purple':
      return 'text-purple-500 dark:text-purple-500'
    case 'blue':
      return 'text-blue-500 dark:text-blue-500'
    case 'lime':
      return 'text-lime-500 dark:text-lime-500'
    case 'gray':
      return 'text-gray-500 dark:text-gray-500'
    case 'red':
      return 'text-red-500 dark:text-red-500'
    case 'cyan':
      return 'text-cyan-500 dark:text-cyan-500'
    case 'teal':
      return 'text-teal-500 dark:text-teal-500'
    case 'orange':
      return 'text-orange-500 dark:text-orange-500'
    default:
      return ''
  }
}

export function getNetworkTextHoverColor(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'hover:text-[#ecae0b]'
    case 'eth':
      return 'hover:text-[#78a5ff]'
    case 'purple':
      return 'hover:text-purple-500'
    case 'blue':
      return 'hover:text-blue-500'
    case 'lime':
      return 'hover:text-lime-500'
    case 'gray':
      return 'hover:text-gray-500'
    case 'red':
      return 'hover:text-red-500'
    case 'cyan':
      return 'hover:text-cyan-500'
    case 'teal':
      return 'hover:text-teal-500'
    case 'orange':
      return 'hover:text-orange-500'
    default:
      return ''
  }
}

export function getNetworkLinkTextColor(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return '!text-gray-800 hover:!text-[#ecae0b] dark:!text-[#ecae0b] dark:hover:!text-[#ecae0b]'
    case 'eth':
      return '!text-[#5170ad] hover:!text-[#78a5ff] dark:hover:!text-[#78a5ff]'
    case 'purple':
      return '!text-purple-500 hover:!text-purple-600 dark:hover:!text-purple-500'
    case 'blue':
      return '!text-blue-500 hover:!text-blue-600 dark:hover:!text-blue-500'
    case 'lime':
      return '!text-lime-500 hover:!text-lime-600 dark:hover:!text-lime-500'
    case 'gray':
      return '!text-gray-500 hover:!text-gray-600 dark:hover:!text-gray-500'
    case 'red':
      return '!text-red-500 hover:!text-red-600 dark:hover:!text-red-500'
    case 'cyan':
      return '!text-cyan-500 hover:!text-red-600 dark:hover:!text-cyan-500'
    case 'teal':
      return '!text-teal-500 hover:!text-teal-600 dark:hover:!text-teal-500'
    case 'orange':
      return '!text-orange-500 hover:!text-orange-600 dark:hover:!text-orange-500'
    default:
      return ''
  }
}

export function getNetworkTextColorContrast(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'text-[#ecae0b]'
    case 'eth':
      return 'text-white'
    case 'purple':
      return 'text-white'
    case 'blue':
      return 'text-white'
    case 'lime':
      return 'text-white'
    case 'gray':
      return 'text-white'
    case 'red':
      return 'text-white'
    case 'cyan':
      return 'text-white'
    case 'teal':
      return 'text-white'
    case 'orange':
      return 'text-white'
    default:
      return ''
  }
}

export function getNetworkTextColorContrastHover(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'group-hover:text-[#ecae0b]'
    case 'eth':
      return 'group-hover:text-white'
    case 'purple':
      return 'group-hover:text-white'
    case 'blue':
      return 'group-hover:text-white'
    case 'lime':
      return 'group-hover:text-white'
    case 'gray':
      return 'group-hover:text-white'
    case 'red':
      return 'group-hover:text-white'
    case 'cyan':
      return 'group-hover:text-white'
    case 'teal':
      return 'group-hover:text-white'
    case 'orange':
      return 'group-hover:text-white'
    default:
      return ''
  }
}

export function getNetworkBgClassName(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'bg-stone-800'
    case 'eth':
      return 'bg-[#5170ad]'
    case 'purple':
      return 'bg-purple-500'
    case 'blue':
      return 'bg-blue-500'
    case 'lime':
      return 'bg-lime-500'
    case 'gray':
      return 'bg-gray-500'
    case 'red':
      return 'bg-red-500'
    case 'cyan':
      return 'bg-cyan-500'
    case 'teal':
      return 'bg-teal-500'
    case 'orange':
      return 'bg-orange-500'
    default:
      return ''
  }
}

export function getNetworkBgClassNameLightDark(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'bg-[#ecae0b] '
    case 'eth':
      return 'bg-[#5170ad] '
    case 'purple':
      return 'bg-purple-500 '
    case 'blue':
      return 'bg-blue-500 '
    case 'lime':
      return 'bg-lime-500 '
    case 'gray':
      return 'bg-gray-500 '
    case 'red':
      return 'bg-red-500 '
    case 'cyan':
      return 'bg-cyan-500 '
    case 'teal':
      return 'bg-teal-500 '
    case 'orange':
      return 'bg-orange-500 '
    default:
      return ''
  }
}

export function getNetworkShadow(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return ' !shadow-yellow-600 hover:shadow-yellow-500'
    case 'eth':
      return ' !shadow-blue-600 hover:shadow-blue-500'
    case 'purple':
      return ' !shadow-purple-600 hover:shadow-purple-500'
    case 'blue':
      return ' !shadow-blue-600 hover:shadow-blue-500'
    case 'lime':
      return ' !shadow-lime-600 hover:shadow-lime-500'
    case 'gray':
      return '  hover:shadow-2xl'
    case 'red':
      return ' !shadow-red-600 hover:shadow-red-500'
    case 'cyan':
      return ' !shadow-cyan-600 hover:shadow-cyan-500'
    case 'teal':
      return ' !shadow-teal-600 hover:shadow-teal-500'
    case 'orange':
      return ' !shadow-orange-500 hover:!shadow-orange-500'
    default:
      return ''
  }
}

export function getNetworkHoverShadow(chainId) {
  switch (NETWORK_COLORS[chainId]) {
    case 'yellow':
      return 'hover:!shadow-yellow-500'
    case 'eth':
      return 'hover:!shadow-blue-500'
    case 'purple':
      return 'hover:!shadow-purple-500'
    case 'blue':
      return 'hover:!shadow-blue-500'
    case 'lime':
      return 'hover:!shadow-lime-500'
    case 'gray':
      return 'hover:!shadow-500'
    case 'red':
      return 'hover:!shadow-red-500'
    case 'cyan':
      return 'hover:!shadow-cyan-500'
    case 'teal':
      return 'hover:!shadow-teal-500'
    case 'orange':
      return 'hover:!shadow-orange-500'
    default:
      return ''
  }
}
