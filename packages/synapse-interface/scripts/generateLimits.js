const fs = require('fs')

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
const ETH = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

// Function to get token decimals
const getTokenDecimals = async (chainId, token) => {
  if (token === ETH) {
    return 18
  }
  const decimals = await new ethers.Contract(
    token,
    ERC20ABI,
    providers[chainId]
  ).decimals()
  return decimals
}

// Main function to generate limits
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
          `Skipping ${originTokenInfo.symbol} (${originTokenInfo.swappableType})`
        )
        continue
      }

      // Get token decimals for origin token
      const originTokenDecimals = await getTokenDecimals(
        originChainId,
        originTokenAddress
      )

      for (const destinationChainId in originTokenInfo.routes) {
        const destinationTokens = originTokenInfo.routes[destinationChainId]

        for (const destinationTokenAddress in destinationTokens) {
          const destinationTokenData =
            destinationTokens[destinationTokenAddress]

          let minOriginValue

          // Iterate through the lower limit values
          for (const limitValue of lowerLimitValues) {
            try {
              const bridgeQuotes = await fetchBridgeQuote(
                originChainId,
                destinationChainId,
                originTokenAddress,
                destinationTokenAddress,
                limitValue
              )

              if (bridgeQuotes.length > 0) {
                const bridgeQuote = bridgeQuotes[0]

                console.log('bridgeQuote: ', bridgeQuote)
                // Save the minOriginValue as the feeAmount of the first bridgeQuote
                minOriginValue = ethers.utils.formatUnits(
                  bridgeQuote.feeAmount,
                  originTokenDecimals
                )
                break // Stop querying if a valid bridge quote is found
              }
            } catch (error) {
              console.error(
                `Failed to fetch bridge quote for ${originTokenAddress} to ${destinationTokenAddress}:`,
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
          } else {
            destinationTokenData.minOriginValue = null
          }
        }
      }
    }
  }

  // Save the result to 'limitsMap.ts'
  prettyPrintTS(bridgeLimitMap, 'LIMITS_MAP', './constants/limitsMap.ts')
}

// Run the generateLimits function
generateLimits()
  .then(() => console.log('Limits generation complete.'))
  .catch((error) => console.error('Error generating limits:', error))
