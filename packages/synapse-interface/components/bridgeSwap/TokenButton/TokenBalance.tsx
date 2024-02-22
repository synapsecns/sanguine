import { Token } from '@/utils/types'
export const TokenBalance = ({
  token,
  parsedBalance,
}: {
  token: Token
  parsedBalance?: string
}) => {
  return (
    <div className="ml-auto mr-5 text-md text-primaryTextColor">
      {parsedBalance && parsedBalance !== '0.0' && (
        <div>
          {parsedBalance}
          <span className="text-md text-secondaryTextColor">
            {' '}
            {token ? token.symbol : ''}
          </span>
        </div>
      )}
    </div>
  )
}