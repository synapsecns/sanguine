import { BigNumber } from 'ethers'
import { Provider } from '@ethersproject/abstract-provider'
import { Zero } from '@ethersproject/constants'

import {
  BridgeRoute,
  FeeConfig,
  Query,
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import { ChainProvider } from '../router'
import { getChainIds } from './api'
import { GasZipModule } from './gasZipModule'

// TODO: figure out if accurate
const MEDIAN_TIME_GAS_ZIP = 30

export class GasZipModuleSet extends SynapseModuleSet {
  public readonly bridgeModuleName = 'GasZip'
  public readonly allEvents = []

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
    // TODO: implement
  }

  /**
   * @inheritdoc SynapseModuleSet.getFeeData
   */
  public async getFeeData(): Promise<{
    feeAmount: BigNumber
    feeConfig: FeeConfig
  }> {
    // TODO: implement
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
    // TODO: implement
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
    destQueryPrecise: Query,
    slipNumerator: number,
    slipDenominator: number
  ): { originQuery: Query; destQuery: Query } {
    // TODO: implement
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
