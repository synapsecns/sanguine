import { Zero } from '@ethersproject/constants'
import { BigNumber, BigNumberish } from 'ethers'

import {
  LZ_EID_MAP,
  SupportedChainId,
  USDT0_ADDRESS_MAP,
  USDT_OFT_ADDRESS_MAP,
} from '../constants'
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
import { OftSendParams, UsdtModule } from './usdtModule'
import { ChainProvider } from '../router'
import { applySlippage, encodeZapData, USER_SIMULATED_ADDRESS } from '../swap'
import { isSameAddress, logExecutionTime } from '../utils'

// TODO: on-chain calls to get amount of block confirmations
const MEDIAN_TIME_USDT = 60
const MEDIAN_TIME_MAP: Record<number, number> = {
  [SupportedChainId.ARBITRUM]: 60,
  [SupportedChainId.BERACHAIN]: 2.5 * 60,
  [SupportedChainId.ETH]: 3.5 * 60,
  // This is not a joke, currently taking a full day to be processed
  [SupportedChainId.HYPEREVM]: 24 * 60 * 60,
  [SupportedChainId.OPTIMISM]: 16 * 60,
  [SupportedChainId.UNICHAIN]: 8 * 60,
}

export class UsdtModuleSet extends SynapseModuleSet {
  public readonly moduleName = 'USDT0'
  public readonly allEvents = []
  public readonly isBridgeV2Supported = true

  public modules: {
    [chainId: number]: UsdtModule
  }

  constructor(chains: ChainProvider[]) {
    super()
    this.modules = {}
    chains.forEach(({ chainId, provider }) => {
      const address = USDT_OFT_ADDRESS_MAP[chainId]
      // Skip chains without a USDT OFT address
      if (!address) {
        return
      }
      this.modules[chainId] = new UsdtModule(provider, address)
    })
  }

  public getModule(chainId: number): SynapseModule | undefined {
    return this.modules[chainId]
  }

  public getEstimatedTime(originChainId: number): number {
    return MEDIAN_TIME_MAP[originChainId] ?? MEDIAN_TIME_USDT
  }

  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    toToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    const originToken = USDT0_ADDRESS_MAP[fromChainId]
    const destToken = USDT0_ADDRESS_MAP[toChainId]
    // Skip chains without a USDT0 address
    if (!originToken || !destToken) {
      return []
    }
    // Check that output token is USDT0 (if provided)
    if (toToken && !isSameAddress(toToken, destToken)) {
      return []
    }
    return [
      {
        originChainId: fromChainId,
        destChainId: toChainId,
        originToken,
        destToken,
      },
    ]
  }

  @logExecutionTime('UsdtModuleSet.getBridgeRouteV2')
  public async getBridgeRouteV2(
    params: GetBridgeRouteV2Parameters
  ): Promise<BridgeRouteV2 | undefined> {
    if (!this.validateBridgeRouteV2Params(params)) {
      return undefined
    }
    const { originSwapRoute, bridgeToken, fromSender, toRecipient, slippage } =
      params
    const quote = await this.getUsdtSendQuote(
      bridgeToken.originChainId,
      originSwapRoute.expectedToAmount,
      bridgeToken.destChainId
    )
    if (!quote) {
      return undefined
    }
    const { expectedToAmount, nativeFee } = quote
    if (expectedToAmount.isZero()) {
      return undefined
    }
    // With no slippage or no swap on origin, the minToAmount is the same as expectedToAmount.
    const hasOriginSlippage = !originSwapRoute.expectedToAmount.eq(
      originSwapRoute.minToAmount
    )
    const minToAmount =
      hasOriginSlippage && slippage
        ? applySlippage(expectedToAmount, slippage)
        : expectedToAmount
    return {
      bridgeToken,
      toToken: bridgeToken.destToken,
      expectedToAmount,
      minToAmount,
      nativeFee,
      zapData: this.getUsdtSendZapData(
        bridgeToken.originChainId,
        originSwapRoute.expectedToAmount,
        bridgeToken.destChainId,
        nativeFee,
        fromSender,
        toRecipient
      ),
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
    // Deadline settings are not supported by USDT bridge
    return {
      originPeriod: 0,
      destPeriod: 0,
    }
  }

  public applySlippage(
    originQueryPrecise: Query,
    destQueryPrecise: Query
  ): { originQuery: Query; destQuery: Query } {
    // Slippage settings are not supported by USDT bridge
    return {
      originQuery: originQueryPrecise,
      destQuery: destQueryPrecise,
    }
  }

  private async getUsdtSendQuote(
    fromChainId: number,
    fromAmount: BigNumberish,
    toChainId: number
  ): Promise<
    { expectedToAmount: BigNumber; nativeFee: BigNumber } | undefined
  > {
    const module = this.modules[fromChainId]
    const toEid = LZ_EID_MAP[toChainId]
    if (!module || !toEid) {
      return undefined
    }
    const params = {
      toEid,
      toRecipient: USER_SIMULATED_ADDRESS,
      amount: fromAmount,
      fromSender: USER_SIMULATED_ADDRESS,
    }
    const [expectedToAmount, nativeFee] = await Promise.all([
      module.getDestinationQuote(params),
      module.getNativeFee(params),
    ])
    return { expectedToAmount, nativeFee }
  }

  private getUsdtSendZapData(
    fromChainId: number,
    fromAmount: BigNumberish,
    toChainId: number,
    nativeFee: BigNumber,
    fromSender?: string,
    toRecipient?: string
  ): string | undefined {
    const module = this.modules[fromChainId]
    const toEid = LZ_EID_MAP[toChainId]
    if (!module || !toEid || !fromSender || !toRecipient) {
      return undefined
    }
    const sendParams: OftSendParams = {
      toEid,
      toRecipient,
      amount: fromAmount,
      fromSender,
    }
    return encodeZapData({
      target: module.address,
      payload: module.populateOftSend(sendParams, nativeFee).data,
      amountPosition: module.getAmountPosition(),
    })
  }
}
