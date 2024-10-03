import * as tokenList from '../constants/bridgeable'
import { logger } from '../middleware/logger'

export const tokenListController = async (req, res) => {
  logger.info(`Successful tokenListController response`, { query: req.query })
  res.json(tokenList)
}
