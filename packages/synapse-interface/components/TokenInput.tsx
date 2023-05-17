import { useCallback, MouseEvent } from 'react'
import InteractiveInputRow from './InteractiveInputRow'
import { formatBNToString } from '@bignumber/format'
import { formatUnits } from '@ethersproject/units'
import { displaySymbol } from '@utils/displaySymbol'
import { Token } from '@types'
import { BigNumber } from 'ethers'
import { cleanNumberInput } from '@utils/cleanNumberInput'

const DEFAULT_VALUE = '0.0'

const TokenInput = ({
  token,
  balanceStr,
  inputValueStr,
  onChange,
  chainId,
  address,
}: {
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
      if (balanceStr === DEFAULT_VALUE) {
        onChange(DEFAULT_VALUE)
        return
      } else {
        const maxStr = formatUnits(balanceStr, token.decimals[chainId])
        maxStr != 'undefined' ? onChange(maxStr) : null
      }
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
export default TokenInput
