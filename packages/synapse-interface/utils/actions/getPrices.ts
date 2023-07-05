import { useState, useEffect, useMemo } from 'react'
import { SYN } from '@constants/tokens/master'
import { WETH } from '@constants/tokens/swapMaster'
import * as ALL_CHAINS from '@constants/chains/master'
import CHAINLINK_AGGREGATOR_ABI from '@abis/chainlinkAggregator.json'
import {
  CHAINLINK_ETH_PRICE_ADDRESSES,
  CHAINLINK_AVAX_PRICE_ADDRESSES,
} from '@constants/chainlink'
import { readContract, fetchBalance } from '@wagmi/core'
import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/sushiMaster'

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
  const ethPriceResult: any = await readContract({
    address: `0x${CHAINLINK_ETH_PRICE_ADDRESSES[ALL_CHAINS.ETH.id].slice(2)}`,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: 1,
  })

  const ethPriceBigInt = ethPriceResult ? BigInt(ethPriceResult) : BigInt('0')
  if (ethPriceBigInt === BigInt('0')) {
    return 0
  } else {
    // Note: BigInt to Number conversion happens here
    return Number(ethPriceBigInt) / Math.pow(10, 8)
  }
}

export const getAvaxPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const avaxPriceResult: any = await readContract({
    address: `0x${CHAINLINK_AVAX_PRICE_ADDRESSES[ALL_CHAINS.ETH.id].slice(2)}`,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: 1,
  })

  const avaxPriceBigInt = avaxPriceResult
    ? BigInt(avaxPriceResult)
    : BigInt('0')

  if (avaxPriceBigInt === BigInt('0')) {
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
        token: `0x${SYN.addresses[ALL_CHAINS.ETH.id].slice(2)}`,
        chainId: ALL_CHAINS.ETH.id,
        address: `0x${SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id].slice(
          2
        )}`,
      })
    )?.value ?? BigInt('0')
  const sushiEthBalance =
    (
      await fetchBalance({
        token: `0x${WETH.addresses[ALL_CHAINS.ETH.id].slice(2)}`,
        chainId: ALL_CHAINS.ETH.id,
        address: `0x${SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id].slice(
          2
        )}`,
      })
    )?.value ?? BigInt('0')

  // Assuming formatUnits(sushiEthBalance, 'ether') converts the balance to Ether (i.e., divides by 10**18)
  const ethBalanceNumber = Number(sushiEthBalance) / Math.pow(10, 18)
  const synBalanceNumber = Number(sushiSynBalance) / Math.pow(10, 18)

  const synPerEth = synBalanceNumber / ethBalanceNumber

  const synPrice: number = ethPrice * synPerEth

  return {
    synBalanceNumber,
    ethBalanceNumber,
    synPrice,
    ethPrice,
  }
}
