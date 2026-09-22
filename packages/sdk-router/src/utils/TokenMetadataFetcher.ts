import { Provider } from '@ethersproject/abstract-provider'
import { Contract } from '@ethersproject/contracts'

import { isNativeToken } from './addressUtils'
import { HYPERCORE_CHAIN_ID, SYN_ADDRESS_MAP } from '../constants'
import { logger } from './logger'
import erc20ABI from '../abi/IERC20Metadata.json'
import { marshallChainToken } from '../rfq/ticker'
import { IERC20Metadata as ERC20 } from '../typechain/IERC20Metadata'

export class TokenMetadataFetcher {
  private providers: {
    [chainId: number]: Provider
  }
  private decimalsCache: {
    [tokenId: string]: number
  }

  constructor(chainIdToProvider: { [chainId: number]: Provider }) {
    this.providers = chainIdToProvider
    this.decimalsCache = {}
  }

  /**
   * Gets the number of decimals for a token.
   * Uses a cache to avoid repeated contract calls.
   * Returns 18 for native tokens.
   *
   * @param chainId - The chain ID of the token.
   * @param token - The token address.
   * @returns The number of decimals for the token.
   */
  public async getTokenDecimals(
    chainId: number,
    token: string
  ): Promise<number> {
    // HyperCore SYN uses its native token identifier and 8-decimal balance units.
    if (
      chainId === HYPERCORE_CHAIN_ID &&
      token.toLowerCase() === SYN_ADDRESS_MAP[HYPERCORE_CHAIN_ID].toLowerCase()
    ) {
      return 8
    }
    if (isNativeToken(token)) {
      return 18
    }

    const tokenId = marshallChainToken({ chainId, token })
    if (this.decimalsCache[tokenId]) {
      return this.decimalsCache[tokenId]
    }

    const provider = this.providers[chainId]
    if (!provider) {
      logger.error(`No provider found for chainId: ${chainId}`)
      return 0
    }

    const tokenContract = new Contract(token, erc20ABI, provider) as ERC20
    try {
      const decimals = await tokenContract.decimals()
      this.decimalsCache[tokenId] = decimals
      return decimals
    } catch (error) {
      logger.error({ error, chainId, token }, 'Error fetching token decimals')
      return 0
    }
  }
}
