import { CHAINS_BY_ID } from '@/constants/chains'
import { Token } from '@/utils/types'

interface Fields {
  fromChainId?: number
  fromToken?: Token
  toChainId?: number
  toToken?: Token
}

export const fromTokenText = ({
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
    return 'All bridgeable tokens'
  }

  if (
    fromChainId &&
    fromToken === null &&
    toChainId === null &&
    toToken === null
  ) {
    return `All ${fromChainName} tokens you can send`
  }

  if (
    fromChainId === null &&
    fromToken &&
    toChainId === null &&
    toToken === null
  ) {
    return 'All bridgeable tokens'
  }

  if (fromChainId && fromToken && toChainId === null && toToken === null) {
    return `${fromChainName} tokens you can send`
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId &&
    toToken === null
  ) {
    return `All tokens you can send to ${toChainName}`
  }

  if (fromChainId && fromToken === null && toChainId && toToken === null) {
    return `${fromChainName} tokens you can send to ${toChainName}`
  }

  if (fromChainId === null && fromToken && toChainId && toToken === null) {
    return `All tokens you can send to ${toChainName}`
  }

  if (fromChainId && fromToken && toChainId && toToken === null) {
    return `${fromChainName} tokens you can send to ${toChainName}`
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId === null &&
    toToken
  ) {
    return `All tokens you can send and receive ${toToken.symbol}`
  }

  if (fromChainId && fromToken === null && toChainId === null && toToken) {
    return `${fromChainName} tokens you can send and receive ${toToken.symbol}`
  }

  if (fromChainId === null && fromToken && toChainId === null && toToken) {
    return `Tokens you send and receive ${toToken.symbol}`
  }

  if (fromChainId === null && fromToken === null && toChainId && toToken) {
    return `Tokens you can send and receive ${toToken.symbol} on ${toChainName}`
  }

  if (fromChainId && fromToken === null && toChainId && toToken) {
    return `${fromChainName} tokens you can send and receive ${toToken.symbol} on ${toChainName}`
  }

  if (fromChainId === null && fromToken && toChainId && toToken) {
    return `Tokens you can send and receive ${toToken.symbol} on ${toChainName}`
  }

  if (fromChainId && fromToken && toChainId && toToken) {
    return `${fromChainName} tokens you can send and receive ${toToken.symbol} on ${toChainName}`
  }
}
