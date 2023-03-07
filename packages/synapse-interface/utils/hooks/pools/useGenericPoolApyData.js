import { formatUnits } from '@ethersproject/units'

import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/lp'

import { MINICHEF_ADDRESSES } from '@constants/minichef'


import { useGenericMiniChefContract } from '@hooks/contracts/useMiniChefContract'
import { useGenericTokenContract } from '@hooks/contracts/useContract'

import { useSynPrices } from '@hooks/useSynPrices'

import {
  useSingleCallResult,
  useSingleContractMultipleMethods
} from '@hooks/multicall'




export function useGenericPoolApyData(chainId, poolToken) {
  const miniChefContract = useGenericMiniChefContract(chainId)
  const poolTokenContract = useGenericTokenContract(chainId, poolToken)

  const synPriceData = useSynPrices()


  const [
    synapsePerSecondResult,
    totalAllocPointsResult,
    poolInfoResult,
  ] = useSingleContractMultipleMethods(
    chainId,
    miniChefContract,
    {
      'synapsePerSecond': [],
      'totalAllocPoint': [],
      'poolInfo': [poolToken.poolId[chainId]]
    },
    { resultOnly: true }
  )

  const lpTokenBalanceResult = useSingleCallResult(
    chainId,
    poolTokenContract,
    'balanceOf',
    [MINICHEF_ADDRESSES[chainId]],
    { resultOnly: true }
  )

  const lpTokenSupplyResult = useSingleCallResult(
    chainId,
    poolTokenContract,
    'totalSupply',
    [],
    { resultOnly: true }
  )

  const synapsePerSecond = synapsePerSecondResult?.[0] ?? 0
  const totalAllocPoints = totalAllocPointsResult?.[0] ?? 1
  const allocPoints = poolInfoResult?.allocPoint ?? 1
  const lpTokenBalance = lpTokenBalanceResult?.balance ?? 0
  const lpTokenSupply = lpTokenSupplyResult?.[0] ?? 0

  let rewardsPerWeek
  try {
    rewardsPerWeek = Number(formatUnits(synapsePerSecond, 'ether')) * 604800
  } catch (e) {
    rewardsPerWeek = 0
  }

  const poolRewardsPerWeek = (allocPoints / totalAllocPoints) * rewardsPerWeek
  if (poolRewardsPerWeek == 0) {
    return
  }


  const synValueInUsd        = synPriceData.synBalanceNumber * synPriceData.synPrice
  const ethValueInUsd        = synPriceData.ethBalanceNumber * synPriceData.ethPrice
  const lpTokenSupplyNumber  = Number(formatUnits(lpTokenSupply, 'ether'))
  const lpTokenBalanceNumber = Number(formatUnits(lpTokenBalance, 'ether'))

  let stakedTvl
  if (SYN_ETH_SUSHI_TOKEN.symbol === poolToken.symbol) {
    const lpTokenUSDValue = (synValueInUsd + ethValueInUsd) / lpTokenSupplyNumber
    stakedTvl = lpTokenBalanceNumber * lpTokenUSDValue
  } else if (poolToken.poolType === "USD") {
    stakedTvl = lpTokenBalanceNumber
  } else if (poolToken.poolType === "ETH") {
    stakedTvl = lpTokenBalanceNumber * synPriceData.ethPrice
  } else {
    stakedTvl = 0
  }

  const usdPerWeek = poolRewardsPerWeek * synPriceData.synPrice

  const weeklyAPR = (usdPerWeek / stakedTvl) * 100
  const yearlyAPR = weeklyAPR * 52
  const decimalAPR = yearlyAPR / 100
  const yearlyCompoundedAPR = 100 * ((1 + decimalAPR / 365) ** 365 - 1)

  return {
    fullCompoundedAPY: _.round(yearlyCompoundedAPR, 2),
    weeklyAPR: _.round(weeklyAPR, 2),
    yearlyAPRUnvested: _.round(yearlyAPR, 2),
  }
}
