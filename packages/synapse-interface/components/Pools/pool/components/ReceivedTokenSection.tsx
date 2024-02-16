import _ from 'lodash'
import { formatBigIntToString } from '@utils/bigint/format'
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
  // If withdrawQuote or outputs in withdrawQuote is undefined, return null
  if (!withdrawQuote || !withdrawQuote.outputs) return null;

  // Cannot handle nativeTokens currently without state reworking to make accessible the pools object
  // const tokensArray = poolTokens ? poolTokens : pool.nativeTokens;  // Using the appropriate tokens array
  const tokensArray = poolTokens
  // If the "ALL" key exists, use its array. Otherwise, convert the outputs object into an array
  const outputsArray = withdrawQuote.outputs.ALL ? withdrawQuote.outputs.ALL : Object.values(withdrawQuote.outputs);

  return (
    <div className="text-center sm:text-left">
      <p className="text-sm font-medium bg-opacity-70 pb-0.5 text-secondaryTextColor text-opacity-50 mb-2">
        You will receive
      </p>
      {tokensArray &&
        outputsArray
          .filter((output) => tokensArray[output.index])  // Using the output's index to filter the tokens
          .map((output) => {
            const token = tokensArray[output.index];  // Using the output's index to fetch the corresponding token
            return (
              // Added a key for mapped elements
              <div className="flex items-center" key={token.addresses[chainId]}>
                <img
                  alt=""
                  className="inline-block -mt-0.5 w-4 mr-1"
                  src={token.icon.src}
                ></img>
                <span className="text-sm text-white">
                  {formatBigIntToString(
                    output.value,  // Adjusted to use output.value directly
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
