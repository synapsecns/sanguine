import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'
import { hexlify } from '@ethersproject/bytes'
import { AddressZero, MaxUint256, Zero } from '@ethersproject/constants'

import fastBridgeV2Abi from '../abi/FastBridgeV2.json'
import previewerAbi from '../abi/SynapseIntentPreviewer.json'
import synapseIntentRouterAbi from '../abi/SynapseIntentRouter.json'
import {
  FastBridgeV2 as FastBridgeV2Contract,
  IFastBridge,
} from '../typechain/FastBridgeV2'
import { SynapseIntentRouter as SIRContract } from '../typechain/SynapseIntentRouter'
import { BigintIsh } from '../constants'
import { SynapseModule, CCTPRouterQuery } from '../module'
import { getMatchingTxLog } from '../utils/logs'
import { adjustValueIfNative, isNativeToken } from '../utils/handleNativeToken'
import { CACHE_TIMES, RouterCache } from '../utils/RouterCache'
import { USER_SIMULATED_ADDRESS } from './engine'
import { decodeSavedBridgeParams } from './paramsV2'
import { StepParams, decodeStepParams } from './steps'
import { decodeZapData, encodeZapData } from './zapData'
import { isSameAddress } from '../utils/addressUtils'

export class SynapseIntentRouter implements SynapseModule {
  static fastBridgeV2Interface = new Interface(fastBridgeV2Abi)
  static previewerInterface = new Interface(previewerAbi)
  static sirInterface = new Interface(synapseIntentRouterAbi)

  public readonly address: string
  public readonly chainId: number
  public readonly provider: Provider

  private readonly fastBridgeV2Contract: FastBridgeV2Contract
  private readonly sirContract: SIRContract
  private readonly tokenZapAddress: string

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
    this.chainId = chainId
    this.provider = provider
    this.address = contracts.sirAddress
    this.tokenZapAddress = contracts.tokenZapAddress

    this.fastBridgeV2Contract = new Contract(
      contracts.fastBridgeV2Address,
      fastBridgeV2Abi,
      provider
    ) as FastBridgeV2Contract

    this.sirContract = new Contract(
      contracts.sirAddress,
      SynapseIntentRouter.sirInterface,
      provider
    ) as SIRContract
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
        this.tokenZapAddress,
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
    // We should have saved neccessary params within dstQuery.rawParams
    if (dstQuery.rawParams.length <= 2) {
      throw new Error('Missing bridge params for FastBridgeV2')
    }
    const { paramsV1, paramsV2 } = decodeSavedBridgeParams(dstQuery.rawParams)
    const dstZapData = decodeZapData(hexlify(paramsV2.zapData))
    if (paramsV1.originSender === AddressZero) {
      throw new Error('Missing sender address for FastBridgeV2')
    }
    if (paramsV1.destRelayRecipient === AddressZero) {
      throw new Error('Missing recipient address for FastBridgeV2')
    }
    // Override the simulated forward address if it was used.
    if (isSameAddress(paramsV1.destRelayRecipient, USER_SIMULATED_ADDRESS)) {
      paramsV1.destRelayRecipient = to
    }
    if (isSameAddress(dstZapData.forwardTo, USER_SIMULATED_ADDRESS)) {
      paramsV2.zapData = encodeZapData({
        ...dstZapData,
        forwardTo: to,
      })
    }
    const bridgeParamsV1: IFastBridge.BridgeParamsStruct = {
      dstChainId,
      sender: paramsV1.originSender,
      to: paramsV1.destRelayRecipient,
      originToken,
      destToken: paramsV1.destRelayToken,
      // Will be set in encodeZapData below
      originAmount: 0,
      destAmount: paramsV1.destRelayAmount,
      sendChainGas: false,
      deadline: dstQuery.deadline,
    }
    const fastBridgeV2CallData =
      this.fastBridgeV2Contract.interface.encodeFunctionData('bridgeV2', [
        bridgeParamsV1,
        paramsV2,
      ])
    // Amount is the 6-th parameter within the FastBridgeV2 call
    const originZapData = encodeZapData({
      target: this.fastBridgeV2Contract.address,
      payload: fastBridgeV2CallData,
      amountPosition: 4 + 32 * 5,
    })
    return {
      token: originToken,
      // Use the full balance for the Zap action
      amount: MaxUint256,
      msgValue: Zero,
      zapData: originZapData,
    }
  }
}
