import { Chain } from 'types'
import { Tooltip } from './Tooltip'
import { Web3Context } from 'providers/Web3Provider'
import { useCallback, useContext } from 'react'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { switchNetwork } from '@/utils/actions/switchNetwork'
import { useValidations } from '@/hooks/useValidations'
import { useWalletState } from '@/state/slices/wallet/hooks'

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

  const { connectedAddress, signer, provider, networkId } =
    web3Context.web3Provider

  const {
    inputAmount,
    debouncedInputAmount,
    originChainId,
    originToken,
    destinationChainId,
    destinationToken,
  } = useBridgeState()

  const { hasEnoughBalance, isInputValid, onSelectedChain, isApproved } =
    useValidations()

  const handleSwitchNetwork = useCallback(async () => {
    switchNetwork(originChainId, provider)
  }, [originChainId, provider])

  const buttonStyle = `
    p-2 text-lg font-sans font-semibold rounded-md w-full 
    bg-[--synapse-button-bg] disabled:bg-[--synapse-root] 
    border border-solid border-[--synapse-button-border] 
    hover:border-[--synapse-focus] disabled:hover:border-[--synapse-button-border] 
    active:opacity-40 disabled:opacity-70 
    text-[--synapse-button-text] disabled:text-[--synapse-secondary] 
    cursor-pointer disabled:cursor-not-allowed
  `

  const tooltipPositionStyle = '-top-8'

  if (!onSelectedChain) {
    return (
      <button className={buttonStyle} onClick={handleSwitchNetwork}>
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
        <button className={buttonStyle} disabled>
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
        <button className={buttonStyle} disabled>
          Send
        </button>
      </Tooltip>
    )
  }
  if (!isValidQuote) {
    return (
      <button className={buttonStyle} disabled>
        Send
      </button>
    )
  }
  return (
    <div data-test-id="bridge-button">
      {isApproved ? (
        <button
          disabled={isBridgePending}
          onClick={!isBridgePending && handleBridge}
          className={buttonStyle}
        >
          {isBridgePending ? 'Confirm in Wallet' : 'Send'}
        </button>
      ) : (
        <Tooltip
          hoverText={isApprovalPending && 'Wallet approval required'}
          positionStyles={tooltipPositionStyle}
        >
          <button
            disabled={isApprovalPending}
            onClick={!isApprovalPending && handleApprove}
            className={buttonStyle}
          >
            {isApprovalPending ? 'Approve in Wallet' : 'Approve & Sign'}
          </button>
        </Tooltip>
      )}
    </div>
  )
}
