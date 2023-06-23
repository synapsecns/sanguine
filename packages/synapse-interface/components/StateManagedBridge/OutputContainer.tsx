import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'

import SelectTokenDropdown from '@/components/input/TokenAmountInput/SelectTokenDropdown'
import { useDispatch } from 'react-redux'
import { setShowToTokenSlideOver } from '@/slices/bridgeDisplaySlice'
import LoadingSpinner from '../ui/tailwind/LoadingSpinner'
import { DestinationChainLabel } from './DestinationChainLabel'
import SwitchButton from '../buttons/SwitchButton'
import { setFromChainId, setToChainId } from '@/slices/bridgeSlice'

export const OutputContainer = ({}) => {
  const dispatch = useDispatch()
  const {
    fromChainId,
    toChainId,
    toToken,
    bridgeQuote,
    toChainIds,
    isLoading,
  } = useSelector((state: RootState) => state.bridge)

  return (
    <div
      className={`
        mt-12
        text-left px-2 sm:px-4 pt-2 pb-4 rounded-xl
        bg-bgLight
      `}
    >
      <div>
        <div className="absolute">
          <div className="-mt-12">
            <SwitchButton
              onClick={() => {
                dispatch(setFromChainId(toChainId))
                dispatch(setToChainId(fromChainId))
              }}
            />
          </div>
        </div>
        <div className="pt-1 pb-3">
          <DestinationChainLabel
            chainId={toChainId}
            chains={toChainIds}
            connectedChainId={fromChainId}
          />
        </div>
      </div>
      <div className="flex h-16 mb-4 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-3 sm:pl-4
            w-full h-20
            rounded-xl
            border border-white border-opacity-20
          `}
        >
          <SelectTokenDropdown
            chainId={toChainId}
            selectedToken={toToken}
            isOrigin={false}
            onClick={() => dispatch(setShowToTokenSlideOver(true))}
          />
          {isLoading ? (
            <LoadingSpinner className="w-2/3 pr-4 ml-4 opacity-50" />
          ) : (
            <input
              pattern="[0-9.]+"
              disabled={true}
              className={`
                ml-4
                focus:outline-none
                bg-transparent
                pr-4
                w-2/3
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
  )
}
