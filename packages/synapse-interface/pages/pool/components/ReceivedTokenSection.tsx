import _ from 'lodash'
import { formatBNToString } from '@utils/bignumber/format'
import { Token } from '@types'

const ReceivedTokenSection = ({
  withdrawQuote,
  poolTokens,
  chainId,
}: {
  withdrawQuote: any
  poolTokens: Token[]
  chainId: number
}) => {
  return (
    <div className="text-center sm:text-left">
      <p className="text-sm font-medium bg-opacity-70 pb-0.5 text-secondaryTextColor text-opacity-50 mb-2">
        You will receive
      </p>
      {poolTokens &&
        poolTokens
          .filter((token) => withdrawQuote.outputs[token.addresses[chainId]])
          .map((token) => {
            return (
              <div className="flex items-center">
                <img
                  alt=""
                  className="inline-block -mt-0.5 w-4 mr-1"
                  src={token.icon.src}
                ></img>
                <span className="text-sm text-white">
                  {formatBNToString(
                    withdrawQuote.outputs[token.addresses[chainId]].value,
                    token.decimals[chainId],
                    6
                  )}
                </span>
                <span className="px-1 text-sm font-medium text-opacity-50 text-secondaryTextColor">
                  {token.symbol}
                </span>
              </div>
            )
          })}
    </div>
  )
}
export default ReceivedTokenSection
