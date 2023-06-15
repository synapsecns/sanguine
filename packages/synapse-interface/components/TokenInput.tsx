import { useCallback, MouseEvent } from 'react'
import InteractiveInputRow from './InteractiveInputRow'
import { displaySymbol } from '@utils/displaySymbol'
import { PoolToken, PoolUserData, Token } from '@types'
import { cleanNumberInput } from '@utils/cleanNumberInput'
import { formatUnits } from '@ethersproject/units'
import { BigNumber } from '@ethersproject/bignumber'

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
  rawBalance: BigNumber
  inputValueStr: string
  onChange: (v: string) => void
  chainId: number
  address: string
}) => {
  const symbol = displaySymbol(chainId, token)

  const onClickMax = useCallback(
    (e: MouseEvent<HTMLButtonElement>) => {
      e.preventDefault()

      const adjustedValue = formatUnits(rawBalance, token.decimals[chainId])

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
      const adjustedValue = formatUnits(
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
