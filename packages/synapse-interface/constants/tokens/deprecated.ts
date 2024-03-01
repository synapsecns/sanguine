import fusdtLogo from '@assets/icons/usdt.svg'

import { Token } from '@/utils/types'
import * as CHAINS from '@/constants/chains/master'

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
