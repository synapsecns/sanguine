import { useMemo } from 'react'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { useAccount, useAccountEffect, useSwitchChain } from 'wagmi'
import { useEffect, useState } from 'react'
import { isAddress } from 'viem'

import { useConnectModal } from '@rainbow-me/rainbowkit'
import { stringToBigInt } from '@/utils/bigint/format'
import { useBridgeDisplayState, useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useAppDispatch } from '@/store/hooks'
import { setIsDestinationWarningAccepted } from '@/slices/bridgeDisplaySlice'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { useBridgeSelections } from './hooks/useBridgeSelections'
import { useBridgeValidations } from './hooks/useBridgeValidations'

export const BridgeTransactionButton = ({
  approveTxn,
  executeBridge,
  isApproved,
  isBridgePaused,
}) => {
  const dispatch = useAppDispatch()
  const [isConnected, setIsConnected] = useState(false)
  const { openConnectModal } = useConnectModal()

  const { chain, isConnected: isConnectedInit } = useAccount()
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
    // fromValue,
    // toToken,
    fromChainId,
    toChainId,
  } = useBridgeState()

  const { isLoading, bridgeQuote } = useBridgeQuoteState()

  const { isWalletPending } = useWalletState()
  const { showDestinationWarning, isDestinationWarningAccepted } =
    useBridgeDisplayState()

  const { fromTokenBalance, debouncedFromValueBigInt } = useBridgeSelections()
  const {
    hasValidInput,
    hasValidQuote,
    hasSufficientBalance,
    doesChainSelectionsMatchBridgeQuote,
    isBridgeQuoteAmountGreaterThanInputForRfq,
    onSelectedChain,
  } = useBridgeValidations()

  // const balances = usePortfolioBalances()
  // const balancesForChain = balances[fromChainId]
  // const balanceForToken = balancesForChain?.find(
  //   (t) => t.tokenAddress === fromToken?.addresses[fromChainId]
  // )?.balance

  // const sufficientBalance = useMemo(() => {
  //   if (!fromChainId || !fromToken || !toChainId || !toToken) return false
  //   return (
  //     stringToBigInt(fromValue, fromToken?.decimals[fromChainId]) <=
  //     balanceForToken
  //   )
  // }, [balanceForToken, fromValue, fromChainId, toChainId, toToken])

  // const fromTokenDecimals: number | undefined =
  //   fromToken && fromToken?.decimals[fromChainId]

  // const fromValueBigInt = useMemo(() => {
  //   return fromTokenDecimals ? stringToBigInt(fromValue, fromTokenDecimals) : 0
  // }, [fromValue, fromTokenDecimals])

  // const bridgeQuoteAmountGreaterThanInputForRfq = useMemo(() => {
  //   return (
  //     bridgeQuote.bridgeModuleName === 'SynapseRFQ' &&
  //     bridgeQuote.outputAmount > fromValueBigInt
  //   )
  // }, [bridgeQuote.outputAmount, fromValueBigInt])

  // const chainSelectionsMatchBridgeQuote = useMemo(() => {
  //   return (
  //     fromChainId === bridgeQuote.originChainId &&
  //     toChainId === bridgeQuote.destChainId
  //   )
  // }, [fromChainId, toChainId, bridgeQuote])

  const isButtonDisabled =
    isLoading ||
    isWalletPending ||
    !hasValidQuote ||
    (destinationAddress && !isAddress(destinationAddress)) ||
    (isConnected && !hasSufficientBalance) ||
    isBridgeQuoteAmountGreaterThanInputForRfq ||
    !doesChainSelectionsMatchBridgeQuote ||
    isBridgePaused

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
  } else if (isLoading) {
    buttonProperties = {
      label: `Bridge ${fromToken?.symbol}`,
      onClick: null,
    }
  } else if (
    !isLoading &&
    hasValidQuote &&
    hasValidInput &&
    bridgeQuote.feeAmount > debouncedFromValueBigInt
  ) {
    buttonProperties = {
      label: `Amount must be greater than fee`,
      onClick: null,
    }
  } else if (
    !isLoading &&
    !doesChainSelectionsMatchBridgeQuote &&
    hasValidInput
  ) {
    buttonProperties = {
      label: 'Please reset chain selection',
      onClick: null,
    }
  } else if (
    !isLoading &&
    isBridgeQuoteAmountGreaterThanInputForRfq &&
    hasValidInput
  ) {
    buttonProperties = {
      label: 'Invalid bridge quote',
      onClick: null,
    }
  } else if (!isConnected && hasValidInput) {
    buttonProperties = {
      label: `Connect Wallet to Bridge`,
      onClick: openConnectModal,
    }
  } else if (!isLoading && isConnected && !hasSufficientBalance) {
    buttonProperties = {
      label: 'Insufficient balance',
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
