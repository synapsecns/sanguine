import { Zero, One } from '@ethersproject/constants'
// import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useGenericPoolApyData } from '@hooks/pools/useGenericPoolApyData'
import { useEthPrice, useAvaxPrice } from '@hooks/usePrices'
import {
  calcBnSum,
  calcIfZero,
  getTokenBalanceInfo,
  getPoolTokenInfoArr,
  MAX_BN_POW,
} from '@utils/poolDataFuncs'
import { STAKING_MAP_TOKENS } from '@constants/tokens'
import { fetchBalance } from '@wagmi/core'
import { PoolTokenObject } from '@types'
import { BigNumber } from 'ethers'

const getBalanceData = async ({
  poolAddress,
  poolToken,
  chainId,
  address,
  lpTokenAddress,
}) => {
  const tokenBalances: PoolTokenObject[] = []
  const userTokenBalances: PoolTokenObject[] = []
  let poolTokenSum = Zero
  let userLpTokenBalance = One
  let lpTokenBalance = One
  for (const token of poolToken?.poolTokens) {
    const isLP = token.addresses[chainId] === lpTokenAddress

    const rawBalance =
      (
        await fetchBalance({
          address: poolAddress,
          chainId,
          token: token.addresses[chainId],
        })
      )?.value ?? Zero

    const rawUserBalance =
      (
        await fetchBalance({
          address,
          chainId,
          token: token.addresses[chainId],
        })
      )?.value ?? Zero
    const balance =
      rawBalance?.mul(BigNumber.from(10)?.pow(18 - token.decimals[chainId])) ??
      Zero
    const userBalance =
      rawUserBalance?.mul(
        BigNumber.from(10)?.pow(18 - token.decimals[chainId])
      ) ?? Zero

    // add to balances
    tokenBalances.push({
      rawBalance,
      balance,
      token,
      isLP,
    })

    // add to user balances
    userTokenBalances.push({
      rawBalance: rawUserBalance,
      balance: userBalance,
      token,
      isLP,
    })

    // set lp variables
    if (isLP) {
      userLpTokenBalance = userBalance
      lpTokenBalance = balance
    }

    // running sum of all tokens in the pool
    if (balance) {
      poolTokenSum = poolTokenSum.add(balance)
    }
  }
  return {
    tokenBalances,
    userTokenBalances,
    poolTokenSum,
    userLpTokenBalance,
    lpTokenBalance,
  }
}
export const useGenericPoolData = async (
  chainId,
  poolName,
  address,
  SynapseSDK
) => {
  const poolToken = STAKING_MAP_TOKENS?.[chainId]?.[poolName]
  const poolTokenAddress = poolToken?.addresses[chainId]
  const poolAddress = poolToken?.swapAddresses[chainId]
  if (!poolTokenAddress || !poolAddress) {
    return null
  }
  // get LP token
  const lpTokenAddress = (await SynapseSDK.getPoolInfo(chainId, poolAddress))
    ?.lpToken
  const {
    tokenBalances,
    userTokenBalances,
    poolTokenSum,
    userLpTokenBalance,
    lpTokenBalance,
  } = await getBalanceData({
    poolAddress,
    poolToken,
    chainId,
    address,
    lpTokenAddress,
  })

  const virtualPriceResult = poolTokenSum?.div(lpTokenBalance)
  const ethPrice = await useEthPrice()
  const avaxPrice = await useAvaxPrice()

  let virtualPrice
  if (lpTokenBalance.isZero()) {
    virtualPrice = MAX_BN_POW
  } else {
    virtualPrice = virtualPriceResult?.[0]
  }

  const { tokenBalancesSum, tokenBalancesUSD } = getTokenBalanceInfo({
    tokenBalances: tokenBalances.map((t) => t.balance),
    prices: {
      ethPrice,
      avaxPrice,
    },
    poolType: poolToken?.poolType,
  })

  // User share data
  const userShare = userLpTokenBalance
    .mul(MAX_BN_POW)
    .div(calcIfZero(lpTokenBalance))

  const userPoolTokenBalances = tokenBalances.map((token) =>
    userShare.mul(token.balance).div(MAX_BN_POW)
  )

  const userPoolTokenBalancesSum = calcBnSum(userPoolTokenBalances)

  const sharedArgs = {
    lpTokenBalance,
    tokenBalancesSum,
  }

  const generalPoolTokens = getPoolTokenInfoArr({
    tokenBalances,
    ...sharedArgs,
  })

  const userPoolTokens = getPoolTokenInfoArr({
    tokenBalances: userTokenBalances,
    ...sharedArgs,
  })

  const poolApyData = (await useGenericPoolApyData(chainId, poolToken)) ?? {}

  const poolDataObj = {
    name: poolName,
    tokens: generalPoolTokens,
    totalLocked: tokenBalancesSum,
    totalLockedUSD: tokenBalancesUSD,
    virtualPrice,
    volume: 'XXX', // TODO
    utilization: 'XXX', // TODO
    apy: poolApyData, //? DIFF apidata
  }

  let userShareData
  if (address) {
    userShareData = {
      name: poolName,
      share: userShare,
      value: userPoolTokenBalancesSum,
      avgBalance: userPoolTokenBalancesSum,
      tokens: userPoolTokens,
      lpTokenBalance: userLpTokenBalance,
      // the code was always doing this, i could not find out why
      lpTokenMinted: userLpTokenBalance,
    }
  } else {
    userShareData = null
  }
  return { poolDataObj, userShareData }

  // return [poolData, userPoolData]
}
