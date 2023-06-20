import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { Address, multicall, erc20ABI } from '@wagmi/core'
import { BRIDGABLE_TOKENS } from '@/constants/tokens'
import { Token } from '../types'
import { AddressZero } from '@ethersproject/constants'
import multicallABI from '@/constants/abis/multicall.json'

//move to constants file later
const MULTICALL3_ADDRESS: Address = '0xcA11bde05977b3631167028862bE2a173976CA11'

interface NetworkSpecificToken {
  token: Token
  symbol: string
  queryAddress: Address
  chainId: number
}

const getQueryableTokensByChain = (): NetworkSpecificToken[] => {
  const tokens = []
  const queryableNetworks = Object.keys(BRIDGABLE_TOKENS)

  queryableNetworks.forEach((chainId: string) => {
    BRIDGABLE_TOKENS[chainId].forEach((token: Token) => {
      const transformedToken = {
        token: token,
        symbol: token.symbol,
        queryAddress:
          token.addresses[Number(chainId) as keyof Token['addresses']],
        chainId: Number(chainId),
      } as NetworkSpecificToken
      tokens.push(transformedToken)
    })
  })

  return tokens
}

export const usePortfolioBalances = () => {}

const useTokenBalances = (
  address: Address,
  tokens: Token[],
  chainId: number
) => {
  const [balances, setBalances] = useState([])

  let calls = []

  useEffect(() => {
    if (!address || chainId === undefined) return
    if (tokens.length === 0) return
    ;(async () => {
      tokens.forEach((token: Token) => {
        const tokenAddress =
          token.addresses[chainId as keyof Token['addresses']]

        switch (tokenAddress) {
          case undefined:
            break
          case AddressZero || '':
            calls.push({
              address: MULTICALL3_ADDRESS,
              abi: multicallABI,
              functionName: 'getEthBalance',
              chainId,
              args: [address],
            })
            break
          default:
            calls.push({
              address: tokenAddress,
              abi: multicallABI,
              functionName: 'balanceOf',
              chainId,
              args: [address],
            })
        }
      })

      const multicallData = await multicall({ contracts: calls })
    })()
  }, [tokens])
}

const useTokenApprovals = () => {}
