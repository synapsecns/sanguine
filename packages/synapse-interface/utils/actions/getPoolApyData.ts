import { formatUnits } from '@ethersproject/units'
import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/sushiMaster'
import { MINICHEF_ADDRESSES } from '@constants/minichef'
import { Token } from '@types'
import { fetchBalance, readContracts, fetchToken, Address } from '@wagmi/core'
import { MINICHEF_ABI } from '@abis/miniChef'
import { getSynPrices } from '@utils/actions/getPrices'
// import { useContractReads } from 'wagmi'
// import Abi from 'viem'
// import { BigNumber } from 'ethers'
// import { Zero, One } from '@ethersproject/constants'

export const getPoolApyData = async (
  chainId: number,
  poolToken: Token,
  prices?: any
) => {
  if (!MINICHEF_ADDRESSES?.[chainId]) {
    console.log('no minichef address found for chainId', chainId)
    return {
      fullCompoundedAPY: 0,
      weeklyAPR: 0,
      yearlyAPRUnvested: 0,
    }
  }
  const minichefAddress: Address = MINICHEF_ADDRESSES[chainId]

  const data = await readContracts({
    contracts: [
      {
        address: minichefAddress,
        abi: MINICHEF_ABI,
        functionName: 'synapsePerSecond',
        chainId,
      },
      {
        address: minichefAddress,
        abi: MINICHEF_ABI,
        functionName: 'totalAllocPoint',
        chainId,
      },
      {
        address: minichefAddress,
        abi: MINICHEF_ABI,
        functionName: 'poolInfo',
        chainId,
        args: [poolToken.poolId[chainId]],
      },
    ],
  })
  const synapsePerSecondResult: any = data[0].result
  const totalAllocPointsResult: any = data[1].result
  const poolInfoResult: any = data[2].result ?? []

  const lpTokenBalanceResult =
    (
      await fetchBalance({
        address: minichefAddress,
        chainId,
        token: poolToken.addresses[chainId] as Address,
      })
    )?.value ?? BigInt(0)

  const lpTokenSupplyResult =
    (
      await fetchToken({
        address: poolToken.addresses[chainId] as Address,
        chainId,
      })
    )?.totalSupply?.value ?? BigInt(0)

  const synPriceData = prices?.synPrices ?? (await getSynPrices())

  const synapsePerSecond: bigint = BigInt(synapsePerSecondResult ?? 0)
  const totalAllocPoints: bigint = BigInt(totalAllocPointsResult ?? 1)
  const allocPoints: bigint = BigInt(poolInfoResult?.allocPoint ?? 1)
  const lpTokenBalance: bigint = BigInt(lpTokenBalanceResult ?? 0)
  const lpTokenSupply: bigint = BigInt(lpTokenSupplyResult ?? 0)

  let rewardsPerWeek
  try {
    rewardsPerWeek = Number(formatUnits(synapsePerSecond, 'ether')) * 604800
  } catch (e) {
    rewardsPerWeek = 0
  }

  const poolRewardsPerWeek =
    (Number(allocPoints) / Number(totalAllocPoints)) * rewardsPerWeek

  if (poolRewardsPerWeek === 0) {
    return {}
  }

  const synValueInUsd = synPriceData.synBalanceNumber * synPriceData.synPrice
  const ethValueInUsd = synPriceData.ethBalanceNumber * synPriceData.ethPrice
  const lpTokenSupplyNumber = Number(
    formatUnits(BigInt(lpTokenSupply), 'ether')
  )
  const lpTokenBalanceNumber = Number(
    formatUnits(BigInt(lpTokenBalance), 'ether')
  )

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
  console.log(`usd per week`, usdPerWeek)

  const weeklyAPR = (usdPerWeek / stakedTvl) * 100
  const yearlyAPR = weeklyAPR * 52
  const decimalAPR = yearlyAPR / 100
  const yearlyCompoundedAPR = 100 * ((1 + decimalAPR / 365) ** 365 - 1)
  const fullCompoundedAPY = Math.round(yearlyCompoundedAPR * 100) / 100
  const fullCompoundedAPYStr = isFinite(fullCompoundedAPY)
    ? fullCompoundedAPY.toFixed(2)
    : '-'

  return {
    fullCompoundedAPY,
    fullCompoundedAPYStr,
    weeklyAPR: Math.round(weeklyAPR * 100) / 100,
    yearlyAPRUnvested: Math.round(yearlyAPR * 100) / 100,
  }
}
