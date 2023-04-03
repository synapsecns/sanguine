import { fetchBalance } from '@wagmi/core'
import { Token } from '@/utils/types'
import { Zero, AddressZero } from '@ethersproject/constants'

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
  let tokensWithBalances: any[] = []
  let zeroTokensWithBalances: any[] = []
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
    console.log('tokenAddr', tokenAddr, AddressZero, tokenAddr === AddressZero)
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
