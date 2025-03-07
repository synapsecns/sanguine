import { BigNumber, Contract, PopulatedTransaction } from 'ethers'
import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'

import synapseIntentRouterAbi from '../abi/SynapseIntentRouter.json'
import {
  BigintIsh,
  SYNAPSE_INTENT_ROUTER_ADDRESS_MAP,
  TOKEN_ZAP_V1_ADDRESS_MAP,
} from '../constants'
import { ChainProvider } from '../router'
import { StepParams } from '../swap'
import { SynapseIntentRouter } from '../typechain/SynapseIntentRouter'
import { adjustValueIfNative, isNativeToken } from '../utils/handleNativeToken'

export class SynapseIntentRouterSet {
  static sirInterface = new Interface(synapseIntentRouterAbi)

  public providers: {
    [chainId: number]: Provider
  }

  private sirCache: {
    [chainId: number]: SynapseIntentRouter
  }

  constructor(chains: ChainProvider[]) {
    this.sirCache = {}
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      this.providers[chainId] = provider
    })
  }

  public async completeIntent(
    chainId: number,
    token: string,
    amount: BigintIsh,
    deadline: BigintIsh,
    steps: StepParams[]
  ): Promise<PopulatedTransaction> {
    const sir = this.getSir(chainId)
    return this._completeIntent(
      chainId,
      token,
      amount,
      deadline,
      steps,
      sir.populateTransaction.completeIntent
    )
  }

  public async completeIntentWithBalanceChecks(
    chainId: number,
    token: string,
    amount: BigintIsh,
    deadline: BigintIsh,
    steps: StepParams[]
  ): Promise<PopulatedTransaction> {
    const sir = this.getSir(chainId)
    return this._completeIntent(
      chainId,
      token,
      amount,
      deadline,
      steps,
      sir.populateTransaction.completeIntentWithBalanceChecks
    )
  }

  public getSir(chainId: number): SynapseIntentRouter {
    const address = this.getSirAddress(chainId)
    const provider = this.providers[chainId]
    if (!provider) {
      throw new Error(`Provider for chain ${chainId} not found`)
    }
    if (!this.sirCache[chainId]) {
      this.sirCache[chainId] = new Contract(
        address,
        synapseIntentRouterAbi,
        provider
      ) as SynapseIntentRouter
    }
    return this.sirCache[chainId]
  }

  public getSirAddress(chainId: number): string {
    const address = SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[chainId]
    if (!address) {
      throw new Error(
        `SynapseIntentRouter address not found for chain ${chainId}`
      )
    }
    return address
  }

  public getTokenZapAddress(chainId: number): string {
    const address = TOKEN_ZAP_V1_ADDRESS_MAP[chainId]
    if (!address) {
      throw new Error(`TokenZap address not found for chain ${chainId}`)
    }
    return address
  }

  private async _completeIntent(
    chainId: number,
    token: string,
    amount: BigintIsh,
    deadline: BigintIsh,
    steps: StepParams[],
    populateTx: (
      zapRecipient: string,
      amountIn: BigintIsh,
      deadline: BigintIsh,
      steps: StepParams[]
    ) => Promise<PopulatedTransaction>
  ): Promise<PopulatedTransaction> {
    const tokenZap = this.getTokenZapAddress(chainId)
    if (steps.length === 0) {
      throw new Error('No steps found')
    }
    if (isNativeToken(token)) {
      steps[0].msgValue = amount
    }
    const populatedTx = await populateTx(tokenZap, amount, deadline, steps)
    // Adjust the tx.value if the initial token is native
    return adjustValueIfNative(populatedTx, token, BigNumber.from(amount))
  }
}
