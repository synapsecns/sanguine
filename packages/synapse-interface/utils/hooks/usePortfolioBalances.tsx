import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address, multicall, erc20ABI, getAccount } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { Token } from '../types'
import { AddressZero } from '@ethersproject/constants'
import multicallABI from '@/constants/abis/multicall.json'
import { getSortedBridgableTokens } from '../actions/getSortedBridgableTokens'
import { ChainId } from '@/constants/chains'
import { sortByTokenBalance } from '../sortTokens'
import { fetchBalance } from '@wagmi/core'
import { BigNumber } from 'ethers'

//move to constants file later
const MULTICALL3_ADDRESS: Address = '0xcA11bde05977b3631167028862bE2a173976CA11'

export const getTokensByChainId = async (chainId: number) => {
  const { address } = getAccount()

  const tokens = BRIDGABLE_TOKENS[chainId]

  return await sortByTokenBalance(tokens, chainId, address)
}

interface TokenBalance {
  token: Token
  balance: BigNumber
}
interface NetworkTokenBalances {
  [index: number]: Array<TokenBalance>
}

export const usePortfolioBalances = () => {
  const [balances, setBalances] = useState<NetworkTokenBalances>({})
  const { address } = getAccount()
  const availableChains = Object.keys(BRIDGABLE_TOKENS)

  useEffect(() => {
    const getData = async () => {
      const balanceLibrary = {}
      availableChains.forEach(async (chainId) => {
        const currentChainId = Number(chainId)
        const tokenBalances: Array<TokenBalance> = await getTokensByChainId(
          currentChainId
        )
        balanceLibrary[currentChainId] = tokenBalances
      })

      setBalances(balanceLibrary)
    }
    getData()
  }, [])

  return balances
}

const useTokenApprovals = () => {}
