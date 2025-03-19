import { BigNumber } from 'ethers'
import { Provider } from '@ethersproject/abstract-provider'
import { AddressZero, Zero } from '@ethersproject/constants'

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
import { getChainIds, getGasZipQuote } from './api'
import { GasZipModule } from './gasZipModule'
import { isNativeToken } from '../utils'
import { BigintIsh } from '../constants'

const MEDIAN_TIME_GAS_ZIP = 30

export class GasZipModuleSet extends SynapseModuleSet {
  public readonly bridgeModuleName = 'Gas.zip'
  public readonly allEvents = []
  // Gas.zip does not support swaps on neither origin nor destination chains.
  public readonly isBridgeV2Supported = false

  public modules: {
    [chainId: number]: GasZipModule
  }
  public providers: {
    [chainId: number]: Provider
  }

  private cachedChainIds: number[]

  constructor(chains: ChainProvider[]) {
    super()
    this.modules = {}
    this.providers = {}
    this.cachedChainIds = []
    chains.forEach(({ chainId, provider }) => {
      this.modules[chainId] = new GasZipModule(chainId, provider)
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
    throw new Error(
      'BridgeRouteV2 is not supported by ' + this.bridgeModuleName
    )
  }

  /**
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  public async getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    originUserAddress?: string
  ): Promise<BridgeRoute[]> {
    // Check that both chains are supported by gas.zip
    const supportedChainIds = await this.getChainIds()
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
    const user = originUserAddress ?? AddressZero
    const quote = await getGasZipQuote(
      originChainId,
      destChainId,
      amountIn,
      user,
      user
    )
    // Check that non-zero amount is returned
    if (quote.amountOut.eq(Zero)) {
      return []
    }
    // Save user address in the origin query raw params
    const originQuery = createNoSwapQuery(tokenIn, BigNumber.from(amountIn))
    originQuery.rawParams = quote.calldata
    const destQuery = createNoSwapQuery(tokenOut, quote.amountOut)
    destQuery.rawParams = user
    const route: BridgeRoute = {
      originChainId,
      destChainId,
      originQuery,
      destQuery,
      bridgeToken: {
        symbol: 'NATIVE',
        token: tokenIn,
      },
      bridgeModuleName: this.bridgeModuleName,
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

  private async getChainIds(): Promise<number[]> {
    if (this.cachedChainIds.length === 0) {
      this.cachedChainIds = await getChainIds()
    }
    return this.cachedChainIds
  }
}
