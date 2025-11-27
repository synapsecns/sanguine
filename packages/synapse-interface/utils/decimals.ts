import { Token } from '@/utils/types'
import { stringToBigInt } from '@/utils/bigint/format'

export const getTokenDecimals = (
  token?: Pick<Token, 'decimals'>,
  chainId?: number
): number | undefined => {
  if (!token || !chainId) return undefined

  return typeof token.decimals === 'number'
    ? token.decimals
    : token.decimals?.[chainId]
}

export const parseTokenAmount = (
  amount?: string,
  token?: Pick<Token, 'decimals'>,
  chainId?: number
): bigint | undefined => {
  const decimals = getTokenDecimals(token, chainId)

  if (!decimals || !amount) return undefined

  return stringToBigInt(amount, decimals)
}
