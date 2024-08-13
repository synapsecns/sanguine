import { useAccount } from 'wagmi'

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

export const OutputContainer = () => {
  const { address } = useAccount()
  const { bridgeQuote, isLoading } = useBridgeState()
  const { showDestinationAddress } = useBridgeDisplayState()

  const showValue =
    bridgeQuote?.outputAmountString === '0'
      ? ''
      : bridgeQuote?.outputAmountString

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
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const ToChainSelector = () => {
  const { toChainId, isWalletPending } = useBridgeState()

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
  const { toToken, isWalletPending } = useBridgeState()

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
