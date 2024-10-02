import * as tokenList from '../constants/bridgeable'
import { logger } from '../middleware/logger'

export const tokenListController = async (_req, res) => {
  logger.info(`Successful tokenListController response`)
  res.json(tokenList)
}
