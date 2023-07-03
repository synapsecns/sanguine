import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address, multicall, erc20ABI, getAccount } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { Token } from '../types'
import { sortByTokenBalance } from '../sortTokens'
import { BigNumber } from 'ethers'
import { AddressZero } from '@ethersproject/constants'

const ROUTER_ADDRESS = '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a'

export interface TokenAndBalance {
  token: Token
  balance: BigNumber
  parsedBalance: string
}
export interface TokenAndAllowance {
  token: Token
  allowance: BigNumber
}

export interface TokenWithBalanceAndAllowance
  extends TokenAndBalance,
    TokenAndAllowance {}

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
  balances: { token: Token; balance: BigNumber; parsedBalance: string }[],
  allowances: { token: Token; allowance: BigNumber }[]
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
        allowance: correspondingAllowance.allowance,
      }
    }
    // if no allowance is matched with corresponding balance
    // e.g native gas tokens
    return {
      token: balance.token,
      balance: balance.balance,
      parsedBalance: balance.parsedBalance,
      allowance: null,
    }
  })
}

export const usePortfolioBalancesAndAllowances =
  (): NetworkTokenBalancesAndAllowances => {
    const [balancesAndAllowances, setBalancesAndAllowances] =
      useState<NetworkTokenBalancesAndAllowances>({})
    const { address } = getAccount()
    const availableChains = Object.keys(BRIDGABLE_TOKENS)

    useEffect(() => {
      if (!address) return
      const fetchBalancesAcrossNetworks = async () => {
        const balanceRecord = {}
        availableChains.forEach(async (chainId) => {
          const currentChainId = Number(chainId)
          const currentChainTokens = BRIDGABLE_TOKENS[chainId]
          const tokenBalances: TokenAndBalance[] = await getTokensByChainId(
            address,
            currentChainTokens,
            currentChainId
          )
          const tokenAllowances: TokenAndAllowance[] = await getTokensAllowance(
            address,
            ROUTER_ADDRESS,
            currentChainTokens,
            currentChainId
          )
          const mergedBalancesAndAllowances: TokenWithBalanceAndAllowance[] =
            mergeBalancesAndAllowances(tokenBalances, tokenAllowances)
          balanceRecord[currentChainId] = mergedBalancesAndAllowances
        })
        setBalancesAndAllowances(balanceRecord)
      }
      fetchBalancesAcrossNetworks()
    }, [])

    return balancesAndAllowances
  }

const getTokensAllowance = async (
  owner: string,
  spender: string,
  tokens: Token[],
  chainId: number
): Promise<any> => {
  const inputs = tokens.map((token: Token) => {
    const tokenAddress = token.addresses[
      chainId as keyof Token['addresses']
    ] as `0x${string}`
    return {
      address: tokenAddress,
      abi: erc20ABI,
      functionName: 'allowance',
      chainId,
      args: [owner, spender],
    }
  })
  const allowances: unknown[] = await multicall({
    contracts: inputs,
    chainId,
  })

  return tokens.map((token: Token, index: number) => {
    return {
      token,
      allowance: allowances[index],
    }
  })
}
