import LoadingDots from '@/components/ui/tailwind/LoadingDots'
import { useSwapState } from '@/slices/swap/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { setSwapToToken } from '@/slices/swap/reducer'
import { useSwapToTokenListArray } from './hooks/useSwapToTokenListArray'
import { joinClassNames } from '@/utils/joinClassNames'

export const SwapOutputContainer = () => {
  const { swapQuote, isLoading } = useSwapState()

  const showValue =
    swapQuote.outputAmountString === '0' ? '' : swapQuote.outputAmountString

  const inputClassName = joinClassNames({
    unset: 'bg-transparent border-none p-0',
    layout: 'flex-1', // required for Swap Output – different from other inputs for some reason
    placeholder: 'placeholder:text-zinc-500 placeholder:dark:text-zinc-400',
    font: 'text-xl md:text-2xl font-medium',
    focus: 'focus:outline-none focus:ring-0 focus:border-none',
  })

  return (
    <BridgeSectionContainer>
      <BridgeAmountContainer>
        <SwapToTokenSelector />
        {isLoading ? (
          <LoadingDots className="opacity-50" />
        ) : (
          <input
            pattern="[0-9.]+"
            readOnly={true}
            disabled={true}
            className={inputClassName}
            placeholder="0.0000"
            value={showValue}
            name="inputRow"
            autoComplete="off"
          />
        )}
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const SwapToTokenSelector = () => {
  const { swapToToken } = useSwapState()

  return (
    <TokenSelector
      dataTestId="swap-destination-token"
      isOrigin={false}
      selectedItem={swapToToken}
      placeholder="Out"
      itemListFunction={useSwapToTokenListArray}
      setFunction={setSwapToToken}
      action="Swap"
    />
  )
}
