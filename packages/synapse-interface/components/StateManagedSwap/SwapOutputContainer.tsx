
import { useSwapState } from '@/slices/swap/hooks'

import LoadingDots from '@tw/LoadingDots'
import { SwapToTokenSelector } from './SwapToTokenSelector'


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
            border border-white border-opacity-20
          `}
        >
          <SwapToTokenSelector />
          <div className="flex ml-4">
            {isLoading ? (
              <LoadingDots className="opacity-50" />
            ) : (
              <input
                pattern="[0-9.]+"
                disabled={true}
                className={`
                  focus:outline-none
                  focus:ring-0
                  focus:border-none
                  border-none
                  p-0
                  bg-transparent
                  max-w-[190px]
                placeholder:text-[#88818C]
                text-white text-opacity-80 text-xl md:text-2xl font-medium
                `}
                placeholder="0.0000"
                value={
                  swapQuote.outputAmountString === '0'
                    ? ''
                    : swapQuote.outputAmountString
                }
                name="inputRow"
                autoComplete="off"
              />
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
