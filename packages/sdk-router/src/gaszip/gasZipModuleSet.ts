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
  Query,
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import { ChainProvider } from '../router'
import { Chains, getChains, getGasZipQuote } from './api'
import { GasZipModule } from './gasZipModule'
import { isNativeToken } from '../utils'

const MEDIAN_TIME_GAS_ZIP = 30

export class GasZipModuleSet extends SynapseModuleSet {
  public readonly moduleName = 'Gas.zip'
  public readonly allEvents = []
  // Gas.zip does not support swaps on neither origin nor destination chains.
  public readonly isBridgeV2Supported = false

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

  public async getBridgeTokenCandidates(): Promise<BridgeTokenCandidate[]> {
    return []
  }

  public async getBridgeRouteV2(): Promise<BridgeRouteV2> {
    throw new Error('BridgeRouteV2 is not supported by ' + this.moduleName)
  }

  /**
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  public async getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigNumberish
  ): Promise<BridgeRoute[]> {
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
}
