import { useCallback } from 'react'
import { useAppDispatch } from '@/state/hooks'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { setInputAmount } from '@/state/slices/bridge/reducer'
import { Warning } from './icons/Warning'
import { Tooltip } from './Tooltip'
import { useCurrentTokenBalance } from '@/hooks/useCurrentTokenBalance'
import { useValidations } from '@/hooks/useValidations'
import { useBridgeState } from '@/state/slices/bridge/hooks'

export const AvailableBalance = ({
  connectedAddress,
}: {
  connectedAddress: string
}) => {
  const dispatch = useAppDispatch()

  const { originChainId, originToken } = useBridgeState()

  const tokenBalance = useCurrentTokenBalance()

  const { hasEnoughBalance } = useValidations()

  const handleAvailableBalanceClick = useCallback(() => {
    const maxAmount: string =
      formatBigIntToString(
        BigInt(tokenBalance.rawBalance ?? 0),
        tokenBalance.decimals ?? 0,
        18
      ) ?? '0.0'
    dispatch(setInputAmount(maxAmount))
  }, [dispatch, tokenBalance, originToken, originChainId])

  if (!connectedAddress) return

  if (!originToken) {
    return (
      <div className="text-sm text-[--synapse-secondary] whitespace-nowrap">
        Select source token
      </div>
    )
  }

  return (
    <div className={`
      row-start-3 col-start-1 col-end-3
      flex items-center gap-1 p-1 text-sm justify-self-end
    `}>
      <div
        onClick={handleAvailableBalanceClick}
        className="cursor-pointer hover:underline active:opacity-40 text-[--synapse-secondary] whitespace-nowrap"
      >
        Available {tokenBalance.parsedBalance ?? '0.0'}
      </div>
      {!hasEnoughBalance && 
        <Tooltip
          hoverText="Amount may not exceed available balance"
          positionStyles="-right-1 -top-8"
        >
          <Warning styles="w-3.5" />
        </Tooltip>
      }
    </div>
  )
}
