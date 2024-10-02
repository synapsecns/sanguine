import * as tokensList from '../constants/bridgeable'
import { CHAINS_ARRAY } from '../constants/chains'
import { logger } from '../middleware/logger'

export const indexController = async (_req, res) => {
  try {
    const tokensWithChains = Object.values(tokensList).map((token: any) => ({
      symbol: token.symbol,
      chains: Object.entries(token.addresses).map(
        ([chainId, tokenAddress]) => ({
          chainId,
          address: tokenAddress,
        })
      ),
    }))

    const payload = {
      message: 'Welcome to the Synapse REST API for swap and bridge quotes',
      availableChains: CHAINS_ARRAY.map((chain) => ({
        name: chain.name,
        id: chain.id,
      })),
      availableTokens: tokensWithChains,
    }

    logger.info(`Successful indexController response`)
    res.json(payload)
  } catch (err) {
    logger.error(`Error in indexController`, {
      error: err.message,
      stack: err.stack,
    })
    res.status(500).json({
      error: 'An unexpected error occurred in /. Please try again later.',
      details: err.message,
    })
  }
}
