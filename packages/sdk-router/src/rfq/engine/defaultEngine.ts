import { Interface } from '@ethersproject/abi'
import { hexlify } from '@ethersproject/bytes'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber, Contract } from 'ethers'
import invariant from 'tiny-invariant'

import defaultActionsAbi from '../../abi/IDefaultActions.json'
import previewerAbi from '../../abi/SynapseIntentPreviewer.json'
import {
  SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP,
  SWAP_QUOTER_V2_ADDRESS_MAP,
} from '../../constants/addresses'
import { BigintIsh } from '../../constants'
import { ChainProvider } from '../../router'
import { SynapseIntentPreviewer as PreviewerContract } from '../../typechain/SynapseIntentPreviewer'
import { IDefaultActionsInterface } from '../../typechain/IDefaultActions'
import { isSameAddress } from '../../utils/addressUtils'
import { decodeZapData, encodeZapData, ZapDataV1 } from '../zapData'
import {
  SwapEngine,
  SwapEngineRoute,
  EmptyRoute,
  Recipient,
  RecipientEntity,
  EngineID,
} from './swapEngine'

export class DefaultEngine implements SwapEngine {
  static defaultActions = new Interface(
    defaultActionsAbi
  ) as IDefaultActionsInterface
  static previewerInterface = new Interface(previewerAbi)

  public readonly id = EngineID.Default

  private contracts: {
    [chainId: number]: {
      previewer: PreviewerContract
      swapQuoter: string
    }
  }

  constructor(chains: ChainProvider[]) {
    invariant(DefaultEngine.defaultActions, 'INTERFACE_UNDEFINED')
    invariant(DefaultEngine.previewerInterface, 'INTERFACE_UNDEFINED')
    this.contracts = {}
    chains.forEach(({ chainId, provider }) => {
      const previewerAddress = SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP[chainId]
      // Skip chains without a SynapseIntentPreviewer address
      if (!previewerAddress) {
        return
      }
      // Swap Quoter must be defined for chains with SynapseIntentPreviewer
      const swapQuoterAddress = SWAP_QUOTER_V2_ADDRESS_MAP[chainId]
      invariant(
        swapQuoterAddress,
        'Swap Quoter address not found for chain ' + chainId
      )
      this.contracts[chainId] = {
        previewer: new Contract(
          previewerAddress,
          DefaultEngine.previewerInterface,
          provider
        ) as PreviewerContract,
        swapQuoter: swapQuoterAddress,
      }
    })
  }

  public async findRoute(
    chainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh,
    finalRecipient: Recipient,
    strictOut: boolean
  ): Promise<SwapEngineRoute> {
    const { previewer, swapQuoter } = this.contracts[chainId]
    if (
      !previewer ||
      !swapQuoter ||
      isSameAddress(tokenIn, tokenOut) ||
      BigNumber.from(amountIn).eq(Zero)
    ) {
      return EmptyRoute
    }
    // Get the quote
    const forwardTo = this.getForwardTo(finalRecipient)
    const { amountOut, steps: stepsOutput } = await previewer.previewIntent(
      swapQuoter,
      forwardTo,
      strictOut,
      tokenIn,
      tokenOut,
      amountIn
    )
    // Remove extra fields before the encoding
    return {
      engineID: this.id,
      expectedAmountOut: amountOut,
      minAmountOut: amountOut,
      steps: stepsOutput.map(({ token, amount, msgValue, zapData }) => ({
        token,
        amount,
        msgValue,
        zapData,
      })),
    }
  }

  public modifyMinAmountOut(
    _chainId: number,
    route: SwapEngineRoute,
    minAmountOut: BigintIsh
  ): SwapEngineRoute {
    const decodedZapData = this.getLastStepZapData(route)
    if (!decodedZapData.payload) {
      throw new Error('modifyMinAmountOut: no payload in the last step zapData')
    }
    let newPayload

    if (this.isSelectorMatching(decodedZapData.payload, 'addLiquidity')) {
      const params = DefaultEngine.defaultActions.decodeFunctionData(
        'addLiquidity',
        decodedZapData.payload
      ) as [BigNumber[], BigNumber, BigNumber]
      // addLiquidity(amounts, minToMint, deadline)
      newPayload = DefaultEngine.defaultActions.encodeFunctionData(
        'addLiquidity',
        [params[0], minAmountOut, params[2]]
      )
    }

    if (
      this.isSelectorMatching(decodedZapData.payload, 'removeLiquidityOneToken')
    ) {
      const params = DefaultEngine.defaultActions.decodeFunctionData(
        'removeLiquidityOneToken',
        decodedZapData.payload
      ) as [BigNumber, BigNumber, BigNumber, BigNumber]
      // removeLiquidityOneToken(tokenAmount, tokenIndex, minAmount, deadline)
      newPayload = DefaultEngine.defaultActions.encodeFunctionData(
        'removeLiquidityOneToken',
        [params[0], params[1], minAmountOut, params[3]]
      )
    }

    if (this.isSelectorMatching(decodedZapData.payload, 'swap')) {
      const params = DefaultEngine.defaultActions.decodeFunctionData(
        'swap',
        decodedZapData.payload
      ) as [BigNumber, BigNumber, BigNumber, BigNumber, BigNumber]
      // swap(tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
      newPayload = DefaultEngine.defaultActions.encodeFunctionData('swap', [
        params[0],
        params[1],
        params[2],
        minAmountOut,
        params[4],
      ])
    }

    if (
      this.isSelectorMatching(decodedZapData.payload, 'deposit') ||
      this.isSelectorMatching(decodedZapData.payload, 'withdraw')
    ) {
      newPayload = decodedZapData.payload
    }

    if (!newPayload) {
      throw new Error(
        'modifyMinAmountOut: no matching payload for the last step'
      )
    }
    // Last step exists after `getLastStepZapData`
    route.minAmountOut = BigNumber.from(minAmountOut)
    route.steps[route.steps.length - 1].zapData = encodeZapData({
      ...decodedZapData,
      payload: newPayload,
    })
    return route
  }

  public modifyRecipient(
    _chainId: number,
    route: SwapEngineRoute,
    finalRecipient: Recipient
  ): SwapEngineRoute {
    const decodedZapData = this.getLastStepZapData(route)
    if (!decodedZapData.forwardTo) {
      throw new Error(
        'modifyRecipient: no forwardTo address in the last step zapData'
      )
    }
    decodedZapData.forwardTo = this.getForwardTo(finalRecipient)
    // Last step exists after `getLastStepZapData`
    route.steps[route.steps.length - 1].zapData = encodeZapData(decodedZapData)
    return route
  }

  private getLastStepZapData(route: SwapEngineRoute): Partial<ZapDataV1> {
    const stepsCount = route.steps.length
    if (stepsCount === 0) {
      throw new Error('getLastStepZapData: no steps')
    }
    const lastStepZapData = hexlify(route.steps[stepsCount - 1].zapData)
    return decodeZapData(lastStepZapData)
  }

  private getForwardTo(recipient: Recipient): string {
    return recipient.entity === RecipientEntity.Self
      ? AddressZero
      : recipient.address
  }

  private isSelectorMatching(
    payload: string,
    functionName:
      | 'addLiquidity'
      | 'deposit'
      | 'removeLiquidityOneToken'
      | 'swap'
      | 'withdraw'
  ): boolean {
    return payload.startsWith(
      DefaultEngine.defaultActions.getSighash(
        DefaultEngine.defaultActions.getFunction(functionName)
      )
    )
  }
}
