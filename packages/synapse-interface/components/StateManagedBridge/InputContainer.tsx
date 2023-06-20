import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '@/store/store'

import { updateFromValue, setShowTokenSlideOver } from '@/slices/bridgeSlice'
import { stringToBigNum } from '@/utils/stringToBigNum'
import SelectTokenDropdown from '@/components/input/TokenAmountInput/SelectTokenDropdown'
import { ChainLabel } from '@/components/ChainLabel'

export const InputContainer = ({}) => {
  const { fromChainId, fromToken, fromChainIds, fromValue } = useSelector(
    (state: RootState) => state.bridge
  )

  const dispatch = useDispatch()

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    let fromValueString = event.target.value
    try {
      let fromValueBigNumber = stringToBigNum(
        fromValueString,
        fromToken.decimals[fromChainId]
      )
      dispatch(updateFromValue(fromValueBigNumber))
    } catch (error) {
      console.error('Invalid value for conversion to BigNumber')
    }
  }

  return (
    <div
      className={`
        text-left px-2 sm:px-4 pt-2 pb-4 rounded-xl
        bg-bgLight
      `}
    >
      <div>
        <div className="pt-1 pb-3">
          <ChainLabel
            isOrigin={true}
            chainId={fromChainId}
            setDisplayType={() => {}}
            onChangeChain={() => {}}
            chains={fromChainIds.map((id) => `${id}`)}
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
            chainId={fromChainId}
            selectedToken={fromToken}
            isOrigin={false}
            onClick={() => dispatch(setShowTokenSlideOver(true))}
          />
          <input
            pattern="[0-9.]+"
            disabled={false}
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
            onChange={handleFromValueChange}
            name="inputRow"
            autoComplete="off"
          />
        </div>
      </div>
    </div>
  )
}
