import { CHAINS_BY_ID } from '@/constants/chains'
import { ChainSelector } from '@/components/ui/ChainSelector'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { useToChainListArray } from './hooks/useToChainListArray'
import { setToChainId, setToToken } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { AmountInput } from '@/components/ui/AmountInput'
import { useToTokenListArray } from '@/components/StateManagedBridge/hooks/useToTokenListArray'

export const OutputContainer = ({}) => {
  const { bridgeQuote, isLoading } = useBridgeState()

  const showValue =
    bridgeQuote?.outputAmountString === '0'
      ? ''
      : bridgeQuote?.outputAmountString

  return (
    <BridgeSectionContainer>
      <div className="flex items-center justify-between">
        <ToChainSelector />
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
  const { toChainId } = useBridgeState()

  return (
    <ChainSelector
      dataTestId="bridge-destination-chain"
      isOrigin={false}
      selectedItem={CHAINS_BY_ID[toChainId]}
      label="To"
      itemListFunction={useToChainListArray}
      setFunction={setToChainId}
      action="Bridge"
    />
  )
}

const ToTokenSelector = () => {
  const { toToken } = useBridgeState()

  return (
    <TokenSelector
      dataTestId="bridge-destination-token"
      isOrigin={false}
      selectedItem={toToken}
      placeholder="In"
      itemListFunction={useToTokenListArray}
      setFunction={setToToken}
      action="Bridge"
    />
  )
}
