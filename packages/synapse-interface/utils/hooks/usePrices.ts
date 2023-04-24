import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import * as ALL_CHAINS from '@constants/chains/master'
import CHAINLINK_AGGREGATOR_ABI from '@abis/chainlinkAggregator.json'
import {
  CHAINLINK_ETH_PRICE_ADDRESSES,
  CHAINLINK_AVAX_PRICE_ADDRESSES,
} from '@constants/chainlink'
import { readContract } from '@wagmi/core'
export const useEthPrice = async (): Promise<BigNumber> => {
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

export const useAvaxPrice = async (): Promise<BigNumber> => {
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
