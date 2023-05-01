import _ from 'lodash'
import { useEffect } from 'react'
import { InformationCircleIcon } from '@heroicons/react/outline'
import { Switch } from '@headlessui/react'
import { useKeyPress } from '@hooks/useKeyPress'
import Tooltip from '@tw/Tooltip'
import Button from '@tw/Button'

import { DeadlineInput } from '@components/input/DeadlineInput'
import { DisplayType } from './BridgeCard'

const SettingsSlideOver = ({
  settings,
  setSettings,
  setDisplayType,
  setDestinationAddress,
  deadlineMinutes,
  setDeadlineMinutes,
}: {
  settings: any
  setSettings: (v: any) => void
  setDisplayType: (v: DisplayType) => void
  setDestinationAddress: (v: string) => void
  deadlineMinutes: string
  setDeadlineMinutes: (v: string) => void
}) => {
  const escPressed = useKeyPress('Escape')

  function onClose() {
    setDisplayType(DisplayType.DEFAULT)
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
          <div className="flex items-center mb-4 text-sm font-light text-white">
            Deadline
            <Tooltip content="Enter deadline in minutes">
              <InformationCircleIcon className="w-4 h-4 ml-1 cursor-pointer text-[#252027] fill-bgLighter" />
            </Tooltip>
          </div>
          <DeadlineInput
            deadlineMinutes={deadlineMinutes}
            setDeadlineMinutes={setDeadlineMinutes}
          />
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
              checked={settings?.expertMode}
              onChange={(updatedExpertMode: any) => {
                if (!updatedExpertMode) {
                  setDestinationAddress('')
                }
                setSettings({
                  ...settings,
                  expertMode: updatedExpertMode,
                })
              }}
              className={`
                bg-gradient-to-r
                ${
                  settings?.expertMode
                    ? ' from-[#FF00FF] to-[#AC8FFF]'
                    : 'from-gray-900 to-gray-900'
                }
                relative inline-flex items-center h-6 rounded-full w-11
                transition-colors focus:outline-none`}
            >
              <span
                className={`
                  ${settings?.expertMode ? 'translate-x-6' : 'translate-x-1'}
                  inline-block w-6 h-6 transform bg-white rounded-full transition-transform
                `}
              />
            </Switch>
          </div>
        </Switch.Group>
        {settings?.expertMode && <WithdrawalWarning onClose={onClose} />}
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

export default SettingsSlideOver
