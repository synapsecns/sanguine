import _ from 'lodash'
import { formatBNToString } from '@utils/bignumber/format'
import { Token } from '@types'
const RecievedTokenSection = ({
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
      {poolTokens
        .filter((token) => withdrawQuote.outputs[token.addresses[chainId]])
        .map((token) => {
          console.log(
            'withdrawQuote.outputs',
            token.name,
            withdrawQuote.outputs[token.addresses[chainId]],
            formatBNToString(
              withdrawQuote.outputs[token.addresses[chainId]],
              token.decimals[chainId],
              2
            )
          )

          return (
            <div className="flex items-center">
              <img
                alt=""
                className="inline-block -mt-0.5 w-4 mr-1"
                src={token.icon.src}
              ></img>
              <span className="text-sm text-white">
                {formatBNToString(
                  withdrawQuote.outputs[token.addresses[chainId]],
                  token.decimals[chainId],
                  2
                )}
              </span>
              <span className="text-sm px-1 font-medium text-secondaryTextColor text-opacity-50">
                {token.symbol}
              </span>
            </div>
          )
        })}
    </div>
  )
}
export default RecievedTokenSection
