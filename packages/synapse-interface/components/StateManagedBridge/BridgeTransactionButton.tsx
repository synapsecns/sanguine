import { useEffect, useState } from 'react'
import { useSelector } from 'react-redux'
import { useAccount } from 'wagmi'
import { isAddress } from 'viem'
import { useConnectModal } from '@rainbow-me/rainbowkit'

import { TransactionButton } from '@/components/buttons/TransactionButton'
import { RootState } from '@/store/store'

import { useBridgeState, useBridgeStatus } from '@/slices/bridge/hooks'

export const BridgeTransactionButton = ({ executeBridge }) => {
  const [isConnected, setIsConnected] = useState(false)
  const { openConnectModal } = useConnectModal()

  const { isConnected: isConnectedInit } = useAccount({
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
    toToken,
    fromChainId,
    toChainId,
    isLoading,
    bridgeQuote,
  } = useBridgeState()

  const { showDestinationAddress } = useSelector(
    (state: RootState) => state.bridgeDisplay
  )

  const {
    hasEnoughBalance,
    hasEnoughApproved,
    hasInputAmount,
    hasSelectedNetwork,
  } = useBridgeStatus()

  const isButtonDisabled =
    isLoading ||
    (destinationAddress && !isAddress(destinationAddress)) ||
    (showDestinationAddress && !destinationAddress) ||
    (isConnected && !hasSelectedNetwork) ||
    (isConnected && !hasEnoughBalance) ||
    (isConnected && !hasEnoughApproved)

  let buttonProperties: {
    label: string | JSX.Element
    pendingLabel?: string | JSX.Element
    onClick: any
    toolTipLabel?: string
  } = {
    label: 'Bridge',
    pendingLabel: 'Bridging',
    onClick: null,
  }

  if (!isConnected) {
    buttonProperties = {
      ...buttonProperties,
      label: 'Connect Wallet',
      onClick: openConnectModal,
    }
  } else if (!fromChainId) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Origin chain required',
    }
  } else if (!fromToken) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Origin token required',
    }
  } else if (!toChainId) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Destination chain required',
    }
  } else if (!toToken) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Destination token required',
    }
  } else if (!hasInputAmount) {
    buttonProperties = {
      ...buttonProperties,
    }
  } else if (isConnected && !hasEnoughBalance) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Insufficient balance',
    }
  } else if (showDestinationAddress && !destinationAddress) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Please add valid destination address',
    }
  } else if (destinationAddress && !isAddress(destinationAddress)) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Invalid destination address',
    }
  } else if (isConnected && bridgeQuote.outputAmount === 0n) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'No route found',
    }
  } else if (isConnected && !hasEnoughApproved) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Token approval required to bridge',
    }
  } else if (isConnected && hasEnoughApproved && hasInputAmount) {
    buttonProperties = {
      ...buttonProperties,
      onClick: executeBridge,
    }
  } else {
    buttonProperties = {
      ...buttonProperties,
    }
  }

  return (
    buttonProperties && (
      <TransactionButton
        style={{
          background: 'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
          borderColor: '#9B6DD7',
        }}
        {...buttonProperties}
        disabled={isButtonDisabled}
        chainId={fromChainId}
      />
    )
  )
}
