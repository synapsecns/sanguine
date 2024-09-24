import express from 'express'
import { check } from 'express-validator'

import { CHAINS_ARRAY } from '../constants/chains'
import { showFirstValidationError } from '../middleware/showFirstValidationError'
import { bridgeLimitsController } from '../controllers/bridgeLimitsController'
import { isTokenSupportedOnChain } from './../utils/isTokenSupportedOnChain'
import { isTokenAddress } from '../utils/isTokenAddress'

const router = express.Router()

/**
 * @openapi
 * /bridgeLimits:
 *   get:
 *     summary: Get min/max origin values for bridge quote
 *     description: Retrieve minimum and maximum bridgeable amounts to bridge from source chain to destination chain.
 *     parameters:
 *       - in: query
 *         name: fromChain
 *         required: true
 *         schema:
 *           type: integer
 *         description: The source chain ID.
 *       - in: query
 *         name: toChain
 *         required: true
 *         schema:
 *           type: integer
 *         description: The destination chain ID.
 *       - in: query
 *         name: fromToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token on the source chain.
 *       - in: query
 *         name: toToken
 *         required: true
 *         schema:
 *           type: string
 *         description: The address of the token on the destination chain.
 *     responses:
 *       200:
 *         description: Successful response containing min and max origin amounts.
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 maxOriginAmount:
 *                   type: string
 *                   description: Maximum amount of tokens that can be bridged from the origin chain.
 *                 minOriginAmount:
 *                   type: string
 *                   description: Minimum amount of tokens that can be bridged from the origin chain.
 *             example:
 *               maxOriginAmount: "999600"
 *               minOriginAmount: "4"
 *       400:
 *         description: Invalid input
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: object
 *                   properties:
 *                     value:
 *                       type: string
 *                     message:
 *                       type: string
 *                     field:
 *                       type: string
 *                     location:
 *                       type: string
 *             example:
 *               error:
 *                 value: "999"
 *                 message: "Unsupported fromChain"
 *                 field: "fromChain"
 *                 location: "query"
 *       500:
 *         description: Server error
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: string
 *                 details:
 *                   type: string
 */
router.get(
  '/',
  [
    check('fromChain')
      .exists()
      .withMessage('fromChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported fromChain'),
    check('toChain')
      .exists()
      .withMessage('toChain is required')
      .isNumeric()
      .custom((value) => CHAINS_ARRAY.some((c) => c.id === Number(value)))
      .withMessage('Unsupported toChain'),
    check('fromToken')
      .exists()
      .withMessage('fromToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid fromToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.fromChain as string)
      )
      .withMessage('Token not supported on specified chain'),
    check('toToken')
      .exists()
      .withMessage('toToken is required')
      .custom((value) => isTokenAddress(value))
      .withMessage('Invalid toToken address')
      .custom((value, { req }) =>
        isTokenSupportedOnChain(value, req.query.toChain as string)
      )
      .withMessage('Token not supported on specified chain'),
  ],
  showFirstValidationError,
  bridgeLimitsController
)

export default router
