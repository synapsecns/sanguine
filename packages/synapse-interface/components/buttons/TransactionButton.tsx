import Button from '@tw/Button'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'
import { TransactionResponse } from '@ethersproject/providers'
import { CSSProperties, useState } from 'react'
import { LoaderIcon } from 'react-hot-toast'

const BASE_PROPERTIES = `
    py-3 px-4
    text-lg
    w-full rounded
    text-white
    bg-[#343036]
    border border-separator
    disabled:cursor-default disabled:opacity-40 disabled:border-transparent
  `

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
      className="relative flex items-center justify-center"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {isHovered && toolTipLabel && (
        <div className="absolute -top-4 z-10 px-2 py-1 bg-[#151315] border border-[#343036] rounded cursor-default text-[#EEEDEF] text-sm">
          {toolTipLabel}
        </div>
      )}
      <Button
        {...props}
        style={style}
        disabled={disabled}
        className={BASE_PROPERTIES}
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
            <span className="opacity-40">{pendingLabel}</span>{' '}
          </div>
        ) : (
          <span>{label}</span>
        )}
      </Button>
    </div>
  )
}
