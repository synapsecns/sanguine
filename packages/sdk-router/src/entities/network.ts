import invariant from 'tiny-invariant'

import { NativeCurrency } from './nativeCurrency'
import { GasToken } from './gas'
export class Network {
  /**
   * Returns the chain ID of the network
   */
  public readonly chainId: number

  /**
   * Returns the name of a network, i.e. a descriptive textual non-unique identifier
   */
  public readonly name?: string

  /**
   * Returns the URL to an icon/logo of the network
   */
  public readonly icon?: string

  /**
   * Returns the native gas token of the network, as the NativeCurrency type
   */
  public readonly currency?: GasToken

  /**
   * Returns an array of RPCs endpoints which could be used to connect to the network
   */
  public readonly rpcs?: string[]

  /**
   * Returns an array of explorers which could display the state of the network
   */
  public readonly explorers?: string[]

  public constructor(
    chainId: number,
    name?: string,
    icon?: string,
    currency?: NativeCurrency,
    rpcs?: string[],
    explorers?: string[]
  ) {
    invariant(Number.isSafeInteger(chainId), 'CHAIN_ID')
    this.chainId = chainId
    this.name = name
    this.icon = icon
    this.currency = currency
    this.rpcs = rpcs
    this.explorers = explorers
  }
}

export type NetworkIsh = Network | string | number
