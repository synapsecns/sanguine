import { useState, useEffect, useMemo } from 'react'
import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import { SYN } from '@constants/tokens/master'
import { WETH } from '@constants/tokens/swapMaster'
import { formatUnits } from '@ethersproject/units'
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
  const [ethPrice, setEthPrice] = useState<BigNumber>(undefined)
  const [avaxPrice, setAvaxPrice] = useState<BigNumber>(undefined)

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

export const getEthPrice = async (): Promise<BigNumber> => {
  // the price result returned by latestAnswer is 8 decimals
  const ethPriceResult: any = await readContract({
    address: `0x${CHAINLINK_ETH_PRICE_ADDRESSES[ALL_CHAINS.ETH.id].slice(2)}`,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: 1,
  })
  const ethPriceBigNumber = BigNumber.from(ethPriceResult?._hex) ?? Zero

  return ethPriceBigNumber.div(BigNumber.from(10).pow(8))
}

export const getAvaxPrice = async (): Promise<BigNumber> => {
  // the price result returned by latestAnswer is 8 decimals
  const avaxPriceResult: any = await readContract({
    address: `0x${CHAINLINK_AVAX_PRICE_ADDRESSES[ALL_CHAINS.ETH.id].slice(2)}`,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: 1,
  })
  const avaxPriceBigNumber = BigNumber.from(avaxPriceResult?._hex) ?? Zero

  return avaxPriceBigNumber.div(BigNumber.from(10).pow(8))
}

export const getSynPrices = async () => {
  const ethPrice: BigNumber = await getEthPrice()
  const sushiSynBalance =
    (
      await fetchBalance({
        token: `0x${SYN.addresses[ALL_CHAINS.ETH.id].slice(2)}`,
        chainId: ALL_CHAINS.ETH.id,
        address: `0x${SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id].slice(
          2
        )}`,
      })
    )?.value ?? Zero
  const sushiEthBalance =
    (
      await fetchBalance({
        token: `0x${WETH.addresses[ALL_CHAINS.ETH.id].slice(2)}`,
        chainId: ALL_CHAINS.ETH.id,
        address: `0x${SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id].slice(
          2
        )}`,
      })
    )?.value ?? Zero

  const ethBalanceNumber = Number(formatUnits(sushiEthBalance, 'ether'))
  const synBalanceNumber = Number(formatUnits(sushiSynBalance, 'ether'))

  const synPerEth = synBalanceNumber / ethBalanceNumber

  const synPrice: number = ethPrice.toNumber() / synPerEth

  return {
    synBalanceNumber,
    ethBalanceNumber,
    synPrice,
    ethPrice: ethPrice.toNumber(),
  }
}
