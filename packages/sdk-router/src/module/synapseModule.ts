import { PopulatedTransaction } from '@ethersproject/contracts'

import { Query } from '../router/query'
import { BigintIsh } from '../constants'

export interface SynapseModule {
  /**
   * Populates the transaction data for a bridge transaction.
   *
   * @param to - The address to send the bridged tokens to.
   * @param destChainId - The ID of the destination chain.
   * @param token - The address of the token to bridge.
   * @param amount - The amount of tokens to bridge.
   * @param originQuery - The Query struct with the information about swap to be executed on the origin chain.
   * @param destQuery - The Query struct with the information about swap to be executed on the destination chain.
   * @returns A promise that resolves to the populated transaction data.
   */
  bridge(
    to: string,
    destChainId: number,
    token: string,
    amount: BigintIsh,
    originQuery: Query,
    destQuery: Query
  ): Promise<PopulatedTransaction>

  /**
   * Returns the Synapse transaction ID for a given transaction hash on the current chain.
   * This is used to track the status of a bridge transaction originating from the current chain.
   *
   * @param txHash - The transaction hash of the bridge transaction.
   * @returns A promise that resolves to the Synapse transaction ID.
   */
  getSynapseTxId(txHash: string): Promise<string>

  /**
   * Checks whether a bridge transaction has been completed on the current chain.
   * This is used to track the status of a bridge transaction originating from another chain, having
   * current chain as the destination chain.
   *
   * @param synapseTxId - The unique Synapse txId of the bridge transaction.
   * @returns A promise that resolves to a boolean indicating whether the bridge transaction has been completed.
   */
  getBridgeTxStatus(synapseTxId: string): Promise<boolean>
}
