import { multicall, erc20ABI, Address } from '@wagmi/core'
import { sortByTokenBalance } from '../sortTokens'
import { Chain, Token } from '../types'
import { BRIDGABLE_TOKENS, POOLS_BY_CHAIN } from '@/constants/tokens'
import { FetchState } from '@/slices/portfolio/actions'

export const ROUTER_ADDRESS = '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a'
export const CCTP_ROUTER_ADDRESS = '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48'

export interface TokenAndBalance {
  token: Token
  tokenAddress: string
  balance: bigint
  parsedBalance: string
  queriedChain: Chain
}

export interface TokenAndAllowance {
  token: Token
  spender: string
  allowance: bigint
}
export interface TokenWithBalanceAndAllowance
  extends TokenAndBalance,
    TokenAndAllowance {}

export interface Allowances {
  [spender: string]: bigint
}

export interface TokenWithBalanceAndAllowances extends TokenAndBalance {
  allowances: Allowances
}

export interface NetworkTokenBalancesAndAllowances {
  [index: number]: TokenWithBalanceAndAllowances[]
}

export const getTokenBalances = async (
  owner: string,
  tokens: Token[],
  chainId: number
): Promise<TokenAndBalance[]> => {
  return await sortByTokenBalance(tokens, chainId, owner)
}

function mergeBalancesAndAllowances(
  balances: TokenAndBalance[],
  allowances: TokenAndAllowance[]
): TokenWithBalanceAndAllowances[] {
  return balances.map((balance: TokenAndBalance) => {
    const tokenAllowances = {}
    const matchedAllowancesByToken: TokenAndAllowance[] = allowances.filter(
      (allowance: TokenAndAllowance) => allowance.token === balance.token
    )

    matchedAllowancesByToken.forEach((spenderAllowance: TokenAndAllowance) => {
      const { spender, allowance } = spenderAllowance
      tokenAllowances[spender] = allowance
    })

    return {
      queriedChain: balance.queriedChain,
      token: balance.token,
      tokenAddress: balance.tokenAddress,
      balance: balance.balance,
      parsedBalance: balance.parsedBalance,
      allowances: tokenAllowances,
    }
  })
}

const getTokensAllowances = async (
  owner: string,
  spender: string,
  tokens: Token[],
  chainId: number
): Promise<any> => {
  const inputs = tokens.map((token: Token) => {
    const tokenAddress = token.addresses[chainId] as Address
    return {
      address: tokenAddress,
      abi: erc20ABI,
      functionName: 'allowance',
      chainId,
      args: [owner, spender],
    }
  })
  const allowancesResponse: {
    error?: any
    result?: any
    status: 'success' | 'failure'
  }[] = await multicall({
    contracts: inputs,
    chainId,
  })

  return tokens.map((token: Token, index: number) => {
    let allowance
    if (allowancesResponse[index].status === 'success') {
      allowance = allowancesResponse[index].result
    } else {
      allowance = null
    }
    return {
      token,
      spender,
      allowance,
    }
  })
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
  balancesAndAllowances: NetworkTokenBalancesAndAllowances
  poolTokenBalances: NetworkTokenBalancesAndAllowances
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
        // getTokensAllowances(
        //   address,
        //   ROUTER_ADDRESS,
        //   currentChainTokens,
        //   currentChainId
        // ),
      ])
      // const mergedBalancesAndAllowances = mergeBalancesAndAllowances(
      //   tokenBalances,
      //   tokenAllowances
      // )
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

export function separateTokensByAllowance(
  tokens: TokenWithBalanceAndAllowances[]
): [TokenWithBalanceAndAllowances[], TokenWithBalanceAndAllowances[]] {
  const tokensWithAllowance: TokenWithBalanceAndAllowances[] = []
  const tokensWithoutAllowance: TokenWithBalanceAndAllowances[] = []

  tokens &&
    tokens.forEach((token: TokenWithBalanceAndAllowances) => {
      // currently separating by bridge allowance
      // update this when incorporating other allowances to order by
      const bridgeAllowance: bigint | null = token.allowances[ROUTER_ADDRESS]
      if (bridgeAllowance === null) {
        tokensWithAllowance.push(token)
      } else if (bridgeAllowance > 0n) {
        tokensWithAllowance.push(token)
      } else {
        tokensWithoutAllowance.push(token)
      }
    })

  return [tokensWithAllowance, tokensWithoutAllowance]
}

export function sortTokensByBalanceDescending(
  tokens: TokenWithBalanceAndAllowances[]
): TokenWithBalanceAndAllowances[] {
  return (
    tokens &&
    tokens.sort(
      (a: TokenWithBalanceAndAllowances, b: TokenWithBalanceAndAllowances) =>
        Number(b.parsedBalance) > Number(a.parsedBalance) ? 1 : -1
    )
  )
}
