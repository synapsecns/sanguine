import { useEffect, useMemo, useState } from 'react'
import { useAccount, useAccountEffect, useSwitchChain } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { useTranslations } from 'next-intl'

import { TransactionButton } from '@/components/buttons/TransactionButton'
import { EMPTY_SWAP_QUOTE, EMPTY_SWAP_QUOTE_ZERO } from '@/constants/swap'
import { stringToBigInt } from '@/utils/bigint/format'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useSwapState } from '@/slices/swap/hooks'
import { SWAP_CHAIN_IDS } from '@/constants/existingSwapRoutes'

export const SwapTransactionButton = ({
  isTyping,
  approveTxn,
  executeSwap,
  isApproved,
  isSwapPaused,
}) => {
  const [isConnected, setIsConnected] = useState(false)
  const { openConnectModal } = useConnectModal()

  const t = useTranslations('Swap')

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
    swapQuote === EMPTY_SWAP_QUOTE ||
    isSwapPaused ||
    isTyping

  let buttonProperties

  const fromTokenDecimals: number | undefined =
    swapFromToken && swapFromToken.decimals[swapChainId]

  const fromValueBigInt = useMemo(() => {
    return fromTokenDecimals
      ? stringToBigInt(swapFromValue, fromTokenDecimals)
      : 0
  }, [swapFromValue, fromTokenDecimals, swapChainId, swapFromToken])

  if (isSwapPaused) {
    buttonProperties = {
      label: t('Swap paused'),
      onClick: null,
    }
  } else if (!swapChainId) {
    buttonProperties = {
      label: t('Please select Origin network'),
      onClick: null,
    }
  } else if (!SWAP_CHAIN_IDS.includes(swapChainId)) {
    buttonProperties = {
      label: t('Swaps are not available on this network'),
      onClick: null,
    }
  } else if (!swapFromToken) {
    buttonProperties = {
      label: t('Please select token'),
      onClick: null,
    }
  } else if (!isConnected && fromValueBigInt > 0) {
    buttonProperties = {
      label: t('Connect Wallet to Swap'),
      onClick: openConnectModal,
    }
  } else if (isConnected && !sufficientBalance) {
    buttonProperties = {
      label: t('Insufficient balance'),
      onClick: null,
    }
  } else if (chain?.id != swapChainId && fromValueBigInt > 0) {
    buttonProperties = {
      label: `${t('Switch to')} ${
        chains.find((c) => c.id === swapChainId).name
      }`,
      onClick: () => switchChain({ chainId: swapChainId }),
      pendingLabel: t('Switching chains'),
    }
  } else if (
    !isApproved &&
    !isLoading &&
    fromValueBigInt > 0 &&
    swapQuote?.quote
  ) {
    buttonProperties = {
      onClick: approveTxn,
      label: `${t('Approve')} ${swapFromToken?.symbol}`,
      pendingLabel: t('Approving'),
    }
  } else {
    buttonProperties = {
      onClick: executeSwap,
      label: `${t('Swap')} ${swapFromToken?.symbol} for ${swapToToken?.symbol}`,
      pendingLabel: t('Swapping'),
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
