import express from 'express'

import { indexController } from '../controllers/indexController'

const router: express.Router = express.Router()

/**
 * @openapi
 * /:
 *   get:
 *     summary: Get API information
 *     description: Retrieve general information about the Synapse REST API, including available chains and tokens
 *     responses:
 *       200:
 *         description: Successful response
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 message:
 *                   type: string
 *                   description: Welcome message for the API
 *                 availableChains:
 *                   type: array
 *                   description: List of available blockchain networks
 *                   items:
 *                     type: object
 *                     properties:
 *                       name:
 *                         type: string
 *                         description: Name of the blockchain network
 *                       id:
 *                         type: integer
 *                         description: Chain ID of the blockchain network
 *                 availableTokens:
 *                   type: array
 *                   description: List of available tokens across different chains
 *                   items:
 *                     type: object
 *                     properties:
 *                       symbol:
 *                         type: string
 *                         description: Token symbol
 *                       chains:
 *                         type: array
 *                         description: List of chains where the token is available
 *                         items:
 *                           type: object
 *                           properties:
 *                             chainId:
 *                               type: string
 *                               description: Chain ID where the token is available
 *                             address:
 *                               type: string
 *                               description: Token contract address on the specific chain
 *             example:
 *               message: "Welcome to the Synapse REST API for swap and bridge quotes"
 *               availableChains:
 *                 - name: "Ethereum"
 *                   id: 1
 *                 - name: "Arbitrum"
 *                   id: 42161
 *               availableTokens:
 *                 - symbol: "USDC"
 *                   chains:
 *                     - chainId: "1"
 *                       address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
 *                     - chainId: "42161"
 *                       address: "0xaf88d065e77c8cc2239327c5edb3a432268e5831"
 *       500:
 *         description: Server error
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: string
 */
router.get('/', indexController)

export default router
