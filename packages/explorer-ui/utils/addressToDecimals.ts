import { TOKEN_HASH_MAP } from '@constants-new/tokens/index'

export function addressToDecimals({ tokenAddress, chainId }) {
  let decimals =
    tokenAddress &&
    chainId &&
    TOKEN_HASH_MAP[chainId][tokenAddress.toLowerCase()]?.decimals[chainId]

  if (decimals === undefined) {
    decimals = 18
  }
  return decimals
}
