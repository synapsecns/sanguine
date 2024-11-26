import express from 'express'

import { tokenListController } from '../controllers/tokenListController'

const router: express.Router = express.Router()

/**
 * @openapi
 * /tokenlist:
 *   get:
 *     summary: Get the list of bridgeable tokens & associated chain metadata
 *     description: Retrieve the complete list of tokens that can be bridged across different chains
 *     responses:
 *       200:
 *         description: Successful response
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               additionalProperties:
 *                 type: object
 *                 properties:
 *                   addresses:
 *                     type: object
 *                     additionalProperties:
 *                       type: string
 *                   decimals:
 *                     type: object
 *                     additionalProperties:
 *                       type: integer
 *                   symbol:
 *                     type: string
 *                   name:
 *                     type: string
 *                   swapableType:
 *                     type: string
 *                   color:
 *                     type: string
 *                   priorityRank:
 *                     type: integer
 *                   routeSymbol:
 *                     type: string
 *                   imgUrl:
 *                     type: string
 *             example:
 *               USDC:
 *                 addresses:
 *                   "1": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
 *                   "10": "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85"
 *                 decimals:
 *                   "1": 6
 *                   "10": 6
 *                 symbol: "USDC"
 *                 name: "USD Coin"
 *                 swapableType: "USD"
 *                 color: "blue"
 *                 priorityRank: 100
 *                 routeSymbol: "USDC"
 *                 imgUrl: "https://example.com/usdc.svg"
 *               USDT:
 *                 addresses:
 *                   "1": "0xdac17f958d2ee523a2206206994597c13d831ec7"
 *                   "10": "0x94b008aA00579c1307B0EF2c499aD98a8ce58e58"
 *                 decimals:
 *                   "1": 6
 *                   "10": 6
 *                 symbol: "USDT"
 *                 name: "USD Tether"
 *                 swapableType: "USD"
 *                 color: "lime"
 *                 priorityRank: 100
 *                 routeSymbol: "USDT"
 *                 imgUrl: "https://example.com/usdt.svg"
 *               NUSD:
 *                 addresses:
 *                   "1": "0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F"
 *                   "10": "0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00"
 *                 decimals:
 *                   "1": 18
 *                   "10": 18
 *                 symbol: "nUSD"
 *                 name: "Synapse nUSD"
 *                 swapableType: "USD"
 *                 color: "purple"
 *                 priorityRank: 500
 *                 routeSymbol: "nUSD"
 *                 imgUrl: "https://example.com/nusd.svg"
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
router.get('/', tokenListController)

export default router
