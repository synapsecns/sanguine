import _ from 'lodash'
import { useEffect, useState } from 'react'
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
import { setDeadlineMinutes, setDestinationAddress } from '@/slices/bridgeSlice'
import { RootState } from '@/store/store'

const SettingsSlideOver = () => {
  const dispatch = useDispatch()
  const escPressed = useKeyPress('Escape')

  const [expertMode, setExpertMode] = useState(false)

  const { deadlineMinutes } = useSelector((state: RootState) => state.bridge)

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
    <div className="max-h-full pb-4 overflow-auto rounded-3xl">
      <div
        className={`
         px-3 md:px-6 rounded-xl text-base focus:outline-none
          overflow-hidden z-10 w-full
          space-y-4
        `}
      >
        <div className="pt-2">
          {/* <div className="flex items-center mb-4 text-sm font-light text-white">
            Deadline
            <Tooltip content="Enter deadline in minutes">
              <InformationCircleIcon className="w-4 h-4 ml-1 cursor-pointer text-[#252027] fill-bgLighter" />
            </Tooltip>
          </div>
          <DeadlineInput deadlineMinutes={deadlineMinutes} /> */}
        </div>
        <div className="text-sm font-light text-white">Options</div>
        <Switch.Group>
          <div className="flex items-center justify-between w-full">
            <Switch.Label className="flex items-center mr-4 text-white">
              Show withdrawal address{' '}
              <Tooltip content="Allows bridging to another address.">
                <InformationCircleIcon className="w-4 h-4 ml-1 cursor-pointer text-[#252027] fill-bgLighter" />
              </Tooltip>
            </Switch.Label>
            <Switch
              checked={expertMode}
              onChange={(selected: boolean) => {
                if (selected) {
                  dispatch(setShowDestinationAddress(true))
                  setExpertMode(true)
                } else {
                  dispatch(setShowDestinationAddress(false))
                  dispatch(setDestinationAddress(null))
                  setExpertMode(false)
                }
              }}
              className={`
                bg-gradient-to-r
                ${
                  expertMode
                    ? ' from-[#FF00FF] to-[#AC8FFF]'
                    : 'from-gray-900 to-gray-900'
                }
                relative inline-flex items-center h-6 rounded-full w-11
                transition-colors focus:outline-none`}
            >
              <span
                className={`
                  ${expertMode ? 'translate-x-6' : 'translate-x-1'}
                  inline-block w-6 h-6 transform bg-white rounded-full transition-transform
                `}
              />
            </Switch>
          </div>
        </Switch.Group>
        {expertMode && <WithdrawalWarning onClose={onClose} />}
      </div>
    </div>
  )
}

const WithdrawalWarning = ({ onClose }: { onClose: any }) => {
  return (
    <div className="w-full p-4 bg-bgLight rounded-xl">
      <div className="flex items-center justify-between space-x-1">
        <div className="w-3/4 text-xs text-white md:text-sm">
          Do not send your funds to a custodial wallet or exchange address!{' '}
          <span className="text-white text-opacity-50">
            It may be impossible to recover your funds.
          </span>
        </div>
        <Button
          className={`
            p-4 rounded-xl
            text-sm font-medium text-white
            bg-bgLighter hover:bg-bgLightest active:bg-bgLightest
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
          py-0.5 rounded-xl
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
