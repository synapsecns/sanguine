import { ChainId } from '@/constants/chains'

const COLOR_NETWORK_MAP = {
  eth: [ChainId.ETH],
  gray: [ChainId.ARBITRUM],
  yellow: [ChainId.BSC],
  green: [],
  sky: [],
  teal: [ChainId.MOONBEAM, ChainId.METIS, ChainId.CANTO],
  blue: [ChainId.FANTOM, ChainId.CRONOS, ChainId.TERRA],
  orange: [ChainId.KLAYTN],
  purple: [ChainId.POLYGON, ChainId.MOONRIVER, ChainId.DOGECHAIN],
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

export const getNetworkButtonBorderHover = (chainId) => {
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
    case 'green':
      return 'hover:!border-green-500'
    default:
      return ''
  }
}
