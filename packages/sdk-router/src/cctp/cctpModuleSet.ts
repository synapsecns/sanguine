import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { CctpModule } from './cctpModule'
import {
  CCTP_V2_EXECUTOR_ADDRESS_MAP,
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
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import { ChainProvider } from '../router'
import { isSameAddress, logExecutionTime } from '../utils'

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
    chains.forEach(({ chainId }) => {
      const address = CCTP_V2_EXECUTOR_ADDRESS_MAP[chainId]
      // Skip chains without an address
      if (!address) {
        return
      }
      this.modules[chainId] = new CctpModule(chainId, address)
    })
  }

  public getModule(chainId: number): CctpModule | undefined {
    return this.modules[chainId]
  }

  public getEstimatedTime(): number {
    return 0
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
}
