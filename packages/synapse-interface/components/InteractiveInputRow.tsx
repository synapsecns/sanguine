import React from 'react'
import Button from '@tw/Button'

const InteractiveInputRow = ({
  title,
  isConnected,
  balanceStr,
  onClickBalance,
  value,
  placeholder,
  onChange,
  disabled,
  icon,
}: {
  title: string
  isConnected: boolean
  balanceStr: string
  onClickBalance: (e) => void
  value: string
  placeholder: string
  onChange: (e) => void
  disabled: boolean
  icon: string
}) => {
  return (
    <div className="flex items-center gap-3 px-2 py-1 rounded-md bg-zinc-50 border border-zinc-200 dark:bg-zinc-800 dark:border-zinc-700">
      <div className="flex items-center gap-2 pl-2 pr-4 py-2 rounded bg-zinc-200 dark:bg-zinc-600 cursor-default min-w-[7rem]">
        <img className="w-8 h-8" src={icon} />
        <h4 className="text-lg">{title}</h4>
      </div>
      <div className="grid flex-grow">
        <input
          autoComplete="off"
          className={`
            focus:outline-none focus:ring-0 focus:border-none
            border-none bg-transparent p-0 block
            placeholder:text-zinc-400
            text-lg md:text-2xl font-medium
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
            className="text-sm text-zinc-400 hover:text-opacity-70 hover:cursor-pointer"
            onClick={onClickBalance}
          >
            {balanceStr} <span className="text-zinc-400">available</span>
          </label>
        )}
      </div>
      {isConnected && (
        <Button
          className={`
            bg-zinc-200 dark:bg-zinc-600
            border border-zinc-300 dark:border-transparent
            px-4 py-1 rounded
            hidden md:block
            ${disabled
              ? 'opacity-60 cursor-not-allowed'
              : 'hover:border-purple-500'
            }
          `}
          onClick={disabled ? undefined : onClickBalance}
        >
          Max
        </Button>
      )}
    </div>
  )
}
export default InteractiveInputRow
