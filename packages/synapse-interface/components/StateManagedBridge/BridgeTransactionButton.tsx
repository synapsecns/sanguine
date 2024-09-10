import { isAddress } from 'viem'
import { useEffect, useState } from 'react'
import { useAccount, useAccountEffect, useSwitchChain } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'

import { useAppDispatch } from '@/store/hooks'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { setIsDestinationWarningAccepted } from '@/slices/bridgeDisplaySlice'
import { useBridgeDisplayState, useBridgeState } from '@/slices/bridge/hooks'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useBridgeValidations } from './hooks/useBridgeValidations'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useConfirmNewBridgePrice } from './hooks/useConfirmNewBridgePrice'

export const BridgeTransactionButton = ({
  approveTxn,
  executeBridge,
  isApproved,
  isBridgePaused,
  isTyping,
  isQuoteStale,
}) => {
  const dispatch = useAppDispatch()
  const { openConnectModal } = useConnectModal()
  const [isConnected, setIsConnected] = useState(false)

  const { isConnected: isConnectedInit } = useAccount()
  const { chains, switchChain } = useSwitchChain()

  useAccountEffect({
    onDisconnect() {
      setIsConnected(false)
    },
  })

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit])

  const {
    destinationAddress,
    fromToken,
    fromChainId,
    toToken,
    toChainId,
    debouncedFromValue,
  } = useBridgeState()
  const { bridgeQuote, isLoading } = useBridgeQuoteState()
  const {
    hasSameSelectionsAsPreviousQuote,
    hasQuoteOutputChanged,
    hasUserConfirmedChange,
    onUserAcceptChange,
  } = useConfirmNewBridgePrice()

  const { isWalletPending } = useWalletState()
  const { showDestinationWarning, isDestinationWarningAccepted } =
    useBridgeDisplayState()

  const {
    hasValidInput,
    hasValidQuote,
    hasSufficientBalance,
    doesBridgeStateMatchQuote,
    isBridgeFeeGreaterThanInput,
    isBridgeQuoteAmountGreaterThanInputForRfq,
    onSelectedChain,
  } = useBridgeValidations()

  const isButtonDisabled =
    isBridgePaused ||
    isQuoteStale ||
    isTyping ||
    isLoading ||
    isWalletPending ||
    !hasValidInput ||
    !doesBridgeStateMatchQuote ||
    isBridgeQuoteAmountGreaterThanInputForRfq ||
    (isConnected && !hasValidQuote) ||
    (isConnected && !hasSufficientBalance) ||
    (destinationAddress && !isAddress(destinationAddress))

  let buttonProperties

  if (isBridgePaused) {
    buttonProperties = {
      label: 'Bridge paused',
      onClick: null,
    }
  } else if (!fromChainId) {
    buttonProperties = {
      label: 'Please select Origin Network',
      onClick: null,
    }
  } else if (!toChainId) {
    buttonProperties = {
      label: 'Please select Destination network',
      onClick: null,
    }
  } else if (!fromToken) {
    buttonProperties = {
      label: `Please select an Origin token`,
      onClick: null,
    }
  } else if (isConnected && !hasSufficientBalance) {
    buttonProperties = {
      label: 'Insufficient balance',
      onClick: null,
    }
  } else if (isLoading && hasSameSelectionsAsPreviousQuote) {
    buttonProperties = {
      label: 'Updating quote',
      onClick: null,
    }
  } else if (isLoading) {
    buttonProperties = {
      label: `Bridge ${fromToken?.symbol}`,
      pendingLabel: `Bridge ${fromToken?.symbol}`,
      onClick: null,
    }
  } else if (!isConnected && hasValidInput) {
    buttonProperties = {
      label: `Connect Wallet to Bridge`,
      onClick: openConnectModal,
    }
  } else if (!isLoading && isBridgeFeeGreaterThanInput && hasValidInput) {
    buttonProperties = {
      label: `Amount must be greater than fee`,
      onClick: null,
    }
  } else if (
    bridgeQuote.bridgeModuleName !== null &&
    !isLoading &&
    !isTyping &&
    !doesBridgeStateMatchQuote &&
    hasValidInput
  ) {
    buttonProperties = {
      label: 'Error in bridge quote',
      onClick: null,
    }

    segmentAnalyticsEvent(`[Bridge] error: state out of sync with quote`, {
      inputAmountForState: debouncedFromValue,
      originChainIdForState: fromChainId,
      originTokenForState: fromToken.symbol,
      originTokenAddressForState: fromToken.addresses[fromChainId],
      destinationChainIdForState: toChainId,
      destinationTokenForState: toToken.symbol,
      destinationTokenAddressForState: toToken.addresses[toChainId],
      bridgeQuote,
    })
  } else if (
    !isLoading &&
    isBridgeQuoteAmountGreaterThanInputForRfq &&
    hasValidInput
  ) {
    buttonProperties = {
      label: 'Invalid bridge quote',
      onClick: null,
    }
  } else if (destinationAddress && !isAddress(destinationAddress)) {
    buttonProperties = {
      label: 'Invalid Destination address',
    }
  } else if (showDestinationWarning && !isDestinationWarningAccepted) {
    buttonProperties = {
      label: 'Confirm destination address',
      onClick: () => dispatch(setIsDestinationWarningAccepted(true)),
      className: '!from-bgLight !to-bgLight',
    }
  } else if (!onSelectedChain && hasValidInput) {
    buttonProperties = {
      label: `Switch to ${chains.find((c) => c.id === fromChainId)?.name}`,
      onClick: () => switchChain({ chainId: fromChainId }),
      pendingLabel: 'Switching chains',
    }
  } else if (hasQuoteOutputChanged && !hasUserConfirmedChange) {
    buttonProperties = {
      label: 'Confirm new quote',
      onClick: () => onUserAcceptChange(),
      className:
        '!border !border-synapsePurple !from-bgLight !to-bgLight !animate-pulse',
    }
  } else if (!isApproved && hasValidInput && hasValidQuote) {
    buttonProperties = {
      onClick: approveTxn,
      label: `Approve ${fromToken?.symbol}`,
      pendingLabel: 'Approving',
    }
  } else {
    buttonProperties = {
      onClick: executeBridge,
      label: `Bridge ${fromToken?.symbol}`,
      pendingLabel: 'Bridging',
    }
  }

  return (
    buttonProperties && (
      <>
        <TransactionButton
          {...buttonProperties}
          disabled={isButtonDisabled}
          chainId={fromChainId}
        />
      </>
    )
  )
}
