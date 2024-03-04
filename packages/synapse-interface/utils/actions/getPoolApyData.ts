import { formatUnits } from '@ethersproject/units'
import { readContracts, Address, erc20ABI } from '@wagmi/core'
import type { Token } from '@/utils/types'
import { MINICHEF_ABI } from '@/constants/abis/miniChef'
import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import { SYN_ETH_SUSHI_TOKEN } from '@/constants/tokens/sushiMaster'
import { getSynPrices } from '@/utils/actions/getPrices'


type PoolInfoResult = readonly [
  accSynapsePerShare: bigint,
  lastRewardTime: bigint,
  allocPoint: bigint
]

type Prices = {
  synPrices: {
    ethBalanceNumber: number
    ethPrice: number
    synBalanceNumber: number
    synPrice: number
  }
  ethPrice: number
  avaxPrice: number
  metisPrice: number
}

export const getPoolApyData = async (
  chainId: number,
  poolToken: Token,
  prices?: Prices
) => {
  if (!MINICHEF_ADDRESSES?.[chainId]) {
    console.log('no minichef address found for chainId', chainId)
    return {
      fullCompoundedAPY: 0,
      weeklyAPR: 0,
      yearlyAPRUnvested: 0,
    }
  }


  const minichefAddress = poolToken.miniChefAddress as Address

  const minichefContractInfo = {
    address: minichefAddress,
    abi:     MINICHEF_ABI,
    chainId,
  }
  const poolTokenContractInfo = {
    address: poolToken.addresses[chainId] as Address,
    abi:     erc20ABI,
    chainId,
  }

  const [synPriceData, data] = await Promise.all([
    prices?.synPrices?.synPrice ? prices?.synPrices : getSynPrices(),
    readContracts({
      contracts: [
        {
          ...minichefContractInfo,
          functionName: 'synapsePerSecond',
        },
        {
          ...minichefContractInfo,
          functionName: 'totalAllocPoint',
        },
        {
          ...minichefContractInfo,
          functionName: 'poolInfo',
          args: [poolToken.poolId[chainId]],
        },
        {
          ...poolTokenContractInfo,
          functionName: 'balanceOf',
          args: [minichefAddress],
        },
        {
          ...poolTokenContractInfo,
          functionName: 'totalSupply',
        },
      ],
    })
  ])

  // const metisPrice = prices?.metisPrice ?? (await getMetisPrice())

  const synapsePerSecond: bigint = data[0].result ?? 0n
  const totalAllocPoints: bigint = data[1].result ?? 1n
  const allocPoints: bigint = (data[2].result as PoolInfoResult)?.[2] ?? 1n
  const lpTokenBalance: bigint = data[3].result ?? 0n
  const lpTokenSupply: bigint = data[4].result ?? 0n

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

  const usePrice = synPriceData.synPrice

  /* Placeholder for non-SYN pools */
  // if (poolToken.chainId === METIS.id) {
  //   usePrice = metisPrice
  // } else {
  //   usePrice = synPriceData.synPrice
  // }

  const usdPerWeek = poolRewardsPerWeek * usePrice

  const weeklyAPR = (usdPerWeek / stakedTvl) * 100
  const yearlyAPR = weeklyAPR * 52
  const decimalAPR = yearlyAPR / 100
  const yearlyCompoundedAPR = 100 * ((1 + decimalAPR / 365) ** 365 - 1)
  const fullCompoundedAPY = Math.round(yearlyCompoundedAPR * 100) / 100
  const fullCompoundedAPYStr = isFinite(fullCompoundedAPY)
    ? fullCompoundedAPY.toFixed(2)
    : '\u2212'

  return {
    fullCompoundedAPY,
    fullCompoundedAPYStr,
    weeklyAPR: Math.round(weeklyAPR * 100) / 100,
    yearlyAPRUnvested: Math.round(yearlyAPR * 100) / 100,
  }
}
