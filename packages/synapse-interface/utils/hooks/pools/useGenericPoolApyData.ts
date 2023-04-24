import { formatUnits } from '@ethersproject/units'
import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/sushiMaster'
import { MINICHEF_ADDRESSES } from '@constants/minichef'
import { useSynPrices } from '@utils/hooks/useSynPrices'
import { Zero, One } from '@ethersproject/constants'
import { fetchBalance, fetchToken, readContract } from '@wagmi/core'
import MINICHEF_ABI from '@abis/miniChef.json'
import { BigNumber } from 'ethers'
export const useGenericPoolApyData = async (chainId, poolToken) => {
  if (!MINICHEF_ADDRESSES?.[chainId]) {
    console.log('no minichef address found for chainId', chainId)
    return {
      fullCompoundedAPY: 0,
      weeklyAPR: 0,
      yearlyAPRUnvested: 0,
    }
  }
  const minichefAddress: `0x${string}` = `0x${MINICHEF_ADDRESSES[chainId].slice(
    2
  )}`

  const synapsePerSecondResult: any = await readContract({
    address: minichefAddress,
    abi: MINICHEF_ABI,
    functionName: 'synapsePerSecond',
    chainId,
  })

  const totalAllocPointsResult: any = await readContract({
    address: minichefAddress,
    abi: MINICHEF_ABI,
    functionName: 'totalAllocPoint',
    chainId,
  })

  const poolInfoResult: any = await readContract({
    address: minichefAddress,
    abi: MINICHEF_ABI,
    functionName: 'poolInfo',
    chainId,
    args: [poolToken.poolId[chainId]],
  })

  const lpTokenBalanceResult =
    (
      await fetchBalance({
        address: minichefAddress,
        chainId,
        token: poolToken.addresses[chainId],
      })
    )?.value ?? Zero

  const lpTokenSupplyResult =
    (
      await fetchToken({
        address: poolToken.addresses[chainId],
        chainId,
      })
    )?.totalSupply?.value ?? Zero

  const synPriceData = await useSynPrices()

  const synapsePerSecond: BigNumber = synapsePerSecondResult ?? Zero
  const totalAllocPoints: BigNumber = totalAllocPointsResult ?? One
  const allocPoints: BigNumber = poolInfoResult?.allocPoint ?? One
  const lpTokenBalance: BigNumber = lpTokenBalanceResult ?? Zero
  const lpTokenSupply: BigNumber = lpTokenSupplyResult ?? Zero

  let rewardsPerWeek
  try {
    rewardsPerWeek = Number(formatUnits(synapsePerSecond, 'ether')) * 604800
  } catch (e) {
    rewardsPerWeek = 0
  }

  const poolRewardsPerWeek =
    allocPoints.div(totalAllocPoints).toNumber() * rewardsPerWeek
  if (poolRewardsPerWeek === 0) {
    console.log("poolRewardsPerWeek === 0, can't calculate APY", chainId)
    return {}
  }

  const synValueInUsd = synPriceData.synBalanceNumber * synPriceData.synPrice
  const ethValueInUsd = synPriceData.ethBalanceNumber * synPriceData.ethPrice
  const lpTokenSupplyNumber = Number(formatUnits(lpTokenSupply, 'ether'))
  const lpTokenBalanceNumber = Number(formatUnits(lpTokenBalance, 'ether'))

  let stakedTvl
  if (SYN_ETH_SUSHI_TOKEN.symbol === poolToken.symbol) {
    const lpTokenUSDValue =
      (synValueInUsd + ethValueInUsd) / lpTokenSupplyNumber
    stakedTvl = lpTokenBalanceNumber * lpTokenUSDValue
  } else if (poolToken.poolType === 'USD') {
    stakedTvl = lpTokenBalanceNumber
  } else if (poolToken.poolType === 'ETH') {
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
    fullCompoundedAPY: Math.round(yearlyCompoundedAPR * 100) / 100,
    weeklyAPR: Math.round(weeklyAPR * 100) / 100,
    yearlyAPRUnvested: Math.round(yearlyAPR * 100) / 100,
  }
}
