import { BigNumber } from '@ethersproject/bignumber'
import { Zero, One, AddressZero } from '@ethersproject/constants'
import { formatBNToPercentString, formatBNToString } from '@bignumber/format'
import { PoolTokenObject } from '@types'

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
  const tokenBalancesUSD = tokenBalancesSum?.mul(priceMultiplier ?? 0)

  return {
    tokenBalancesSum,
    tokenBalancesUSD,
  }
}

export const getPoolTokenInfoArr = ({
  tokenBalances,
  lpTotalSupply,
  tokenBalancesSum,
}: {
  tokenBalances: PoolTokenObject[]
  chainId: number
  lpTotalSupply: BigNumber
  tokenBalancesSum: BigNumber
}) => {
  return tokenBalances.map((poolToken) => ({
    symbol: poolToken.token.symbol,
    percent: poolToken.balance.isZero()
      ? '0'
      : formatBNToPercentString(
          poolToken.balance
            .mul(10 ** 5)
            .div(lpTotalSupply.isZero() ? One : tokenBalancesSum),
          5
        ),
    balance: poolToken.balance,
    balanceStr: formatBNToString(poolToken.balance, 18, 4),
    token: poolToken.token,
    isLp: poolToken.isLP,
    rawBalance: poolToken.rawBalance,
  }))
}
