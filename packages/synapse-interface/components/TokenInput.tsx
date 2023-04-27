import InteractiveInputRow from './InteractiveInputRow'
import { formatBNToString } from '@bignumber/format'
import { formatUnits } from '@ethersproject/units'
import { cleanNumberInput } from '@utils/cleanNumberInput'
import { displaySymbol } from '@utils/displaySymbol'
import { Token } from '@types'
import { BigNumber } from 'ethers'
const TokenInput = ({
  token,
  max,
  inputValue,
  onChange,
  chainId,
  address,
}: {
  token: Token
  max: string
  inputValue: any
  onChange: (v: string) => void
  chainId: number
  address: string
}) => {
  const symbol = displaySymbol(chainId, token)

  const onClickMax = (e) => {
    e.preventDefault()
    const maxStr = formatUnits(max, token.decimals[chainId])
    if (maxStr != 'undefined') {
      onChange(maxStr)
    }
  }

  const onChangeInput = (e) => {
    if (cleanNumberInput(e.target.value)) {
      onChange(e.target.value)
    }
  }

  let balanceStr
  if (max && max != '') {
    balanceStr = formatBNToString(
      BigNumber.from(max),
      token.decimals[chainId],
      4
    )
  } else {
    balanceStr = '0.0'
  }
  console.log('token.icon.sr', token.icon.src)
  return (
    <div className="items-center">
      <div className="w-full">
        <InteractiveInputRow
          title={symbol}
          isConnected={address !== undefined}
          balanceStr={balanceStr}
          onClickBalance={onClickMax}
          value={inputValue[symbol]}
          placeholder={'0.0000'}
          onChange={onChangeInput}
          disabled={inputValue == ''}
          showButton={false}
          icon={token.icon.src}
          token={token}
          isPending={false}
          onClickEnter={() => {}}
        />
      </div>
    </div>
  )
}
export default TokenInput
