import { useEffect, useMemo, useState } from 'react'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'

import { TransactionButton } from '@/components/buttons/TransactionButton'
import { EMPTY_SWAP_QUOTE, EMPTY_SWAP_QUOTE_ZERO } from '@/constants/swap'
import { stringToBigInt } from '@/utils/bigint/format'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useSwapState } from '@/slices/swap/hooks'
import { SWAP_CHAIN_IDS } from '@/constants/existingSwapRoutes'

export const SwapTransactionButton = ({
  approveTxn,
  executeSwap,
  isApproved,
}) => {
  const [isConnected, setIsConnected] = useState(false)
  const { openConnectModal } = useConnectModal()

  const { chain } = useNetwork()
  const { chains, switchNetwork } = useSwitchNetwork()

  const { isConnected: isConnectedInit } = useAccount({
    onDisconnect() {
      setIsConnected(false)
    },
  })

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit])

  const {
    swapChainId,
    swapFromToken,
    swapToToken,
    swapFromValue,
    isLoading,
    swapQuote,
  } = useSwapState()

  const balances = usePortfolioBalances()
  const balancesForChain = balances[swapChainId]
  const balanceForToken = balancesForChain?.find(
    (t) => t.tokenAddress === swapFromToken?.addresses[swapChainId]
  )?.balance

  const sufficientBalance = useMemo(() => {
    if (!swapChainId || !swapFromToken || !swapToToken) return false
    return (
      stringToBigInt(swapFromValue, swapFromToken?.decimals[swapChainId]) <=
      balanceForToken
    )
  }, [balanceForToken, swapFromValue, swapChainId, swapFromToken, swapToToken])

  const isButtonDisabled =
    (isLoading && !isApproved) ||
    (isConnected && !sufficientBalance) ||
    swapQuote === EMPTY_SWAP_QUOTE_ZERO ||
    swapQuote === EMPTY_SWAP_QUOTE

  let buttonProperties

  const fromTokenDecimals: number | undefined =
    swapFromToken && swapFromToken.decimals[swapChainId]

  const fromValueBigInt = useMemo(() => {
    return fromTokenDecimals
      ? stringToBigInt(swapFromValue, fromTokenDecimals)
      : 0
  }, [swapFromValue, fromTokenDecimals, swapChainId, swapFromToken])

  if (!swapChainId) {
    buttonProperties = {
      label: 'Please select Origin network',
      onClick: null,
    }
  } else if (!SWAP_CHAIN_IDS.includes(swapChainId)) {
    buttonProperties = {
      label: 'Swaps are not available on this network',
      onClick: null,
    }
  } else if (!swapFromToken) {
    buttonProperties = {
      label: `Please select token`,
      onClick: null,
    }
  } else if (!isConnected && fromValueBigInt > 0) {
    buttonProperties = {
      label: `Connect Wallet to Swap`,
      onClick: openConnectModal,
    }
  } else if (isConnected && !sufficientBalance) {
    buttonProperties = {
      label: 'Insufficient balance',
      onClick: null,
    }
  } else if (chain?.id != swapChainId && fromValueBigInt > 0) {
    buttonProperties = {
      label: `Switch to ${chains.find((c) => c.id === swapChainId).name}`,
      onClick: () => switchNetwork(swapChainId),
      pendingLabel: 'Switching chains',
    }
  } else if (
    !isApproved &&
    !isLoading &&
    fromValueBigInt > 0 &&
    swapQuote?.quote
  ) {
    buttonProperties = {
      onClick: approveTxn,
      label: `Approve ${swapFromToken?.symbol}`,
      pendingLabel: 'Approving',
    }
  } else {
    buttonProperties = {
      onClick: executeSwap,
      label: `Swap ${swapFromToken?.symbol} for ${swapToToken?.symbol}`,
      pendingLabel: 'Swapping',
    }
  }

  return (
    buttonProperties && (
      <TransactionButton
        {...buttonProperties}
        disabled={isButtonDisabled}
        chainId={swapChainId}
      />
    )
  )
}
