import { useCallback } from 'react'

import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { Warning } from '@/components/icons/Warning'
import { Tooltip } from '@/components/ui/Tooltip'
import { useCurrentTokenBalance } from '@/hooks/useCurrentTokenBalance'
import { useValidations } from '@/hooks/useValidations'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { useWalletState } from '@/state/slices/wallet/hooks'
import { FetchState } from '@/state/slices/wallet/reducer'

export const AvailableBalance = ({
  connectedAddress,
  nativeSafeMax,
  setInputAmount,
}: {
  connectedAddress: string
  nativeSafeMax: {
    amountWei: bigint | null
    fillAmount: string | null
    isClickable: boolean
    isNativeOriginToken: boolean
    labelAmount: string | null
    status: 'idle' | 'loading' | 'ready' | 'unavailable'
  }
  setInputAmount: React.Dispatch<React.SetStateAction<string>>
}) => {
  const { originChainId, originToken } = useBridgeState()
  const { balancesFetchStatus } = useWalletState()
  const isFetchingBalance = balancesFetchStatus === FetchState.LOADING
  const tokenBalance = useCurrentTokenBalance()
  const { hasEnoughBalance } = useValidations()

  const handleAvailableBalanceClick = useCallback(() => {
    if (nativeSafeMax.isNativeOriginToken) {
      if (nativeSafeMax.isClickable && nativeSafeMax.fillAmount) {
        setInputAmount(nativeSafeMax.fillAmount)
      }

      return
    }

    const maxAmount: string =
      formatBigIntToString(
        BigInt(tokenBalance.rawBalance ?? 0),
        tokenBalance.decimals ?? 0,
        18
      ) ?? '0.0'
    setInputAmount(maxAmount)
  }, [
    nativeSafeMax,
    setInputAmount,
    tokenBalance.rawBalance,
    tokenBalance.decimals,
  ])

  if (!connectedAddress) {
    return
  }

  if (!originToken) {
    return (
      <div className="text-sm text-[--synapse-secondary] whitespace-nowrap">
        Select source token
      </div>
    )
  }

  if (!originChainId) {
    return (
      <div className="text-sm text-[--synapse-secondary] whitespace-nowrap">
        Select source chain
      </div>
    )
  }

  const nativeBalanceLabel =
    nativeSafeMax.status === 'ready'
      ? `Bridgeable ${nativeSafeMax.labelAmount ?? '0.0'}`
      : `Available ${tokenBalance.parsedBalance ?? '0.0'}`

  const balanceLabel = nativeSafeMax.isNativeOriginToken
    ? nativeBalanceLabel
    : `Available ${tokenBalance.parsedBalance ?? '0.0'}`

  const isClickable =
    !isFetchingBalance &&
    (!nativeSafeMax.isNativeOriginToken || nativeSafeMax.isClickable)

  return (
    <div
      className={`
      row-start-3 col-start-1 col-end-3
      flex items-center gap-1.5 px-1 text-sm justify-self-end
    `}
    >
      {isFetchingBalance ? (
        <div>loading...</div>
      ) : (
        <div
          onClick={isClickable ? handleAvailableBalanceClick : undefined}
          className={`text-[--synapse-secondary] whitespace-nowrap ${
            isClickable
              ? 'cursor-pointer hover:underline active:opacity-40'
              : 'cursor-default'
          }`}
        >
          {balanceLabel}
        </div>
      )}
      {tokenBalance.parsedBalance && !hasEnoughBalance && (
        <Tooltip
          hoverText="Amount may not exceed available balance"
          positionStyles="-right-1 -top-8"
        >
          <Warning styles="w-3.5" />
        </Tooltip>
      )}
    </div>
  )
}
