import { useSynapseContext } from '../providers/SynapseProvider'

/*
  useBridgeQuote hook
  - Gets quote data from the Synapse SDK (from the imported provider)
  - Calculates slippage by subtracting fee from input amount (checks to ensure proper num of decimals are in use - ask someone about stable swaps if you want to learn more)
  - Wraps SynapseSDK calls in Promise wrapper to be cancellable when new request is made
  */
export const useBridgeQuote = () => {
  const SynapseSDK = useSynapseContext()

  return <></>
}
