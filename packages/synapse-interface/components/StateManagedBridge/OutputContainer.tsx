import { useAccount } from 'wagmi'
import { useMemo } from 'react'

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
import { useBridgeValidations } from './hooks/useBridgeValidations'

interface OutputContainerProps {
  isQuoteStale: boolean
}

export const OutputContainer = ({ isQuoteStale }: OutputContainerProps) => {
  const { address } = useAccount()
  const { bridgeQuote, isLoading } = useBridgeQuoteState()
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

  const inputClassName = isQuoteStale ? 'opacity-50' : undefined

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
          className={inputClassName}
        />
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
