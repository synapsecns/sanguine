import { isAddress } from 'viem'
import { useEffect, useState } from 'react'
import { useAccount, useAccountEffect, useSwitchChain } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { setIsDestinationWarningAccepted } from '@/slices/bridgeDisplaySlice'
import { useBridgeDisplayState, useBridgeState } from '@/slices/bridge/hooks'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { useBridgeValidations } from './hooks/useBridgeValidations'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useConfirmNewBridgePrice } from './hooks/useConfirmNewBridgePrice'
import { HYPERLIQUID } from '@/constants/chains/master'
import { HYPERLIQUID_MINIMUM_DEPOSIT } from '@/constants'

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

  const t = useTranslations('Bridge')

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
  const { isPendingConfirmChange, onUserAcceptChange } =
    useConfirmNewBridgePrice()

  const { isWalletPending } = useWalletState()
  const { showDestinationWarning, isDestinationWarningAccepted } =
    useBridgeDisplayState()

  const hasHyperliquidMinDeposit =
    toChainId === HYPERLIQUID.id
      ? Number(debouncedFromValue) > HYPERLIQUID_MINIMUM_DEPOSIT
        ? true
        : false
      : true

  const {
    hasValidInput,
    hasValidQuote,
    hasSufficientBalance,
    doesBridgeStateMatchQuote,
    onSelectedChain,
  } = useBridgeValidations()

  const isButtonDisabled =
    isBridgePaused ||
    isTyping ||
    isLoading ||
    isWalletPending ||
    !hasValidInput ||
    !doesBridgeStateMatchQuote ||
    (isConnected && !hasValidQuote) ||
    (isConnected && !hasSufficientBalance) ||
    (isConnected && isQuoteStale) ||
    (destinationAddress && !isAddress(destinationAddress)) ||
    !hasHyperliquidMinDeposit

  let buttonProperties

  if (isBridgePaused) {
    buttonProperties = {
      label: t('Bridge paused'),
      onClick: null,
    }
  } else if (!fromChainId) {
    buttonProperties = {
      label: t('Please select Origin Network'),
      onClick: null,
    }
  } else if (!toChainId) {
    buttonProperties = {
      label: t('Please select Destination network'),
      onClick: null,
    }
  } else if (!fromToken) {
    buttonProperties = {
      label: t('Please select an Origin token'),
      onClick: null,
    }
  } else if (isConnected && !hasSufficientBalance) {
    buttonProperties = {
      label: t('Insufficient balance'),
      onClick: null,
    }
  } else if (isLoading && hasValidQuote) {
    buttonProperties = {
      label: isPendingConfirmChange
        ? t('Confirm new quote')
        : t('Bridge {symbol}', { symbol: fromToken?.symbol }),
      pendingLabel: t('Bridge {symbol}', { symbol: fromToken?.symbol }),
      onClick: null,
      className: `
      ${
        isPendingConfirmChange
          ? '!outline !outline-1 !outline-synapsePurple !outline-offset-[-1px] !from-bgLight !to-bgLight'
          : '!bg-gradient-to-r !from-fuchsia-500 !to-purple-500 dark:!to-purple-600'
      }
      !opacity-100`,
    }
  } else if (isLoading) {
    buttonProperties = {
      label: t('Bridge {symbol}', { symbol: fromToken?.symbol }),
      pendingLabel: t('Bridge {symbol}', { symbol: fromToken?.symbol }),
      onClick: null,
    }
  } else if (!isConnected && hasValidInput) {
    buttonProperties = {
      label: t('Connect Wallet to Bridge'),
      onClick: openConnectModal,
    }
  } else if (!hasHyperliquidMinDeposit) {
    buttonProperties = {
      label: `${HYPERLIQUID_MINIMUM_DEPOSIT} USDC Minimum`,
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
      label: t('Error in bridge quote'),
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
  } else if (destinationAddress && !isAddress(destinationAddress)) {
    buttonProperties = {
      label: t('Invalid Destination address'),
    }
  } else if (showDestinationWarning && !isDestinationWarningAccepted) {
    buttonProperties = {
      label: t('Confirm destination address'),
      onClick: () => dispatch(setIsDestinationWarningAccepted(true)),
      className: '!from-bgLight !to-bgLight',
    }
  } else if (!onSelectedChain && hasValidInput) {
    buttonProperties = {
      label: t('Switch to {chainName}', {
        chainName: chains.find((c) => c.id === fromChainId)?.name,
      }),
      onClick: () => switchChain({ chainId: fromChainId }),
      pendingLabel: t('Switching chains'),
    }
  } else if (isApproved && hasValidQuote && isPendingConfirmChange) {
    buttonProperties = {
      label: t('Confirm new quote'),
      onClick: () => onUserAcceptChange(),
      className:
        '!outline !outline-1 !outline-synapsePurple !outline-offset-[-1px] !from-bgLight !to-bgLight',
    }
  } else if (!isApproved && hasValidInput && hasValidQuote) {
    buttonProperties = {
      onClick: approveTxn,
      label: t('Approve {symbol}', { symbol: fromToken?.symbol }),
      pendingLabel: t('Approving'),
    }
  } else {
    buttonProperties = {
      onClick: executeBridge,
      label: t('Bridge {symbol}', { symbol: fromToken?.symbol }),
      pendingLabel: t('Bridging'),
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
