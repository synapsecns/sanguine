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
      style={
        disabled
          ? {}
          : {
              background:
                'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
            }
      }
      className={`
              w-full self-center
              rounded-sm my-2 px-4 py-3
              text-white text-opacity-100
              border border-purple-500
              disabled:border-white/10
              hover:opacity-80  disabled:text-[#88818C]
              bg-gradient-to-r disabled:from-bgBase/20 disabled:to-bgBase/20
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
