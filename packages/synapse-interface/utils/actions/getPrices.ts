import { useState, useEffect, useMemo } from 'react'
import { readContract, fetchBalance, Address } from '@wagmi/core'
import { SYN } from '@constants/tokens/bridgeable'
import { WETH } from '@constants/tokens/exceptions'
import * as ALL_CHAINS from '@constants/chains/master'
import {
  CHAINLINK_ETH_PRICE_ADDRESSES,
  CHAINLINK_AVAX_PRICE_ADDRESSES,
} from '@constants/chainlink'
import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/sushiMaster'
import CHAINLINK_AGGREGATOR_ABI from '@abis/chainlinkAggregator.json'

export const usePrices = (connectedChainId: number) => {
  const [synPrices, setSynPrices] = useState<any>(undefined)
  const [ethPrice, setEthPrice] = useState<number>(undefined)
  const [avaxPrice, setAvaxPrice] = useState<number>(undefined)

  useEffect(() => {
    ;(async () => {
      try {
        const data = await getSynPrices()
        setSynPrices(data)
      } catch (err) {
        console.log('Could not get syn prices', err)
      }

      try {
        const data = await getEthPrice()
        setEthPrice(data)
      } catch (err) {
        console.log('Could not get eth prices', err)
      }

      try {
        const data = await getAvaxPrice()
        setAvaxPrice(data)
      } catch (err) {
        console.log('Could not get avax prices', err)
      }
    })()
  }, [connectedChainId])

  return useMemo(() => {
    const prices = { synPrices, ethPrice, avaxPrice }
    return prices
  }, [synPrices, ethPrice, avaxPrice])
}

export const getEthPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const ethPriceResult: bigint = (await readContract({
    address: `0x${CHAINLINK_ETH_PRICE_ADDRESSES[ALL_CHAINS.ETH.id].slice(2)}`,
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
  const avaxPriceResult: bigint = (await readContract({
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

export const getSynPrices = async () => {
  const ethPrice: number = await getEthPrice()
  const sushiSynBalance =
    (
      await fetchBalance({
        token: SYN.addresses[ALL_CHAINS.ETH.id] as Address,
        chainId: ALL_CHAINS.ETH.id,
        address: SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id] as Address,
      })
    )?.value ?? 0n

  const sushiEthBalance =
    (
      await fetchBalance({
        token: WETH.addresses[ALL_CHAINS.ETH.id] as Address,
        chainId: ALL_CHAINS.ETH.id,
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
