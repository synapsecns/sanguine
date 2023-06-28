import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address, multicall, erc20ABI, getAccount } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { Token } from '../types'
import { sortByTokenBalance } from '../sortTokens'
import { BigNumber } from 'ethers'

export const getTokensByChainId = async (
  chainId: number
): Promise<TokenBalance[]> => {
  const { address } = getAccount()
  const tokens = BRIDGABLE_TOKENS[chainId]
  return await sortByTokenBalance(tokens, chainId, address)
}

interface TokenBalance {
  token: Token
  balance: BigNumber
}
interface NetworkTokenBalances {
  [index: number]: TokenBalance[]
}

export const usePortfolioBalances = () => {
  const [balances, setBalances] = useState<NetworkTokenBalances>({})
  const availableChains = Object.keys(BRIDGABLE_TOKENS)

  useEffect(() => {
    const fetchBalances = async () => {
      const balanceLibrary = {}
      availableChains.forEach(async (chainId) => {
        const currentChainId = Number(chainId)
        const tokenBalances: TokenBalance[] = await getTokensByChainId(
          currentChainId
        )
        balanceLibrary[currentChainId] = tokenBalances
      })
      setBalances(balanceLibrary)
    }
    fetchBalances()
  }, [])

  return balances
}

const useTokenApprovals = () => {}
