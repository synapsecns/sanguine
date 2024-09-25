import { check } from 'express-validator'

import { tokenSymbolToToken } from '../utils/tokenSymbolToToken'

export const validateTokens = (chainParam, tokenParam, paramName) => {
  return check(tokenParam)
    .isString()
    .exists()
    .withMessage(`${paramName} is required`)
    .custom((value, { req }) => {
      const chain = req.query[chainParam]
      const tokenInfo = tokenSymbolToToken(chain, value)
      if (!tokenInfo) {
        throw new Error(`Invalid ${paramName} symbol`)
      }
      if (!req.res.locals.tokenInfo) {
        req.res.locals.tokenInfo = {}
      }
      req.res.locals.tokenInfo[paramName] = tokenInfo
      return true
    })
}
