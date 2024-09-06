import { useAccount } from 'wagmi'
import { useMemo, useEffect, useState } from 'react'

import { ChainSelector } from '@/components/ui/ChainSelector'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { AmountInput } from '@/components/ui/AmountInput'
import { useToChainListArray } from '@/components/StateManagedBridge/hooks/useToChainListArray'
import { useToTokenListArray } from '@/components/StateManagedBridge/hooks/useToTokenListArray'
import { DestinationAddressInput } from '@/components/StateManagedBridge/DestinationAddressInput'
import { CHAINS_BY_ID } from '@/constants/chains'
import { setToChainId, setToToken } from '@/slices/bridge/reducer'
import { useBridgeDisplayState, useBridgeState } from '@/slices/bridge/hooks'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { BridgeQuote } from '@/utils/types'
import { useBridgeValidations } from './hooks/useBridgeValidations'
import { useConfirmNewBridgePrice } from './hooks/useConfirmNewBridgePrice'

export const OutputContainer = () => {
  const { address } = useAccount()
  const { bridgeQuote, previousBridgeQuote, isLoading } = useBridgeQuoteState()
  const { showDestinationAddress } = useBridgeDisplayState()
  const { hasValidInput, hasValidQuote } = useBridgeValidations()

  const showValue = useMemo(() => {
    if (!hasValidInput) {
      return ''
    } else if (hasValidQuote) {
      return bridgeQuote?.outputAmountString
    } else {
      return ''
    }
  }, [bridgeQuote, hasValidInput, hasValidQuote])

  const { hasSameSelectionsAsPreviousQuote } = useConfirmNewBridgePrice()

  return (
    <BridgeSectionContainer>
      <div className="flex items-center justify-between">
        <ToChainSelector />
        {showDestinationAddress ? (
          <DestinationAddressInput connectedAddress={address} />
        ) : null}
      </div>

      <BridgeAmountContainer>
        <ToTokenSelector />
        <AmountInput
          disabled={true}
          showValue={showValue}
          isLoading={isLoading}
        />
        {hasValidQuote && !isLoading && (
          <AnimatedCircle bridgeQuoteId={bridgeQuote?.id} />
        )}
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const ToChainSelector = () => {
  const { toChainId } = useBridgeState()
  const { isWalletPending } = useWalletState()

  return (
    <ChainSelector
      dataTestId="bridge-destination-chain"
      isOrigin={false}
      selectedItem={CHAINS_BY_ID[toChainId]}
      label="To"
      itemListFunction={useToChainListArray}
      setFunction={setToChainId}
      action="Bridge"
      disabled={isWalletPending}
    />
  )
}

const ToTokenSelector = () => {
  const { toToken } = useBridgeState()
  const { isWalletPending } = useWalletState()

  return (
    <TokenSelector
      dataTestId="bridge-destination-token"
      isOrigin={false}
      selectedItem={toToken}
      placeholder="In"
      itemListFunction={useToTokenListArray}
      setFunction={setToToken}
      action="Bridge"
      disabled={isWalletPending}
    />
  )
}

const AnimatedCircle = ({ bridgeQuoteId }) => {
  const [animationKey, setAnimationKey] = useState(0) // Key to force re-render

  useEffect(() => {
    // Whenever the `quote` prop changes, restart the animation by updating the key
    setAnimationKey((prevKey) => prevKey + 1)
  }, [bridgeQuoteId])

  return (
    <svg
      key={animationKey} // Update key to trigger re-render
      width="36"
      height="36"
      viewBox="-12 -12 24 24"
      stroke="currentcolor"
      strokeOpacity=".33"
      fill="none"
      className="-rotate-90 -scale-y-100"
    >
      {/* Inner circle remains visible, no opacity animation */}
      <circle r="8" />

      {/* Outer circle with stroke animation */}
      <circle r="8" strokeDasharray="1" pathLength="1">
        <animate
          attributeName="stroke-dashoffset"
          values="1; 1; 2" // Stays outlined after completion
          dur="15s"
          keyTimes="0; .67; 1"
          fill="freeze" // Ensures the final state is retained
        />
      </circle>
    </svg>
  )
}
