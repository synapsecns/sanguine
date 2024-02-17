import Button from '@tw/Button'
import ButtonLoadingDots from '@/components/buttons/ButtonLoadingDots'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'
import { TransactionResponse } from '@ethersproject/providers'
import { CSSProperties } from 'react'

const BASE_PROPERTIES = `
    w-full rounded-sm px-4 py-3
    text-white text-opacity-100 transition-all
    hover:opacity-80 disabled:opacity-50 disabled:text-[#88818C]
    border border-fuchsia-600
    disabled:border-zinc-600 disabled:bg-zinc-800 disabled:cursor-default
  `

export const TransactionButton = ({
  className,
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
  className?: string
  onClick: () => Promise<TransactionResponse | any>
  pendingLabel?: string
  label: string
  onSuccess?: () => void
  chainId?: number
  style?: CSSProperties
  toolTipLabel?: string
  disabled?: boolean
}) => {
  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()

  return (
    <div className="relative flex items-center justify-center group">
      {toolTipLabel && (
        <div className="absolute z-10 px-2 py-1 text-sm text-white transition-opacity duration-150 ease-in-out bg-black border border-black rounded opacity-0 cursor-default -top-6 group-hover:opacity-100">
          {toolTipLabel}
        </div>
      )}
      <Button
        {...props}
        style={style}
        disabled={disabled}
        className={`
          ${className}
          ${BASE_PROPERTIES}
          ${!disabled && 'bg-custom-gradient'}
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
            <ButtonLoadingDots className="mr-8" />
            <span className="opacity-30">{pendingLabel}</span>{' '}
          </div>
        ) : (
          <span>{label}</span>
        )}
      </Button>
    </div>
  )
}
