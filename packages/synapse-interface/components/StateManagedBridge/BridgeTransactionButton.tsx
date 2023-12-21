import { useSelector } from 'react-redux'
import { useMemo } from 'react'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { EMPTY_BRIDGE_QUOTE, EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'
import { RootState } from '@/store/store'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useEffect, useState } from 'react'
import { isAddress } from '@ethersproject/address'
import {} from 'wagmi'

import {
  useConnectModal,
  useAccountModal,
  useChainModal,
} from '@rainbow-me/rainbowkit'
import { stringToBigInt } from '@/utils/bigint/format'
import { useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { PAUSED_FROM_CHAIN_IDS, PAUSED_TO_CHAIN_IDS } from '@/constants/chains'

export const BridgeTransactionButton = ({
  approveTxn,
  executeBridge,
  isApproved,
}) => {
  const [isConnected, setIsConnected] = useState(false)
  const { openConnectModal } = useConnectModal()

  const { chain } = useNetwork()
  const { chains, error, pendingChainId, switchNetwork } = useSwitchNetwork()

  const { address, isConnected: isConnectedInit } = useAccount({
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
    toToken,
    fromChainId,
    toChainId,
    isLoading,
    bridgeQuote,
  } = useBridgeState()

  const { showDestinationAddress } = useSelector(
    (state: RootState) => state.bridgeDisplay
  )

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

  const isButtonDisabled =
    isLoading ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE ||
    (destinationAddress && !isAddress(destinationAddress)) ||
    (showDestinationAddress && !destinationAddress) ||
    (isConnected && !sufficientBalance) ||
    PAUSED_FROM_CHAIN_IDS.includes(fromChainId) ||
    PAUSED_TO_CHAIN_IDS.includes(toChainId)

  let buttonProperties

  const fromTokenDecimals: number | undefined =
    fromToken && fromToken?.decimals[fromChainId]

  const fromValueBigInt = useMemo(() => {
    return fromTokenDecimals ? stringToBigInt(fromValue, fromTokenDecimals) : 0
  }, [fromValue, fromTokenDecimals])

  if (!fromChainId) {
    buttonProperties = {
      label: 'Please select Origin network',
      onClick: null,
    }
  } else if (!toChainId) {
    buttonProperties = {
      label: 'Please select Destination network',
      onClick: null,
    }
  } else if (
    PAUSED_FROM_CHAIN_IDS.includes(fromChainId) ||
    PAUSED_TO_CHAIN_IDS.includes(toChainId)
  ) {
    buttonProperties = {
      label: `Bridge unavailable due to network issues`,
      onClick: null,
    }
  } else if (!fromToken) {
    buttonProperties = {
      label: `Unsupported Network`,
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
  } else if (!isConnected && fromValueBigInt > 0) {
    buttonProperties = {
      label: `Connect Wallet to Bridge`,
      onClick: openConnectModal,
    }
  } else if (isConnected && !sufficientBalance) {
    buttonProperties = {
      label: 'Insufficient balance',
      onClick: null,
    }
  } else if (showDestinationAddress && !destinationAddress) {
    buttonProperties = {
      label: 'Please add valid destination address',
    }
  } else if (destinationAddress && !isAddress(destinationAddress)) {
    buttonProperties = {
      label: 'Invalid destination address',
    }
  } else if (chain?.id != fromChainId && fromValueBigInt > 0) {
    buttonProperties = {
      label: `Switch to ${chains.find((c) => c.id === fromChainId)?.name}`,
      onClick: () => switchNetwork(fromChainId),
      pendingLabel: 'Switching chains',
    }
  } else if (!isApproved) {
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
      <TransactionButton
        {...buttonProperties}
        disabled={isButtonDisabled}
        chainId={fromChainId}
      />
    )
  )
}
