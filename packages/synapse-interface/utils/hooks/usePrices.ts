import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import * as ALL_CHAINS from '@constants/chains/master'
import CHAINLINK_AGGREGATOR_ABI from '@abis/chainlinkAggregator.json'
import {
  CHAINLINK_ETH_PRICE_ADDRESSES,
  CHAINLINK_AVAX_PRICE_ADDRESSES,
} from '@constants/chainlink'
import { readContract } from '@wagmi/core'

export const useEthPrice = async () => {
  // the price result returned by latestAnswer is 8 decimals
  const ethPriceResult = await readContract({
    address: `0x${CHAINLINK_ETH_PRICE_ADDRESSES[ALL_CHAINS.ETH.id].slice(2)}`,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
  })
  console.log('ethPriceResult', ethPriceResult)
  const bnEthPrice = ethPriceResult?.[0] ?? Zero

  return bnEthPrice.div(BigNumber.from(10).pow(8))
}

export const useAvaxPrice = async () => {
  // the price result returned by latestAnswer is 8 decimals
  const avaxPriceResult = await readContract({
    address: `0x${CHAINLINK_AVAX_PRICE_ADDRESSES[ALL_CHAINS.ETH.id].slice(2)}`,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
  })
  console.log('avaxPriceResult', avaxPriceResult)
  const bnAvaxPrice = avaxPriceResult?.[0] ?? Zero

  return bnAvaxPrice.div(BigNumber.from(10).pow(8))
}
