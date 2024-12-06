import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'
import { AddressZero, MaxUint256, Zero } from '@ethersproject/constants'

import fastBridgeV2Abi from '../abi/FastBridgeV2.json'
import previewerAbi from '../abi/SynapseIntentPreviewer.json'
import synapseIntentRouterAbi from '../abi/SynapseIntentRouter.json'
import tokenZapV1Abi from '../abi/TokenZapV1.json'
import {
  FastBridgeV2 as FastBridgeV2Contract,
  IFastBridge,
  IFastBridgeV2,
} from '../typechain/FastBridgeV2'
import { SynapseIntentPreviewer as PreviewerContract } from '../typechain/SynapseIntentPreviewer'
import { SynapseIntentRouter as SIRContract } from '../typechain/SynapseIntentRouter'
import { TokenZapV1 as TokenZapV1Contract } from '../typechain/TokenZapV1'
import { BigintIsh } from '../constants'
import { SynapseModule, CCTPRouterQuery } from '../module'
import { getMatchingTxLog } from '../utils/logs'
import { adjustValueIfNative, isNativeToken } from '../utils/handleNativeToken'
import { CACHE_TIMES, RouterCache } from '../utils/RouterCache'
import { StepParams, encodeStepParams, decodeStepParams } from './steps'

export class SynapseIntentRouter implements SynapseModule {
  static fastBridgeV2Interface = new Interface(fastBridgeV2Abi)
  static previewerInterface = new Interface(previewerAbi)
  static sirInterface = new Interface(synapseIntentRouterAbi)
  static tokenZapV1Interface = new Interface(tokenZapV1Abi)

  public readonly address: string
  public readonly chainId: number
  public readonly provider: Provider

  private readonly fastBridgeV2Contract: FastBridgeV2Contract
  private readonly previewerContract: PreviewerContract
  private readonly sirContract: SIRContract
  private readonly swapQuoterAddress: string
  private readonly tokenZapContract: TokenZapV1Contract

  // All possible events emitted by the FastBridgeV2 contract in the origin transaction (in alphabetical order)
  private readonly originEvents = ['BridgeRequested']

  constructor(
    chainId: number,
    provider: Provider,
    contracts: {
      fastBridgeV2Address: string
      previewerAddress: string
      sirAddress: string
      swapQuoterAddress: string
      tokenZapAddress: string
    }
  ) {
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    invariant(contracts.fastBridgeV2Address, 'ADDRESS_UNDEFINED')
    invariant(contracts.previewerAddress, 'ADDRESS_UNDEFINED')
    invariant(contracts.sirAddress, 'ADDRESS_UNDEFINED')
    invariant(contracts.swapQuoterAddress, 'ADDRESS_UNDEFINED')
    invariant(contracts.tokenZapAddress, 'ADDRESS_UNDEFINED')
    invariant(SynapseIntentRouter.fastBridgeV2Interface, 'INTERFACE_UNDEFINED')
    invariant(SynapseIntentRouter.previewerInterface, 'INTERFACE_UNDEFINED')
    invariant(SynapseIntentRouter.sirInterface, 'INTERFACE_UNDEFINED')
    invariant(SynapseIntentRouter.tokenZapV1Interface, 'INTERFACE_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.address = contracts.sirAddress
    this.swapQuoterAddress = contracts.swapQuoterAddress

    this.fastBridgeV2Contract = new Contract(
      contracts.fastBridgeV2Address,
      fastBridgeV2Abi,
      provider
    ) as FastBridgeV2Contract

    this.previewerContract = new Contract(
      contracts.previewerAddress,
      SynapseIntentRouter.previewerInterface,
      provider
    ) as PreviewerContract

    this.sirContract = new Contract(
      contracts.sirAddress,
      SynapseIntentRouter.sirInterface,
      provider
    ) as SIRContract

    this.tokenZapContract = new Contract(
      contracts.tokenZapAddress,
      SynapseIntentRouter.tokenZapV1Interface,
      provider
    ) as TokenZapV1Contract
  }

  /**
   * @inheritdoc SynapseModule.bridge
   */
  public async bridge(
    to: string,
    destChainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: CCTPRouterQuery,
    destQuery: CCTPRouterQuery
  ): Promise<PopulatedTransaction> {
    // Merge the preparation and final steps
    const steps: StepParams[] = [
      ...decodeStepParams(originQuery.rawParams),
      await this.getFinalStep(to, destChainId, originQuery.tokenOut, destQuery),
    ]
    if (isNativeToken(token)) {
      steps[0].msgValue = BigNumber.from(amount)
    }
    // Get data for the complete intent transaction
    const populatedTransaction =
      await this.sirContract.populateTransaction.completeIntentWithBalanceChecks(
        this.tokenZapContract.address,
        amount,
        originQuery.minAmountOut,
        originQuery.deadline,
        steps
      )
    // Adjust the tx.value if the initial token is native
    return adjustValueIfNative(
      populatedTransaction,
      token,
      BigNumber.from(amount)
    )
  }

  /**
   * @inheritdoc SynapseModule.getSynapseTxId
   */
  public async getSynapseTxId(txHash: string): Promise<string> {
    // TODO: this should support older instances of FastBridge to track legacy txs
    const fastBridgeLog = await getMatchingTxLog(
      this.provider,
      txHash,
      this.fastBridgeV2Contract,
      this.originEvents
    )
    // transactionId always exists in the log as we are using the correct ABI
    const parsedLog =
      this.fastBridgeV2Contract.interface.parseLog(fastBridgeLog)
    return parsedLog.args.transactionId
  }

  /**
   * @inheritdoc SynapseModule.getBridgeTxStatus
   */
  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    // TODO: this should support older instances of FastBridge to track legacy txs
    return this.fastBridgeV2Contract.bridgeRelays(synapseTxId)
  }

  // ══════════════════════════════════════════════ FAST BRIDGE V2 ═══════════════════════════════════════════════════

  public async getOriginAmountOut(
    tokenIn: string,
    rfqTokens: string[],
    amountIn: BigintIsh
  ): Promise<CCTPRouterQuery[]> {
    // TODO: this should be multicalled?
    return Promise.all(
      rfqTokens.map(async (tokenOut) => {
        // Get a quote and steps for the intent
        const { amountOut, steps: stepsOutput } =
          await this.previewerContract.previewIntent(
            this.swapQuoterAddress,
            tokenIn,
            tokenOut,
            amountIn
          )
        // Remove extra fields before the encoding
        const steps: StepParams[] = stepsOutput.map(
          ({ token, amount, msgValue, zapData }) => ({
            token,
            amount,
            msgValue,
            zapData,
          })
        )
        return {
          // To preserve consistency with other modules, router adapter is not set for a no-op intent
          routerAdapter: steps.length > 0 ? this.address : AddressZero,
          tokenOut,
          minAmountOut: amountOut,
          deadline: MaxUint256,
          rawParams: encodeStepParams(steps),
        }
      })
    )
  }

  /**
   * @returns The protocol fee rate, multiplied by 1_000_000 (e.g. 1 basis point = 100).
   */
  @RouterCache(CACHE_TIMES.TEN_MINUTES)
  public async getProtocolFeeRate(): Promise<BigNumber> {
    return this.fastBridgeV2Contract.protocolFeeRate()
  }

  // ═════════════════════════════════════════════════ SIR TOOLS ═════════════════════════════════════════════════════

  private async getFinalStep(
    to: string,
    dstChainId: number,
    originToken: string,
    dstQuery: CCTPRouterQuery
  ): Promise<StepParams> {
    const bridgeParamsV1: IFastBridge.BridgeParamsStruct = {
      dstChainId,
      // TODO
      sender: AddressZero,
      to,
      originToken,
      destToken: dstQuery.tokenOut,
      // Will be set by the TokenZap
      originAmount: 0,
      destAmount: dstQuery.minAmountOut,
      sendChainGas: false,
      deadline: dstQuery.deadline,
    }
    // TODO: non-default values
    const bridgeParamsV2: IFastBridgeV2.BridgeParamsV2Struct = {
      quoteRelayer: AddressZero,
      quoteExclusivitySeconds: 0,
      quoteId: '0x',
      zapNative: 0,
      zapData: '0x',
    }
    const fastBridgeV2CallData =
      this.fastBridgeV2Contract.interface.encodeFunctionData('bridgeV2', [
        bridgeParamsV1,
        bridgeParamsV2,
      ])
    const zapData = await this.tokenZapContract.encodeZapData(
      // target
      this.fastBridgeV2Contract.address,
      // payload
      fastBridgeV2CallData,
      // amount position: 6-th parameter
      4 + 32 * 5
    )
    return {
      token: originToken,
      // Use the full balance for the Zap action
      amount: MaxUint256,
      msgValue: Zero,
      zapData,
    }
  }
}
