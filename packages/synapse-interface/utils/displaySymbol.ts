import { SYNJEWEL } from '@constants/tokens/bridgeable'
import * as CHAINS from '@constants/chains/master'

import { Token } from '@/utils/types'
/**
 * In our `constants/tokens/master.tsx` file, we combine multiple stablecoin
 * contracts like USDC/USDT into multiple networks.
 * However, on chains like Avalanche, their contract addresses correspond to
 * other symbols (USDC.e/USDT.e). This helper function decorates a symbol
 * according to what it should reflect. We do this so we don't break any
 * downstream logic from tokens.
 *
 *
 * CHECK IF CAN DELETE ONCE SWAP IMPLEMENTED
 */

export const displaySymbol = (chainId: number, token: Token) => {
  if (!token?.symbol) {
    return ''
  }
  if (token.symbol === SYNJEWEL.symbol) {
    return 'synJEWEL'
  } else if (chainId === CHAINS.AVALANCHE.id) {
    switch (token.symbol) {
      // case 'USDC':
      //   return 'USDC.e'
      // case 'USDT':
      //   return 'USDT.e'
      case 'DAI':
        return 'DAI.e'
      default:
        return token.symbol
    }
  } else {
    return token.symbol
  }
}
