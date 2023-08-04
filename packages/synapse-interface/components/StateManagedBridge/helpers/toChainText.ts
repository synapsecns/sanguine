import { CHAINS_BY_ID } from '@/constants/chains'
import { Token } from '@/utils/types'

interface Fields {
  fromChainId?: number
  fromToken?: Token
  toChainId?: number
  toToken?: Token
}

export const toChainText = ({
  fromChainId,
  fromToken,
  toChainId,
  toToken,
}: Fields) => {
  const fromChainName = CHAINS_BY_ID[fromChainId]?.name

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId === null &&
    toToken === null
  ) {
    return `Chains you can receive at`
  }

  if (
    fromChainId &&
    fromToken === null &&
    toChainId === null &&
    toToken === null
  ) {
    return `Chains you can receive at from ${fromChainName}`
  }

  if (
    fromChainId === null &&
    fromToken &&
    toChainId === null &&
    toToken === null
  ) {
    return `Chains you can receive ${fromToken.symbol}`
  }

  if (fromChainId && fromToken && toChainId === null && toToken === null) {
    return `Chains you can receive ${fromToken.symbol} on ${fromChainName}`
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId &&
    toToken === null
  ) {
    return `Chains you can receive at`
  }

  if (fromChainId && fromToken === null && toChainId && toToken === null) {
    return `Chains you can receive from ${fromChainName}`
  }

  if (fromChainId === null && fromToken && toChainId && toToken === null) {
    return `Chains you can receive ${fromToken.symbol} at`
  }

  if (fromChainId && fromToken && toChainId && toToken === null) {
    return `Chains you can receive ${fromToken.symbol} on ${fromChainName}`
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId === null &&
    toToken
  ) {
    return `Chains you can receive ${toToken.symbol}`
  }

  if (fromChainId && fromToken === null && toChainId === null && toToken) {
    return `Chains you can receive ${toToken.symbol} from ${fromChainName}`
  }

  if (fromChainId === null && fromToken && toChainId === null && toToken) {
    return `Chains you can receive ${toToken.symbol} by ${fromToken.symbol}`
  }

  if (fromChainId === null && fromToken === null && toChainId && toToken) {
    return `Chains you can receive ${toToken.symbol}`
  }

  if (fromChainId && fromToken === null && toChainId && toToken) {
    return `Chains you can receive ${toToken.symbol} from ${fromChainName}`
  }

  if (fromChainId === null && fromToken && toChainId && toToken) {
    return `Chains you can receive ${toToken.symbol} from ${fromToken.symbol}`
  }

  if (fromChainId && fromToken && toChainId && toToken) {
    return `Chains you can receive ${toToken.symbol} by ${fromToken.symbol} on ${fromChainName}`
  }
}
