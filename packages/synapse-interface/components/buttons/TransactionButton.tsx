import Button from '@tw/Button'
import ButtonLoadingDots from '@/components/buttons/ButtonLoadingDots'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'
import { TransactionResponse } from '@ethersproject/providers'
import { CSSProperties } from 'react'

const BASE_PROPERTIES = `
    w-full rounded-md py-3 text-lg
    text-white transition-all
    hover:opacity-80 disabled:opacity-50 disabled:cursor-default
    disabled:text-black disabled:dark:text-white
    disabled:from-zinc-300 disabled:to-zinc-300
    disabled:dark:from-zinc-700 disabled:dark:to-zinc-700
    bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
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
  ...props
}: {
  className?: string
  onClick: () => Promise<TransactionResponse | any>
  pendingLabel: string
  label: string
  onSuccess?: () => void
  chainId?: number
  style?: CSSProperties
  disabled?: boolean
}) => {
  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()

  return (
    <Button
      {...props}
      style={style}
      disabled={disabled}
      className={`
        ${className}
        ${BASE_PROPERTIES}
        ${isPending && 'from-[#622e71] to-[#564071]'}
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
  )
}
