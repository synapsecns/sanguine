import { BigNumber } from 'ethers'
import { multicall, Address } from '@wagmi/core'
import { Zero, AddressZero } from '@ethersproject/constants'

import multicallABI from '../constants/abis/multicall.json'
import erc20ABI from '../constants/abis/erc20.json'
import { Token } from '@/utils/types'

export const sortByVisibilityRank = (tokens: Token[]) => {
  if (tokens === undefined) {
    return []
  }

  return Object.values(tokens).sort(
    (a, b) => b.visibilityRank - a.visibilityRank
  )
}

const sortArrayByBalance = (array) => {
  return array.sort((a, b) => {
    const balanceA = BigInt(a.balance || '')
    const balanceB = BigInt(b.balance || '')

    if (balanceA < balanceB) {
      return 1
    } else if (balanceA > balanceB) {
      return -1
    } else {
      return 0
    }
  })
}

export const sortByTokenBalance = async (
  tokens: Token[],
  chainId: number,
  address: any
) => {
  const tokensWithBalances: any[] = []
  const multicallInputs = []
  let multicallData

  if (chainId === undefined || !address) {
    tokens.map((token) => {
      tokensWithBalances.push({
        token,
        balance: Zero,
      })
    })
  } else {
    tokens.map((token) => {
      const tokenAddress = token.addresses[chainId as keyof Token['addresses']]
      const tokenAbi = erc20ABI
      // deterministic multicall3 address on all eth chains
      const multicallAddress: Address = `0xcA11bde05977b3631167028862bE2a173976CA11`

      if (tokenAddress === undefined) {
        return
      }

      if (tokenAddress === AddressZero || tokenAddress === '') {
        multicallInputs.push({
          address: multicallAddress,
          abi: multicallABI,
          functionName: 'getEthBalance',
          chainId,
          args: [address],
        })
      } else {
        const formattedTokenAddress: Address = `0x${tokenAddress.slice(2)}`
        multicallInputs.push({
          address: formattedTokenAddress,
          abi: tokenAbi,
          functionName: 'balanceOf',
          chainId,
          args: [address],
        })
      }
    })
  }

  if (multicallInputs.length > 0) {
    multicallData = await multicall({
      contracts: multicallInputs,
    })
    return sortArrayByBalance(
      multicallData.map((tokenBalance: BigNumber | undefined, index) => ({
        token: tokens[index],
        balance: tokenBalance,
      }))
    )
  }

  return tokensWithBalances
}
