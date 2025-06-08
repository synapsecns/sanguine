import { Interface } from '@ethersproject/abi'
import { Provider } from '@ethersproject/abstract-provider'
import { MaxUint256, Zero } from '@ethersproject/constants'
import { BigNumberish, Contract, PopulatedTransaction } from 'ethers'
import { uuidv7 } from 'uuidv7'

import synapseIntentRouterAbi from '../abi/SynapseIntentRouter.json'
import {
  SYNAPSE_INTENT_ROUTER_ADDRESS_MAP,
  TOKEN_ZAP_V1_ADDRESS_MAP,
} from '../constants'
import { BridgeRouteV2 } from '../module'
import { ChainProvider } from '../router'
import { StepParams, SwapEngineRoute } from '../swap'
import { SynapseIntentRouter } from '../typechain/SynapseIntentRouter'
import { BridgeQuoteV2 } from '../types'
import {
  isNativeToken,
  calculateDeadline,
  logExecutionTime,
  TEN_MINUTES,
  stringifyPopulatedTransaction,
} from '../utils'

const FULL_BALANCE = MaxUint256

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

  @logExecutionTime('SynapseIntentRouterSet.finalizeBridgeRouteV2')
  public async finalizeBridgeRouteV2(
    fromToken: string,
    fromAmount: BigNumberish,
    originSwapRoute: SwapEngineRoute,
    bridgeRoute: BridgeRouteV2,
    originDeadline?: number
  ): Promise<BridgeQuoteV2> {
    const fromChainId = bridgeRoute.bridgeToken.originChainId
    const moduleNames: string[] = []
    if (originSwapRoute.steps.length > 0) {
      moduleNames.push(originSwapRoute.engineName)
    }
    const tx = bridgeRoute.zapData
      ? await this.completeIntentWithBalanceChecks(
          fromChainId,
          fromToken,
          fromAmount,
          originDeadline ?? calculateDeadline(TEN_MINUTES),
          [
            ...originSwapRoute.steps,
            {
              token: bridgeRoute.bridgeToken.originToken,
              amount: FULL_BALANCE,
              msgValue: bridgeRoute.nativeFee,
              zapData: bridgeRoute.zapData,
            },
          ]
        )
      : undefined
    const bridgeQuoteV2: BridgeQuoteV2 = {
      id: uuidv7(),
      fromChainId,
      fromToken,
      fromAmount: fromAmount.toString(),
      toChainId: bridgeRoute.bridgeToken.destChainId,
      toToken: bridgeRoute.toToken,
      expectedToAmount: bridgeRoute.expectedToAmount.toString(),
      minToAmount: bridgeRoute.minToAmount.toString(),
      routerAddress: this.getSirAddress(fromChainId),
      // These will be filled by the corresponding bridge module
      estimatedTime: 0,
      moduleNames,
      gasDropAmount: '0',
      nativeFee: bridgeRoute.nativeFee.toString(),
      tx: stringifyPopulatedTransaction(tx),
    }
    return bridgeQuoteV2
  }

  public async completeIntent(
    chainId: number,
    token: string,
    amount: BigNumberish,
    deadline: BigNumberish,
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
    amount: BigNumberish,
    deadline: BigNumberish,
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
    amount: BigNumberish,
    deadline: BigNumberish,
    steps: StepParams[],
    populateTx: (
      zapRecipient: string,
      amountIn: BigNumberish,
      deadline: BigNumberish,
      steps: StepParams[]
    ) => Promise<PopulatedTransaction>
  ): Promise<PopulatedTransaction> {
    const tokenZap = this.getTokenZapAddress(chainId)
    if (steps.length === 0) {
      throw new Error('No steps found')
    }
    if (isNativeToken(token)) {
      // TODO: msgValue might be always set at this point
      steps[0].msgValue = amount
    }
    // Use the total msgValue of all steps as the tx value
    const txValue = steps.reduce((acc, step) => acc.add(step.msgValue), Zero)
    if (isNativeToken(token)) {
      // Adjust the native token amount to match the tx value, but keep the amount for the first step
      steps[0].amount = amount
      amount = txValue
    }
    const populatedTx = await populateTx(tokenZap, amount, deadline, steps)
    populatedTx.value = txValue
    return populatedTx
  }
}
