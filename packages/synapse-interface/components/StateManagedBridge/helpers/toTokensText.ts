import { CHAINS_BY_ID } from '@/constants/chains'
import { Token } from '@/utils/types'

interface Fields {
  fromChainId?: number
  fromToken?: Token
  toChainId?: number
  toToken?: Token
}

export const toTokenText = ({
  fromChainId,
  fromToken,
  toChainId,
  toToken,
}: Fields) => {
  const fromChainName = CHAINS_BY_ID[fromChainId]?.name
  const toChainName = CHAINS_BY_ID[toChainId]?.name

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId === null &&
    toToken === null
  ) {
    return 'All tokens you can receive'
  }

  if (
    fromChainId &&
    fromToken === null &&
    toChainId === null &&
    toToken === null
  ) {
    return `Tokens you can receive from ${fromChainName}`
  }

  if (
    fromChainId === null &&
    fromToken &&
    toChainId === null &&
    toToken === null
  ) {
    return `Tokens you can receive by ${fromToken.symbol}`
  }

  if (fromChainId && fromToken && toChainId === null && toToken === null) {
    return `Tokens you can receive by ${fromToken.symbol} on ${fromChainName}`
  }

  if (fromChainId && fromToken && toChainId === null && toToken) {
    return `Tokens you can receive by ${fromToken.symbol} on ${fromChainName}`
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId &&
    toToken === null
  ) {
    return `Tokens you can receive on ${toChainName}`
  }

  if (fromChainId && fromToken === null && toChainId && toToken === null) {
    return `Tokens you can receive on ${toChainName} from ${fromChainName}`
  }

  if (fromChainId === null && fromToken && toChainId && toToken === null) {
    return `Tokens you can receive on ${toChainName} by ${fromToken.symbol}`
  }

  if (fromChainId && fromToken && toChainId && toToken === null) {
    return `Tokens you can receive on ${toChainName} by ${fromToken.symbol} on ${fromChainName}`
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId === null &&
    toToken
  ) {
    return `All tokens you can receive`
  }

  if (fromChainId && fromToken === null && toChainId === null && toToken) {
    return `Tokens you can receive from ${fromChainName}`
  }

  if (fromChainId === null && fromToken && toChainId === null && toToken) {
    return `Tokens you can receive by ${fromToken.symbol}`
  }

  if (fromChainId === null && fromToken === null && toChainId && toToken) {
    return `Tokens you can receive on ${toChainName}`
  }

  if (fromChainId && fromToken === null && toChainId && toToken) {
    return `Tokens you can receive on ${toChainName} from ${fromChainName}`
  }

  if (fromChainId === null && fromToken && toChainId && toToken) {
    return `Tokens you can receive on ${toChainName} by ${fromToken.symbol}`
  }

  if (fromChainId && fromToken && toChainId && toToken) {
    return `Tokens you can receive on ${toChainName} by ${fromToken.symbol} on ${fromChainName}`
  }
}
