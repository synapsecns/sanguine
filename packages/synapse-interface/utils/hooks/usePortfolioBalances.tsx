import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address, multicall, erc20ABI, getAccount } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { Token } from '../types'
import { sortByTokenBalance, TokenAndBalance } from '../sortTokens'
import { BigNumber } from 'ethers'

interface NetworkTokenBalances {
  [index: number]: TokenAndBalance[]
}

export const getTokensByChainId = async (
  chainId: number
): Promise<TokenAndBalance[]> => {
  const { address } = getAccount()
  const tokens = BRIDGABLE_TOKENS[chainId]
  return await sortByTokenBalance(tokens, chainId, address)
}

export const usePortfolioBalances = () => {
  const [balances, setBalances] = useState<NetworkTokenBalances>({})
  const availableChains = Object.keys(BRIDGABLE_TOKENS)

  useEffect(() => {
    const fetchBalancesAcrossNetworks = async () => {
      const balanceRecord = {}
      availableChains.forEach(async (chainId) => {
        const currentChainId = Number(chainId)
        const tokenBalances: TokenAndBalance[] = await getTokensByChainId(
          currentChainId
        )
        balanceRecord[currentChainId] = tokenBalances
      })
      setBalances(balanceRecord)
    }
    fetchBalancesAcrossNetworks()
  }, [])

  return balances
}
