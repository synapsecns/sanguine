import { WETHE, ETH } from '@constants/tokens/bridgeable'
import { AVWETH, WETH } from '@constants/tokens/auxilliary'

/**
 * @param {Token} matchCoin the coin in the pool/ token of poolTokens
 * @param {Token} compareCoin the coin to compare (fromCoin/toCoin)
 */
export const matchSymbolWithinPool = (matchCoin, compareCoin) => {
  let compareSymbol = compareCoin.symbol
  if (compareSymbol === WETHE.symbol) {
    compareSymbol = AVWETH.symbol
  } else if (compareSymbol === ETH.symbol) {
    compareSymbol = WETH.symbol
  }
  return matchCoin.symbol === compareSymbol
}
