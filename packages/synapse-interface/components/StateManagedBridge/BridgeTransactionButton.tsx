import { useMemo } from 'react'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { EMPTY_BRIDGE_QUOTE, EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'
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
    fromValue,
    debouncedFromValue,
    toToken,
    fromChainId,
    toChainId,
    isLoading,
    bridgeQuote,
  } = useBridgeState()

  const { isWalletPending } = useWalletState()
  const { showDestinationWarning, isDestinationWarningAccepted } =
    useBridgeDisplayState()

  const balances = usePortfolioBalances()
  const balancesForChain = balances[fromChainId]
  const balanceForToken = balancesForChain?.find(
    (t) => t.tokenAddress === fromToken?.addresses[fromChainId]
  )?.balance

  const sufficientBalance = useMemo(() => {
    if (!fromChainId || !fromToken || !toChainId || !toToken) return false
    return (
      stringToBigInt(fromValue, fromToken?.decimals[fromChainId]) <=
      balanceForToken
    )
  }, [balanceForToken, fromValue, fromChainId, toChainId, toToken])

  const fromTokenDecimals: number | undefined =
    fromToken && fromToken?.decimals[fromChainId]

  const fromValueBigInt = useMemo(() => {
    return fromTokenDecimals ? stringToBigInt(fromValue, fromTokenDecimals) : 0
  }, [fromValue, fromTokenDecimals])

  const bridgeQuoteAmountGreaterThanInputForRfq = useMemo(() => {
    return (
      bridgeQuote.bridgeModuleName === 'SynapseRFQ' &&
      bridgeQuote.outputAmount > fromValueBigInt
    )
  }, [bridgeQuote.outputAmount, fromValueBigInt])

  const stringifiedBridgeQuote = useMemo(() => {
    return constructStringifiedBridgeSelections(
      bridgeQuote.inputAmountForQuote,
      bridgeQuote.originChainId,
      bridgeQuote.originTokenForQuote,
      bridgeQuote.destChainId,
      bridgeQuote.destTokenForQuote
    )
  }, [
    bridgeQuote.inputAmountForQuote,
    bridgeQuote.originChainId,
    bridgeQuote.originTokenForQuote,
    bridgeQuote.destChainId,
    bridgeQuote.destTokenForQuote,
  ])

  const stringifiedBridgeState = useMemo(() => {
    return constructStringifiedBridgeSelections(
      debouncedFromValue,
      fromChainId,
      fromToken,
      toChainId,
      toToken
    )
  }, [debouncedFromValue, fromChainId, fromToken, toChainId, toToken])

  const bridgeStateMatchesQuote = useMemo(() => {
    return stringifiedBridgeQuote === stringifiedBridgeState
  }, [stringifiedBridgeQuote, stringifiedBridgeState])

  const isButtonDisabled =
    isLoading ||
    isWalletPending ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE ||
    (destinationAddress && !isAddress(destinationAddress)) ||
    (isConnected && !sufficientBalance) ||
    bridgeQuoteAmountGreaterThanInputForRfq ||
    !bridgeStateMatchesQuote ||
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
    bridgeQuote?.feeAmount === 0n &&
    fromValueBigInt > 0
  ) {
    buttonProperties = {
      label: `Amount must be greater than fee`,
      onClick: null,
    }
  } else if (!isLoading && !bridgeStateMatchesQuote && fromValueBigInt > 0) {
    buttonProperties = {
      label: 'Error in bridge quote',
      onClick: null,
    }
  } else if (
    !isLoading &&
    bridgeQuoteAmountGreaterThanInputForRfq &&
    fromValueBigInt > 0
  ) {
    buttonProperties = {
      label: 'Invalid bridge quote',
      onClick: null,
    }
  } else if (!isConnected && fromValueBigInt > 0) {
    buttonProperties = {
      label: `Connect Wallet to Bridge`,
      onClick: openConnectModal,
    }
  } else if (!isLoading && isConnected && !sufficientBalance) {
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
  } else if (chain?.id != fromChainId && fromValueBigInt > 0) {
    buttonProperties = {
      label: `Switch to ${chains.find((c) => c.id === fromChainId)?.name}`,
      onClick: () => switchChain({ chainId: fromChainId }),
      pendingLabel: 'Switching chains',
    }
  } else if (!isApproved && fromValueBigInt > 0 && bridgeQuote?.destQuery) {
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

const constructStringifiedBridgeSelections = (
  originAmount,
  originChainId,
  originToken,
  destChainId,
  destToken
) => {
  const state = {
    originAmount,
    originChainId,
    originToken,
    destChainId,
    destToken,
  }
  return JSON.stringify(state)
}
