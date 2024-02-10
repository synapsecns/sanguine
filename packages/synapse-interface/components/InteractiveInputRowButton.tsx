import React from 'react'
import Button from '@tw/Button'
import ButtonLoadingDots from './buttons/ButtonLoadingDots'

export const InteractiveInputRowButton = ({
  title,
  disabled,
  isPending,
  onClick,
  buttonLabel,
  loadingLabel,
}: {
  title: string
  disabled: boolean
  isPending: boolean
  onClick: (e) => void
  buttonLabel?: string
  loadingLabel?: string
}) => {
  return (
    <Button
      className={`
        w-full self-center
        rounded-sm my-2 px-4 py-3
        text-white text-opacity-100
        border border-purple-500
        disabled:border-white/10
        hover:opacity-80  disabled:text-[#88818C]
        bg-gradient-to-r disabled:from-bgBase/20 disabled:to-bgBase/20
        mt-5
        ${isPending && '!from-[#622e71] !to-[#564071]'}
        ${!disabled && 'from-[rgba(128, 0, 255, 0.2)] to-[rgba(255, 0, 191, 0.2)]'}
      `}
      disabled={disabled}
      onClick={onClick}
    >
      {isPending ? (
        <>
          {loadingLabel ? (
            <div className="flex items-center justify-center space-x-5 animate-pulse">
              <ButtonLoadingDots className="mr-2" />
              <span>{loadingLabel}</span>
            </div>
          ) : (
            <ButtonLoadingDots />
          )}
        </>
      ) : (
        <span>{buttonLabel ?? title}</span>
      )}
    </Button>
  )
}
