import { Zero, One } from '@ethersproject/constants'
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
import { PoolTokenObject, Token, PoolUserData, PoolData } from '@types'
import { BigNumber } from 'ethers'

import { getVirtualPrice } from './getPoolFee'

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
        address: lpTokenAddress as `0x${string}`,
        chainId,
      })
    )?.totalSupply?.value ?? Zero

  const tokens: Token[] = [...pool.poolTokens, pool]
  for (const token of tokens) {
    const isLP = token.addresses[chainId] === lpTokenAddress

    const rawBalance = (
      await fetchBalance({
        address: address as `0x${string}`,
        chainId,
        token: token.addresses[chainId] as `0x${string}`,
      })
    )?.value

    // TODO: this is to support virtual price calcs, which needs to get updated
    // as a contract call
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
): Promise<PoolData | PoolUserData> => {
  const poolAddress = pool?.swapAddresses[chainId]
  if (!poolAddress || !pool || (!address && user)) {
    return null
  }

  // TODO: Check if we even need sdk call here since lp token is hardcoded
  // const lpTokenAddress =
  //   (await SynapseSDK.getPoolInfo(chainId, poolAddress))?.lpToken ??
  //   pool?.addresses[chainId]

  const lpTokenAddress = pool?.addresses[chainId]

  const { tokenBalances, lpTokenBalance, lpTotalSupply } = await getBalanceData(
    {
      pool,
      chainId,
      address: user ? address : poolAddress,
      lpTokenAddress,
    }
  )

  const virtualPrice = await getVirtualPrice(poolAddress, chainId)

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
      nativeTokens: pool.nativeTokens,
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
