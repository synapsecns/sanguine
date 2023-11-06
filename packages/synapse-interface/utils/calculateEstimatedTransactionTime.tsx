import { Chain, Token } from './types'
import { Address } from 'viem'
import { BRIDGE_REQUIRED_CONFIRMATIONS } from '@/constants/bridge'
import { tokenAddressToToken } from '@/constants/tokens'
import { ARBITRUM, ETH } from '@/constants/chains/master'
import { USDC } from '@/constants/tokens/bridgeable'
import { CHAINS_BY_ID } from '@/constants/chains'

// Utility function to fetch estimated transaction time
// Returned as a number, in seconds
export const calculateEstimatedTransactionTime = ({
  originChainId,
  originTokenAddress,
}: {
  originChainId: number
  originTokenAddress: Address
}): number => {
  const originChain: Chain = CHAINS_BY_ID[originChainId]
  const originToken: Token = tokenAddressToToken(
    originChainId,
    originTokenAddress
  )
  const baseEstimatedCompletionInSeconds: number =
    (BRIDGE_REQUIRED_CONFIRMATIONS[originChainId] * originChain?.blockTime) /
    1000

  let estimatedCompletionInSeconds: number

  // Specific to CCTP Transactions
  if (originChainId === ARBITRUM.id || originChainId === ETH.id) {
    const isCCTP: boolean = originTokenAddress === USDC.addresses[originChainId]
    const attestationTime: number = 13 * 60

    estimatedCompletionInSeconds =
      baseEstimatedCompletionInSeconds + attestationTime

    return isCCTP
      ? estimatedCompletionInSeconds
      : baseEstimatedCompletionInSeconds
  }

  return baseEstimatedCompletionInSeconds
}
