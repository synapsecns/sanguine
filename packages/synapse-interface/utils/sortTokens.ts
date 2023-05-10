import { fetchBalance, erc20ABI } from '@wagmi/core'
import { Zero, AddressZero } from '@ethersproject/constants'

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

export const _sortTokenByBalance = async (
  tokens: Token[],
  chainId: number,
  address: any
) => {
  const i = 0
  const tokensWithBalances: any[] = []
  const zeroTokensWithBalances: any[] = []

  const multicallInputs = []

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
      const abi = erc20ABI
      const functionName = 'balanceOf'

      multicallInputs.push({
        address: tokenAddress,
        abi,
        functionName,
      })
    })
  }
}
