import React from 'react'
import Button from '@tw/Button'
import ButtonLoadingDots from '@/components/buttons/ButtonLoadingDots'


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
      fancy={true}
      className={`
        self-center mt-5
        ${isPending && '!from-[#622e71] !to-[#564071]'}
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
