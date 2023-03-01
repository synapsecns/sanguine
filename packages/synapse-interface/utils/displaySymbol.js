import { SYNJEWEL } from '@constants/tokens/basic'
import { ChainId } from '@constants/networks'

/**
 * @param {Token} token
 * In our `constants/tokens/basic.js` file, we combine multiple stablecoin
 * contracts like USDC/USDT into multiple networks.
 * However, on chains like Avalanche, their contract addresses correspond to
 * other symbols (USDC.e/USDT.e). This helper function decorates a symbol
 * according to what it should reflect. We do this so we don't break any
 * downstream logic from tokens.
 */

export function displaySymbol(chainId, token) {
  if (token.symbol === SYNJEWEL.symbol) {
    return 'synJEWEL'
  } else if (chainId === ChainId.AVALANCHE) {
    switch (token.symbol) {
      case 'USDC':
        return 'USDC.e'
      case 'USDT':
        return 'USDT.e'
      case 'DAI':
        return 'DAI.e'
      default:
        return token.symbol
    }
  } else {
    return token.symbol
  }
}
