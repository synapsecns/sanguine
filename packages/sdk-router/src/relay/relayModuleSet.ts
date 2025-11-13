import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { RELAY_ADDRESS_MAP, TOKEN_ZAP_V1_ADDRESS_MAP } from '../constants'
import {
  BridgeRoute,
  BridgeRouteV2,
  BridgeTokenCandidate,
  FeeConfig,
  GetBridgeRouteV2Parameters,
  GetBridgeTokenCandidatesParameters,
  Query,
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import { RelayModule } from './relayModule'
import { ChainProvider } from '../router'
import {
  ETH_NATIVE_TOKEN_ADDRESS,
  isNativeToken,
  logExecutionTime,
} from '../utils'
import {
  addressToCurrency,
  getQuote,
  isStepActionable,
  QuoteRequest,
  StepData,
  TradeType,
} from './api'
import { encodeZapData, toBasisPoints, USER_SIMULATED_ADDRESS } from '../swap'

export class RelayModuleSet extends SynapseModuleSet {
  public readonly moduleName = 'Relay'
  public readonly allEvents = []
  public readonly isBridgeV2Supported = true

  public modules: {
    [chainId: number]: RelayModule
  }

  constructor(chains: ChainProvider[]) {
    super()
    this.modules = {}
    chains.forEach(({ chainId }) => {
      const address = RELAY_ADDRESS_MAP[chainId]
      // Skip chains without a relay address
      if (!address) {
        return
      }
      this.modules[chainId] = new RelayModule(chainId, address)
    })
  }

  public getModule(chainId: number): SynapseModule | undefined {
    return this.modules[chainId]
  }

  public getEstimatedTime(): number {
    // TODO: implement
    return 0
  }

  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    toToken,
    fromToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    // Check that both chains are supported
    if (!this.getModule(fromChainId) || !this.getModule(toChainId)) {
      return []
    }
    // Relay supports origin swaps natively
    return [
      {
        originChainId: fromChainId,
        destChainId: toChainId,
        originToken: fromToken,
        destToken: toToken ?? ETH_NATIVE_TOKEN_ADDRESS,
      },
    ]
  }

  @logExecutionTime('RelayModuleSet.getBridgeRouteV2')
  public async getBridgeRouteV2(
    params: GetBridgeRouteV2Parameters
  ): Promise<BridgeRouteV2 | undefined> {
    const tokenZap = TOKEN_ZAP_V1_ADDRESS_MAP[params.bridgeToken.originChainId]
    if (!this.validateBridgeRouteV2Params(params) || !tokenZap) {
      return undefined
    }
    const quoteRequest: QuoteRequest = {
      user: tokenZap,
      originChainId: params.bridgeToken.originChainId,
      destinationChainId: params.bridgeToken.destChainId,
      originCurrency: addressToCurrency(params.originSwapRoute.toToken),
      destinationCurrency: addressToCurrency(params.toToken),
      amount: params.originSwapRoute.expectedToAmount.toString(),
      tradeType: TradeType.ExactInput,
      recipient: params.toRecipient ?? USER_SIMULATED_ADDRESS,
      refundTo: params.fromSender ?? USER_SIMULATED_ADDRESS,
      refundOnOrigin: true,
      useReceiver: true,
      explicitDeposit: true,
      usePermit: false,
      slippageTolerance: params.slippage
        ? toBasisPoints(params.slippage).toString()
        : undefined,
    }
    const quote = await getQuote(quoteRequest)
    if (!quote) {
      return undefined
    }
    const relaySteps = quote.steps.filter(isStepActionable)
    // Only a single step is expected
    if (relaySteps.length !== 1 || relaySteps[0].items.length !== 1) {
      return undefined
    }
    const relayData = relaySteps[0].items[0].data
    const ethAmountIn = isNativeToken(params.originSwapRoute.toToken)
      ? BigNumber.from(params.originSwapRoute.expectedToAmount)
      : Zero
    const txValue = BigNumber.from(relayData.value)
    if (txValue.lt(ethAmountIn)) {
      return undefined
    }
    const expectedToAmount = BigNumber.from(quote.details.currencyOut.amount)
    if (expectedToAmount.isZero()) {
      return undefined
    }
    return {
      bridgeToken: {
        ...params.bridgeToken,
        destToken: params.toToken,
      },
      toToken: params.toToken,
      expectedToAmount,
      minToAmount: BigNumber.from(quote.details.currencyOut.minimumAmount),
      nativeFee: txValue.sub(ethAmountIn),
      estimatedTime: quote.details.timeEstimate,
      zapData: this.getZapData(params, relayData),
    }
  }

  public async getBridgeRoutes(): Promise<BridgeRoute[]> {
    // Bridge V1 is not supported
    return []
  }

  public async getFeeData(): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    // Bridge V1 is not supported
    return {
      feeAmount: Zero,
      feeConfig: {
        bridgeFee: 0,
        minFee: Zero,
        maxFee: Zero,
      },
    }
  }

  public getDefaultPeriods(): {
    originPeriod: number
    destPeriod: number
  } {
    // Bridge V1 is not supported
    return {
      originPeriod: 0,
      destPeriod: 0,
    }
  }

  public applySlippage(
    originQueryPrecise: Query,
    destQueryPrecise: Query
  ): { originQuery: Query; destQuery: Query } {
    // Bridge V1 is not supported
    return {
      originQuery: originQueryPrecise,
      destQuery: destQueryPrecise,
    }
  }

  private getZapData(
    params: GetBridgeRouteV2Parameters,
    relayData: StepData
  ): string | undefined {
    if (!params.fromSender || !params.toRecipient) {
      return undefined
    }
    return encodeZapData({
      target: relayData.to,
      payload: relayData.data,
    })
  }
}
