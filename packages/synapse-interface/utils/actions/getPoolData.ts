// import { Zero, One } from '@ethersproject/constants'
import { getEthPrice, getAvaxPrice } from '@utils/actions/getPrices'
// import {
//   commifyBnToString,
//   commifyBnWithDefault,
//   formatBNToString,
// } from '@bignumber/format'
import {
  // calcBnSum,
  // calcIfZero,
  getTokenBalanceInfo,
  // getPoolTokenInfoArr,
  // MAX_BN_POW,
} from '@utils/poolDataFuncs'
import { fetchBalance, fetchToken } from '@wagmi/core'
import { Token, PoolUserData, PoolData } from '@types'
// import { BigNumber } from 'ethers'

import { getVirtualPrice } from './getPoolFee'
// import { formatBigIntToString } from '@/utils/bigint/format'
import { commifyBigIntToString } from '@utils/bigint/format'

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
  const tokenBalances = []
  let poolTokenSum = 0n
  let lpTokenBalance = 1n
  const lpTotalSupply =
    (
      await fetchToken({
        address: lpTokenAddress as `0x${string}`,
        chainId,
      })
    )?.totalSupply?.value ?? 0n

  const tokens: Token[] = [...pool.poolTokens, pool]
  for (const token of tokens) {
    const isLP = token.addresses[chainId] === lpTokenAddress

    const rawBalanceResult = await fetchBalance({
      address: address as `0x${string}`,
      chainId,
      token: token.addresses[chainId] as `0x${string}`,
    })
    // console.log(rawBalanceResult?.value)
    // add to balances
    tokenBalances.push({
      rawBalance: rawBalanceResult?.value ?? 0n,
      balance: rawBalanceResult?.formatted ?? '0',
      token,
      isLP,
    })

    // set lp variables
    if (isLP) {
      lpTokenBalance = rawBalanceResult?.value
      continue
    }
    // running sum of all tokens in the pool
    if (rawBalanceResult?.formatted) {
      poolTokenSum =
        poolTokenSum + BigInt(Math.round(Number(rawBalanceResult?.formatted)))
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
  // const poolTokensMatured = getPoolTokenInfoArr({
  //   tokenBalances: tokenBalances.filter((t) => !t.isLP),
  //   ...{
  //     lpTotalSupply,
  //     tokenBalancesSum,
  //   },
  //   chainId,
  // })
  const poolTokensMatured = tokenBalances.filter((token) => !token.isLP)
  if (user) {
    // const MAX_BN_POW_BIGINT = 1000000000000000000n;
    // const power = 18n;
    // const base = 10n;
    // console.log("ebfore erorr")
    // console.log(base ** power);

    // const userShare = (lpTokenBalance * MAX_BN_POW_BIGINT) / (lpTokenBalance === 0n ? 1n : lpTokenBalance);
    // const userPoolTokenBalances = tokenBalances.map((token) => (userShare * token.rawBalance) / MAX_BN_POW_BIGINT);
    // const userPoolTokenBalancesSum = userPoolTokenBalances.reduce((sum, b) => sum + b, 0n);

    return {
      name: pool.name,
      share: 0n,
      value: 0n,
      tokens: poolTokensMatured,
      lpTokenBalance,
      // lpTokenBalanceStr: formatBigIntToString(lpTokenBalance, 18, 4),
      lpTokenBalanceStr: lpTokenBalance.toString(),
      nativeTokens: pool.nativeTokens,
    }
  }

  // const standardUnits = pool.priceUnits ?? ''
  // const displayDecimals = standardUnits === 'ETH' ? 3 : 0

  return {
    name: pool.name,
    tokens: poolTokensMatured,
    totalLocked: tokenBalancesSum,
    // totalLockedStr: commifyBnWithDefault(tokenBalancesSum, displayDecimals),
    totalLockedStr: tokenBalancesSum,
    totalLockedUSD: tokenBalancesUSD,
    // totalLockedUSDStr: commifyBnToString(tokenBalancesUSD, 0),
    totalLockedUSDStr: tokenBalancesUSD,
    virtualPrice,
    virtualPriceStr: commifyBigIntToString(virtualPrice.result, 18, 5),
  }
}

export const getSinglePoolData = async (
  chainId: number,
  pool: Token,
  prices?: any
): Promise<PoolData> => {
  const poolAddress = pool?.swapAddresses[chainId]

  if (!pool || !poolAddress) {
    return null
  }

  const lpTokenAddress = pool?.addresses[chainId]

  const { tokenBalances, lpTokenBalance, lpTotalSupply } = await getBalanceData(
    {
      pool,
      chainId,
      address: poolAddress,
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

  // const poolTokensMatured = getPoolTokenInfoArr({
  //   tokenBalances: tokenBalances.filter((t) => !t.isLP),
  //   ...{
  //     lpTotalSupply,
  //     tokenBalancesSum,
  //   },
  //   chainId,
  // })

  // need to filter out lp tokens somewhere for display
  const poolTokensMatured = tokenBalances

  // console.log(`poolTokenMatured`, poolTokensMatured)
  // const standardUnits = pool.priceUnits ?? ''
  // const displayDecimals = standardUnits === 'ETH' ? 3 : 0

  return {
    name: pool.name,
    tokens: poolTokensMatured,
    totalLocked: tokenBalancesSum,
    // totalLockedStr: commifyBnWithDefault(tokenBalancesSum, displayDecimals),
    totalLockedStr: tokenBalancesSum,
    totalLockedUSD: tokenBalancesUSD,
    // totalLockedUSDStr: commifyBnToString(tokenBalancesUSD, 0),
    totalLockedUSDStr: tokenBalancesUSD,
    virtualPrice,
    // virtualPriceStr: commifyBnToString(virtualPrice, 5),
    virtualPriceStr: virtualPrice.result.toString(),
  }
}

export const getPoolUserData = async (
  chainId: number,
  pool: Token,
  address: string,
  prices?: any
): Promise<PoolUserData> => {
  const poolAddress = pool?.swapAddresses[chainId]
  if (!poolAddress || !pool || !address) {
    return null
  }

  const lpTokenAddress = pool?.addresses[chainId]

  const { tokenBalances, lpTokenBalance, lpTotalSupply } = await getBalanceData(
    {
      pool,
      chainId,
      address: address,
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
  // const poolTokensMatured = getPoolTokenInfoArr({
  //   tokenBalances: tokenBalances.filter((t) => !t.isLP),
  //   ...{
  //     lpTotalSupply,
  //     tokenBalancesSum,
  //   },
  //   chainId,
  // })

  // need to filter out LP tokens somewhere for display
  const poolTokensMatured = tokenBalances
  // const MAX_BN_POW_BIGINT = 1000000000000000000n;
  // const power = 18n;
  // const base = 10n;
  // console.log("ebfore erorr")
  // console.log(base ** power);

  // const userShare = (lpTokenBalance * MAX_BN_POW_BIGINT) / (lpTokenBalance === 0n ? 1n : lpTokenBalance);
  // const userPoolTokenBalances = tokenBalances.map((token) => (userShare * token.rawBalance) / MAX_BN_POW_BIGINT);
  // const userPoolTokenBalancesSum = userPoolTokenBalances.reduce((sum, b) => sum + b, 0n);

  return {
    name: pool.name,
    share: 0n,
    value: 0n,
    tokens: poolTokensMatured,
    lpTokenBalance,
    // lpTokenBalanceStr: formatBigIntToString(lpTokenBalance, 18, 4),
    lpTokenBalanceStr: lpTokenBalance.toString(),
    nativeTokens: pool.nativeTokens,
  }
}
