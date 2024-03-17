import { useCallback, useContext } from 'react'
import { Chain } from 'types'
import { Web3Context } from 'providers/Web3Provider'

import { Tooltip } from '@/components/ui/Tooltip'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { switchNetwork } from '@/utils/actions/switchNetwork'
import { useValidations } from '@/hooks/useValidations'

interface BridgeButtonProps {
  originChain: Chain
  isValidQuote: boolean
  handleApprove: () => any
  handleBridge: () => any
  isApprovalPending: boolean
  isBridgePending: boolean
}

export const BridgeButton = ({
  originChain,
  isValidQuote,
  handleApprove,
  handleBridge,
  isApprovalPending,
  isBridgePending,
}: BridgeButtonProps) => {
  const web3Context = useContext(Web3Context)

  const { connectedAddress, provider } = web3Context.web3Provider

  const { originChainId } = useBridgeState()

  const { hasEnoughBalance, isInputValid, onSelectedChain, isApproved } =
    useValidations()

  const handleSwitchNetwork = useCallback(async () => {
    switchNetwork(originChainId, provider)
  }, [originChainId, provider])

  const buttonClassName = `
    p-2 text-lg font-sans font-medium tracking-wide rounded-md w-full
    border border-solid border-[--synapse-button-border]
    hover:border-[--synapse-focus] disabled:hover:border-[--synapse-button-border]
    active:opacity-70 disabled:opacity-40
    text-[--synapse-button-text]
    cursor-pointer disabled:cursor-not-allowed
  `

  const buttonStyle = {
    background: 'var(--synapse-button-bg)',
  }

  const tooltipPositionStyle = '-top-8'

  if (!provider || !connectedAddress) {
    return (
      <Tooltip hoverText="Connect Wallet" positionStyles={tooltipPositionStyle}>
        <button className={buttonClassName} style={buttonStyle} disabled>
          Connect Wallet
        </button>
      </Tooltip>
    )
  }

  if (!originChainId) {
    return (
      <Tooltip
        hoverText="Select origin chain"
        positionStyles={tooltipPositionStyle}
      >
        <button className={buttonClassName} style={buttonStyle} disabled>
          Select origin chain
        </button>
      </Tooltip>
    )
  }

  if (!onSelectedChain) {
    return (
      <button
        className={buttonClassName}
        style={buttonStyle}
        onClick={handleSwitchNetwork}
      >
        Connect to {originChain?.name}
      </button>
    )
  }
  if (!hasEnoughBalance) {
    return (
      <Tooltip
        hoverText="Amount may not exceed available balance"
        positionStyles={tooltipPositionStyle}
      >
        <button className={buttonClassName} style={buttonStyle} disabled>
          Send
        </button>
      </Tooltip>
    )
  }
  if (!isInputValid) {
    return (
      <Tooltip
        hoverText="Enter valid amount"
        positionStyles={tooltipPositionStyle}
      >
        <button className={buttonClassName} style={buttonStyle} disabled>
          Send
        </button>
      </Tooltip>
    )
  }
  if (!isValidQuote) {
    return (
      <button className={buttonClassName} style={buttonStyle} disabled>
        Send
      </button>
    )
  }
  return (
    <div data-test-id="bridge-button">
      {isApproved ? (
        <button
          disabled={isBridgePending}
          onClick={!isBridgePending ? handleBridge : () => null}
          className={buttonClassName}
          style={buttonStyle}
        >
          {isBridgePending ? 'Confirm in Wallet' : 'Send'}
        </button>
      ) : (
        <Tooltip
          hoverText={isApprovalPending ? 'Wallet approval required' : null}
          positionStyles={tooltipPositionStyle}
        >
          <button
            disabled={isApprovalPending}
            onClick={!isApprovalPending ? handleApprove : () => null}
            className={buttonClassName}
            style={buttonStyle}
          >
            {isApprovalPending ? 'Approve in Wallet' : 'Approve & Sign'}
          </button>
        </Tooltip>
      )}
    </div>
  )
}
