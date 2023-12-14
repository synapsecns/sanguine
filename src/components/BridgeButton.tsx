import { Chain } from 'types'
import { Tooltip } from './Tooltip'

interface BridgeButtonProps {
  originChain: Chain
  isApproved: boolean
  isValidQuote: boolean
  isValidAmount: boolean
  isWrongNetwork: boolean
  isInputGreaterThanBalance: boolean
  handleApprove: () => any
  handleBridge: () => any
  handleSwitchNetwork: () => Promise<any>
  isApprovalPending: boolean
  isBridgePending: boolean
  approveError: string
  bridgeError: string
}

const BridgeError = ({
  error,
  type,
}: {
  error: string
  type: 'bridge' | 'approve'
}) => {
  if (error) {
    return (
      <div
        data-test-id={`${type}-error`}
        className="text-sm text-right text-red-500"
      >
        {error}
      </div>
    )
  }
}

export const BridgeButton = ({
  originChain,
  isApproved,
  isValidQuote,
  isValidAmount,
  isWrongNetwork,
  isInputGreaterThanBalance,
  handleApprove,
  handleBridge,
  handleSwitchNetwork,
  isApprovalPending,
  isBridgePending,
  approveError,
  bridgeError,
}: BridgeButtonProps) => {
  if (isWrongNetwork) {
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
  if (isInputGreaterThanBalance) {
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
  if (!isValidAmount) {
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
