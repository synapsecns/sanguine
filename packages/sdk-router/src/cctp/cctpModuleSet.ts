import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'
import { parseUnits } from 'ethers/lib/utils'

import {
  CircleFeesRequest as ChainDomains,
  ExecutorQuoteRequest,
  ExecutorQuoteResponse,
  getCircleFastAllowance,
  getCircleFees,
  getExecutorQuote,
} from './api'
import { CctpModule } from './cctpModule'
import {
  addressToBytes32,
  evmChainIdToWormholeChainId,
  serializeGasInstruction,
} from './utils'
import {
  CCTP_DOMAIN_MAP,
  CCTP_V2_EXECUTOR_ADDRESS_MAP,
  isChainIdSupported,
  MEDIAN_TIME_BLOCK,
  MEDIAN_TIME_CCTP_V2_FAST,
  MEDIAN_TIME_CCTP_V2_FAST_MAP,
  TOKEN_ZAP_V1_ADDRESS_MAP,
  USDC_ADDRESS_MAP,
} from '../constants'
import {
  BridgeRoute,
  BridgeRouteV2,
  BridgeTokenCandidate,
  FeeConfig,
  GetBridgeRouteV2Parameters,
  GetBridgeTokenCandidatesParameters,
  Query,
  SynapseModuleSet,
} from '../module'
import { ChainProvider } from '../router'
import { applySlippage, encodeZapData } from '../swap'
import { isSameAddress, logExecutionTime } from '../utils'

const FAST_FINALITY_THRESHOLD = 1000
/**
 * @dev We use 5x buffer to account for other people trying to use the same allowance.
 */
const ALLOWANCE_BUFFER = 5
/**
 * @dev We use 2x the expected fee to account for positive slippage,
 * as the fee is paid as percentage from the final amount.
 */
const MAX_FEE_MULTIPLIER = 2

export class CctpModuleSet extends SynapseModuleSet {
  public readonly moduleName = 'CCTP'
  public readonly allEvents = []
  public readonly isBridgeV2Supported = true

  public modules: {
    [chainId: number]: CctpModule
  }

  constructor(chains: ChainProvider[]) {
    super()
    this.modules = {}
    chains.forEach(({ chainId, provider }) => {
      const address = CCTP_V2_EXECUTOR_ADDRESS_MAP[chainId]
      // Skip chains without an address
      if (!address) {
        return
      }
      this.modules[chainId] = new CctpModule(chainId, provider, address)
    })
  }

  public getModule(chainId: number): CctpModule | undefined {
    return this.modules[chainId]
  }

  public getEstimatedTime(originChainId: number, destChainId?: number): number {
    const originEstimate =
      MEDIAN_TIME_CCTP_V2_FAST_MAP[originChainId] || MEDIAN_TIME_CCTP_V2_FAST
    const destEstimate =
      destChainId && isChainIdSupported(destChainId)
        ? MEDIAN_TIME_BLOCK[destChainId]
        : 0
    return originEstimate + destEstimate
  }

  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    toToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    // Check that both chains are supported
    if (!this.getModule(fromChainId) || !this.getModule(toChainId)) {
      return []
    }
    // USDC must exist on both chains
    const originToken = USDC_ADDRESS_MAP[fromChainId]
    const destToken = USDC_ADDRESS_MAP[toChainId]
    if (!originToken || !destToken) {
      return []
    }
    // If a specific token is requested, it must be USDC
    if (toToken && !isSameAddress(toToken, destToken)) {
      return []
    }
    return [
      {
        originChainId: fromChainId,
        originToken,
        destChainId: toChainId,
        destToken,
      },
    ]
  }

  @logExecutionTime('CctpModuleSet.getBridgeRouteV2')
  public async getBridgeRouteV2(
    params: GetBridgeRouteV2Parameters
  ): Promise<BridgeRouteV2 | undefined> {
    const tokenZap = TOKEN_ZAP_V1_ADDRESS_MAP[params.bridgeToken.originChainId]
    if (!this.validateBridgeRouteV2Params(params) || !tokenZap) {
      return undefined
    }
    const { originSwapRoute, bridgeToken, slippage } = params
    const domains = {
      sourceDomainId: CCTP_DOMAIN_MAP[bridgeToken.originChainId],
      destDomainId: CCTP_DOMAIN_MAP[bridgeToken.destChainId],
    }
    if (
      domains.sourceDomainId === undefined ||
      domains.destDomainId === undefined
    ) {
      return undefined
    }
    const request = this.getExecutorQuoteRequest(bridgeToken)
    const [quote, fastAllowance, fastFinalityFee] = await Promise.all([
      getExecutorQuote(request),
      getCircleFastAllowance(),
      this.getFastFinalityFee(domains),
    ])
    if (!quote || !fastAllowance || fastFinalityFee === null) {
      return undefined
    }
    // Check if the fast USDC allowance is sufficient
    if (
      parseUnits(fastAllowance.allowance.toFixed(6), 6).lt(
        originSwapRoute.expectedToAmount.mul(ALLOWANCE_BUFFER)
      )
    ) {
      // TODO: fallback to slow USDC transfer if insufficient
      return undefined
    }
    const expectedFee = this.getFee(
      originSwapRoute.expectedToAmount,
      fastFinalityFee
    )
    const expectedToAmount = originSwapRoute.expectedToAmount.sub(expectedFee)
    // With no slippage or no swap on origin, the minToAmount is the same as expectedToAmount.
    const hasOriginSlippage = !originSwapRoute.expectedToAmount.eq(
      originSwapRoute.minToAmount
    )
    const minToAmount =
      hasOriginSlippage && slippage
        ? applySlippage(expectedToAmount, slippage)
        : expectedToAmount
    const zapData = await this.getZapData(
      params,
      domains,
      expectedFee,
      request,
      quote
    )
    return {
      bridgeToken,
      toToken: bridgeToken.destToken,
      expectedToAmount,
      minToAmount,
      nativeFee: BigNumber.from(quote.estimatedCost),
      // Will be filled using getEstimatedTime
      estimatedTime: 0,
      zapData,
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

  private async getFastFinalityFee(
    domains: ChainDomains
  ): Promise<number | null> {
    const fees = await getCircleFees(domains)
    if (!fees || fees.length === 0) {
      return null
    }
    // Extract the fee for the fast finality threshold
    const feeInfo = fees.find(
      (fee) => fee.finalityThreshold === FAST_FINALITY_THRESHOLD
    )
    return feeInfo?.minimumFee ?? null
  }

  private getExecutorQuoteRequest(
    bridgeToken: BridgeTokenCandidate
  ): ExecutorQuoteRequest {
    return {
      srcChain: evmChainIdToWormholeChainId(bridgeToken.originChainId),
      dstChain: evmChainIdToWormholeChainId(bridgeToken.destChainId),
      relayInstructions: serializeGasInstruction({}),
    }
  }

  private async getZapData(
    params: GetBridgeRouteV2Parameters,
    domains: ChainDomains,
    expectedFee: BigNumber,
    request: ExecutorQuoteRequest,
    quote: ExecutorQuoteResponse
  ): Promise<string | undefined> {
    const module = this.getModule(params.bridgeToken.originChainId)
    if (!module || !params.fromSender || !params.toRecipient) {
      return undefined
    }
    const tx = await module.contract.populateTransaction.depositForBurn(
      0, // amount - will be populated within TokenZap using amountPosition
      request.dstChain, // destinationChain (WH)
      domains.destDomainId, // destinationDomain (CCTP)
      addressToBytes32(params.toRecipient), // mintRecipient
      params.bridgeToken.originToken, // burnToken
      addressToBytes32(AddressZero), // destinationCaller
      expectedFee.mul(MAX_FEE_MULTIPLIER), // maxFee (inflated to account for positive slippage)
      FAST_FINALITY_THRESHOLD, // minFinalityThreshold
      {
        refundAddress: params.fromSender,
        signedQuote: quote.signedQuote,
        instructions: request.relayInstructions,
      }, // executorArgs
      {
        dbps: 0,
        payee: AddressZero,
      } // feeArgs
    )
    return encodeZapData({
      target: module.address,
      payload: tx.data,
      amountPosition: 4, // first argument of depositForBurn (function selector is 4 bytes)
    })
  }

  private getFee(amount: BigNumber, fee: number): BigNumber {
    return amount.mul(fee).div(10000)
  }
}
