import Button from '@tw/Button'
import ButtonLoadingSpinner from '@components/buttons/ButtonLoadingSpinner'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'
import { TransactionResponse } from '@ethersproject/providers'
import ExplorerToastLink from '@components/ExplorerToastLink'
import toast from 'react-hot-toast'
import { AddressZero } from '@ethersproject/constants'

const BASE_PROPERTIES = `
    w-full rounded-lg my-2 px-4 py-3
    text-white text-opacity-100 transition-all
    hover:opacity-80 disabled:opacity-50 disabled:text-[#88818C]
    disabled:from-bgLight disabled:to-bgLight
    bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
  `

const disabledClass = `opacity-30 cursor-default`

export const TransactionButton = ({
  className,
  onClick,
  pendingLabel,
  label,
  onSuccess,
  disabled,
  chainId,
  ...props
}: {
  className?: string
  onClick: () => Promise<TransactionResponse | any>
  pendingLabel: string
  label: string
  onSuccess?: () => void
  chainId?: number
  disabled?: boolean
}) => {
  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()

  return (
    <Button
      {...props}
      disabled={disabled}
      className={`
        ${className}
        ${BASE_PROPERTIES}
        ${disabled && disabledClass}
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
          <ButtonLoadingSpinner className="mr-2" />
          <span className="opacity-30">{pendingLabel}</span>{' '}
        </div>
      ) : (
        <span>{label}</span>
      )}
    </Button>
  )
}
