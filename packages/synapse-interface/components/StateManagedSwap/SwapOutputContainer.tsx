import { useSwapState } from '@/slices/swap/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { setSwapToToken } from '@/slices/swap/reducer'
import { useSwapToTokenListArray } from './hooks/useSwapToTokenListArray'
import { AmountInput } from '@/components/ui/AmountInput'
import { useWalletState } from '@/slices/wallet/hooks'

export const SwapOutputContainer = () => {
  const { swapQuote, isLoading } = useSwapState()

  const showValue =
    swapQuote.outputAmountString === '0' ? '' : swapQuote.outputAmountString

  return (
    <BridgeSectionContainer>
      <BridgeAmountContainer>
        <SwapToTokenSelector />
        <AmountInput
          disabled={true}
          showValue={showValue}
          isLoading={isLoading}
        />
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const SwapToTokenSelector = () => {
  const { swapToToken } = useSwapState()
  const { isWalletPending } = useWalletState()

  return (
    <TokenSelector
      dataTestId="swap-destination-token"
      isOrigin={false}
      selectedItem={swapToToken}
      placeholder="Out"
      itemListFunction={useSwapToTokenListArray}
      setFunction={setSwapToToken}
      action="Swap"
      disabled={isWalletPending}
    />
  )
}
