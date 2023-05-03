import { BigNumber } from '@ethersproject/bignumber'
import { formatBNToString } from '@utils/bignumber/format'
import { commify } from '@ethersproject/units'
import { displaySymbol } from '@utils/displaySymbol'
import { Token } from '@types'

const TokenBalance = ({
  token,
  chainId,
  tokenBalance,
}: {
  token: Token
  chainId: number
  tokenBalance: BigNumber
}) => {
  const formattedBalance = commify(
    formatBNToString(
      tokenBalance,
      token?.decimals?.[chainId as keyof Token['decimals']],
      2
    )
  )
  return (
    <div className="ml-auto mr-5 text-lg text-white">
      {!tokenBalance.eq(0) && (
        <p>
          {formattedBalance}
          <span className="text-sm opacity-80">
            {' '}
            {token ? displaySymbol(chainId, token) : ''}
          </span>
        </p>
      )}
    </div>
  )
}

export default TokenBalance
