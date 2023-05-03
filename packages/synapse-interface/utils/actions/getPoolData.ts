import { Zero, One } from '@ethersproject/constants'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import { getEthPrice, getAvaxPrice } from '@utils/actions/getPrices'
import {
  commifyBnToString,
  commifyBnWithDefault,
  formatBNToString,
} from '@bignumber/format'
import {
  calcBnSum,
  calcIfZero,
  getTokenBalanceInfo,
  getPoolTokenInfoArr,
  MAX_BN_POW,
} from '@utils/poolDataFuncs'
import { fetchBalance, fetchToken } from '@wagmi/core'
import { PoolTokenObject, Token } from '@types'
import { BigNumber } from 'ethers'

const getBalanceData = async ({
  pool,
  chainId,
  address,
  lpTokenAddress,
}: {
  pool: Token
  chainId: number
  address: string
  lpTokenAddress: string
}) => {
  const tokenBalances: PoolTokenObject[] = []
  let poolTokenSum = Zero
  let lpTokenBalance = One
  const lpTotalSupply =
    (
      await fetchToken({
        address: `0x${lpTokenAddress.slice(2)}`,
        chainId,
      })
    )?.totalSupply?.value ?? Zero

  const tokens: Token[] = [...pool.poolTokens, pool]
  for (const token of tokens) {
    const isLP = token.addresses[chainId] === lpTokenAddress

    const rawBalance = (
      await fetchBalance({
        address: `0x${address.slice(2)}`,
        chainId,
        token: `0x${token.addresses[chainId].slice(2)}`,
      })
    )?.value

    const balance = rawBalance.mul(
      BigNumber.from(10).pow(18 - token.decimals[chainId])
    )

    // add to balances
    tokenBalances.push({
      rawBalance,
      balance,
      token,
      isLP,
    })

    // set lp variables
    if (isLP) {
      lpTokenBalance = balance
      continue
    }
    // running sum of all tokens in the pool
    if (balance) {
      poolTokenSum = poolTokenSum.add(balance)
    }
  }
  return {
    tokenBalances,
    poolTokenSum,
    lpTokenBalance,
    lpTotalSupply,
  }
}
export const getPoolData = async (
  chainId: number,
  pool: Token,
  address: string,
  user: boolean,
  prices?: any
) => {
  const poolAddress = pool?.swapAddresses[chainId]
  if (!poolAddress || !pool || (!address && user)) {
    return null
  }

  // TODO: Check if we even need contract call here since lp token is hardcoded
  // const lpTokenAddress =
  //   (await SynapseSDK.getPoolInfo(chainId, poolAddress))?.lpToken ??
  //   pool?.addresses[chainId]

  const lpTokenAddress = pool?.addresses[chainId]

  const { tokenBalances, poolTokenSum, lpTokenBalance, lpTotalSupply } =
    await getBalanceData({
      pool,
      chainId,
      address: user ? address : poolAddress,
      lpTokenAddress,
    })

  const virtualPrice = lpTotalSupply.isZero()
    ? MAX_BN_POW
    : calculateExchangeRate(lpTotalSupply, 18, poolTokenSum, 18)

  const ethPrice = prices?.ethPrice ?? (await getEthPrice())
  const avaxPrice = prices?.avaxPrice ?? (await getAvaxPrice())

  const { tokenBalancesSum, tokenBalancesUSD } = getTokenBalanceInfo({
    tokenBalances: tokenBalances.filter((t) => !t.isLP).map((t) => t.balance),
    prices: {
      ethPrice,
      avaxPrice,
    },
    poolType: pool?.poolType,
  })
  const poolTokensMatured = getPoolTokenInfoArr({
    tokenBalances: tokenBalances.filter((t) => !t.isLP),
    ...{
      lpTotalSupply,
      tokenBalancesSum,
    },
    chainId,
  })
  if (user) {
    const userShare = lpTokenBalance
      .mul(MAX_BN_POW)
      .div(calcIfZero(lpTokenBalance))
    const userPoolTokenBalances = tokenBalances.map((token) =>
      userShare.mul(token.balance).div(MAX_BN_POW)
    )
    const userPoolTokenBalancesSum = calcBnSum(userPoolTokenBalances)

    return {
      name: pool.name,
      share: userShare,
      value: userPoolTokenBalancesSum,
      tokens: poolTokensMatured,
      lpTokenBalance,
      lpTokenBalanceStr: formatBNToString(lpTokenBalance, 18, 4),
    }
  }

  const standardUnits = pool.priceUnits ?? ''
  const displayDecimals = standardUnits === 'ETH' ? 3 : 0
  return {
    name: pool.name,
    tokens: poolTokensMatured,
    totalLocked: tokenBalancesSum,
    totalLockedStr: commifyBnWithDefault(tokenBalancesSum, displayDecimals),
    totalLockedUSD: tokenBalancesUSD,
    totalLockedUSDStr: commifyBnToString(tokenBalancesUSD, 0),
    virtualPrice,
    virtualPriceStr: commifyBnToString(virtualPrice, 5),
  }
}
