import { useSelector } from 'react-redux'
import { useAccount } from 'wagmi'
import { isAddress } from '@ethersproject/address'

import { TransactionButton } from '@/components/buttons/TransactionButton'
import { RootState } from '@/store/store'
import { useBridgeState } from '@/slices/bridge/hooks'
import { PAUSED_FROM_CHAIN_IDS, PAUSED_TO_CHAIN_IDS } from '@/constants/chains'
import { useBridgeStatus } from '@/utils/hooks/useBridgeStatus'

export const BridgeTransactionButton = ({
  approveTxn,
  executeBridge,
  isApproved,
}) => {
  const { isConnected } = useAccount()

  const {
    destinationAddress,
    fromToken,
    toToken,
    fromChainId,
    toChainId,
    isLoading,
    bridgeQuote,
  } = useBridgeState()

  const { hasEnoughBalance, hasInputAmount, onSelectedChain, hasValidRoute } =
    useBridgeStatus()

  const { showDestinationAddress } = useSelector(
    (state: RootState) => state.bridgeDisplay
  )

  const isButtonDisabled =
    isLoading ||
    !onSelectedChain ||
    !hasValidRoute ||
    (destinationAddress && !isAddress(destinationAddress)) ||
    (showDestinationAddress && !destinationAddress) ||
    (isConnected && !hasEnoughBalance) ||
    PAUSED_FROM_CHAIN_IDS.includes(fromChainId) ||
    PAUSED_TO_CHAIN_IDS.includes(toChainId)

  let buttonProperties: {
    label: string
    pendingLabel?: string
    onClick: any
    toolTipLabel?: string
  } = {
    label: `Bridge ${fromToken?.symbol ?? ''}`,
    pendingLabel: 'Bridging',
    onClick: null,
  }

  if (!isConnected) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Connect Wallet',
    }
  } else if (!onSelectedChain && hasInputAmount) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Please switch chains',
    }
  } else if (!fromChainId) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Please select Origin network',
    }
  } else if (!fromToken) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Please select Origin token',
    }
  } else if (!toChainId) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Please select Destination network',
    }
  } else if (!toToken) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Please select Destination token',
    }
  } else if (
    PAUSED_FROM_CHAIN_IDS.includes(fromChainId) ||
    PAUSED_TO_CHAIN_IDS.includes(toChainId)
  ) {
    buttonProperties = {
      ...buttonProperties,
      label: `Bridge unavailable`,
    }
  } else if (!hasInputAmount) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Please enter an amount',
    }
  } else if (!isLoading && bridgeQuote?.feeAmount === 0n && hasInputAmount) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Amount must be greater than fee',
    }
  } else if (isConnected && !hasEnoughBalance) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Insufficient balance',
    }
  } else if (showDestinationAddress && !destinationAddress) {
    buttonProperties = {
      ...buttonProperties,
      label: 'Please add valid destination address',
    }
  } else if (destinationAddress && !isAddress(destinationAddress)) {
    buttonProperties = {
      ...buttonProperties,
      label: 'Invalid destination address',
    }
  } else if (!isApproved) {
    buttonProperties = {
      ...buttonProperties,
      onClick: approveTxn,
      label: `Approve ${fromToken?.symbol}`,
      pendingLabel: 'Approving',
    }
  } else {
    buttonProperties = {
      ...buttonProperties,
      onClick: executeBridge,
      label: `Bridge ${fromToken?.symbol}`,
      pendingLabel: 'Bridging',
    }
  }

  return (
    buttonProperties && (
      <TransactionButton
        {...buttonProperties}
        disabled={isButtonDisabled}
        chainId={fromChainId}
      />
    )
  )
}
