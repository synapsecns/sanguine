import { BigNumber } from 'ethers'
import { fetchBalance, multicall } from '@wagmi/core'
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
// move to utils
export const sortByTokenBalance = async (
  tokens: Token[],
  chainId: number,
  address: any
) => {
  let i = 0
  const tokensWithBalances: any[] = []
  const zeroTokensWithBalances: any[] = []
  // go through all tokens and retrieve token balances
  while (i < tokens.length) {
    if (chainId === undefined || address === undefined) {
      tokensWithBalances.push({
        token: tokens[i],
        balance: Zero,
      })
      i++
      continue
    }
    const tokenAddr = tokens[i].addresses[chainId as keyof Token['addresses']]

    let rawTokenBalance: any
    // Check for native token
    if (tokenAddr === '' || tokenAddr === AddressZero) {
      const data = await fetchBalance({
        address,
        chainId,
      })
      rawTokenBalance = data
    } else if (tokenAddr?.length > 0) {
      const data = await fetchBalance({
        address,
        token: `0x${tokenAddr.slice(2)}`,
        chainId,
      })
      rawTokenBalance = data
    }

    // manages two the array of tokens with zero balances and non-zero balances
    if (rawTokenBalance) {
      if (rawTokenBalance?.value._hex !== '0x00') {
        zeroTokensWithBalances.push({
          token: tokens[i],
          balance: rawTokenBalance.value,
        })
      } else {
        tokensWithBalances.push({
          token: tokens[i],
          balance: rawTokenBalance.value,
        })
      }
    }
    i++
  }

  return zeroTokensWithBalances.concat(tokensWithBalances)
}

export const _sortByTokenBalance = async (
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
      const multicallAddress = '0xcA11bde05977b3631167028862bE2a173976CA11'

      if (tokenAddress === AddressZero) {
        multicallInputs.push({
          address: multicallAddress,
          abi: multicallABI,
          functionName: 'getEthBalance',
          chainId,
          args: [address],
        })
      } else {
        multicallInputs.push({
          address: tokenAddress,
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

    return multicallData.map((tokenBalance: BigNumber | undefined, index) => ({
      token: tokens[index],
      balance: tokenBalance,
    }))
  }

  return tokensWithBalances
}
