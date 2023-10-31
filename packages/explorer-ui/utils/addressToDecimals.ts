import { TOKEN_HASH_MAP } from '@constants/tokens/basic'

export function addressToDecimals({ tokenAddress, chainId }) {
  const decimals =
    tokenAddress &&
    chainId &&
    TOKEN_HASH_MAP[chainId][tokenAddress.toLowerCase()]?.decimals[chainId]

  return decimals
}
