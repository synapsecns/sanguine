import { Provider } from '@ethersproject/abstract-provider'
import { Zero } from '@ethersproject/constants'
import { BigNumber, BigNumberish } from 'ethers'

import { GASZIP_ADDRESS_MAP } from '../constants'
import {
  BridgeRoute,
  BridgeRouteV2,
  BridgeTokenCandidate,
  createNoSwapQuery,
  FeeConfig,
  GetBridgeRouteV2Parameters,
  GetBridgeTokenCandidatesParameters,
  Query,
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import { ChainProvider } from '../router'
import {
  Chains,
  getChains,
  getGasZipBlockHeightMap,
  getGasZipQuote,
} from './api'
import { GasZipModule } from './gasZipModule'
import { applySlippage, encodeZapData } from '../swap'
import {
  ETH_NATIVE_TOKEN_ADDRESS,
  isNativeToken,
  isSameAddress,
  logExecutionTime,
  logger,
} from '../utils'

const MEDIAN_TIME_GAS_ZIP = 30
const GAS_ZIP_MAX_BLOCK_AGE_MS = 5 * 60 * 1000

export class GasZipModuleSet extends SynapseModuleSet {
  public readonly moduleName = 'Gas.zip'
  public readonly allEvents = []
  public readonly isBridgeV2Supported = true

  public modules: {
    [chainId: number]: GasZipModule
  }
  public providers: {
    [chainId: number]: Provider
  }

  private cachedChains: Chains

  constructor(chains: ChainProvider[]) {
    super()
    this.modules = {}
    this.providers = {}
    this.cachedChains = {}
    chains.forEach(({ chainId, provider }) => {
      const address = GASZIP_ADDRESS_MAP[chainId]
      // Skip chains without a GasZip address
      if (!address) {
        return
      }
      this.modules[chainId] = new GasZipModule(chainId, provider, address)
      this.providers[chainId] = provider
    })
  }

  /**
   * @inheritdoc SynapseModuleSet.getModule
   */
  public getModule(chainId: number): SynapseModule | undefined {
    return this.modules[chainId]
  }

  /**
   * @inheritdoc SynapseModuleSet.getEstimatedTime
   */
  public getEstimatedTime(): number {
    return MEDIAN_TIME_GAS_ZIP
  }

  /**
   * @inheritdoc SynapseModuleSet.getGasDropAmount
   */
  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  @logExecutionTime('GasZipModuleSet.getBridgeTokenCandidates')
  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    toToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    // Check that both chains are supported by gas.zip
    const supportedChainIds = await this.getAllChainIds()
    if (
      !supportedChainIds.includes(fromChainId) ||
      !supportedChainIds.includes(toChainId)
    ) {
      return []
    }
    // Check that output token is native (if provided)
    if (toToken && !isNativeToken(toToken)) {
      return []
    }
    return [
      {
        originChainId: fromChainId,
        destChainId: toChainId,
        originToken: ETH_NATIVE_TOKEN_ADDRESS,
        destToken: ETH_NATIVE_TOKEN_ADDRESS,
      },
    ]
  }

  @logExecutionTime('GasZipModuleSet.getBridgeRouteV2')
  public async getBridgeRouteV2({
    originSwapRoute,
    bridgeToken,
    toToken,
    toRecipient,
    slippage,
    allowMultipleTxs,
  }: GetBridgeRouteV2Parameters): Promise<BridgeRouteV2 | undefined> {
    if (
      !this.getModule(bridgeToken.originChainId) ||
      !this.getModule(bridgeToken.destChainId)
    ) {
      return undefined
    }
    if (!allowMultipleTxs && !isSameAddress(bridgeToken.destToken, toToken)) {
      return undefined
    }
    const syncedPromise = this.checkBlockHeights(
      bridgeToken.originChainId,
      bridgeToken.destChainId
    )
    const quote = await getGasZipQuote(
      bridgeToken.originChainId,
      bridgeToken.destChainId,
      originSwapRoute.expectedToAmount
    )
    const expectedToAmount = quote.amountOut
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
    const route: BridgeRouteV2 = {
      bridgeToken,
      toToken: bridgeToken.destToken,
      expectedToAmount,
      minToAmount,
      zapData: await this.getGasZipZapData(
        bridgeToken.originChainId,
        originSwapRoute.expectedToAmount,
        bridgeToken.destChainId,
        toRecipient
      ),
    }
    // Verify that both chains are up to date before returning the route
    const synced = await syncedPromise
    if (!synced) {
      return undefined
    }
    return route
  }

  /**
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  @logExecutionTime('GasZipModuleSet.getBridgeRoutes')
  public async getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigNumberish
  ): Promise<BridgeRoute[]> {
    const syncedPromise = this.checkBlockHeights(originChainId, destChainId)
    // Check that both chains are supported by gas.zip
    const supportedChainIds = await this.getAllChainIds()
    if (
      !supportedChainIds.includes(originChainId) ||
      !supportedChainIds.includes(destChainId)
    ) {
      return []
    }
    // Check that both tokens are native assets
    if (!isNativeToken(tokenIn) || !isNativeToken(tokenOut)) {
      return []
    }
    const destGasZipChain = await this.getGasZipId(destChainId)
    if (!destGasZipChain) {
      return []
    }
    const quote = await getGasZipQuote(originChainId, destChainId, amountIn)
    // Check that non-zero amount is returned
    if (quote.amountOut.eq(Zero)) {
      return []
    }
    // Save destination gas.zip chain id in the destination query raw params
    const originQuery = createNoSwapQuery(tokenIn, BigNumber.from(amountIn))
    const destQuery = createNoSwapQuery(tokenOut, quote.amountOut)
    destQuery.rawParams = '0x' + destGasZipChain.toString(16)
    const route: BridgeRoute = {
      originChainId,
      destChainId,
      originQuery,
      destQuery,
      bridgeToken: {
        symbol: 'NATIVE',
        token: tokenIn,
      },
      bridgeModuleName: this.moduleName,
    }
    // Verify that both chains are up to date before returning the route
    const synced = await syncedPromise
    if (!synced) {
      return []
    }
    return [route]
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  public async getFeeData(): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    // There's no good way to determine the fee for gas.zip
    return {
      feeAmount: Zero,
      feeConfig: {
        bridgeFee: 0,
        minFee: BigNumber.from(0),
        maxFee: BigNumber.from(0),
      },
    }
  }

  /**
   * @inheritdoc SynapseModuleSet.getDefaultPeriods
   */
  public getDefaultPeriods(): {
    originPeriod: number
    destPeriod: number
  } {
    // Deadline settings are not supported by gas.zip
    return {
      originPeriod: 0,
      destPeriod: 0,
    }
  }

  /**
   * @inheritdoc SynapseModuleSet.applySlippage
   */
  public applySlippage(
    originQueryPrecise: Query,
    destQueryPrecise: Query
  ): { originQuery: Query; destQuery: Query } {
    // Slippage settings are not supported by gas.zip
    return {
      originQuery: originQueryPrecise,
      destQuery: destQueryPrecise,
    }
  }

  private async getAllChainIds(): Promise<number[]> {
    const chains = await this.getChains()
    return chains.chains?.map((chain) => chain.chain) ?? []
  }

  private async getGasZipId(chainId: number): Promise<number | undefined> {
    const chains = await this.getChains()
    return chains.chains?.find((chain) => chain.chain === chainId)?.short
  }

  private async getChains(): Promise<Chains> {
    if (!this.cachedChains.chains) {
      this.cachedChains = await getChains()
    }
    return this.cachedChains
  }

  private async getGasZipZapData(
    fromChainId: number,
    fromAmount: BigNumberish,
    toChainId: number,
    toRecipient?: string
  ): Promise<string | undefined> {
    const module = this.modules[fromChainId]
    if (!module || !toRecipient) {
      return undefined
    }
    const gasZipId = await this.getGasZipId(toChainId)
    if (!gasZipId) {
      return undefined
    }
    return encodeZapData({
      target: module.address,
      payload: module.populateGasZipTransaction(
        toRecipient,
        gasZipId,
        fromAmount
      ).data,
    })
  }

  /**
   * Checks if the latest block heights reported by gas.zip are within the maximum age.
   * Both chains must be up to date to enable the bridge.
   */
  private async checkBlockHeights(
    originChainId: number,
    destChainId: number
  ): Promise<boolean> {
    const blockHeightMap = await getGasZipBlockHeightMap()
    const [originSynced, destSynced] = await Promise.all([
      this.checkBlockHeight(originChainId, blockHeightMap.get(originChainId)),
      this.checkBlockHeight(destChainId, blockHeightMap.get(destChainId)),
    ])
    return originSynced && destSynced
  }

  /**
   * Checks if the block height is within the maximum age for a chain.
   */
  private async checkBlockHeight(
    chainId: number,
    blockHeight?: number
  ): Promise<boolean> {
    if (!blockHeight) {
      logger.info(`Gas.zip block height not found for chain ${chainId}`)
      return false
    }
    const provider = this.providers[chainId]
    if (!provider) {
      logger.info(`Provider not found for chain ${chainId}`)
      return false
    }
    let block
    try {
      block = await provider.getBlock(blockHeight)
    } catch (error) {
      logger.error(
        `Block height ${blockHeight} for chain ${chainId} not found: ${error}`
      )
      return false
    }
    if (!block) {
      logger.error(
        `Null block for chain ${chainId} at block height ${blockHeight}`
      )
      return false
    }
    const blockAge = Date.now() - block.timestamp * 1000
    const result = 0 <= blockAge && blockAge <= GAS_ZIP_MAX_BLOCK_AGE_MS
    if (!result) {
      logger.info(
        `Block height ${blockHeight} for chain ${chainId} is too old: ${blockAge} ms (allowed: 0 .. ${GAS_ZIP_MAX_BLOCK_AGE_MS} ms)`
      )
    }
    return result
  }
}
