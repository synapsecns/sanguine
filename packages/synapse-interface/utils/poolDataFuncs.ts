import { BigNumber } from '@ethersproject/bignumber'
import { Zero, One, AddressZero } from '@ethersproject/constants'
import { formatBNToPercentString, formatBNToString } from '@bignumber/format'
import { formatBigIntToPercentString } from '@/utils/bigint/format'
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
  // const tokenBalancesSum = tokenBalances.reduce((a, b) => Number(a) + Number(b), 0)
  // const priceMultiplier = getPriceMultiplier({ prices, poolType })
  // const tokenBalancesUSD = tokenBalancesSum * (priceMultiplier ?? 0)

  return {
    tokenBalancesSum: 0,
    tokenBalancesUSD: 0,
  }
}

import { formatUnits } from '@ethersproject/units'

function formatBigIntUnits(value: bigint, decimals = 18) {
  const stringValue = value.toString();
  const decimalPointIndex = stringValue.length - decimals;

  if (decimalPointIndex <= 0) {
    return '0.' + stringValue.padStart(decimals, '0');
  }

  return (
    stringValue.slice(0, decimalPointIndex) +
    '.' +
    stringValue.slice(decimalPointIndex)
  );
}


export const getPoolTokenInfoArr = ({
  tokenBalances,
  lpTotalSupply,
  tokenBalancesSum,
}: {
  tokenBalances: { rawBalance: bigint, balance: string, token: any, isLP: boolean }[]
  chainId: number
  lpTotalSupply: bigint
  tokenBalancesSum: bigint
}) => {
  console.log(tokenBalances)
  return tokenBalances.map((poolToken) => ({
    symbol: poolToken.token.symbol,
    // percent: poolToken.rawBalance === 0n
    //   ? '0'
    //   : formatBigIntToPercentString(
    //       (poolToken.rawBalance * 10n ** 5n) /
    //       (lpTotalSupply === 0n ? 1n : tokenBalancesSum),
    //       5
    //     ),
    percent: 0,
    balance: poolToken.balance,
    balanceStr: poolToken.balance,
    token: poolToken.token,
    isLp: poolToken.isLP,
    rawBalance: poolToken.rawBalance,
  }))
}

