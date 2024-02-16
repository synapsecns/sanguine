import LoadingDots from '@tw/LoadingDots'
import { LoadingHelix } from '@tw/LoadingHelix'
import { OutputNumber } from '@/components/bridgeSwap/OutputNumber'
import { LoadingInfinity } from '../ui/tailwind/LoadingInfinity'


export const GenericOutputContainer = ({
  chainSelector,
  tokenSelector,
  isLoading,
  quote
} : {
  chainSelector?: any,
  tokenSelector: any,
  isLoading?: boolean,
  quote?: any
}) => {
  return (
    <div className="mt-[1.125rem] p-md text-left rounded-md bg-bgBase/10 ring-1 ring-white/10">
      {chainSelector &&
        <div className="flex items-center justify-between mb-3">
            {chainSelector}
        </div>
      }
      <div className="flex h-16 space-x-2" /**was mb-2 on out */>
        <div
          className={`
            flex flex-grow items-center
            pl-md
            w-full h-16
            rounded-md
            border border-transparent
          `}
        >
          {tokenSelector}
          <div className="flex ml-4">
            {/* <LoadingHelix /> */}
            {isLoading ? (
              // <div className="animate-pulse">
              //   <OutputNumber quote={{outputAmountString:"0.0000"}}/>
              // </div>
              // <LoadingDots className="opacity-50" />
              <LoadingHelix />
            ) : (
              <OutputNumber quote={quote}/>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}