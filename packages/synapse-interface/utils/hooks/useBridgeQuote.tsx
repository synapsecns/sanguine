import { Address } from 'wagmi'
import { useSynapseContext } from '../providers/SynapseProvider'
import { BigNumberish } from 'ethers'

interface useBridgeQuoteProps {
  fromChainId: number
  toChainId: number
  fromTokenAddress: string | Address
  toTokenAddress: string | Address
  inputValue: BigNumberish
}

/*
  useBridgeQuote hook
  - Gets quote data from the Synapse SDK (from the imported provider)
  - Calculates slippage by subtracting fee from input amount (checks to ensure proper num of decimals are in use - ask someone about stable swaps if you want to learn more)
  - Wraps SynapseSDK calls in Promise wrapper to be cancellable when new request is made
  */
export const useBridgeQuote = ({
  fromChainId,
  toChainId,
  fromTokenAddress,
  toTokenAddress,
  inputValue,
}: useBridgeQuoteProps) => {
  const SynapseSDK = useSynapseContext()

  return <></>
}
