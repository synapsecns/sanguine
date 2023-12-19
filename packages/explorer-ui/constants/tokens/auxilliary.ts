import avaxLogo from '@assets/icons/avax.svg'
import avwethLogo from '@assets/icons/avweth.svg'
import ethLogo from '@assets/icons/eth.svg'
import mimLogo from '@assets/icons/mim.svg'
import usdcLogo from '@assets/icons/usdc.svg'
import usdtLogo from '@assets/icons/usdt.svg'

import { Token } from '../../utils/types/index'
import * as CHAINS from '../chains/master'

export const AVWETH = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21',
  },
  decimals: 18,
  symbol: 'AVWETH',
  name: 'Aave Wrapped ETH',
  logo: avwethLogo,
  swapableType: 'ETH',
  color: 'cyan',
  priorityRank: 2,
  routeSymbol: 'AVWETH',
})

export const KLAYTN_oUSDT = new Token({
  addresses: {
    [CHAINS.KLAYTN.id]: '0xceE8FAF64bB97a73bb51E115Aa89C17FfA8dD167',
  },
  decimals: {
    [CHAINS.KLAYTN.id]: 6,
  },
  symbol: 'orbitUSDT',
  name: 'Orbit Bridged USDT',
  logo: usdtLogo,
  swapableType: 'KLAYTN_USDT',
  swapableOn: [CHAINS.KLAYTN.id],
  priorityRank: 6,
  routeSymbol: 'KLAYTN_oUSDT',
})

export const MIM = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x82f0b8b456c1a451378467398982d4834b6829c1',
    [CHAINS.ARBITRUM.id]: '0xfea7a6a0b346362bf88a9e4a88416b77a57d6c2a',
  },
  decimals: 18,
  symbol: 'MIM',
  name: 'Magic Internet Money',
  logo: mimLogo,
  swapableType: 'USD',
  color: 'indigo',
  priorityRank: 6,
  routeSymbol: 'MIM',
})

export const MULTIAVAX = new Token({
  addresses: {
    [CHAINS.HARMONY.id]: '0xb12c13e66ade1f72f71834f2fc5082db8c091358',
  },
  decimals: 18,
  symbol: 'multiAVAX',
  name: 'AnySwap Wrapped AVAX',
  logo: avaxLogo,
  swapableType: 'AVAX',
  swapableOn: [CHAINS.HARMONY.id],
  color: 'red',
  priorityRank: 3,
  routeSymbol: 'MULTIAVAX',
})

export const FANTOMUSDC = new Token({
  visibilityRank: 101,
  addresses: {
    [CHAINS.FANTOM.id]: '0x04068da6c83afcfa0e13ba15a6696662335d5b75',
  },
  decimals: {
    [CHAINS.FANTOM.id]: 6,
  },
  symbol: 'USDC',
  name: 'USD Coin',
  logo: usdcLogo,
  swapableType: 'USD',
  swapableOn: [],
  color: 'blue',
  priorityRank: 1,
  routeSymbol: 'USDC',
})

export const FANTOMUSDT = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x049d68029688eabf473097a2fc38ef61633a3c7a',
  },
  swapExceptions: {},
  decimals: {
    [CHAINS.FANTOM.id]: 6,
  },
  symbol: 'USDT',
  name: 'USD Tether',
  logo: usdtLogo,
  color: 'lime',
  swapableType: 'USD',
  swapableOn: [],
  visibilityRank: 100,
  priorityRank: 1,
  routeSymbol: 'USDT',
})

export const FANTOMETH = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x74b23882a30290451A17c44f4F05243b6b58C76d',
  },
  decimals: 18,
  symbol: 'ETH',
  name: 'Ethereum',
  logo: ethLogo,
  isNative: true,
  swapableType: 'ETH',
  color: 'sky',
  visibilityRank: 101,
  priorityRank: 2,
  swapableOn: [],
  routeSymbol: 'ETH',
})
