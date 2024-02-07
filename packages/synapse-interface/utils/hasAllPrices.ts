export const hasAllPrices = (prices) => {
  if (
    prices.ethPrice === null ||
    prices.avaxPrice === null ||
    prices.metisPrice === null
  ) {
    return false
  }

  const synPrices = prices.synPrices
  if (
    synPrices.ethBalanceNumber === null ||
    synPrices.ethPrice === null ||
    synPrices.synBalanceNumber === null ||
    synPrices.synPrice === null
  ) {
    return false
  }

  return true
}
