import { useState, useMemo } from 'react'
import { useNetwork } from 'wagmi'
import { multicall, erc20ABI, getAccount, Address } from '@wagmi/core'
import { sortByTokenBalance } from '../sortTokens'
import { Token } from '../types'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { FetchState } from '@/slices/portfolio/reducer'

export const ROUTER_ADDRESS = '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a'

export interface TokenAndBalance {
  token: Token
  balance: bigint
  parsedBalance: string
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
  allowances: Allowances[]
}

export interface NetworkTokenBalancesAndAllowances {
  [index: number]: TokenWithBalanceAndAllowance[]
}

export const getTokensByChainId = async (
  owner: string,
  tokens: Token[],
  chainId: number
): Promise<TokenAndBalance[]> => {
  return await sortByTokenBalance(tokens, chainId, owner)
}

function mergeBalancesAndAllowances(
  balances: { token: Token; balance: bigint; parsedBalance: string }[],
  allowances: { token: Token; spender: Address; allowance: bigint }[]
): TokenWithBalanceAndAllowance[] {
  return balances.map((balance) => {
    const correspondingAllowance = allowances.find(
      (item2) => item2.token === balance.token
    )
    if (correspondingAllowance) {
      return {
        token: balance.token,
        balance: balance.balance,
        parsedBalance: balance.parsedBalance,
        spender: correspondingAllowance.spender,
        allowance: correspondingAllowance.allowance,
      }
    }
    // if no allowance is matched with corresponding balance
    // e.g native gas tokens
    return {
      token: balance.token,
      balance: balance.balance,
      parsedBalance: balance.parsedBalance,
      spender: correspondingAllowance.spender,
      allowance: null,
    }
  })
}

const getTokensAllowance = async (
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
  status: FetchState
  error?: any | undefined
}> => {
  const balanceRecord = {}
  const availableChains: string[] = Object.keys(BRIDGABLE_TOKENS)
  const isSingleNetworkCall: boolean = typeof chainId === 'number'

  const filteredChains: string[] = availableChains.filter((chain: string) => {
    return isSingleNetworkCall ? Number(chain) === chainId : chain !== '2000' // need to figure out whats wrong with Dogechain
  })

  try {
    const balancePromises = filteredChains.map(async (chainId) => {
      const currentChainId = Number(chainId)
      const currentChainTokens = BRIDGABLE_TOKENS[chainId]
      const [tokenBalances, tokenAllowances] = await Promise.all([
        getTokensByChainId(address, currentChainTokens, currentChainId),
        getTokensAllowance(
          address,
          ROUTER_ADDRESS,
          currentChainTokens,
          currentChainId
        ),
      ])
      const mergedBalancesAndAllowances = mergeBalancesAndAllowances(
        tokenBalances,
        tokenAllowances
      )
      return { currentChainId, mergedBalancesAndAllowances }
    })
    const balances = await Promise.all(balancePromises)
    balances.forEach(({ currentChainId, mergedBalancesAndAllowances }) => {
      balanceRecord[currentChainId] = mergedBalancesAndAllowances
    })

    return { balancesAndAllowances: balanceRecord, status: FetchState.VALID }
  } catch (error) {
    console.error('error from fetch:', error)
    return { balancesAndAllowances: {}, status: FetchState.INVALID, error }
  }
}
