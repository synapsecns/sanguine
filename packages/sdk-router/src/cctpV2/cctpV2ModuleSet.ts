import { Zero } from '@ethersproject/constants'
import { BigNumber, BigNumberish } from 'ethers'

import {
  CCTP_V2_DOMAIN_MAP,
  CCTP_V2_FORWARD_SERVICE_FEE_USDC,
  CCTP_V2_TOKEN_MESSENGER_ADDRESS_MAP,
  CCTP_V2_USDC_ADDRESS_MAP,
  SupportedChainId,
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
import { encodeZapData } from '../swap'
import { isSameAddress, logExecutionTime } from '../utils'
import { CctpV2Fee, getBurnUSDCFees } from './api'
import { CCTPv2Module } from './cctpV2Module'

const BPS_DENOMINATOR = 10_000
const BURN_MAX_FEE_NUMERATOR = 11
const BURN_MAX_FEE_DENOMINATOR = 10
const FAST_FINALITY_THRESHOLD = 1000
const STANDARD_FINALITY_THRESHOLD = 2000

const CCTP_V2_STANDARD_ESTIMATED_TIME_BY_SOURCE_CHAIN: Record<number, number> =
  {
    [SupportedChainId.ETH]: 1080,
    [SupportedChainId.ARBITRUM]: 1080,
    [SupportedChainId.BASE]: 1080,
    [SupportedChainId.OPTIMISM]: 1080,
    [SupportedChainId.AVALANCHE]: 30,
    [SupportedChainId.POLYGON]: 30,
  }

const CCTP_V2_FAST_ESTIMATED_TIME_BY_SOURCE_CHAIN: Record<number, number> = {
  [SupportedChainId.ETH]: 630,
  [SupportedChainId.ARBITRUM]: 630,
  [SupportedChainId.BASE]: 630,
  [SupportedChainId.OPTIMISM]: 630,
  [SupportedChainId.AVALANCHE]: 30,
  [SupportedChainId.POLYGON]: 30,
}

export class CCTPv2ModuleSet extends SynapseModuleSet {
  public readonly moduleName = 'CCTPv2'
  public readonly allEvents = []
  public readonly isBridgeV2Supported = true

  public modules: {
    [chainId: number]: CCTPv2Module
  }

  constructor(chains: ChainProvider[]) {
    super()
    this.modules = {}
    chains.forEach(({ chainId }) => {
      const address = CCTP_V2_TOKEN_MESSENGER_ADDRESS_MAP[chainId]
      if (!address) {
        return
      }
      this.modules[chainId] = new CCTPv2Module(chainId, address)
    })
  }

  public getModule(chainId: number): SynapseModule | undefined {
    return this.modules[chainId]
  }

  public getModuleWithAddress(): SynapseModule | undefined {
    // This module is bridge-v2-only and must not be used by V1 `bridge()` flow.
    return undefined
  }

  public getEstimatedTime(originChainId: number): number {
    return this.getEstimatedTimeFallback(originChainId) ?? 0
  }

  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    toToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    const originToken = CCTP_V2_USDC_ADDRESS_MAP[fromChainId]
    const destToken = CCTP_V2_USDC_ADDRESS_MAP[toChainId]
    const originDomain = CCTP_V2_DOMAIN_MAP[fromChainId]
    const destDomain = CCTP_V2_DOMAIN_MAP[toChainId]
    if (
      !originToken ||
      !destToken ||
      originDomain === undefined ||
      destDomain === undefined
    ) {
      return []
    }
    if (
      !CCTP_V2_TOKEN_MESSENGER_ADDRESS_MAP[fromChainId] ||
      !CCTP_V2_TOKEN_MESSENGER_ADDRESS_MAP[toChainId] ||
      !this.getModule(fromChainId) ||
      !this.getModule(toChainId)
    ) {
      return []
    }
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

  @logExecutionTime('CCTPv2ModuleSet.getBridgeRouteV2')
  public async getBridgeRouteV2(
    params: GetBridgeRouteV2Parameters
  ): Promise<BridgeRouteV2 | undefined> {
    if (!this.validateBridgeRouteV2Params(params)) {
      return undefined
    }
    const { bridgeToken, originSwapRoute, toToken, allowMultipleTxs } = params
    const fromModule = this.modules[bridgeToken.originChainId]
    const originDomain = CCTP_V2_DOMAIN_MAP[bridgeToken.originChainId]
    const destDomain = CCTP_V2_DOMAIN_MAP[bridgeToken.destChainId]
    const originUsdc = CCTP_V2_USDC_ADDRESS_MAP[bridgeToken.originChainId]
    const destUsdc = CCTP_V2_USDC_ADDRESS_MAP[bridgeToken.destChainId]
    if (
      !fromModule ||
      originDomain === undefined ||
      destDomain === undefined ||
      !originUsdc ||
      !destUsdc ||
      !CCTP_V2_TOKEN_MESSENGER_ADDRESS_MAP[bridgeToken.originChainId] ||
      !CCTP_V2_TOKEN_MESSENGER_ADDRESS_MAP[bridgeToken.destChainId]
    ) {
      return undefined
    }
    if (
      !isSameAddress(bridgeToken.originToken, originUsdc) ||
      !isSameAddress(bridgeToken.destToken, destUsdc)
    ) {
      return undefined
    }
    if (!allowMultipleTxs && !isSameAddress(toToken, destUsdc)) {
      return undefined
    }
    const feeEntries = await getBurnUSDCFees(originDomain, destDomain)
    const selectedFee = feeEntries && this.getSlowestFeeEntry(feeEntries)
    if (!selectedFee) {
      return undefined
    }
    const estimatedTime = this.getEstimatedTimeByFinalityThreshold(
      bridgeToken.originChainId,
      selectedFee.finalityThreshold
    )
    if (estimatedTime === undefined) {
      return undefined
    }
    const amountInExpected = BigNumber.from(originSwapRoute.expectedToAmount)
    const amountInMin = BigNumber.from(originSwapRoute.minToAmount)
    const protocolFee = this.getProtocolFeeBudget(
      amountInExpected,
      selectedFee.minimumFee
    )
    if (protocolFee === undefined) {
      return undefined
    }
    const forwardingFee = this.getForwardingFeeBudget(
      bridgeToken.destChainId,
      selectedFee.forwardFee
    )
    if (forwardingFee === undefined) {
      return undefined
    }
    const maxFee = protocolFee.add(forwardingFee)
    const maxFeeForBurn = this.getBurnMaxFee(maxFee)
    if (amountInMin.lte(maxFeeForBurn)) {
      return undefined
    }
    const expectedToAmount = amountInExpected.sub(maxFee)
    const minToAmount = amountInMin.sub(maxFee)
    const routeToToken = destUsdc
    return {
      bridgeToken: {
        ...bridgeToken,
        destToken: destUsdc,
      },
      toToken: routeToToken,
      expectedToAmount,
      minToAmount,
      nativeFee: Zero,
      estimatedTime,
      zapData: this.getBurnZapData({
        ...params,
        toToken: routeToToken,
        maxFee: maxFeeForBurn,
        minFinalityThreshold: selectedFee.finalityThreshold,
      }),
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
    // Deadline settings are not supported by Circle CCTP V2.
    return {
      originPeriod: 0,
      destPeriod: 0,
    }
  }

  public applySlippage(
    originQueryPrecise: Query,
    destQueryPrecise: Query
  ): { originQuery: Query; destQuery: Query } {
    // Slippage settings are not supported by Circle CCTP V2.
    return {
      originQuery: originQueryPrecise,
      destQuery: destQueryPrecise,
    }
  }

  private getProtocolFeeBudget(
    amount: BigNumber,
    minimumFeeBps: number
  ): BigNumber | undefined {
    const minimumFee = this.parseDecimalToFraction(minimumFeeBps)
    if (!minimumFee) {
      return undefined
    }
    const denominator = minimumFee.denominator.mul(BPS_DENOMINATOR)
    return amount
      .mul(minimumFee.numerator)
      .add(denominator.sub(1))
      .div(denominator)
  }

  private parseDecimalToFraction(value: number):
    | {
        numerator: BigNumber
        denominator: BigNumber
      }
    | undefined {
    if (!Number.isFinite(value) || value < 0) {
      return undefined
    }
    const valueString = value.toString()
    if (valueString.toLowerCase().includes('e')) {
      return undefined
    }
    const [wholePart = '', fractionalPart = ''] = valueString.split('.')
    if (!/^\d+$/.test(wholePart) || !/^\d*$/.test(fractionalPart)) {
      return undefined
    }
    const decimals = fractionalPart.length
    const numerator = BigNumber.from(
      (
        `${wholePart}${fractionalPart}`.replace(/^0+(?=\d)/, '') || '0'
      ).toString()
    )
    return {
      numerator,
      denominator: BigNumber.from(10).pow(decimals),
    }
  }

  private getForwardingFeeBudget(
    destChainId: number,
    forwardFee?: Record<string, number>
  ): BigNumber | undefined {
    if (forwardFee !== undefined) {
      if (!forwardFee || typeof forwardFee !== 'object') {
        return undefined
      }
      // Circle forwarding tiers are key-based; unknown-only keysets fail closed.
      const parseTierFee = (tier: string): BigNumber | undefined => {
        const feeValue = forwardFee[tier]
        if (Number.isInteger(feeValue) && feeValue >= 0) {
          return BigNumber.from(feeValue)
        }
        return undefined
      }
      const hasMed = Object.hasOwn(forwardFee, 'med')
      const hasMedium = Object.hasOwn(forwardFee, 'medium')
      if (hasMed) {
        return parseTierFee('med')
      }
      if (hasMedium) {
        return parseTierFee('medium')
      }
      const highFee = parseTierFee('high')
      if (highFee) {
        return highFee
      }
      return parseTierFee('low')
    }
    const fallbackFee =
      CCTP_V2_FORWARD_SERVICE_FEE_USDC.perChainOverrides[destChainId] ??
      CCTP_V2_FORWARD_SERVICE_FEE_USDC.defaultFee
    return BigNumber.from(fallbackFee)
  }

  private getBurnMaxFee(maxFee: BigNumber): BigNumber {
    return maxFee
      .mul(BURN_MAX_FEE_NUMERATOR)
      .add(BURN_MAX_FEE_DENOMINATOR - 1)
      .div(BURN_MAX_FEE_DENOMINATOR)
  }

  private getSlowestFeeEntry(feeEntries: CctpV2Fee[]): CctpV2Fee | undefined {
    return feeEntries.reduce<CctpV2Fee | undefined>((slowest, entry) => {
      if (!slowest || entry.finalityThreshold > slowest.finalityThreshold) {
        return entry
      }
      return slowest
    }, undefined)
  }

  private getEstimatedTimeByFinalityThreshold(
    originChainId: number,
    finalityThreshold: number
  ): number | undefined {
    // Circle CCTP V2 finality modes: >=2000 is standard, [1000, 2000) is fast.
    if (finalityThreshold >= STANDARD_FINALITY_THRESHOLD) {
      return CCTP_V2_STANDARD_ESTIMATED_TIME_BY_SOURCE_CHAIN[originChainId]
    }
    if (finalityThreshold >= FAST_FINALITY_THRESHOLD) {
      return CCTP_V2_FAST_ESTIMATED_TIME_BY_SOURCE_CHAIN[originChainId]
    }
    return undefined
  }

  private getEstimatedTimeFallback(originChainId: number): number | undefined {
    return (
      CCTP_V2_STANDARD_ESTIMATED_TIME_BY_SOURCE_CHAIN[originChainId] ??
      CCTP_V2_FAST_ESTIMATED_TIME_BY_SOURCE_CHAIN[originChainId]
    )
  }

  private getBurnZapData({
    bridgeToken,
    originSwapRoute,
    fromSender,
    toRecipient,
    maxFee,
    minFinalityThreshold,
  }: GetBridgeRouteV2Parameters & {
    maxFee: BigNumberish
    minFinalityThreshold: number
  }): string | undefined {
    const module = this.modules[bridgeToken.originChainId]
    const destinationDomain = CCTP_V2_DOMAIN_MAP[bridgeToken.destChainId]
    if (
      !module ||
      destinationDomain === undefined ||
      !fromSender ||
      !toRecipient
    ) {
      return undefined
    }
    const burnParams = {
      destinationDomain,
      mintRecipient: toRecipient,
      burnToken: bridgeToken.originToken,
      maxFee,
      minFinalityThreshold,
    }
    return encodeZapData({
      target: module.address,
      payload: module.populateDepositForBurnWithHook({
        ...burnParams,
        amount: originSwapRoute.expectedToAmount,
      }).data,
      amountPosition: module.getAmountPosition(burnParams),
    })
  }
}
