import { Token } from '@/utils/types'

export const getTokenOnChain = (chainId: number, token: Token) => {
  return {
    ...token,
    address: token?.addresses?.[chainId],
    wrapperAddress: token?.wrapperAddresses?.[chainId],
    swapAddress: token?.swapAddresses?.[chainId],
    swapWrapperAddress: token?.swapWrapperAddresses?.[chainId],
    swapDepositAddress: token?.swapDepositAddresses?.[chainId],
    swapEthAddress: token?.swapEthAddresses?.[chainId],
    poolId: token?.poolId?.[chainId as keyof Token['poolId']],
    decimals: token?.decimals?.[chainId as keyof Token['decimals']],
  }
}
