import { useSelector } from 'react-redux'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { EMPTY_BRIDGE_QUOTE, EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'
import { RootState } from '../../store/store'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useEffect, useState } from 'react'
import { isAddress } from '@ethersproject/address'

import {
  useConnectModal,
  useAccountModal,
  useChainModal,
} from '@rainbow-me/rainbowkit'

export const BridgeTransactionButton = ({
  approveTxn,
  executeBridge,
  isApproved,
}) => {
  // TODO: This is only implemented this way to fix a Next Hydration Error
  const [isConnected, setIsConnected] = useState(false) // Initialize to false
  const { openConnectModal } = useConnectModal()

  const { chain } = useNetwork()
  const { chains, error, pendingChainId, switchNetwork } = useSwitchNetwork()

  const { address, isConnected: isConnectedInit } = useAccount({
    onDisconnect() {
      setIsConnected(false);
    },
  });

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit]);

  // Get state from Redux store
  const {
    destinationAddress,
    fromToken,
    fromValue,
    fromChainId,
    isLoading,
    bridgeQuote,
  } = useSelector((state: RootState) => state.bridge)

  const { showDestinationAddress } = useSelector(
    (state: RootState) => state.bridgeDisplay
  )

  const isButtonDisabled =
    isLoading ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE ||
    (destinationAddress && !isAddress(destinationAddress)) ||
    (showDestinationAddress && !destinationAddress)

  let buttonProperties

  if (!isLoading && bridgeQuote?.feeAmount?.eq(0) && fromValue.gt(0)) {
    buttonProperties = {
      label: `Amount must be greater than fee`,
      onClick: null,
    }
  } else if (!isConnected && fromValue.gt(0)) {
    buttonProperties = {
      label: `Connect Wallet to Bridge`,
      onClick: openConnectModal,
    }
  } else if (showDestinationAddress && !destinationAddress) {
    buttonProperties = {
      label: 'Please add valid destination address',
    }
  } else if (destinationAddress && !isAddress(destinationAddress)) {
    buttonProperties = {
      label: 'Invalid destination address',
    }
  } else if (chain?.id != fromChainId && fromValue.gt(0)) {
    buttonProperties = {
      label: `Switch to ${chains.find((c) => c.id === fromChainId).name}`,
      onClick: () => switchNetwork(fromChainId),
      pendingLabel: 'Switching chains',
    }
  } else if (!isApproved) {
    buttonProperties = {
      onClick: approveTxn,
      label: `Approve ${fromToken.symbol}`,
      pendingLabel: 'Approving',
    }
  } else {
    buttonProperties = {
      onClick: executeBridge,
      label: `Bridge ${fromToken.symbol}`,
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
