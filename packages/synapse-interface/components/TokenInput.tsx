import { useCallback, MouseEvent, useMemo } from 'react'
import InteractiveInputRow from './InteractiveInputRow'
import type { Token } from '@types'
import { cleanNumberInput } from '@utils/cleanNumberInput'
import { formatBigIntToString } from '@/utils/bigint/format'
import { useAccount } from 'wagmi'
import {
  usePoolDataState,
  usePoolUserDataState,
  usePoolWithdrawState,
} from '@/slices/pool/hooks'

export const DepositTokenInput = ({
  token,
  balanceStr,
  rawBalance,
  inputValueStr,
  onChange,
  chainId,
  address,
}: {
  token: Token
  balanceStr: string
  rawBalance: bigint
  inputValueStr: string
  onChange: (v: string) => void
  chainId: number
  address: string
}) => {
  const onClickMax = useCallback(
    (e: MouseEvent<HTMLButtonElement>) => {
      e.preventDefault()

      const adjustedValue =
        rawBalance === 0n
          ? formatBigIntToString(rawBalance, token.decimals[chainId], 5)
          : formatBigIntToString(rawBalance, token.decimals[chainId])

      onChange(adjustedValue)
    },
    [onChange, balanceStr, token]
  )

  return (
    <div className="items-center">
      <div className="w-full">
        <InteractiveInputRow
          title={token.symbol}
          isConnected={address !== undefined}
          balanceStr={balanceStr}
          onClickBalance={onClickMax}
          value={inputValueStr}
          placeholder={'0.0000'}
          onChange={(e) => onChange(cleanNumberInput(e.target.value))}
          disabled={rawBalance === 0n}
          icon={token?.icon?.src}
        />
      </div>
    </div>
  )
}

export const WithdrawTokenInput = ({
  onChange,
}: {
  onChange: (v: string) => void
}) => {
  const { pool } = usePoolDataState()
  const { poolUserData } = usePoolUserDataState()
  const { inputValue } = usePoolWithdrawState()
  const { address } = useAccount()

  const balanceStr = useMemo(() => {
    return poolUserData?.lpTokenBalance
      ? formatBigIntToString(
          poolUserData?.lpTokenBalance,
          pool.decimals[pool.chainId],
          5
        )
      : '0.00000'
  }, [inputValue])

  const onClickMax = useCallback(
    (e: MouseEvent<HTMLButtonElement>) => {
      e.preventDefault()
      const adjustedValue = formatBigIntToString(
        poolUserData?.lpTokenBalance,
        pool.decimals[pool.chainId]
      )

      onChange(adjustedValue)
    },
    [onChange, balanceStr, pool]
  )

  return (
    <div className="items-center">
      <div className="w-full">
        <InteractiveInputRow
          title={pool.symbol}
          isConnected={address !== undefined}
          balanceStr={balanceStr}
          onClickBalance={onClickMax}
          value={inputValue}
          placeholder={'0.0000'}
          onChange={(e) => onChange(cleanNumberInput(e.target.value))}
          disabled={poolUserData?.lpTokenBalance === 0n}
          icon={pool?.icon?.src}
        />
      </div>
    </div>
  )
}
