import Button from '@tw/Button'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'
import { TransactionResponse } from '@ethersproject/providers'
import { CSSProperties, useState } from 'react'
import { LoaderIcon } from 'react-hot-toast'

const BASE_PROPERTIES = `
    h-[64px]
    w-full rounded-sm px-4
    text-white
    bg-[#343036]
    border border-separator
  `

const disabledClass = `cursor-default opacity-40 border-transparent`

export const TransactionButton = ({
  onClick,
  pendingLabel,
  label,
  onSuccess,
  disabled,
  chainId,
  style,
  toolTipLabel,
  ...props
}: {
  onClick: () => Promise<TransactionResponse | any>
  pendingLabel?: string | JSX.Element
  label: string | JSX.Element
  onSuccess?: () => void
  chainId?: number
  style?: CSSProperties
  toolTipLabel?: string
  disabled?: boolean
}) => {
  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()
  const [isHovered, setIsHovered] = useState(false)

  return (
    <div
      className="relative flex flex-col items-center justify-center"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {isHovered && toolTipLabel && (
        <div className="absolute -top-3 z-10 flex justify-center items-center pl-2 pr-2 pt-1 pb-1 bg-[#151315] border border-[#343036] rounded-sm h-[29px] ">
          <div className="text-center text-[#EEEDEF] text-sm">
            {toolTipLabel}
          </div>
        </div>
      )}
      <Button
        {...props}
        style={style}
        disabled={disabled}
        className={`
          ${BASE_PROPERTIES}
          ${disabled && disabledClass}
        `}
        onClick={async () => {
          const tx = await pendingTxWrapFunc(onClick())
          if (tx?.hash || tx?.transactionHash || tx?.status === 1) {
            onSuccess?.()
          }
        }}
      >
        {isPending ? (
          <div className="inline-flex items-center justify-center">
            <LoaderIcon className="mr-2" />
            <span className="opacity-30">{pendingLabel}</span>{' '}
          </div>
        ) : (
          <span>{label}</span>
        )}
      </Button>
    </div>
  )
}
