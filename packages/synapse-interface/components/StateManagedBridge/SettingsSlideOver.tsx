import _ from 'lodash'
import { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { InformationCircleIcon } from '@heroicons/react/outline'
import { Switch } from '@headlessui/react'
import { useKeyPress } from '@hooks/useKeyPress'
import Tooltip from '@tw/Tooltip'
import Button from '@tw/Button'

import {
  setShowDestinationAddress,
  setShowSettingsSlideOver,
} from '@/slices/bridgeDisplaySlice'
import {
  setDeadlineMinutes,
  setDestinationAddress,
} from '@/slices/bridge/reducer'
import { RootState } from '@/store/store'

const SettingsSlideOver = () => {
  const dispatch = useDispatch()
  const escPressed = useKeyPress('Escape')

  const { showDestinationAddress } = useSelector(
    (state: RootState) => state.bridgeDisplay
  )

  function onClose() {
    dispatch(setShowSettingsSlideOver(false))
  }

  function escFunc() {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escFunc, [escPressed])

  return (
    <div className="max-h-full pb-4 pt-2 overflow-auto rounded-lg">
      <div
        className={`
         px-3 md:px-6 rounded-md text-base focus:outline-none
          overflow-hidden z-10 w-full
          space-y-4
        `}
      >
        <div className="pt-2"></div>
        <div className="text-sm font-light text-white">Options</div>
        {/* @ts-ignore */}
        <Switch.Group>
          <div className="flex items-center justify-between w-full">
            <Switch.Label className="flex items-center mr-4 text-white">
              Show withdrawal address{' '}
              <Tooltip content="Allows bridging to another address.">
                <div className="inline-block mt-1">
                  <InformationCircleIcon className="w-4 h-4 ml-1 cursor-pointer text-white/20 hover:text-white fill-transparent" />
                </div>
              </Tooltip>
            </Switch.Label>
            <Switch
              checked={showDestinationAddress}
              onChange={(selected: boolean) => {
                if (selected) {
                  dispatch(setShowDestinationAddress(true))
                } else {
                  dispatch(setShowDestinationAddress(false))
                  dispatch(setDestinationAddress(null))
                }
              }}
              className={`
                bg-gradient-to-r
                ${
                  showDestinationAddress
                    ? ' from-[#FF00FF] to-[#AC8FFF]'
                    : 'from-gray-900 to-gray-900'
                }
                relative inline-flex items-center h-6 rounded-full w-11
                transition-colors focus:outline-none`}
            >
              <span
                className={`
                  ${showDestinationAddress ? 'translate-x-6' : 'translate-x-1'}
                  inline-block w-6 h-6 transform bg-white rounded-full transition-transform
                `}
              />
            </Switch>
          </div>
        </Switch.Group>
        {showDestinationAddress && <WithdrawalWarning onClose={onClose} />}
      </div>
    </div>
  )
}

const WithdrawalWarning = ({ onClose }: { onClose: any }) => {
  return (
    <div className="w-full p-4 bg-slate-900/50 rounded-md">
      <div className="flex items-center justify-between space-x-1">
        <div className="w-3/4 text-xs text-white md:text-sm">
          Do not send your funds to a custodial wallet or exchange address!{' '}
          <span className="text-white text-opacity-50">
            It may be impossible to recover your funds.
          </span>
        </div>
        <Button
          className={`
            p-4 rounded-md
            text-sm font-medium text-white
            bg-bgBase/10 hover:bg-bgBase/20 active:bg-bgBase/30
            ring-1 ring-white/20
          `}
          onClick={onClose}
        >
          Okay!
        </Button>
      </div>
    </div>
  )
}

const DeadlineInput = ({ deadlineMinutes }: { deadlineMinutes: number }) => {
  const dispatch = useDispatch()

  return (
    <div className="flex h-16 pb-4 space-x-2 text-left">
      <div
        className={`
          flex flex-grow items-center
          h-14 w-full
        bg-bgLight
          border border-transparent
          hover:border-gradient-br-magenta-melrose-bgLight hover:border-solid
          focus-within:border-gradient-br-magenta-melrose-bgLight focus-within:border-solid
          pl-1
          py-0.5 rounded-md
        `}
      >
        <input
          pattern="[0-9.]+"
          className={`
              ml-4 mr-4
              focus:outline-none
              bg-transparent
              w-[300px]
              sm:min-w-[300px]
              max-w-[calc(100%-92px)]
              sm:w-full
              text-lg
             placeholder-[#716e74]
             text-white text-opacity-80
            `}
          placeholder="Custom deadline..."
          onChange={(e) => dispatch(setDeadlineMinutes(Number(e.target.value)))}
          value={deadlineMinutes}
        />
        <span className="hidden text-lg text-white md:block opacity-30">
          mins
        </span>
      </div>
    </div>
  )
}

export default SettingsSlideOver
