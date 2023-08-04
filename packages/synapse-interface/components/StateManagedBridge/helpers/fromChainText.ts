import { CHAINS_BY_ID } from '@/constants/chains'
import { Token } from '@/utils/types'

interface Fields {
  fromChainId?: number
  fromToken?: Token
  toChainId?: number
  toToken?: Token
}

export const fromChainText = ({
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
    return 'All chains you can bridge from'
  }

  if (
    fromChainId &&
    fromToken === null &&
    toChainId === null &&
    toToken === null
  ) {
    return 'All chains you can bridge from'
  }

  if (
    fromChainId === null &&
    fromToken &&
    toChainId === null &&
    toToken === null
  ) {
    return `Chains you can bridge ${fromToken.symbol} from`
  }

  if (fromChainId && fromToken && toChainId === null && toToken === null) {
    return `Chains you can bridge ${fromToken.symbol} from`
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId &&
    toToken === null
  ) {
    return `Chains you can bridge to ${toChainName}`
  }

  if (fromChainId && fromToken === null && toChainId && toToken === null) {
    return `Chains you can bridge to ${toChainName}`
  }

  if (fromChainId === null && fromToken && toChainId && toToken === null) {
    return `Chains you can bridge ${fromToken.symbol} to ${toChainName}`
  }

  if (fromChainId && fromToken && toChainId && toToken === null) {
    return `Chains you can bridge ${fromToken.symbol} on ${fromChainName} to ${toChainName}`
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId === null &&
    toToken
  ) {
    return `Chains you can bridge from for ${toToken.symbol}`
  }

  if (fromChainId && fromToken === null && toChainId === null && toToken) {
    return `Chains you can bridge from for ${toToken.symbol}`
  }

  if (fromChainId === null && fromToken && toChainId === null && toToken) {
    return `Chains you can bridge ${fromToken.symbol} to ${toToken.symbol}`
  }

  if (fromChainId === null && fromToken === null && toChainId && toToken) {
    return `Chains you can bridge to ${toChainName} for ${toToken.symbol}`
  }

  if (fromChainId && fromToken === null && toChainId && toToken) {
    return `Chains you can bridge to ${toChainName} for ${toToken.symbol}`
  }

  if (fromChainId === null && fromToken && toChainId && toToken) {
    return `Chains you can bridge ${fromToken.symbol} to ${toToken.symbol} on ${toChainName}`
  }

  if (fromChainId && fromToken && toChainId && toToken) {
    return `Chains you can bridge ${fromToken.symbol} to ${toToken.symbol} on ${toChainName}`
  }
}
