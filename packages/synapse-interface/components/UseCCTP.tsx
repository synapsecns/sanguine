import { useMemo } from 'react'

import { useBridgeState } from '@/slices/bridge/hooks'
import { setExcludeCCTP } from '@/slices/bridge/reducer'
import { useAppDispatch } from '@/store/hooks'

export const UseCCTP = () => {
  const dispatch = useAppDispatch()
  const { fromChainId, fromToken, excludeCCTP } = useBridgeState()

  const canBeCCTP = useMemo(() => {
    return fromChainId && fromToken && fromToken?.swapableType === 'USD'
  }, [fromChainId, fromToken])

  const handleRadioChange = (event) => {
    const value: boolean = String(event.target.value) === 'true'
    dispatch(setExcludeCCTP(value))
  }

  return canBeCCTP ? (
    <div className="flex justify-between mt-4 text-sm text-white border border-white rounded-md bg-bgLight border-opacity-20">
      <div
        className={`flex justify-center p-2 w-1/2 ${
          excludeCCTP ? 'bg-bgBase rounded-l-md' : ''
        }`}
      >
        <div className="flex flex-col">
          <div className="flex items-center space-x-2">
            <input
              className="cursor-pointer"
              type="radio"
              id="synapse-router-radio"
              name="cctp"
              value="true"
              checked={excludeCCTP}
              onChange={handleRadioChange}
            />
            <div>5-7 min</div>
          </div>
          <div className="text-sm text-secondaryTextColor">~0.02%</div>
        </div>
      </div>
      <div
        className={`flex justify-center p-2 w-1/2 ${
          !excludeCCTP ? 'bg-bgBase rounded-r-md' : ''
        }`}
      >
        <div className="flex flex-col">
          <div className="flex items-center space-x-2">
            <input
              className="cursor-pointer"
              type="radio"
              id="cctp-radio"
              name="cctp"
              value="false"
              checked={!excludeCCTP}
              onChange={handleRadioChange}
            />
            <div>15-25 min</div>
          </div>
          <div className="text-secondaryTextColor">Zero slippage</div>
        </div>
      </div>
    </div>
  ) : null
}
