import invariant from 'tiny-invariant'

import { Currency } from './currency'
import { NativeCurrency } from './nativeCurrency'
import { Token } from './token'
import { WETH9 } from './weth9'

/**
 * Ether is the main usage of a 'native' currency, i.e. for Ethereum mainnet and all testnets
 */
export class GasToken extends NativeCurrency {
  public constructor(
    chainId: number,
    decimals: number,
    symbol: string,
    name: string
  ) {
    super(chainId, decimals, symbol, name)
  }

  public get wrapped(): Token {
    const weth9 = WETH9[this.chainId]
    invariant(!!weth9, 'WRAPPED')
    return weth9
  }

  // private static _etherCache: { [chainId: number]: GasToken } = {}

  // public static onChain(chainId: number): GasToken {
  //   return this._etherCache[chainId] ?? (this._etherCache[chainId] = new GasToken(chainId))
  // }

  public equals(other: Currency): boolean {
    return other.isNative && other.chainId === this.chainId
  }
}
