import { useBridgeState } from '@/slices/bridge/hooks'

import LoadingDots from '@tw/LoadingDots'

import { ToChainSelector } from './ToChainSelector'
import { ToTokenSelector } from './ToTokenSelector'
import { BridgeSwapOutputNumber } from '@/components/BridgeSwapOutputNumber'


export const OutputContainer = () => {
  const { bridgeQuote, isLoading } = useBridgeState()

  return (
    <div className="mt-[1.125rem] p-md text-left rounded-md bg-bgBase/10 ring-1 ring-white/10">
      <div className="flex items-center justify-between mb-3">
        <ToChainSelector />
      </div>

      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-md
            w-full h-16
            rounded-md
            border border-transparent
          `}
        >
          <ToTokenSelector />
          <div className="flex ml-4">
            {isLoading ? (
              <LoadingDots className="opacity-50" />
            ) : (
              <BridgeSwapOutputNumber
                quote={bridgeQuote}
              />
            )}
          </div>
        </div>
      </div>
    </div>
  )
}

