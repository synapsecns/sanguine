import { BigNumber } from 'ethers'
import { Provider } from '@ethersproject/abstract-provider'
import { AddressZero, Zero } from '@ethersproject/constants'
import { Contract } from '@ethersproject/contracts'

import erc20ABI from '../../abi/IERC20Metadata.json'
import { IERC20Metadata as ERC20 } from '../../typechain/IERC20Metadata'
import { AddressMap, BigintIsh } from '../../constants'
import { isSameAddress } from '../../utils/addressUtils'
import { fetchWithTimeout } from '../api'
import {
  applySlippage,
  EmptyRoute,
  EngineID,
  isCorrectSlippage,
  Recipient,
  Slippage,
  SwapEngine,
  SwapEngineRoute,
} from './swapEngine'
import { StepParams } from '../steps'
import { AMOUNT_NOT_PRESENT, encodeZapData } from '../zapData'
import { ChainProvider } from '../../router'
import { isNativeToken } from '../../utils/handleNativeToken'

const PARASWAP_API_URL = 'https://api.paraswap.io/swap'
const PARASWAP_API_TIMEOUT = 1000

const MAX_SLIPPAGE = 9999

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

const EmptyParaSwapResponse: ParaSwapResponse = {
  priceRoute: {
    destAmount: '0',
  },
  txParams: {
    from: '',
    to: '',
    value: '',
    data: '',
    gasPrice: '',
    chainId: 0,
  },
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

  public async findRoute(
    chainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    finalRecipient: Recipient,
    slippage: Slippage
  ): Promise<SwapEngineRoute> {
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
      slippage: MAX_SLIPPAGE,
      version: '6.2',
    }
    const response = await this.getResponse(request)
    const expectedAmountOut = BigNumber.from(response.priceRoute.destAmount)
    if (expectedAmountOut.eq(Zero)) {
      return EmptyRoute
    }
    const minAmountOut = applySlippage(expectedAmountOut, slippage)
    return {
      engineID: this.id,
      expectedAmountOut,
      minAmountOut,
      steps: [
        this.generateParaSwapStep(
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

  public applySlippage(
    _chainId: number,
    route: SwapEngineRoute
  ): SwapEngineRoute {
    console.log(
      'Custom slippage settings are not supported for ParaSwap at the moment, default 0.1% will be used'
    )
    return route
  }

  public async getResponse(
    request: ParaSwapRequest
  ): Promise<ParaSwapResponse> {
    try {
      if (request.slippage > MAX_SLIPPAGE) {
        request.slippage = MAX_SLIPPAGE
      }
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
        return EmptyParaSwapResponse
      }
      return response.json()
    } catch (error) {
      console.error('Error fetching ParaSwap response:', error)
      return EmptyParaSwapResponse
    }
  }

  private generateParaSwapStep(
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    response: ParaSwapResponse,
    finalRecipient: Recipient,
    minAmountOut: BigNumber
  ): StepParams {
    if (isSameAddress(finalRecipient.address, AddressZero)) {
      throw new Error('Missing recipient address for ParaSwap')
    }
    const zapData = encodeZapData({
      target: response.txParams.to,
      payload: response.txParams.data,
      amountPosition: AMOUNT_NOT_PRESENT,
      finalToken: tokenOut,
      forwardTo: finalRecipient.address,
      minFwdAmount: minAmountOut,
    })
    return {
      token: tokenIn,
      amount: BigNumber.from(amountIn),
      msgValue: BigNumber.from(response.txParams.value),
      zapData,
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
