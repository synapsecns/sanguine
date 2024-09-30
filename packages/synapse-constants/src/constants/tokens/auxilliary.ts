import { Token } from '../../types'
import * as CHAINS from '../chains/master'

export const AVWETH = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21',
  },
  decimals: 18,
  symbol: 'AVWETH',
  name: 'Aave Wrapped ETH',
  logo: 'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/weth.19fa93ab.svg',
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
  logo: 'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/usdt.3c9cd2f8.svg',
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
  logo: 'https://docs.abracadabra.money/~gitbook/image?url=https%3A%2F%2F2388475231-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-legacy-files%2Fo%2Fassets%252F-Mc9U0yE30Tc9xb3mVGA%252F-McF-MUtKQnjF8iBPrv5%252F-McF8_XaAL4kyHrICXUw%252FMIM%2520Logo%2520PNG.png%3Falt%3Dmedia%26token%3D12b05d37-765f-494d-809b-4deab52bd212&width=768&dpr=2&quality=100&sign=7d2ede98&sv=1',
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
  logo: 'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/avax.9d53cbf0.svg',
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
  logo: 'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/usdc.d5dcb030.svg',
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
  logo: 'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/usdt.3c9cd2f8.svg',
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
  logo: 'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/eth.b3692688.svg',
  isNative: true,
  swapableType: 'ETH',
  color: 'sky',
  visibilityRank: 101,
  priorityRank: 2,
  swapableOn: [],
  routeSymbol: 'ETH',
})
