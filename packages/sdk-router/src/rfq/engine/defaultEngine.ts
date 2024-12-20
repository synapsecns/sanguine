import { Interface } from '@ethersproject/abi'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber, Contract } from 'ethers'
import invariant from 'tiny-invariant'

import defaultActionsAbi from '../../abi/IDefaultActions.json'
import previewerAbi from '../../abi/SynapseIntentPreviewer.json'
import {
  SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP,
  SWAP_QUOTER_V2_ADDRESS_MAP,
} from '../../constants/addresses'
import { ChainProvider } from '../../router'
import { SynapseIntentPreviewer as PreviewerContract } from '../../typechain/SynapseIntentPreviewer'
import { IDefaultActionsInterface } from '../../typechain/IDefaultActions'
import { isSameAddress } from '../../utils/addressUtils'
import {
  SwapEngine,
  SwapEngineRoute,
  EmptyRoute,
  Recipient,
  RecipientEntity,
  EngineID,
  toWei,
  applySlippage,
  isCorrectSlippage,
  RouteInput,
  SlippageMax,
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

  public async findRoute(input: RouteInput): Promise<SwapEngineRoute> {
    const { chainId, tokenIn, tokenOut, amountIn, finalRecipient, slippage } =
      input
    const { previewer, swapQuoter } = this.contracts[chainId]
    if (
      !previewer ||
      !swapQuoter ||
      isSameAddress(tokenIn, tokenOut) ||
      BigNumber.from(amountIn).eq(Zero) ||
      !isCorrectSlippage(slippage)
    ) {
      return EmptyRoute
    }
    // Get the quote
    const forwardTo = this.getForwardTo(finalRecipient)
    const { amountOut, steps: stepsOutput } = await previewer.previewIntent(
      swapQuoter,
      forwardTo,
      // slippage settings are applied when generating the zap data as minFwdAmount
      toWei(SlippageMax),
      tokenIn,
      tokenOut,
      amountIn
    )
    const minAmountOut = applySlippage(amountOut, slippage)
    // Remove extra fields before the encoding
    return {
      engineID: this.id,
      expectedAmountOut: amountOut,
      minAmountOut,
      steps: stepsOutput.map(({ token, amount, msgValue, zapData }) => ({
        token,
        amount,
        msgValue,
        zapData,
      })),
    }
  }

  // TODO: findRoutes

  private getForwardTo(recipient: Recipient): string {
    return recipient.entity === RecipientEntity.Self
      ? AddressZero
      : recipient.address
  }
}
