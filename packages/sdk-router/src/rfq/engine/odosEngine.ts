import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'

import { fetchWithTimeout } from '../api'
import {
  EMPTY_SWAP_API_RESPONSE,
  generateAPIRoute,
  SwapAPIResponse,
} from './response'
import {
  SwapEngine,
  EngineID,
  SwapEngineRoute,
  RouteInput,
  isCorrectSlippage,
  EmptyRoute,
  toPercentFloat,
} from './swapEngine'
import { AddressMap } from '../../constants'
import { isSameAddress } from '../../utils/addressUtils'

const ODOS_API_URL = 'https://api.odos.xyz/sor'
const ODOS_API_TIMEOUT = 2000

type OdosQuoteRequest = {
  chainId: number
  inputTokens: {
    amount: string
    tokenAddress: string
  }[]
  outputTokens: {
    proportion: number
    tokenAddress: string
  }[]
  userAddr: string
  slippageLimitPercent: number
  referralCode?: number
  simple?: boolean
}

type OdosQuoteResponse = {
  pathId: string
  outAmounts: string[]
}

type OdosAssembleRequest = {
  userAddr: string
  pathId: string
  receiver?: string
}

type OdosAssembleResponse = {
  transaction: {
    chainId: number
    gas: number
    gasPrice: number
    value: string
    to: string
    from: string
    data: string
  }
}

export class OdosEngine implements SwapEngine {
  readonly id: EngineID = EngineID.Odos

  private readonly tokenZapAddressMap: AddressMap

  constructor(tokenZapAddressMap: AddressMap) {
    this.tokenZapAddressMap = tokenZapAddressMap
  }

  public async findRoute(input: RouteInput): Promise<SwapEngineRoute> {
    const { chainId, tokenIn, tokenOut, amountIn, slippage } = input
    const tokenZap = this.tokenZapAddressMap[chainId]
    if (
      !tokenZap ||
      isSameAddress(tokenIn, tokenOut) ||
      BigNumber.from(amountIn).eq(Zero) ||
      !isCorrectSlippage(slippage)
    ) {
      return EmptyRoute
    }
    const request: OdosQuoteRequest = {
      chainId,
      inputTokens: [
        {
          amount: amountIn.toString(),
          tokenAddress: tokenIn,
        },
      ],
      outputTokens: [
        {
          proportion: 1,
          tokenAddress: tokenOut,
        },
      ],
      userAddr: tokenZap,
      slippageLimitPercent: toPercentFloat(slippage),
      simple: true,
    }
    const response = await this.getResponse(request)
    return generateAPIRoute(input, this.id, response)
  }

  public async getResponse(
    request: OdosQuoteRequest
  ): Promise<SwapAPIResponse> {
    try {
      if (
        request.inputTokens.length !== 1 ||
        request.outputTokens.length !== 1
      ) {
        console.error({ request }, 'Multi-token swaps not supported')
        return EMPTY_SWAP_API_RESPONSE
      }
      // Get a quote with the pathID first
      const response = await fetchWithTimeout(
        `${ODOS_API_URL}/quote/v2`,
        ODOS_API_TIMEOUT,
        {
          method: 'POST',
          body: JSON.stringify(request),
          headers: {
            'Content-Type': 'application/json',
          },
        }
      )
      if (!response.ok) {
        console.error({ request, response }, 'Error fetching Odos response')
        return EMPTY_SWAP_API_RESPONSE
      }
      const odosQuoteResponse: OdosQuoteResponse = await response.json()
      if (
        odosQuoteResponse.outAmounts.length !== 1 ||
        !odosQuoteResponse.pathId
      ) {
        console.error({ request, response }, 'Invalid Odos response')
        return EMPTY_SWAP_API_RESPONSE
      }
      const amountOut = odosQuoteResponse.outAmounts[0]
      if (amountOut === '0') {
        console.info({ request, response }, 'Zero amount out')
        return EMPTY_SWAP_API_RESPONSE
      }
      // Once we verified the amount out, we can assemble the transaction
      const assembleRequest: OdosAssembleRequest = {
        userAddr: request.userAddr,
        pathId: odosQuoteResponse.pathId,
      }
      const assembleResponse = await fetchWithTimeout(
        `${ODOS_API_URL}/assemble`,
        ODOS_API_TIMEOUT,
        {
          method: 'POST',
          body: JSON.stringify(assembleRequest),
          headers: {
            'Content-Type': 'application/json',
          },
        }
      )
      if (!assembleResponse.ok) {
        console.error({ assembleResponse }, 'Error fetching Odos response')
        return EMPTY_SWAP_API_RESPONSE
      }
      const odosAssembleResponse: OdosAssembleResponse =
        await assembleResponse.json()
      return {
        amountOut,
        transaction: odosAssembleResponse.transaction,
      }
    } catch (error) {
      console.error({ request, error }, 'Error fetching Odos response')
      return EMPTY_SWAP_API_RESPONSE
    }
  }
}
