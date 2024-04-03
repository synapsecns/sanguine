import { readContract, getBalance } from '@wagmi/core'
import { type Address } from 'viem'

import { SYN, WETH } from '@/constants/tokens/bridgeable'
import * as ALL_CHAINS from '@/constants/chains/master'
import {
  CHAINLINK_ETH_PRICE_ADDRESSES,
  CHAINLINK_AVAX_PRICE_ADDRESSES,
  CHAINLINK_METIS_PRICE_ADDRESSES,
  CHAINLINK_GMX_PRICE_ADDRESSES,
} from '@/constants/chainlink'
import { SYN_ETH_SUSHI_TOKEN } from '@/constants/tokens/sushiMaster'
import CHAINLINK_AGGREGATOR_ABI from '@/constants/abis/chainlinkAggregator.json'
import { wagmiConfig } from '@/wagmiConfig'

export const getEthPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const ethPriceResult: bigint = (await readContract(wagmiConfig, {
    address: CHAINLINK_ETH_PRICE_ADDRESSES[ALL_CHAINS.ETH.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: 1,
  })) as bigint

  const ethPriceBigInt = ethPriceResult ?? 0n
  if (ethPriceBigInt === 0n) {
    return 0
  } else {
    // Note: BigInt to Number conversion happens here
    return Number(ethPriceBigInt) / Math.pow(10, 8)
  }
}

export const getAvaxPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const avaxPriceResult: bigint = (await readContract(wagmiConfig, {
    address: CHAINLINK_AVAX_PRICE_ADDRESSES[ALL_CHAINS.ETH.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: 1,
  })) as bigint

  const avaxPriceBigInt = avaxPriceResult ?? 0n

  if (avaxPriceBigInt === 0n) {
    return 0
  } else {
    return Number(avaxPriceBigInt) / Math.pow(10, 8)
  }
}

export const getMetisPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const metisPriceResult: bigint = (await readContract(wagmiConfig, {
    address: CHAINLINK_METIS_PRICE_ADDRESSES[ALL_CHAINS.METIS.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: ALL_CHAINS.METIS.id as any,
  })) as bigint

  const metisPriceBigInt = metisPriceResult ?? 0n

  if (metisPriceBigInt === 0n) {
    return 0
  } else {
    return Number(metisPriceBigInt) / Math.pow(10, 8)
  }
}

export const getSynPrices = async () => {
  const ethPrice: number = await getEthPrice()
  const sushiSynBalance =
    (
      await getBalance(wagmiConfig, {
        token: SYN.addresses[ALL_CHAINS.ETH.id] as Address,
        chainId: ALL_CHAINS.ETH.id as any,
        address: SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id] as Address,
      })
    )?.value ?? 0n

  const sushiEthBalance =
    (
      await getBalance(wagmiConfig, {
        token: WETH.addresses[ALL_CHAINS.ETH.id] as Address,
        chainId: ALL_CHAINS.ETH.id as any,
        address: SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id] as Address,
      })
    )?.value ?? 0n

  const synBalanceNumber = Number(sushiSynBalance) / Math.pow(10, 18)
  const ethBalanceNumber = Number(sushiEthBalance) / Math.pow(10, 18)
  const synPerEth = synBalanceNumber / ethBalanceNumber
  const synPrice: number = ethPrice * (1 / synPerEth)

  return {
    synBalanceNumber,
    ethBalanceNumber,
    synPrice,
    ethPrice,
  }
}

export const getGmxPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const gmxPriceResult: bigint = (await readContract(wagmiConfig, {
    address: CHAINLINK_GMX_PRICE_ADDRESSES[ALL_CHAINS.ARBITRUM.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: ALL_CHAINS.ARBITRUM.id as any,
  })) as bigint

  const gmxPriceBigInt = gmxPriceResult ?? 0n

  if (gmxPriceBigInt === 0n) {
    return 0
  } else {
    return Number(gmxPriceBigInt) / Math.pow(10, 8)
  }
}
