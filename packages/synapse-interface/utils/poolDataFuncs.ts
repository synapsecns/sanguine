import { BigNumber } from '@ethersproject/bignumber'
import { Zero, One, AddressZero } from '@ethersproject/constants'
import { formatBNToPercentString } from '@bignumber/format'

export const MAX_BN_POW = BigNumber.from(10).pow(18)

export const getPriceMultiplier = ({ poolType, prices }) => {
  switch (poolType) {
    case 'ETH':
      return prices.ethPrice
    case 'AVAX':
      return prices.avaxPrice
    default:
      return 1
  }
}

/**
 * @param {BigNumber[]} arr
 */
export const calcBnSum = (arr) => {
  return arr.reduce((sum, b) => sum.add(b), Zero)
}

/**
 * @param {BigNumber} lpb
 */
export const calcIfZero = (lpb) => {
  if (lpb.isZero()) {
    return One
  } else {
    return lpb
  }
}

export const getBalanceInfo = async ({ lpTokenContract, account }) => {
  const arr = Promise.all([
    lpTokenContract.balanceOf(account || AddressZero),
    lpTokenContract.totalSupply(),
  ])

  return arr
}

export const getTokenBalanceInfo = ({ tokenBalances, poolType, ...rest }) => {
  const tokenBalancesSum = calcBnSum(tokenBalances)
  const priceMultiplier = getPriceMultiplier({ ...rest, poolType })
  const tokenBalancesUSD = tokenBalancesSum?.mul(priceMultiplier ?? 0)

  return {
    tokenBalancesSum,
    tokenBalancesUSD,
  }
}

export const getPoolTokenInfoArr = ({
  poolTokenBalances,
  tokenBalances,
  totalLpTokenBalance,
  tokenBalancesSum,
  poolTokens,
}) => {
  return poolTokens.map((token, i) => ({
    symbol: token.symbol,
    percent: formatBNToPercentString(
      tokenBalances[i]
        .mul(10 ** 5)
        .div(totalLpTokenBalance.isZero() ? One : tokenBalancesSum),
      5
    ),
    value: poolTokenBalances[i],
  }))
}
