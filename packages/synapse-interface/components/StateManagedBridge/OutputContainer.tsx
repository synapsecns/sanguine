import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'

import LoadingSpinner from '../ui/tailwind/LoadingSpinner'
import ToChainSelect from './ToChainSelect'
import ToTokenSelect from './ToTokenSelect'

export const OutputContainer = ({}) => {
  const { bridgeQuote, isLoading } = useSelector(
    (state: RootState) => state.bridge
  )

  return (
    <div
      className={`
        mt-6
        text-left px-2 sm:px-4 pt-4 pb-1 rounded-xl
        bg-bgLight
      `}
    >
      <ToChainSelect />
      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-3 sm:pl-4
            w-full h-16
            rounded-xl
            border border-white border-opacity-20
          `}
        >
          <ToTokenSelect />
          <div className="flex ml-4 min-w-[190px]">
            {isLoading ? (
              <LoadingSpinner className="opacity-50" />
            ) : (
              <input
                pattern="[0-9.]+"
                disabled={true}
                className={`
                focus:outline-none
                bg-transparent
                max-w-[190px]
               placeholder:text-[#88818C]
               text-white text-opacity-80 text-lg md:text-2xl lg:text-2xl font-medium
              `}
                placeholder="0.0000"
                value={
                  bridgeQuote.outputAmountString === '0'
                    ? ''
                    : bridgeQuote.outputAmountString
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
