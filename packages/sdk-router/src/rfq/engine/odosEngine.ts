import { Zero } from '@ethersproject/constants'

import { fetchWithTimeout } from '../api'
import { EMPTY_SWAP_API_RESPONSE, SwapAPIResponse } from './response'
import { SwapEngine, EngineID, SwapEngineRoute } from './swapEngine'

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

  public async findRoute(): Promise<SwapEngineRoute> {
    // TODO
    return {
      engineID: this.id,
      expectedAmountOut: Zero,
      minAmountOut: Zero,
      steps: [],
    }
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
