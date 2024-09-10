import _ from 'lodash'
import { useEffect } from 'react'
import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import { InformationCircleIcon } from '@heroicons/react/outline'
import Tooltip from '@tw/Tooltip'
import { Switch } from '@headlessui/react'
import { useKeyPress } from '@hooks/useKeyPress'
import {
  setShowDestinationAddress,
  setShowSettingsSlideOver,
} from '@/slices/bridgeDisplaySlice'
import {
  setDeadlineMinutes,
  setDestinationAddress,
} from '@/slices/bridge/reducer'
import { useBridgeDisplayState } from '@/slices/bridge/hooks'

const SettingsSlideOver = () => {
  const dispatch = useAppDispatch()
  const t = useTranslations('Settings')

  const { showDestinationAddress } = useBridgeDisplayState()

  const onClose = () => {
    dispatch(setShowSettingsSlideOver(false))
  }

  const escPressed = useKeyPress('Escape', true)

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escFunc, [escPressed])

  return (
    <div className="max-h-full pb-4 overflow-auto rounded-lg">
      <div
        className={`
          px-3 rounded-md text-base space-y-4 z-10 w-full
          overflow-hidden focus:outline-none md:px-6
        `}
      >
        <div className="pt-2"></div>
        <div className="text-sm font-light text-white">{t('Options')}</div>
        {/* @ts-ignore */}
        <Switch.Group>
          <div className="flex items-center justify-between w-full">
            <Switch.Label className="flex items-center mr-4 text-white">
              {t('Show withdrawal address')}{' '}
              <Tooltip content={t('Allows bridging to another address')}>
                <InformationCircleIcon className="w-4 h-4 ml-1 cursor-pointer text-[#252027] fill-bgLighter" />
              </Tooltip>
            </Switch.Label>
            <Switch
              checked={showDestinationAddress}
              onChange={(selected: boolean) => {
                if (selected) {
                  dispatch(setShowDestinationAddress(true))
                  onClose()
                } else {
                  dispatch(setShowDestinationAddress(false))
                  dispatch(setDestinationAddress(null))
                }
              }}
              className={`
                relative inline-flex items-center h-6 rounded-full w-11
                transition-colors bg-gradient-to-r focus:outline-none
                ${
                  showDestinationAddress
                    ? ' from-[#FF00FF] to-[#AC8FFF]'
                    : 'from-gray-900 to-gray-900'
                }
              `}
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
      </div>
    </div>
  )
}

const DeadlineInput = ({ deadlineMinutes }: { deadlineMinutes: number }) => {
  const dispatch = useAppDispatch()

  const t = useTranslations('Time')

  return (
    <div className="flex h-16 pb-4 space-x-2 text-left">
      <div
        className={`
          flex flex-grow items-center h-14 w-full pl-1 py-0.5
          bg-bgLight border border-transparent rounded-md
          hover:border-gradient-br-magenta-melrose-bgLight hover:border-solid
          focus-within:border-gradient-br-magenta-melrose-bgLight focus-within:border-solid
        `}
      >
        <input
          pattern="[0-9.]+"
          className={`
              ml-4 mr-4 text-lg text-white text-opacity-80
              bg-transparent w-[300px] max-w-[calc(100%-92px)]
              sm:min-w-[300px] sm:w-full focus:outline-none
             placeholder-[#716e74]
            `}
          placeholder="Custom deadline..."
          onChange={(e) => dispatch(setDeadlineMinutes(Number(e.target.value)))}
          value={deadlineMinutes}
        />
        <span className="hidden text-lg text-white md:block opacity-30">
          {t('min')}
        </span>
      </div>
    </div>
  )
}

export default SettingsSlideOver
