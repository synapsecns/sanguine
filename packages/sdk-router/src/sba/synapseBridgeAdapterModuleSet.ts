import { Zero } from '@ethersproject/constants'
import { BigNumber, BigNumberish } from 'ethers'
import NodeCache from 'node-cache'

import { MEDIAN_TIME_BLOCK } from '../constants'
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
import { ChainProvider } from '../router'
import { encodeZapData } from '../swap'
import { isSameAddress, logExecutionTime, logger } from '../utils'
import {
  getSbaChainMetadata,
  SBA_DEFAULT_DESTINATION_BLOCKS,
  SBA_ESTIMATED_TIME_CACHE_TTL,
} from './metadata'
import { getSbaRemoteToken, getSbaSupportedTokens } from './supportedTokens'
import {
  SynapseBridgeAdapterBridgeParams,
  SynapseBridgeAdapterModule,
} from './synapseBridgeAdapterModule'

export class SynapseBridgeAdapterModuleSet extends SynapseModuleSet {
  public readonly moduleName = 'SynapseBridge'
  public readonly allEvents = ['TokenSent', 'TokenReceived']
  public readonly isBridgeV2Supported = true

  public modules: {
    [chainId: number]: SynapseBridgeAdapterModule
  }

  private estimatedTimeCache: NodeCache

  constructor(chains: ChainProvider[]) {
    super()
    this.modules = {}
    this.estimatedTimeCache = new NodeCache({
      stdTTL: SBA_ESTIMATED_TIME_CACHE_TTL,
    })
    chains.forEach(({ chainId, provider }) => {
      const metadata = getSbaChainMetadata(chainId)
      if (!metadata) {
        return
      }
      this.modules[chainId] = new SynapseBridgeAdapterModule(
        chainId,
        provider,
        metadata.adapterAddress
      )
    })
  }

  public getModule(chainId: number): SynapseModule | undefined {
    return this.modules[chainId]
  }

  public getEstimatedTime(fromChainId: number, toChainId?: number): number {
    const cachedValue =
      toChainId &&
      this.estimatedTimeCache.get<number>(
        this.getCacheKey(fromChainId, toChainId)
      )
    return cachedValue ?? this.getFallbackEstimatedTime(fromChainId, toChainId)
  }

  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    toToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    const originModule = this.modules[fromChainId]
    const originMetadata = getSbaChainMetadata(fromChainId)
    const destMetadata = getSbaChainMetadata(toChainId)
    if (
      !originModule ||
      !originMetadata ||
      !this.modules[toChainId] ||
      !destMetadata
    ) {
      return []
    }
    return getSbaSupportedTokens(fromChainId, toChainId, toToken)
  }

  @logExecutionTime('SynapseBridgeAdapterModuleSet.getBridgeRouteV2')
  public async getBridgeRouteV2(
    params: GetBridgeRouteV2Parameters
  ): Promise<BridgeRouteV2 | undefined> {
    if (!this.validateBridgeRouteV2Params(params)) {
      return undefined
    }
    const { bridgeToken, originSwapRoute, fromSender, toRecipient } = params
    const originModule = this.modules[bridgeToken.originChainId]
    const destMetadata = getSbaChainMetadata(bridgeToken.destChainId)
    if (!originModule || !destMetadata) {
      return undefined
    }
    try {
      const artifactDestToken = getSbaRemoteToken(
        bridgeToken.originChainId,
        bridgeToken.originToken,
        bridgeToken.destChainId
      )
      if (
        !artifactDestToken ||
        !isSameAddress(artifactDestToken, bridgeToken.destToken)
      ) {
        return undefined
      }
      const cacheKey = this.getCacheKey(
        bridgeToken.originChainId,
        bridgeToken.destChainId
      )
      const estimatedTimePromise = this.estimatedTimeCache.has(cacheKey)
        ? Promise.resolve(undefined)
        : originModule.getEstimatedTime(bridgeToken.destChainId)
      const nativeFee = await originModule.getNativeFee(destMetadata.lzEid)
      const expectedToAmount = originSwapRoute.expectedToAmount
      if (expectedToAmount.isZero()) {
        return undefined
      }
      const estimatedTime = await estimatedTimePromise
      if (estimatedTime) {
        this.estimatedTimeCache.set(cacheKey, estimatedTime)
      }
      return {
        bridgeToken,
        toToken: artifactDestToken,
        expectedToAmount,
        minToAmount:
          originSwapRoute.steps.length === 0
            ? expectedToAmount
            : originSwapRoute.minToAmount,
        nativeFee,
        estimatedTime,
        zapData: this.getBridgeZapData(
          bridgeToken.originChainId,
          destMetadata.lzEid,
          bridgeToken.originToken,
          originSwapRoute.expectedToAmount,
          nativeFee,
          fromSender,
          toRecipient
        ),
      }
    } catch (error) {
      logger.error(
        `Failed to get SBA route for ${bridgeToken.originChainId} -> ${bridgeToken.destChainId}: ${error}`
      )
      return undefined
    }
  }

  public async getBridgeRoutes(): Promise<BridgeRoute[]> {
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
    return {
      originPeriod: 0,
      destPeriod: 0,
    }
  }

  public applySlippage(
    originQueryPrecise: Query,
    destQueryPrecise: Query
  ): { originQuery: Query; destQuery: Query } {
    return {
      originQuery: originQueryPrecise,
      destQuery: destQueryPrecise,
    }
  }

  private getBridgeZapData(
    fromChainId: number,
    dstEid: number,
    token: string,
    amount: BigNumberish,
    nativeFee: BigNumber,
    fromSender?: string,
    toRecipient?: string
  ): string | undefined {
    const module = this.modules[fromChainId]
    if (!module || !fromSender || !toRecipient) {
      return undefined
    }
    const bridgeParams: SynapseBridgeAdapterBridgeParams = {
      dstEid,
      toRecipient,
      token,
      amount,
    }
    return encodeZapData({
      target: module.address,
      payload: module.populateBridgeERC20(bridgeParams, nativeFee).data,
      amountPosition: module.getAmountPosition(),
    })
  }

  private getFallbackEstimatedTime(
    fromChainId: number,
    toChainId?: number
  ): number {
    const fromMetadata = getSbaChainMetadata(fromChainId)
    const fromBlockTime =
      MEDIAN_TIME_BLOCK[fromChainId as keyof typeof MEDIAN_TIME_BLOCK]
    if (!fromMetadata) {
      throw new Error(`No SBA metadata found for chain ${fromChainId}`)
    }
    if (!fromBlockTime) {
      throw new Error(`No median block time found for chain ${fromChainId}`)
    }
    const destBlockTime = toChainId
      ? MEDIAN_TIME_BLOCK[toChainId as keyof typeof MEDIAN_TIME_BLOCK]
      : fromBlockTime
    if (!destBlockTime) {
      throw new Error(`No median block time found for chain ${toChainId}`)
    }
    return Math.ceil(
      fromMetadata.originBlockConfirmations * fromBlockTime +
        SBA_DEFAULT_DESTINATION_BLOCKS * destBlockTime
    )
  }

  private getCacheKey(fromChainId: number, toChainId: number): string {
    return `${fromChainId}-${toChainId}`
  }
}
