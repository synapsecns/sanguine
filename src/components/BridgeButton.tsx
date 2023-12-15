import { Chain } from 'types'
import { Tooltip } from './Tooltip'
import { Web3Context } from 'providers/Web3Provider'
import { useCallback, useContext } from 'react'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { switchNetwork } from '@/utils/actions/switchNetwork'
import { useValidations } from '@/hooks/useValidations'

interface BridgeButtonProps {
  originChain: Chain
  isApproved: boolean
  isValidQuote: boolean
  handleApprove: () => any
  handleBridge: () => any
  isApprovalPending: boolean
  isBridgePending: boolean
  approveError: string
  bridgeError: string
}

export const BridgeButton = ({
  originChain,
  isApproved,
  isValidQuote,
  handleApprove,
  handleBridge,
  isApprovalPending,
  isBridgePending,
  approveError,
  bridgeError,
}: BridgeButtonProps) => {
  const web3Context = useContext(Web3Context)

  const { connectedAddress, signer, provider, networkId } =
    web3Context.web3Provider

  const {
    inputAmount,
    debouncedInputAmount,
    originChainId,
    originToken,
    destinationChainId,
    destinationToken,
    balances,
  } = useBridgeState()

  const { hasEnoughBalance, isInputValid, onSelectedChain } = useValidations()

  const handleSwitchNetwork = useCallback(async () => {
    switchNetwork(originChainId, provider)
  }, [originChainId, provider])

  if (!onSelectedChain) {
    return (
      <button
        onClick={handleSwitchNetwork}
        className={`
          rounded-md w-full p-2 font-semibold 
          bg-[--synapse-bg-surface] border border-[--synapse-accent] 
          active:opacity-40
        `}
      >
        Connect to {originChain?.name}
      </button>
    )
  }
  if (!hasEnoughBalance) {
    return (
      <Tooltip hoverText="Amount may not exceed available balance">
        <button
          onClick={() => null}
          className={`
            rounded-md w-full p-2 font-semibold 
            bg-[--synapse-bg-surface] text-[--synapse-text-secondary]
            border border-[--synapse-border]
          `}
        >
          Send
        </button>
      </Tooltip>
    )
  }
  if (!isInputValid) {
    return (
      <Tooltip hoverText="Enter valid amount">
        <button
          onClick={() => null}
          className={`
            rounded-md w-full p-2 font-semibold 
            bg-[--synapse-bg-surface] text-[--synapse-text-secondary]
            border border-[--synapse-border]
          `}
        >
          Send
        </button>
      </Tooltip>
    )
  }
  if (!isValidQuote) {
    return (
      <button
        onClick={() => null}
        className={`
          rounded-md w-full p-2 font-semibold 
          bg-[--synapse-bg-surface] text-[--synapse-text-secondary]
          border border-[--synapse-border]
        `}
      >
        Send
      </button>
    )
  }
  return (
    <div data-test-id="bridge-button">
      {isApproved ? (
        <button
          onClick={!isBridgePending ? handleBridge : () => null}
          className={`
              rounded-md w-full p-2  font-semibold 
              bg-[--synapse-bg-surface] 
              border border-[--synapse-accent] 
              hover:border-[--synapse-brand]
              ${isBridgePending && 'opacity-40 border-[--synapse-border]'}
            `}
        >
          {isBridgePending ? 'Confirm in Wallet' : 'Send'}
        </button>
      ) : (
        <Tooltip hoverText={isApprovalPending && 'Wallet approval required'}>
          <button
            onClick={!isApprovalPending ? handleApprove : () => null}
            className={`
              rounded-md w-full p-2 font-semibold 
              bg-[--synapse-bg-surface] border border-[--synapse-accent] 
              hover:border-[--synapse-brand]
              ${isApprovalPending && 'opacity-40 border-[--synapse-border]'}
            `}
          >
            {isApprovalPending ? 'Approve in Wallet' : 'Approve & Sign'}
          </button>
        </Tooltip>
      )}
    </div>
  )
}
