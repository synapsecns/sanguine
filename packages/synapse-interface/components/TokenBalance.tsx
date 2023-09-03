import { formatBigIntToString } from '@utils/bigint/format'
import { commify } from '@ethersproject/units'
import { Token } from '@types'

const TokenBalance = ({
  token,
  chainId,
  tokenBalance,
}: {
  token: Token
  chainId: number
  tokenBalance: bigint
}) => {
  const formattedBalance =
    tokenBalance &&
    commify(
      formatBigIntToString(
        tokenBalance,
        token?.decimals?.[chainId as keyof Token['decimals']],
        3
      )
    )
  return (
    <div className="ml-auto mr-5 text-lg text-white">
      {!(tokenBalance === 0n) && (
        <div>
          {formattedBalance}
          <span className="text-sm opacity-80">
            {' '}
            {token ? token.symbol : ''}
          </span>
        </div>
      )}
    </div>
  )
}

export default TokenBalance
