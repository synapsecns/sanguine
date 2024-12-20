import { BigNumber } from 'ethers'
import { Provider } from '@ethersproject/abstract-provider'
import { Zero } from '@ethersproject/constants'
import { Contract } from '@ethersproject/contracts'

import erc20ABI from '../../abi/IERC20Metadata.json'
import { IERC20Metadata as ERC20 } from '../../typechain/IERC20Metadata'
import { AddressMap } from '../../constants'
import { isSameAddress } from '../../utils/addressUtils'
import { fetchWithTimeout } from '../api'
import {
  applySlippage,
  EmptyRoute,
  EngineID,
  isCorrectSlippage,
  RouteInput,
  SlippageMax,
  SwapEngine,
  SwapEngineRoute,
  toBasisPoints,
} from './swapEngine'
import {
  SwapAPIResponse,
  EMPTY_SWAP_API_RESPONSE,
  generateAPIStep,
} from './response'
import { ChainProvider } from '../../router'
import { isNativeToken } from '../../utils/handleNativeToken'

const PARASWAP_API_URL = 'https://api.paraswap.io/swap'
const PARASWAP_API_TIMEOUT = 2000

export type ParaSwapRequest = {
  srcToken: string
  srcDecimals: number
  destToken: string
  destDecimals: number
  amount: string
  side: string
  userAddress: string
  network: string
  slippage: number
  version: string
}

export type ParaSwapResponse = {
  priceRoute: {
    destAmount: string
  }
  txParams: {
    from: string
    to: string
    value: string
    data: string
    gasPrice: string
    chainId: number
  }
}

export class ParaSwapEngine implements SwapEngine {
  public readonly id: EngineID = EngineID.ParaSwap

  private readonly tokenZapAddressMap: AddressMap

  private providers: {
    [chainId: number]: Provider
  }

  constructor(chains: ChainProvider[], tokenZapAddressMap: AddressMap) {
    this.providers = {}
    this.tokenZapAddressMap = tokenZapAddressMap
    chains.forEach(({ chainId, provider }) => {
      this.providers[chainId] = provider
    })
  }

  public async findRoute(input: RouteInput): Promise<SwapEngineRoute> {
    const { chainId, tokenIn, tokenOut, amountIn, finalRecipient, slippage } =
      input
    const tokenZap = this.tokenZapAddressMap[chainId]
    if (
      !tokenZap ||
      isSameAddress(tokenIn, tokenOut) ||
      BigNumber.from(amountIn).eq(Zero) ||
      !isCorrectSlippage(slippage)
    ) {
      return EmptyRoute
    }
    const srcDecimals = await this.getTokenDecimals(chainId, tokenIn)
    const destDecimals = await this.getTokenDecimals(chainId, tokenOut)
    if (srcDecimals === 0 || destDecimals === 0) {
      return EmptyRoute
    }
    const request = {
      srcToken: tokenIn,
      srcDecimals,
      destToken: tokenOut,
      destDecimals,
      amount: amountIn.toString(),
      side: 'SELL',
      userAddress: tokenZap,
      network: chainId.toString(),
      // slippage settings are applied when generating the zap data as minFwdAmount
      slippage: toBasisPoints(SlippageMax),
      version: '6.2',
    }
    const response = await this.getResponse(request)
    const expectedAmountOut = BigNumber.from(response.amountOut)
    if (expectedAmountOut.eq(Zero)) {
      return EmptyRoute
    }
    const minAmountOut = applySlippage(expectedAmountOut, slippage)
    return {
      engineID: this.id,
      expectedAmountOut,
      minAmountOut,
      steps: [
        generateAPIStep(
          tokenIn,
          tokenOut,
          amountIn,
          response,
          finalRecipient,
          minAmountOut
        ),
      ],
    }
  }

  // TODO: findRoutes

  public async getResponse(request: ParaSwapRequest): Promise<SwapAPIResponse> {
    try {
      // Stringify every value in the request
      const params = new URLSearchParams(
        Object.entries(request).map(([k, v]) => {
          return [k, v.toString()]
        })
      )
      const url = `${PARASWAP_API_URL}?${params.toString()}`
      const response = await fetchWithTimeout(url, PARASWAP_API_TIMEOUT)
      if (!response.ok) {
        console.error('Error fetching ParaSwap response:', response)
        return EMPTY_SWAP_API_RESPONSE
      }
      const paraswapResponse: ParaSwapResponse = await response.json()
      return {
        amountOut: paraswapResponse.priceRoute.destAmount,
        transaction: paraswapResponse.txParams,
      }
    } catch (error) {
      console.error('Error fetching ParaSwap response:', error)
      return EMPTY_SWAP_API_RESPONSE
    }
  }

  private async getTokenDecimals(
    chainId: number,
    token: string
  ): Promise<number> {
    // TODO: cache, move to utils
    if (isNativeToken(token)) {
      return 18
    }
    const provider = this.providers[chainId]
    if (!provider) {
      console.error('No provider found for chainId', chainId)
      return 0
    }
    const tokenContract = new Contract(token, erc20ABI, provider) as ERC20
    try {
      return tokenContract.decimals()
    } catch (error) {
      console.error('Error fetching token decimals:', error)
      return 0
    }
  }
}
