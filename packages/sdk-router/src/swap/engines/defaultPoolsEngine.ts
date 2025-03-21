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
import { isSameAddress, logger } from '../../utils'
import { EngineID, toWei, SlippageMax, getForwardTo } from '../core'
import {
  RouteInput,
  SwapEngine,
  SwapEngineRoute,
  getEmptyRoute,
} from '../models'

export class DefaultPoolsEngine implements SwapEngine {
  static previewerInterface = new Interface(previewerAbi)

  public readonly id = EngineID.DefaultPools

  private contracts: {
    [chainId: number]: {
      previewer: PreviewerContract
      swapQuoter: string
    }
  }

  constructor(chains: ChainProvider[]) {
    invariant(DefaultPoolsEngine.previewerInterface, 'INTERFACE_UNDEFINED')
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
          DefaultPoolsEngine.previewerInterface,
          provider
        ) as PreviewerContract,
        swapQuoter: swapQuoterAddress,
      }
    })
  }

  public async getQuote(input: RouteInput): Promise<SwapEngineRoute> {
    // TODO: timeout
    const { chainId, fromToken, toToken, fromAmount, toRecipient } = input
    const { previewer, swapQuoter } = this.contracts[chainId]
    if (
      !previewer ||
      !swapQuoter ||
      isSameAddress(fromToken, toToken) ||
      BigNumber.from(fromAmount).eq(Zero)
    ) {
      return getEmptyRoute(this.id)
    }
    // Get the quote
    const forwardTo = getForwardTo(toRecipient)
    // Note: restrictComplexity is not supported by the on-chain previewer
    const { amountOut, steps: stepsOutput } = await previewer.previewIntent(
      swapQuoter,
      forwardTo,
      // slippage settings are applied when generating the zap data as minFinalAmount
      toWei(SlippageMax),
      fromToken,
      toToken,
      fromAmount
    )
    // Remove extra fields before the encoding
    return {
      engineID: this.id,
      engineName: EngineID[this.id],
      chainId,
      fromToken,
      toToken,
      fromAmount: BigNumber.from(fromAmount),
      expectedToAmount: amountOut,
      steps: stepsOutput.map(({ token, amount, msgValue, zapData }) => ({
        token,
        amount,
        msgValue,
        zapData,
      })),
    }
  }

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
