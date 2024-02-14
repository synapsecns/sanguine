
import { useSwapState } from '@/slices/swap/hooks'

import LoadingDots from '@tw/LoadingDots'
import { SwapToTokenSelector } from './SwapToTokenSelector'
import { BridgeSwapOutputNumber } from '@/components/BridgeSwapOutputNumber'


export const SwapOutputContainer = () => {
  const { swapQuote, isLoading } = useSwapState()


  return (
    <div className="mt-[1.125rem] p-md text-left rounded-md bg-bgBase/10 ring-1 ring-white/10">
      <div className="flex h-16 mt-2 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-md
            w-full h-16
            rounded-md
            border border-transparent
          `}
        >
          <SwapToTokenSelector />
          <div className="flex ml-4">
            {isLoading ? (
              <LoadingDots className="opacity-50" />
            ) : (
              <BridgeSwapOutputNumber
                quote={swapQuote}
              />
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
