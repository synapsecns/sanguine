import mimLogo from '@assets/icons/mim.svg'
import usdtLogo from '@assets/icons/usdt.svg'
import usdcLogo from '@assets/icons/usdc.svg'

import * as CHAINS from '../chains/master'
import { Token } from '../../utils/types/index'

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
})

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
  logo: usdcLogo,
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
  logo: usdtLogo,
  color: 'lime',
  description: `
    USDT mirrors the price of the U.S. dollar, issued by a Hong Kong-based company Tether.
    The token’s peg to the USD is achieved via maintaining a sum of dollars in reserves equal
    to the number of USDT in circulation.
    `,
  swapableType: 'USD',
  priorityRank: 1,
})
