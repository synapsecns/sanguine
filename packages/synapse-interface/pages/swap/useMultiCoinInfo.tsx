
/**
   * @param {Token} primaryCoin
   * @param {Token} secondaryCoin
   * @returns {{poolName: string, otherCoin:Token}}
   */
export function getInfoMultiCoin(primaryCoin, secondaryCoin, priorityRanking) {
  const { poolName, poolTokens } = priorityRanking[primaryCoin.symbol][0]


  const coinSymbolsInPool = poolTokens.map(i => i.symbol)

  let otherCoin

  if (coinSymbolsInPool.includes(secondaryCoin.symbol)) {
    otherCoin = secondaryCoin
  } else {
    otherCoin = poolTokens.filter(i => i.symbol != primaryCoin.symbol)[0]
  }

  return ({
    poolName,
    otherCoin
  })
}

