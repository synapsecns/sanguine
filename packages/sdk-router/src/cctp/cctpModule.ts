import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { Contract, PopulatedTransaction } from '@ethersproject/contracts'

import cctpV2WithExecutorAbi from '../abi/ICCTPv2WithExecutor.json'
import { SynapseModule } from '../module'
import { getExecutorTxStatus } from './api'
import { evmChainIdToWormholeChainId } from './utils'
import { ICCTPv2WithExecutor } from '../typechain/ICCTPv2WithExecutor'
import { marshallChainHash, unmarshallChainHash } from '../utils'

export class CctpModule implements SynapseModule {
  readonly address: string
  readonly chainId: number

  readonly contract: ICCTPv2WithExecutor

  constructor(chainId: number, provider: Provider, address: string) {
    this.chainId = chainId
    this.address = address

    this.contract = new Contract(
      address,
      new Interface(cctpV2WithExecutorAbi),
      provider
    ) as ICCTPv2WithExecutor
  }

  public async bridge(): Promise<PopulatedTransaction> {
    throw new Error('bridge V1 not supported')
  }

  public async getSynapseTxId(txHash: string): Promise<string> {
    // Encode origin chainId and txHash
    return marshallChainHash({
      chainId: this.chainId,
      hash: txHash,
    })
  }

  public async getBridgeTxStatus(synapseTxId: string): Promise<boolean> {
    // This will be called using destination chain module, so we need to restore the origin chain id
    const { chainId, hash } = unmarshallChainHash(synapseTxId)
    const whChainId = evmChainIdToWormholeChainId(chainId)
    const response = await getExecutorTxStatus({
      txHash: hash,
      chainId: whChainId,
    })
    if (!response) {
      return false
    }
    // Find transaction that was either submitted by requested executor (submitted) or someone else (nonce already used)
    return response.some(
      (tx) =>
        tx.txHash === hash &&
        (tx.status === 'submitted' ||
          tx.failureCause === 'evm_cctp_nonce_already_used')
    )
  }
}
