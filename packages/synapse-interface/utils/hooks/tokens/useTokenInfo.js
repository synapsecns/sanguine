import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'

/**
 * @param {Token} token
 */
export function useTokenInfo(token) {
  const { chainId } = useActiveWeb3React()
  return {
    ...token,
    address:            token?.addresses?.[chainId],
    wrapperAddress:     token?.wrapperAddresses?.[chainId],
    swapAddress:        token?.swapAddresses?.[chainId],
    swapWrapperAddress: token?.swapWrapperAddresses?.[chainId],
    swapDepositAddress: token?.swapDepositAddresses?.[chainId],
    swapEthAddress:     token?.swapEthAddresses?.[chainId],
    poolId:             token?.poolId?.[chainId],
    decimals:           token?.decimals?.[chainId]
  }
}


/**
 * @param {number} chainId
 * @param {Token} token
 */
export function getTokenOnChain(chainId, token) {
  return {
    ...token,
    address:            token?.addresses?.[chainId],
    wrapperAddress:     token?.wrapperAddresses?.[chainId],
    swapAddress:        token?.swapAddresses?.[chainId],
    swapWrapperAddress: token?.swapWrapperAddresses?.[chainId],
    swapDepositAddress: token?.swapDepositAddresses?.[chainId],
    swapEthAddress:     token?.swapEthAddresses?.[chainId],
    poolId:             token?.poolId?.[chainId],
    decimals:           token?.decimals?.[chainId]
  }
}

