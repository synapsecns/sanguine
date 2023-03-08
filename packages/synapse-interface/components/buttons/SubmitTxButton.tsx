import Button from '@tw/Button'

import ButtonLoadingSpinner from '@components/buttons/ButtonLoadingSpinner'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'

const BASE_PROPERTIES = `
    w-full rounded-lg my-2 px-4 py-3
    text-white text-opacity-100 transition-all
    hover:opacity-80 disabled:opacity-100 disabled:text-[#88818C]
    disabled:from-bgLight disabled:to-bgLight
    bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
  `

export function TransactionButton({
  className,
  onClick,
  pendingLabel,
  label,
  onSuccess,
  disabled,
  ...props
}: {
  className?: string
  onClick: () => Promise<void>
  pendingLabel: string
  label: string
  onSuccess?: () => void
  disabled?: boolean
}) {
  const { isPending, pendingTxWrapFunc } = usePendingTxWrapper()

  return (
    <Button
      {...props}
      disabled={disabled}
      className={`${BASE_PROPERTIES} ${className} ${
        isPending && 'from-[#622e71] to-[#564071]'
      }`}
      onClick={async () => {
        const tx = await pendingTxWrapFunc(onClick())
        if (tx?.status === 1) {
          onSuccess?.()
        }
      }}
    >
      {isPending ? (
        <div className="inline-flex items-center justify-center">
          <ButtonLoadingSpinner className="mr-2" />
          <span className="opacity-30">{pendingLabel}</span>{' '}
        </div>
      ) : (
        <span>{label}</span>
      )}
    </Button>
  )
}
