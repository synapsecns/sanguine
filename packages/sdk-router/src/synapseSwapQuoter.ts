import { Interface } from '@ethersproject/abi'
import { Contract } from '@ethersproject/contracts'
import { Provider } from '@ethersproject/abstract-provider'
import invariant from 'tiny-invariant'

import abi from './abi/SynapseSwapQuoter.json'
import { BigintIsh } from './constants'
import {
  SynapseSwapQuoter as SynapseSwapQuoterContract,
  LimitedTokenStruct,
} from './typechain/SynapseSwapQuoter'
export type LimitedTokenStructType = LimitedTokenStruct
export class SynapseSwapQuoter {
  public static INTERFACE: Interface = new Interface(abi)
  public readonly chainId: number
  public readonly provider: Provider
  public readonly swapQuoterContract: SynapseSwapQuoterContract

  constructor(chainId: number, provider: Provider, swapQuoterAddress: string) {
    invariant(chainId !== undefined, 'CHAIN_ID_UNDEFINED')
    invariant(provider !== undefined, 'PROVIDER_UNDEFINED')
    invariant(swapQuoterAddress !== undefined, 'SWAP_QUOTER_ADDRESS_UNDEFINED')

    this.chainId = chainId
    this.provider = provider
    this.swapQuoterContract = new Contract(
      swapQuoterAddress,
      SynapseSwapQuoter.INTERFACE,
      provider
    ) as SynapseSwapQuoterContract
  }

  public async getAmountOut(
    tokenIn: LimitedTokenStruct,
    tokenOut: string,
    amountIn: BigintIsh
  ): Promise<any> {
    return this.swapQuoterContract.getAmountOut(tokenIn, tokenOut, amountIn)
  }
}
