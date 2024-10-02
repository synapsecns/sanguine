import { Token } from '../../types'
import * as CHAINS from '../chains/master'

export const SwapUSDC = new Token({
  visibilityRank: 1,
  addresses: {
    [CHAINS.AVALANCHE.id]: '0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e',
  },
  decimals: {
    [CHAINS.AVALANCHE.id]: 6,
  },
  symbol: 'USDC',
  name: 'USD Circle',
  logo: 'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/usdc.d5dcb030.svg',
  description: `
    USD Coin (known by its ticker USDC) is a stablecoin that is pegged to the
    U.S. dollar on a 1:1 basis. Every unit of this cryptocurrency in circulation
    is backed up by $1 that is held in reserve
    `,
  swapableType: 'USD',

  color: 'blue',
  priorityRank: 1,
})

export const SwapUSDT = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7',
  },

  decimals: {
    [CHAINS.AVALANCHE.id]: 6,
  },
  symbol: 'USDT',
  name: 'USD Tether',
  logo: 'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/usdt.3c9cd2f8.svg',
  color: 'lime',
  description: `
    USDT mirrors the price of the U.S. dollar, issued by a Hong Kong-based company Tether.
    The token’s peg to the USD is achieved via maintaining a sum of dollars in reserves equal
    to the number of USDT in circulation.
    `,
  swapableType: 'USD',
  priorityRank: 1,
})
