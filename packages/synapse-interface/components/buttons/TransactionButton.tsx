import Button from '@tw/Button'
import ButtonLoadingDots from '@/components/buttons/ButtonLoadingDots'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'
import { TransactionResponse } from '@ethersproject/providers'
import { CSSProperties } from 'react'
import { joinClassNames } from '@/utils/joinClassNames'

const BASE_PROPERTIES = joinClassNames({
  flex: 'flex justify-center items-center',
  space: 'w-full rounded-md px-4 py-3 my-1',
  hover: 'hover:opacity-80',
  disabled: 'disabled:opacity-50 disabled:cursor-not-allowed',
  background: 'bg-zinc-400 dark:bg-bgLight',
  gradient: 'enabled:bg-gradient-to-r',
})

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
        ${
          isPending
            ? 'from-fuchsia-400 dark:from-fuchsia-900 to-purple-400 dark:to-purple-900'
            : 'from-fuchsia-500 to-purple-500 dark:to-purple-600'
        }
      `}
      onClick={async () => {
        const tx = await pendingTxWrapFunc(onClick())
        if (tx?.hash || tx?.transactionHash || tx?.status === 1) {
          onSuccess?.()
        }
      }}
    >
      {isPending ? (
        <>
          <ButtonLoadingDots className="mr-8" />
          <span className="opacity-30">{pendingLabel}</span>{' '}
        </>
      ) : (
        <>{label}</>
      )}
    </Button>
  )
}
