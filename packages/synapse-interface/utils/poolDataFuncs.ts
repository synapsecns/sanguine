const getPriceMultiplier = ({ poolType, prices }) => {
  switch (poolType) {
    case 'ETH':
      return prices.ethPrice
    case 'AVAX':
      return prices.avaxPrice
    default:
      return 1
  }
}

export const getTokenBalanceInfo = ({ tokenBalances, poolType, prices }) => {
  const tokenBalancesSum = tokenBalances.reduce(
    (a, b) => Number(a) + Number(b),
    0
  )

  const priceMultiplier = getPriceMultiplier({ prices, poolType })
  const tokenBalancesUSD = tokenBalancesSum * (priceMultiplier ?? 0)

  return {
    tokenBalancesSum,
    tokenBalancesUSD,
  }
}

export const getPoolTokenInfoArr = ({
  tokenBalances,
  tokenBalancesSum,
}: {
  tokenBalances: {
    rawBalance: bigint
    balance: string
    token: any
    isLP: boolean
  }[]
  tokenBalancesSum: number
}) => {
  return tokenBalances.map((poolToken) => {
    const {
      balance,
      token,
      token: { symbol },
      isLP,
      rawBalance,
    } = poolToken

    const rawPercent = Number(balance) / tokenBalancesSum
    const percent =
      tokenBalancesSum !== 0 ? `${(100 * rawPercent).toFixed(2)}%` : '-'

    return {
      symbol,
      percent,
      balance,
      balanceStr: balance,
      token,
      isLP,
      rawBalance,
    }
  })
}
