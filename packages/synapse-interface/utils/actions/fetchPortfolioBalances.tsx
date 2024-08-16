import { sortByTokenBalance } from '../sortTokens'
import { Chain, Token } from '../types'
import {
  BRIDGABLE_TOKENS,
  POOLS_BY_CHAIN,
  NON_BRIDGEABLE_GAS_TOKENS,
} from '@/constants/tokens'
import { FetchState } from '@/slices/portfolio/actions'

export interface TokenAndBalance {
  token: Token
  tokenAddress: string
  balance: bigint
  parsedBalance: string
  queriedChain: Chain
}

export interface NetworkTokenBalances {
  [index: number]: TokenAndBalance[]
}

export const getTokenBalances = async (
  owner: string,
  tokens: Token[],
  chainId: number
): Promise<TokenAndBalance[]> => {
  return await sortByTokenBalance(tokens, chainId, owner)
}

/**
 * @param address: wallet address to fetch balance of
 * @param chainId?: specific network to fetch balances from
 * @returns addresses' token balances and allowances
 * If specifying chainId parameter, function will only fetch from single network
 * If chainId is undefined, function will fetch from all supported networks
 */
export const fetchPortfolioBalances = async (
  address: string,
  chainId?: number | undefined | null
): Promise<{
  balances: NetworkTokenBalances
  poolTokenBalances: NetworkTokenBalances
  status: FetchState
  error?: any | undefined
}> => {
  const balanceRecord = {}
  const poolTokenBalances = {}

  const availableChains: string[] = Object.keys(BRIDGABLE_TOKENS)
  const isSingleNetworkCall: boolean = typeof chainId === 'number'

  const filteredChains: string[] = availableChains.filter((chain: string) => {
    return !isSingleNetworkCall || Number(chain) === chainId
  })

  try {
    const balancePromises = filteredChains.map(async (chainId) => {
      const currentChainId = Number(chainId)

      return (async () => {
        try {
          let currentChainTokens

          currentChainTokens = BRIDGABLE_TOKENS[chainId]

          if (POOLS_BY_CHAIN[chainId]) {
            currentChainTokens = currentChainTokens.concat(
              POOLS_BY_CHAIN[chainId]
            )
          }

          // Reconstruct shape of Token to batch fetching balances
          if (NON_BRIDGEABLE_GAS_TOKENS[chainId]) {
            const currentChainGasTokens = NON_BRIDGEABLE_GAS_TOKENS[
              chainId
            ].map((gasToken) => [
              {
                ...gasToken,
                chainId: currentChainId,
                addresses: { [currentChainId]: gasToken.address },
              },
            ])
            currentChainTokens = currentChainTokens.concat(
              ...currentChainGasTokens
            )
          }
          console.log('currentChainTokens', currentChainId, currentChainTokens)
          const [tokenBalances] = await Promise.all([
            getTokenBalances(address, currentChainTokens, currentChainId),
          ])
          console.log('tokenBalances', currentChainId, tokenBalances)
          return { currentChainId, tokenBalances }
        } catch (error) {
          console.error(
            `Error fetching balances for chainId ${chainId}:`,
            error
          )
          return null
        }
      })()
    })

    const balances = await Promise.all(balancePromises)
    balances.forEach(({ currentChainId, tokenBalances }) => {
      balanceRecord[currentChainId] = tokenBalances.filter(
        (entry) => !entry.token.poolName
      )
      poolTokenBalances[currentChainId] = tokenBalances.filter(
        (entry) => typeof entry.token.poolName === 'string'
      )
    })

    return {
      balances: balanceRecord,
      status: FetchState.VALID,
      poolTokenBalances,
    }
  } catch (error) {
    console.error('error from fetch:', error)
    return {
      balances: {},
      status: FetchState.INVALID,
      error,
      poolTokenBalances: {},
    }
  }
}

export function sortTokensByBalanceDescending(
  tokens: TokenAndBalance[]
): TokenAndBalance[] {
  return (
    tokens &&
    tokens.sort((a: TokenAndBalance, b: TokenAndBalance) =>
      Number(b.parsedBalance) > Number(a.parsedBalance) ? 1 : -1
    )
  )
}
