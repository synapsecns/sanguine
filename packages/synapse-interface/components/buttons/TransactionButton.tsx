import { CSSProperties } from 'react'
import { TransactionResponse } from '@ethersproject/providers'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'
import Button from '@tw/Button'
import ButtonLoadingDots from '@/components/buttons/ButtonLoadingDots'


const disabledClass = `opacity-30 cursor-default`

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
  onClick?: () => Promise<TransactionResponse | any>
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
      fancy={true}
      style={style}
      disabled={disabled}
      className={`
        ${className}
        ${disabled && disabledClass}
        ${isPending && 'from-[#622e71] to-[#564071] hover:opacity-100'}
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
