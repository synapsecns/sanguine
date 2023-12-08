import { Chain } from 'types'

interface BridgeButtonProps {
  originChain: Chain
  isApproved: boolean
  isDisabled: boolean
  isWrongNetwork: boolean
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
  isDisabled,
  isWrongNetwork,
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
        className="rounded-md w-full bg-[--synapse-bg-surface] font-semibold border border-[--synapse-border] p-2 hover:border-[--synapse-brand] active:opacity-40"
        onClick={handleSwitchNetwork}
      >
        Connect to {originChain.name} Network
      </button>
    )
  }
  return (
    <div data-test-id="bridge-button">
      {isApproved ? (
        <div>
          <BridgeError error={bridgeError} type="bridge" />
          <button
            className="rounded-md w-full bg-[--synapse-bg-surface] font-semibold border border-[--synapse-border] p-2 hover:border-[--synapse-brand] active:opacity-40"
            onClick={handleBridge}
          >
            {isBridgePending ? 'Pending Bridge...' : 'Bridge'}
          </button>
        </div>
      ) : (
        <div>
          <BridgeError error={approveError} type="approve" />
          <button
            className="rounded-md w-full bg-[--synapse-bg-surface] font-semibold border border-[--synapse-border] p-2 hover:border-[--synapse-brand] active:opacity-40"
            onClick={handleApprove}
          >
            {isApprovalPending ? 'Pending Approval...' : 'Approve'}
          </button>
        </div>
      )}
    </div>
  )
}
