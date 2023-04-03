import { WETHE, AVWETH, ETH, WETH } from '@constants/tokens/basic'


/**
 * @param {Token} matchCoin the coin in the pool/ token of poolTokens
 * @param {Token} compareCoin the coin to compare (fromCoin/toCoin)
 */
export function matchSymbolWithinPool(matchCoin, compareCoin) {
  let compareSymbol = compareCoin.symbol
  if (compareSymbol === WETHE.symbol) {
    compareSymbol = AVWETH.symbol
  } else if (compareSymbol === ETH.symbol) {
    compareSymbol = WETH.symbol
  }
  return (matchCoin.symbol === compareSymbol)
}
