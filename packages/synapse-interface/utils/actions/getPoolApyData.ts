import { formatUnits } from '@ethersproject/units'
import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/sushiMaster'
import { MINICHEF_ADDRESSES } from '@constants/minichef'
import { Token } from '@types'
import { BigNumber } from 'ethers'
import { Zero, One } from '@ethersproject/constants'
import { fetchBalance, readContracts, fetchToken } from '@wagmi/core'
import MINICHEF_ABI from '@abis/miniChef.json'
import { getSynPrices } from '@utils/actions/getPrices'
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
  const minichefAddress: `0x${string}` = `0x${MINICHEF_ADDRESSES[chainId].slice(
    2
  )}`

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
  const synapsePerSecondResult: any = data[0]
  const totalAllocPointsResult: any = data[1]
  const poolInfoResult: any = data[2] ?? []

  const lpTokenBalanceResult =
    (
      await fetchBalance({
        address: minichefAddress,
        chainId,
        token: `0x${poolToken.addresses[chainId].slice(2)}`,
      })
    )?.value ?? Zero

  const lpTokenSupplyResult =
    (
      await fetchToken({
        address: `0x${poolToken.addresses[chainId].slice(2)}`,
        chainId,
      })
    )?.totalSupply?.value ?? Zero

  const synPriceData = prices?.synPrices ?? (await getSynPrices())

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
    (allocPoints.toNumber() / totalAllocPoints.toNumber()) * rewardsPerWeek
  if (poolRewardsPerWeek === 0) {
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
