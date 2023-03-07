import CHAINLINK_AGGREGATOR_ABI from '@abis/chainlinkAggregator.json'

import { CHAINLINK_ETH_PRICE_ADDRESSES, CHAINLINK_AVAX_PRICE_ADDRESSES } from '@constants/chainlink'
import { ChainId } from '@constants/networks'

import { useGenericContract } from '@hooks/contracts/useContract'




export function useChainlinkEthPriceContract() {
  const chainlinkContract = useGenericContract(
    ChainId.ETH,
    CHAINLINK_ETH_PRICE_ADDRESSES[ChainId.ETH],
    CHAINLINK_AGGREGATOR_ABI
  )

  return chainlinkContract
}


export function useChainlinkAvaxPriceContract() {
  const chainlinkContract = useGenericContract(
    ChainId.ETH,
    CHAINLINK_AVAX_PRICE_ADDRESSES[ChainId.ETH],
    CHAINLINK_AGGREGATOR_ABI
  )

  return chainlinkContract
}




