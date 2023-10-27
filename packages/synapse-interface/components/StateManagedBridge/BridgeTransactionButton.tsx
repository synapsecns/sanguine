import { useEffect, useState } from 'react'
import { useSelector } from 'react-redux'
import { useAccount } from 'wagmi'
import { isAddress } from 'viem'
import { useConnectModal } from '@rainbow-me/rainbowkit'

import { TransactionButton } from '@/components/buttons/TransactionButton'
import { EMPTY_BRIDGE_QUOTE, EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'
import { RootState } from '@/store/store'

import { useBridgeState, useBridgeStatus } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'

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
  } else if (!toChainId) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Destination chain required',
    }
  } else if (!fromToken) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Origin token required',
    }
  } else if (!toToken) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Destination token required',
    }
  } else if (!hasInputAmount) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Input value required',
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
      toolTipLabel: 'Please add valid destination address',
    }
  } else if (destinationAddress && !isAddress(destinationAddress)) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Invalid destination address',
    }
  } else if (!hasSelectedNetwork && hasInputAmount) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: `Switch to ${CHAINS_BY_ID[fromChainId].name}`,
    }
  } else if (isConnected && !hasEnoughApproved) {
    buttonProperties = {
      ...buttonProperties,
      toolTipLabel: 'Token approval required to bridge',
    }
  } else if (isConnected && hasEnoughApproved && hasInputAmount) {
    buttonProperties = {
      onClick: executeBridge,
      label: (
        <div className="flex flex-col space-y-1">
          <div>Bridge</div>
          <div className="text-sm text-[#bfbcc2]">
            {fromValue} {fromToken?.symbol} to {CHAINS_BY_ID[toChainId]?.name}
          </div>
        </div>
      ),
      pendingLabel: 'Bridging',
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
          background:
            'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
          border: '1px solid #9B6DD7',
          borderRadius: '4px',
        }}
        {...buttonProperties}
        disabled={isButtonDisabled}
        chainId={fromChainId}
      />
    )
  )
}
