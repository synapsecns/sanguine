import { Zero } from '@ethersproject/constants'
import { BigNumber, Contract } from 'ethers'

import {
  HYPERCORE_CHAIN_ID,
  LZ_EID_MAP,
  SupportedChainId,
  SYN_ADDRESS_MAP,
  SYN_COMPOSER_ADDRESS,
  SYN_OFT_ADDRESS_MAP,
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
import {
  AMOUNT_NOT_PRESENT,
  encodeZapData,
  USER_SIMULATED_ADDRESS,
} from '../swap'
import { isSameAddress, logger } from '../utils'
import {
  getSynComposer,
  quoteSynHyperCore,
  SynModule,
  SynSendParams,
} from './synModule'
import { getSynBridgeStatus } from './synStatus'

export class SynModuleSet extends SynapseModuleSet {
  public readonly moduleName = 'SYN'
  public readonly allEvents = []
  public readonly isBridgeV2Supported = true
  public modules: { [chainId: number]: SynModule } = {}
  public composer?: Contract

  constructor(chains: ChainProvider[]) {
    super()
    chains.forEach(({ chainId, provider }) => {
      const address = SYN_OFT_ADDRESS_MAP[chainId]
      if (address) {
        this.modules[chainId] = new SynModule(chainId, provider, address)
      }
      if (chainId === SupportedChainId.HYPEREVM) {
        this.composer = getSynComposer(provider)
      }
    })
  }

  public getModule(chainId: number): SynModule | undefined {
    // HyperCore uses the HyperEVM adapter/provider, never its own EVM RPC.
    return this.modules[
      chainId === HYPERCORE_CHAIN_ID ? SupportedChainId.HYPEREVM : chainId
    ]
  }

  public getBridgeTxStatus(
    destChainId: number,
    txHash: string
  ): Promise<boolean> {
    return getSynBridgeStatus(
      destChainId,
      txHash,
      this.modules[SupportedChainId.HYPEREVM]?.oftContract.provider
    )
  }

  public getEstimatedTime(fromChainId: number): number {
    // Wiring requires 64 Ethereum blocks or 100 HyperEVM blocks, plus execution.
    return fromChainId === SupportedChainId.ETH ? 840 : 160
  }

  public async getGasDropAmount(): Promise<BigNumber> {
    return Zero
  }

  public async getBridgeTokenCandidates({
    fromChainId,
    toChainId,
    fromToken,
    toToken,
  }: GetBridgeTokenCandidatesParameters): Promise<BridgeTokenCandidate[]> {
    const supported =
      (fromChainId === SupportedChainId.ETH &&
        [SupportedChainId.HYPEREVM, HYPERCORE_CHAIN_ID].includes(toChainId)) ||
      (fromChainId === SupportedChainId.HYPEREVM &&
        toChainId === SupportedChainId.ETH)
    const originToken = SYN_ADDRESS_MAP[fromChainId]
    const destToken = SYN_ADDRESS_MAP[toChainId]
    // SYN-only: origin swaps and all other pathways are outside this integration.
    if (
      !supported ||
      !this.getModule(fromChainId) ||
      !this.getModule(toChainId) ||
      !isSameAddress(fromToken, originToken) ||
      (toToken && !isSameAddress(toToken, destToken))
    ) {
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

  public async getBridgeRouteV2(
    params: GetBridgeRouteV2Parameters
  ): Promise<BridgeRouteV2 | undefined> {
    if (!this.validateBridgeRouteV2Params(params)) {
      return undefined
    }
    const { bridgeToken, originSwapRoute, fromSender, toRecipient } = params
    const candidates = await this.getBridgeTokenCandidates({
      fromChainId: bridgeToken.originChainId,
      toChainId: bridgeToken.destChainId,
      fromToken: bridgeToken.originToken,
      toToken: params.toToken,
    })
    if (
      !candidates.length ||
      originSwapRoute.steps.length > 0 ||
      !originSwapRoute.minToAmount.eq(originSwapRoute.expectedToAmount)
    ) {
      return undefined
    }
    const hyperCore = bridgeToken.destChainId === HYPERCORE_CHAIN_ID
    const recipient = toRecipient || USER_SIMULATED_ADDRESS
    if (hyperCore && toRecipient) {
      const [user, composer] = await Promise.all([
        this.composer!.coreUserExists(toRecipient),
        this.composer!.coreUserExists(SYN_COMPOSER_ADDRESS),
      ])
      if (!user.exists) {
        throw Object.assign(
          new Error(
            'Activate the recipient account on HyperCore before bridging SYN.'
          ),
          { code: 'HYPERCORE_ACCOUNT_INACTIVE' }
        )
      }
      if (!composer.exists) {
        throw new Error('SYN bridging to HyperCore is temporarily unavailable.')
      }
    }
    const module = this.modules[bridgeToken.originChainId]
    const sendParams: SynSendParams = {
      toEid:
        LZ_EID_MAP[
          hyperCore ? SupportedChainId.HYPEREVM : bridgeToken.destChainId
        ],
      toRecipient: recipient,
      amount: originSwapRoute.expectedToAmount,
      fromSender: fromSender || USER_SIMULATED_ADDRESS,
      hyperCoreRecipient: hyperCore ? recipient : undefined,
    }
    try {
      const expectedOftAmount = await module.getDestinationQuote(sendParams)
      if (expectedOftAmount.isZero()) {
        return undefined
      }
      const expectedToAmount = hyperCore
        ? await quoteSynHyperCore(this.composer!, expectedOftAmount)
        : expectedOftAmount
      if (expectedToAmount.isZero()) {
        return undefined
      }
      sendParams.minAmount = expectedOftAmount
      const nativeFee = await module.getNativeFee(sendParams)
      const zapData =
        fromSender && toRecipient
          ? encodeZapData({
              target: module.address,
              payload: module.populateOftSend(sendParams, nativeFee).data,
              // SYN routes have no origin swap: keep the quoted amount fixed, even if
              // TokenZap already holds tokens. Refund OFT dust to avoid SIR__UnspentFunds.
              amountPosition: AMOUNT_NOT_PRESENT,
              ...(originSwapRoute.expectedToAmount.gt(expectedOftAmount)
                ? {
                    finalToken: bridgeToken.originToken,
                    forwardTo: fromSender,
                  }
                : {}),
            })
          : undefined
      return {
        bridgeToken,
        toToken: bridgeToken.destToken,
        expectedToAmount,
        minToAmount: expectedToAmount,
        nativeFee,
        zapData,
      }
    } catch (error) {
      logger.error(
        `Failed to quote SYN for ${bridgeToken.originChainId} -> ${bridgeToken.destChainId}: ${error}`
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
      feeConfig: { bridgeFee: 0, minFee: Zero, maxFee: Zero },
    }
  }
  public getDefaultPeriods(): { originPeriod: number; destPeriod: number } {
    return { originPeriod: 0, destPeriod: 0 }
  }
  public applySlippage(
    originQuery: Query,
    destQuery: Query
  ): { originQuery: Query; destQuery: Query } {
    return { originQuery, destQuery }
  }
}
