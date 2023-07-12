import { useCallback, MouseEvent } from 'react'
import InteractiveInputRow from './InteractiveInputRow'
import { displaySymbol } from '@utils/displaySymbol'
import { PoolToken, PoolUserData, Token } from '@types'
import { cleanNumberInput } from '@utils/cleanNumberInput'
import { formatBigIntToString } from '@/utils/bigint/format'

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
  const symbol = displaySymbol(chainId, token)

  const onClickMax = useCallback(
    (e: MouseEvent<HTMLButtonElement>) => {
      e.preventDefault()

      const adjustedValue =
        rawBalance === 0n
          ? formatBigIntToString(rawBalance, token.decimals[chainId], 4)
          : formatBigIntToString(rawBalance, token.decimals[chainId])

      onChange(adjustedValue)
    },
    [onChange, balanceStr, token]
  )

  return (
    <div className="items-center">
      <div className="w-full">
        <InteractiveInputRow
          title={symbol}
          isConnected={address !== undefined}
          balanceStr={balanceStr}
          onClickBalance={onClickMax}
          value={inputValueStr}
          placeholder={'0.0000'}
          onChange={(e) => onChange(cleanNumberInput(e.target.value))}
          // disabled={inputValueStr == ''}
          disabled={false}
          showButton={false}
          icon={token?.icon?.src}
          token={token}
          isPending={false}
          onClickEnter={() => {}}
        />
      </div>
    </div>
  )
}

export const WithdrawTokenInput = ({
  poolUserData,
  token,
  balanceStr,
  inputValueStr,
  onChange,
  chainId,
  address,
}: {
  poolUserData?: PoolUserData
  token: Token
  balanceStr: string
  inputValueStr: string
  onChange: (v: string) => void
  chainId: number
  address: string
}) => {
  const symbol = displaySymbol(chainId, token)

  const onClickMax = useCallback(
    (e: MouseEvent<HTMLButtonElement>) => {
      e.preventDefault()
      const adjustedValue = formatBigIntToString(
        poolUserData?.lpTokenBalance,
        token.decimals[chainId]
      )

      onChange(adjustedValue)
    },
    [onChange, balanceStr, token]
  )

  return (
    <div className="items-center">
      <div className="w-full">
        <InteractiveInputRow
          title={symbol}
          isConnected={address !== undefined}
          balanceStr={balanceStr}
          onClickBalance={onClickMax}
          value={inputValueStr}
          placeholder={'0.0000'}
          onChange={(e) => onChange(cleanNumberInput(e.target.value))}
          // disabled={inputValueStr == ''}
          disabled={false}
          showButton={false}
          icon={token?.icon?.src}
          token={token}
          isPending={false}
          onClickEnter={() => {}}
        />
      </div>
    </div>
  )
}
