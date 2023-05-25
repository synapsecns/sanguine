import React from 'react'
import Button from '@tw/Button'
import ButtonLoadingSpinner from '@components/buttons/ButtonLoadingSpinner'
import { getMenuItemBgForCoin } from '@styles/tokens'
import { Token } from '@types'

const InteractiveInputRow = ({
  title,
  isConnected,
  balanceStr,
  onClickBalance,
  value,
  placeholder,
  onChange,
  disabled,
  isPending,
  onClickEnter,
  buttonLabel,
  loadingLabel,
  icon,
  showButton = true,
  token,
  buttonWidth = 'w-2/5',
}: {
  title: string
  isConnected: boolean
  balanceStr: string
  onClickBalance: (e) => void
  value: string
  placeholder: string
  onChange: (e) => void
  disabled: boolean
  isPending: boolean
  onClickEnter: (e) => void
  buttonLabel?: string
  loadingLabel?: string
  icon: string
  showButton: boolean
  token?: Token
  buttonWidth?: string
}) => {
  let width = 'w-40'

  if (title && title.length > 6) {
    width = 'w-48'
  }

  return (
    <div className="mt-4">
      <div className="border-none rounded-xl">
        <div className="flex space-x-2">
          <div
            className={`
              flex flex-grow items-center
              pl-3 sm:pl-4
              w-full h-20
              rounded-lg
              border border-white border-opacity-20
              transform-gpu transition-all duration-75
              hover:border-opacity-30
            `}
          >
            <div className="sm:mt-[-1px] ">
              <div
                className={`
                group rounded-xl
                ${getMenuItemBgForCoin(token?.color)}
              `}
              >
                <div
                  className={`flex justify-center md:justify-start items-center rounded-lg py-1.5 px-2 ${width}`}
                >
                  <div className="self-center flex-shrink-0 hidden mr-2 sm:block">
                    <div
                      className={`
                      relative flex p-1 rounded-full
                    `}
                    >
                      <img className="w-8 h-8 " src={icon} />
                    </div>
                  </div>
                  <div className="text-left">
                    <h4 className="text-lg font-medium text-white">
                      <span className="">{title}</span>
                    </h4>
                  </div>
                </div>
              </div>
            </div>
            <div
              className={`
                flex flex-grow items-center
                mx-3 w-full h-16
                border-none
                relative overflow-hidden
              `}
            >
              <input
                autoComplete="off"
                className={`
                    ${isConnected ? '-mt-2' : '-mt-0'}
                    focus:outline-none bg-transparent
                    w-[300px] sm:min-w-[170px] sm:w-full scrollbar-none
                  placeholder:text-[#88818C] text-white
                    text-opacity-80 text-lg md:text-2xl lg:text-2xl font-medium
                    overflow-hidden
                `}
                value={value}
                placeholder={placeholder}
                onChange={(e) => {
                  onChange(e)
                }}
                name="inputRow"
              />
              {isConnected && (
                <label
                  htmlFor="inputRow"
                  className="absolute bottom-0 text-xs text-white transition-all duration-150 hover:text-opacity-70 hover:cursor-pointer transform-gpu"
                  onClick={onClickBalance}
                >
                  {balanceStr}
                  <span className="text-opacity-50 text-secondaryTextColor">
                    {' '}
                    available
                  </span>
                </label>
              )}
            </div>
          </div>
        </div>
        {showButton && (
          <Button
            className={`
              ${buttonWidth}
              rounded-xl my-2 px-4 py-3
              max-w-content disabled:bg-[#353038]
              text-white text-opacity-100
              transform-gpu transition-all duration-200
              hover:opacity-80 disabled:opacity-100 disabled:text-[#88818C]
              disabled:from-bgLight disabled:to-bgLight
              bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
              mt-5
              ${isPending && 'from-[#622e71] to-[#564071]'}
            `}
            disabled={disabled}
            onClick={(e) => {
              onClickEnter(e)
            }}
          >
            {isPending ? (
              <>
                {loadingLabel ? (
                  <span className="animate-pulse">
                    <ButtonLoadingSpinner className="mr-2" />
                    {loadingLabel}
                  </span>
                ) : (
                  <ButtonLoadingSpinner />
                )}
              </>
            ) : (
              <span>{buttonLabel ?? title}</span>
            )}
          </Button>
        )}
      </div>
    </div>
  )
}
export default InteractiveInputRow
