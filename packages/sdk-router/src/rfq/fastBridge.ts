import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'

import fastBridgeAbi from '../abi/FastBridge.json'
import {
  FastBridge as FastBridgeContract,
  IFastBridge,
} from '../typechain/FastBridge'
import { BigintIsh } from '../constants'
import { SynapseModule, Query } from '../module'
import { getMatchingTxLog } from '../utils/logs'

// Define type alias
type BridgeParams = IFastBridge.BridgeParamsStruct

export class FastBridge implements SynapseModule {
  static fastBridgeInterface = new Interface(fastBridgeAbi)

  public readonly address: string
  public readonly chainId: number
  public readonly provider: Provider

  private readonly fastBridgeContract: FastBridgeContract

  // All possible events emitted by the FastBridge contract in the origin transaction (in alphabetical order)
  private readonly originEvents = ['BridgeRequested']

  constructor(chainId: number, provider: Provider, address: string) {
    invariant(chainId, 'CHAIN_ID_UNDEFINED')
    invariant(provider, 'PROVIDER_UNDEFINED')
    invariant(address, 'ADDRESS_UNDEFINED')
    invariant(FastBridge.fastBridgeInterface, 'INTERFACE_UNDEFINED')
    this.chainId = chainId
    this.provider = provider
    this.address = address
    this.fastBridgeContract = new Contract(
      address,
      FastBridge.fastBridgeInterface,
      provider
    ) as FastBridgeContract
  }

  /**
   * @inheritdoc SynapseModule.bridge
   */
  public async bridge(
    to: string,
    destChainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction> {
    // TODO: remove this once swaps on origin are supported
    invariant(
      BigNumber.from(amount).eq(originQuery.minAmountOut),
      'AMOUNT_MISMATCH'
    )
    invariant(
      token.toLowerCase() === originQuery.tokenOut.toLowerCase(),
      'TOKEN_MISMATCH'
    )
    // TODO: encode sendChainGas into destQuery.rawParams
    const bridgeParams: BridgeParams = {
      dstChainId: destChainId,
      to,
      originToken: token,
      destToken: destQuery.tokenOut,
      originAmount: amount,
      destAmount: destQuery.minAmountOut,
      sendChainGas: false,
      deadline: destQuery.deadline,
    }
    return this.fastBridgeContract.populateTransaction.bridge(bridgeParams)
  }

  /**
   * @inheritdoc SynapseModule.getSynapseTxId
   */
  public async getSynapseTxId(txHash: string): Promise<string> {
    const fastBridgeLog = await getMatchingTxLog(
      this.provider,
      txHash,
      this.fastBridgeContract,
      this.originEvents
    )
    // transactionId always exists in the log as we are using the correct ABI
    const parsedLog = this.fastBridgeContract.interface.parseLog(fastBridgeLog)
    return parsedLog.args.transactionId
  }

  /**
   * @inheritdoc SynapseModule.getBridgeTxStatus
   */
  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    return this.fastBridgeContract.bridgeRelays(synapseTxId)
  }
}
