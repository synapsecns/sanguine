import { useMemo, useCallback } from 'react'
import { useAppDispatch } from '@/state/hooks'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'
import { setInputAmount } from '@/state/slices/bridge/reducer'
import { BridgeableToken } from 'types'
import { Warning } from './icons/Warning'
import { Tooltip } from './Tooltip'

export const AvailableBalance = ({
  originChainId,
  originToken,
  tokenBalance,
  connectedAddress,
  hasEnoughBalance,
}: {
  originChainId: number
  originToken: BridgeableToken
  tokenBalance: {
    rawBalance: bigint
    parsedBalance: string
    decimals: number
  }
  connectedAddress: string
  hasEnoughBalance: boolean
}) => {
  const dispatch = useAppDispatch()

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
      <div className="text-xs text-[--synapse-text-secondary] whitespace-nowrap">
        Select source token
      </div>
    )
  }

  return (
    <div
      onClick={handleAvailableBalanceClick}
      className="flex ml-px text-xs cursor-pointer hover:underline active:opacity-40 text-[--synapse-text-secondary] whitespace-nowrap"
    >
      Available {tokenBalance.parsedBalance ?? '0.0'}
      {!hasEnoughBalance && (
        <Tooltip hoverText="Amount may not exceed available balance">
          <Warning styles="w-3" />
        </Tooltip>
      )}
    </div>
  )
}
