import { sortByTokenBalance } from '../sortTokens'
import { Chain, Token } from '../types'
import { BRIDGABLE_TOKENS, POOLS_BY_CHAIN } from '@/constants/tokens'
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
  balancesAndAllowances: NetworkTokenBalances
  poolTokenBalances: NetworkTokenBalances
  status: FetchState
  error?: any | undefined
}> => {
  const balanceRecord = {}
  const poolTokenBalances = {}
  const availableChains: string[] = Object.keys(BRIDGABLE_TOKENS)
  const isSingleNetworkCall: boolean = typeof chainId === 'number'

  const filteredChains: string[] = availableChains.filter((chain: string) => {
    return isSingleNetworkCall ? Number(chain) === chainId : chain !== '2000' // need to figure out whats wrong with Dogechain
  })

  try {
    const balancePromises = filteredChains.map(async (chainId) => {
      let currentChainTokens
      const currentChainId = Number(chainId)

      if (POOLS_BY_CHAIN[chainId]) {
        currentChainTokens = BRIDGABLE_TOKENS[chainId].concat(
          POOLS_BY_CHAIN[chainId]
        )
      } else {
        currentChainTokens = BRIDGABLE_TOKENS[chainId]
      }

      const [tokenBalances] = await Promise.all([
        getTokenBalances(address, currentChainTokens, currentChainId),
      ])
      return { currentChainId, tokenBalances }
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
      balancesAndAllowances: balanceRecord,
      status: FetchState.VALID,
      poolTokenBalances,
    }
  } catch (error) {
    console.error('error from fetch:', error)
    return {
      balancesAndAllowances: {},
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
