import { getEthPrice, getAvaxPrice } from '@utils/actions/getPrices'
import { getPoolTokenInfoArr, getTokenBalanceInfo } from '@utils/poolDataFuncs'
import { Address, fetchBalance, fetchToken } from '@wagmi/core'
import { Token, PoolData } from '@types'

import { getCorePoolData } from './getCorePoolData'

export const getBalanceData = async ({
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
        address: lpTokenAddress as Address,
        chainId,
      })
    )?.totalSupply?.value ?? 0n

  const tokens: Token[] = [...pool.poolTokens, pool]
  for (const token of tokens) {
    const isLP = token.addresses[chainId] === lpTokenAddress

    const rawBalanceResult = await fetchBalance({
      address: address as Address,
      chainId,
      token: token.addresses[chainId] as Address,
    })

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

  const { tokenBalances } = await getBalanceData({
    pool,
    chainId,
    address: poolAddress,
    lpTokenAddress,
  })

  const { swapFee, virtualPrice } = await getCorePoolData(poolAddress, chainId)

  const ethPrice = prices?.ethPrice ?? (await getEthPrice())
  const avaxPrice = prices?.avaxPrice ?? (await getAvaxPrice())

  const {
    tokenBalancesSum,
    tokenBalancesUSD,
  }: { tokenBalancesSum: number; tokenBalancesUSD: number } =
    getTokenBalanceInfo({
      tokenBalances: tokenBalances.filter((t) => !t.isLP).map((t) => t.balance),
      prices: {
        ethPrice,
        avaxPrice,
      },
      poolType: pool?.poolType,
    })

  const poolTokensMatured = getPoolTokenInfoArr({
    tokenBalances: tokenBalances.filter((t) => !t.isLP),
    tokenBalancesSum,
  })

  return {
    name: pool.name,
    tokens: poolTokensMatured,
    totalLocked: tokenBalancesSum,
    totalLockedUSD: tokenBalancesUSD,
    virtualPrice,
    swapFee,
  }
}
