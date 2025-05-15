import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { USDT_OFT_ADDRESS_MAP } from '../constants'
import {
  BridgeRoute,
  BridgeRouteV2,
  BridgeTokenCandidate,
  FeeConfig,
  Query,
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import { UsdtModule } from './usdtModule'
import { ChainProvider } from '../router'
import { logExecutionTime } from '../utils'

// TODO: use actual estimates
const MEDIAN_TIME_USDT = 90

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

  public getEstimatedTime(): number {
    return MEDIAN_TIME_USDT
  }

  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  public async getBridgeTokenCandidates(): Promise<BridgeTokenCandidate[]> {
    // TODO: implement
    return []
  }

  @logExecutionTime('UsdtModuleSet.getBridgeRouteV2')
  public async getBridgeRouteV2(): Promise<BridgeRouteV2 | undefined> {
    // TODO: implement
    return undefined
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
}
