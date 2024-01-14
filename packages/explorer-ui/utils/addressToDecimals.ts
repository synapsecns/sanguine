import { TOKEN_HASH_MAP } from 'synapse-constants/dist'

export function addressToDecimals({ tokenAddress, chainId }) {
  let decimals =
    tokenAddress &&
    chainId &&
    TOKEN_HASH_MAP[chainId][tokenAddress]?.decimals[chainId]

  if (decimals === undefined) {
    decimals = 18
  }
  return decimals
}
