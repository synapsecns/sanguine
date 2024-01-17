import React from 'react'
import Button from '@tw/Button'
import ButtonLoadingDots from './buttons/ButtonLoadingDots'

export const InteractiveInputRowButton = ({
  title,
  disabled,
  isPending,
  onClickEnter,
  buttonLabel,
  loadingLabel,
}: {
  title: string
  disabled: boolean
  isPending: boolean
  onClickEnter: (e) => void
  buttonLabel?: string
  loadingLabel?: string
}) => {
  return (
    <Button
      className={`
        w-full
        rounded py-3
        hover:opacity-80
        disabled:opacity-40
        border border-fuchsia-500 disabled:border-zinc-500
        disabled:cursor-not-allowed
        bg-fuchsia-100 disabled:bg-transparent
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
