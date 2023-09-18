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
    <div className="flex flex-col bg-[#353038] text-white text-sm p-3 rounded-md mt-4 space-y-2">
      <div className="mb-2 font-bold">Transaction routing type</div>
      <div className="flex space-x-2">
        <div>
          <input
            type="radio"
            id="synapse-router-radio"
            name="cctp"
            value="true"
            checked={excludeCCTP}
            onChange={handleRadioChange}
          />
        </div>
        <div className="flex justify-between w-full">
          <div>Synapse Router</div>
        </div>
      </div>
      <div className="flex space-x-2">
        <input
          type="radio"
          id="cctp-radio"
          name="cctp"
          value="false"
          checked={!excludeCCTP}
          onChange={handleRadioChange}
        />
        <div className="flex justify-between w-full">
          <div>Circle cross-chain transfer protocol (CCTP)</div>
        </div>
      </div>
    </div>
  ) : null
}
