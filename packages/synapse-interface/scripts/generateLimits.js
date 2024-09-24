const { ethers } = require('ethers')

const { fetchBridgeQuote } = require('./utils/fetchBridgeQuote')
const ERC20ABI = require('./abi/IERC20Metadata.json')
const { prettyPrintTS } = require('./utils/prettyPrintTs')
const bridgeLimitMap = require('./data/bridgeLimitMap.json')
const providers = require('./data/providers.json')

// Provider setup
Object.keys(providers).forEach((chainId) => {
  providers[chainId] = new ethers.providers.JsonRpcProvider(providers[chainId])
})

// Constants
const lowerLimitValues = ['0.01', '0.1', '1', '10']
const upperLimitValues = ['20000000', '1000000']

const generateLimits = async () => {
  for (const originChainId in bridgeLimitMap) {
    for (const originTokenAddress in bridgeLimitMap[originChainId]) {
      const originTokenInfo = bridgeLimitMap[originChainId][originTokenAddress]

      // Only proceed if the swappableType is 'USD' or 'ETH'
      if (
        originTokenInfo.swappableType !== 'USD' &&
        originTokenInfo.swappableType !== 'ETH'
      ) {
        console.log(
          `Skipping originChainId ${originChainId} ${originTokenInfo.symbol} swappableType: (${originTokenInfo.swappableType})`
        )
        continue
      }

      for (const destinationChainId in originTokenInfo.routes) {
        const destinationTokens = originTokenInfo.routes[destinationChainId]

        for (const destinationTokenAddress in destinationTokens) {
          const destinationTokenData =
            destinationTokens[destinationTokenAddress]

          let minOriginValue

          // Iterate through the lower limit values
          for (const limitValue of lowerLimitValues) {
            try {
              const bridgeQuotes = await retryFetchBridgeQuote(
                originChainId,
                destinationChainId,
                originTokenAddress,
                destinationTokenAddress,
                limitValue
              )

              if (bridgeQuotes && bridgeQuotes.length > 0) {
                const minBridgeAmountQuote = bridgeQuotes.reduce(
                  (minQuote, currentQuote) => {
                    const currentFee = currentQuote.bridgeFeeFormatted
                    const minFee = minQuote ? minQuote.bridgeFeeFormatted : null

                    return !minFee ||
                      parseFloat(currentFee) < parseFloat(minFee)
                      ? currentQuote
                      : minQuote
                  },
                  null
                )

                minOriginValue = minBridgeAmountQuote.bridgeFeeFormatted

                break
              }
            } catch (error) {
              console.error(
                `Failed to fetch bridge quote for ${originChainId} ${originTokenAddress} to ${destinationChainId} ${destinationTokenAddress}:`,
                error
              )
            }
          }

          let maxOriginValue

          for (const limitValue of upperLimitValues) {
            try {
              const bridgeQuotes = await retryFetchBridgeQuote(
                originChainId,
                destinationChainId,
                originTokenAddress,
                destinationTokenAddress,
                limitValue
              )

              if (bridgeQuotes && bridgeQuotes.length > 0) {
                const maxAmountOutQuote = bridgeQuotes.reduce(
                  (maxQuote, currentQuote) => {
                    const currentMaxAmountOut = currentQuote.maxAmountOutStr
                    const bestMaxAmountOut = maxQuote
                      ? maxQuote.maxAmountOutStr
                      : null

                    return !bestMaxAmountOut ||
                      parseFloat(currentMaxAmountOut) <
                        parseFloat(bestMaxAmountOut)
                      ? maxQuote
                      : currentQuote
                  },
                  null
                )

                maxOriginValue = maxAmountOutQuote.maxAmountOutStr

                break
              }
            } catch (error) {
              console.error(
                `Failed to fetch bridge quote for ${originChainId} ${originTokenAddress} to ${destinationChainId} ${destinationTokenAddress}:`,
                error
              )
            }
          }

          // Update the bridgeLimitMap with the minOriginValue
          if (minOriginValue) {
            console.log(
              `OriginChainId: ${originChainId} OriginToken: ${originTokenInfo.symbol} ${originTokenAddress} DestinationChainId:${destinationChainId} DestinationToken: ${destinationTokenData.symbol} ${destinationTokenAddress} minOriginValue ${minOriginValue}`
            )
            destinationTokenData.minOriginValue = minOriginValue
          }
          if (maxOriginValue) {
            console.log(
              `OriginChainId: ${originChainId} OriginToken: ${originTokenInfo.symbol} ${originTokenAddress} DestinationChainId:${destinationChainId} DestinationToken: ${destinationTokenData.symbol} ${destinationTokenAddress} maxOriginValue ${maxOriginValue}`
            )
            destinationTokenData.maxOriginValue = maxOriginValue
          }
        }
      }
    }
  }

  // Save the result to 'limitsMap.ts'
  prettyPrintTS(
    bridgeLimitMap,
    'BRIDGE_LIMITS_MAP',
    './constants/bridgeLimitsMap.ts'
  )
}

const retryFetchBridgeQuote = async (
  originChainId,
  destinationChainId,
  originTokenAddress,
  destinationTokenAddress,
  limitValue,
  maxRetries = 3
) => {
  let attempt = 0
  while (attempt < maxRetries) {
    try {
      const bridgeQuotes = await fetchBridgeQuote(
        originChainId,
        destinationChainId,
        originTokenAddress,
        destinationTokenAddress,
        limitValue
      )
      return bridgeQuotes
    } catch (error) {
      attempt++
      console.error(
        `Attempt ${attempt} failed for ${originChainId} ${originTokenAddress} to ${destinationChainId} ${destinationTokenAddress}:`,
        error
      )
      if (attempt === maxRetries) {
        throw new Error(
          `Failed after ${maxRetries} attempts for ${originChainId} ${originTokenAddress} to ${destinationChainId} ${destinationTokenAddress}`
        )
      }
    }
  }
}

// Run the generateLimits function
generateLimits()
  .then(() => console.log('Limits generation complete.'))
  .catch((error) => console.error('Error generating limits:', error))
