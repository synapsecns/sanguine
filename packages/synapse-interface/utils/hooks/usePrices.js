import { BigNumber } from "@ethersproject/bignumber"
import { Zero } from "@ethersproject/constants"
import { ChainId } from "@constants/networks"
import {
  useChainlinkEthPriceContract,
  useChainlinkAvaxPriceContract,
} from "@hooks/contracts/useChainlinkPriceContracts"

import { useSingleCallResult } from "@hooks/multicall"

/**
 * @returns {BigNumber}
 */
export function useEthPrice() {
  const chainlinkEthPriceContract = useChainlinkEthPriceContract()
  // the price result returned by latestAnswer is 8 decimals
  const ethPriceResult = useSingleCallResult(
    ChainId.ETH,
    chainlinkEthPriceContract,
    'latestAnswer',
    [],
    { resultOnly: true }
  )
  const bnEthPrice = ethPriceResult?.[0] ?? Zero

  return bnEthPrice.div( BigNumber.from(10).pow(8) )
}

/**
 * @returns {BigNumber}
 */
export function useAvaxPrice() {
  const chainlinkAvaxPriceContract = useChainlinkAvaxPriceContract()
  // the price result returned by latestAnswer is 8 decimals
  const avaxPriceResult = useSingleCallResult(
    ChainId.ETH,
    chainlinkAvaxPriceContract,
    'latestAnswer',
    [],
    { resultOnly: true }
  )
  const bnAvaxPrice = avaxPriceResult?.[0] ?? Zero

  return bnAvaxPrice.div(BigNumber.from(10).pow(8))
}