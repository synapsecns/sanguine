import usdbLogo from '../assets/icons/usdc.svg'
import fusdtLogo from '../assets/icons/usdt.svg'
import { Token } from '../types'

import * as CHAINS from '../chains/master'

export const USDB = new Token({
  addresses: {
    [CHAINS.BNB.id]: '0xc8699abbba90c7479dedccef19ef78969a2fc608',
  },
  decimals: 18,
  symbol: 'USDB',
  name: 'USDB',
  logo: usdbLogo,
  docUrl: '',
  swapableType: 'USDB',
  priorityRank: 6,
  routeSymbol: 'USDB',
})

export const FUSDT = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x049d68029688eabf473097a2fc38ef61633a3c7a',
  },
  decimals: {
    [CHAINS.FANTOM.id]: 6,
  },
  symbol: 'fUSDT',
  name: 'Frapped USDT',
  logo: fusdtLogo,
  color: 'lime',
  swapableType: 'USD',
  swapableOn: [CHAINS.FANTOM.id],
  visibilityRank: 100,
  priorityRank: 3,
  routeSymbol: 'fUSDT',
})
