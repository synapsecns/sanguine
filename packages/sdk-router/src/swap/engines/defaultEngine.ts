import { Interface } from '@ethersproject/abi'
import { Zero } from '@ethersproject/constants'
import { BigNumber, Contract } from 'ethers'
import invariant from 'tiny-invariant'

import previewerAbi from '../../abi/SynapseIntentPreviewer.json'
import {
  SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP,
  SWAP_QUOTER_V2_ADDRESS_MAP,
} from '../../constants/addresses'
import { ChainProvider } from '../../router'
import { SynapseIntentPreviewer as PreviewerContract } from '../../typechain/SynapseIntentPreviewer'
import { isSameAddress } from '../../utils/addressUtils'
import { logger, logExecutionTime } from '../../utils/logger'
import { EngineID, toWei, SlippageMax, getForwardTo } from '../core'
import {
  RouteInput,
  SwapEngine,
  SwapEngineRoute,
  getEmptyRoute,
} from '../models'

export class DefaultEngine implements SwapEngine {
  static previewerInterface = new Interface(previewerAbi)

  public readonly id = EngineID.Default

  private contracts: {
    [chainId: number]: {
      previewer: PreviewerContract
      swapQuoter: string
    }
  }

  constructor(chains: ChainProvider[]) {
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

  @logExecutionTime('DefaultEngine.getQuote')
  public async getQuote(input: RouteInput): Promise<SwapEngineRoute> {
    // TODO: timeout
    const { chainId, tokenIn, tokenOut, amountIn, finalRecipient } = input
    const { previewer, swapQuoter } = this.contracts[chainId]
    if (
      !previewer ||
      !swapQuoter ||
      isSameAddress(tokenIn, tokenOut) ||
      BigNumber.from(amountIn).eq(Zero)
    ) {
      return getEmptyRoute(this.id)
    }
    // Get the quote
    const forwardTo = getForwardTo(finalRecipient)
    // Note: restrictComplexity is not supported by the on-chain previewer
    const { amountOut, steps: stepsOutput } = await previewer.previewIntent(
      swapQuoter,
      forwardTo,
      // slippage settings are applied when generating the zap data as minFinalAmount
      toWei(SlippageMax),
      tokenIn,
      tokenOut,
      amountIn
    )
    // Remove extra fields before the encoding
    return {
      engineID: this.id,
      engineName: EngineID[this.id],
      chainId,
      expectedAmountOut: amountOut,
      steps: stepsOutput.map(({ token, amount, msgValue, zapData }) => ({
        token,
        amount,
        msgValue,
        zapData,
      })),
    }
  }

  @logExecutionTime('DefaultEngine.generateRoute')
  public async generateRoute(
    _input: RouteInput,
    quote: SwapEngineRoute
  ): Promise<SwapEngineRoute> {
    if (quote.engineID !== this.id || !quote.steps) {
      logger.error({ quote }, 'DefaultEngine: unexpected quote')
      return getEmptyRoute(this.id)
    }
    return quote
  }
}
