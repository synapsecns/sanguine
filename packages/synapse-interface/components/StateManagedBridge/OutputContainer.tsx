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
          <AnimatedProgressCircle bridgeQuoteId={bridgeQuote?.id} />
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

const AnimatedProgressCircle = ({ bridgeQuoteId }) => {
  const [animationKey, setAnimationKey] = useState(0)

  useEffect(() => {
    setAnimationKey((prevKey) => prevKey + 1)
  }, [bridgeQuoteId])

  return (
    <svg
      key={animationKey}
      width="36"
      height="36"
      viewBox="-12 -12 24 24"
      stroke="currentcolor"
      strokeOpacity=".33"
      fill="none"
      className="-rotate-90 -scale-y-100"
    >
      <circle r="8">
        <animate
          attributeName="opacity"
          values="0; 0; 1"
          dur="15s"
          fill="freeze"
          keyTimes="0; .5; 1"
        />
      </circle>
      <circle r="8" strokeDasharray="1" pathLength="1">
        <animate
          attributeName="stroke-dashoffset"
          values="1; 1; 2"
          dur="15s"
          keyTimes="0; .67; 1"
          fill="freeze"
        />
      </circle>
    </svg>
  )
}
