import { erc20ABI, multicall } from '@wagmi/core'
import type { Token, PoolData } from '@types'
import { formatBigIntToString } from '@utils/bigint/format'
import { getPoolTokenInfoArr, getTokenBalanceInfo } from '@utils/poolDataFuncs'
import { getEthPrice, getAvaxPrice } from '@utils/actions/getPrices'

import lpTokenABI from '@/constants/abis/lpToken.json'
import { SWAP_ABI } from '@/constants/abis/swap'

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
  const tokens: Token[] = [...pool?.poolTokens, pool]
  const tokenBalances = []
  let poolTokenSum = 0n
  let lpTokenBalance = 1n

  const multicallInputs = []

  const one = {
    address: lpTokenAddress,
    abi: lpTokenABI,
    functionName: 'totalSupply',
    chainId,
  }

  multicallInputs.push(one)

  tokens?.forEach((token, index) => {
    const isLP = token.addresses[chainId] === lpTokenAddress
    // Use pool's getTokenBalance for pool tokens, if the address is the pool itself
    // to exclude the unclaimed admin fees
    if (address === pool.swapAddresses[chainId] && !isLP) {
      multicallInputs.push({
        address: pool.swapAddresses[chainId],
        abi: SWAP_ABI,
        functionName: 'getTokenBalance',
        chainId,
        args: [index],
      })
    } else {
      multicallInputs.push({
        address: token.addresses[chainId],
        abi: erc20ABI,
        functionName: 'balanceOf',
        chainId,
        args: [address],
      })
    }
  })
  const two = {
    address: pool.swapAddresses[chainId],
    abi: SWAP_ABI,
    functionName: 'swapStorage',
    chainId,
  }

  const three = {
    address: pool.swapAddresses[chainId],
    abi: SWAP_ABI,
    functionName: 'getVirtualPrice',
    chainId,
  }

  multicallInputs.push(two)
  multicallInputs.push(three)

  const multicallData: any[] = await multicall({
    contracts: multicallInputs,
    chainId,
  }).catch((error) => {
    console.error('Multicall failed:', error)
    return []
  })

  const lpTotalSupply = multicallData[0].result ?? 0n

  tokens.forEach((token, index) => {
    const isLP = token.addresses[chainId] === lpTokenAddress

    tokenBalances.push({
      rawBalance: multicallData[index + 1].result ?? 0n,
      balance: formatBigIntToString(
        multicallData[index + 1].result,
        token.decimals[chainId]
      ),
      token,
      isLP,
    })

    if (isLP) {
      lpTokenBalance = multicallData[index + 1].result
    }

    poolTokenSum = poolTokenSum + multicallData[index + 1].result
  })

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
  }
}
