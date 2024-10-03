//@ts-ignore
import { TOKEN_HASH_MAP } from '@synapsecns/synapse-constants'

export const addressToSymbol = ({ tokenAddress, chainId }) => {
  if (
    tokenAddress === '0x53f7c5869a859F0AeC3D334ee8B4Cf01E3492f21' &&
    chainId === 43114
  ) {
    return 'AVWETH'
  }
  if (
    tokenAddress.toLowerCase() === '0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee'
  ) {
    return 'ETH'
  }

  const symbol =
    tokenAddress && chainId && TOKEN_HASH_MAP[chainId][tokenAddress]?.symbol
  return symbol
}
