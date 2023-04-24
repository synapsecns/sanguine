import { BigNumber } from '@ethersproject/bignumber'
import { Zero, One, AddressZero } from '@ethersproject/constants'
import { formatBNToPercentString } from '@bignumber/format'
import { PoolTokenObject } from '@types'

export const MAX_BN_POW = BigNumber.from(10).pow(18)

export const getPriceMultiplier = ({ poolType, prices }) => {
  console.log('poolType', poolType, prices)
  switch (poolType) {
    case 'ETH':
      return prices.ethPrice
    case 'AVAX':
      return prices.avaxPrice
    default:
      return 1
  }
}

export const calcBnSum = (arr) => {
  return arr.reduce((sum, b) => sum.add(b), Zero)
}

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

export const getTokenBalanceInfo = ({ tokenBalances, poolType, prices }) => {
  const tokenBalancesSum = calcBnSum(tokenBalances)
  const priceMultiplier = getPriceMultiplier({ prices, poolType })
  console.log('priceMultiplier', tokenBalancesSum, priceMultiplier)
  const tokenBalancesUSD = tokenBalancesSum?.mul(priceMultiplier ?? 0)

  return {
    tokenBalancesSum,
    tokenBalancesUSD,
  }
}

export const getPoolTokenInfoArr = ({
  tokenBalances,
  lpTokenBalance,
  tokenBalancesSum,
}: {
  tokenBalances: PoolTokenObject[]
  lpTokenBalance: BigNumber
  tokenBalancesSum: BigNumber
}) => {
  return tokenBalances.map((poolToken) => ({
    symbol: poolToken.token.symbol,
    percent: formatBNToPercentString(
      poolToken.balance
        .mul(10 ** 5)
        .div(lpTokenBalance.isZero() ? One : tokenBalancesSum),
      5
    ),
    value: poolToken.balance,
  }))
}
