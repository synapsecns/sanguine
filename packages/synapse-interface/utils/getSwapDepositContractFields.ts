import type { Token } from '@types'
import { SWAP_ABI } from '@/constants/abis/swap'
import { SWAP_ETH_WRAPPER_ABI } from '@/constants/abis/swapEthWrapper'
import { AV_SWAP_WRAPPER_ABI } from '@/constants/abis/avSwapWrapper'

export const getSwapDepositContractFields = (pool: Token, chainId: number) => {
  if (pool?.swapWrapperAddresses?.[chainId]) {
    return {
      poolAddress: pool.swapWrapperAddresses[chainId],
      abi:         AV_SWAP_WRAPPER_ABI,
      swapType:    'AV_SWAP'
    }
  } else if (pool?.swapEthAddresses?.[chainId]) {
    return {
      poolAddress: pool.swapEthAddresses[chainId],
      abi:         SWAP_ETH_WRAPPER_ABI,
      swapType:    'SWAP_ETH'
    }
  } else if (pool?.swapAddresses?.[chainId]) {
    return {
      poolAddress: pool?.swapAddresses[chainId],
      abi:         SWAP_ABI,
      swapType:    'SWAP'
    }
  } else {
    return {
      poolAddress: undefined,
      abi:         undefined,
      swapType:    undefined
    }
  }

}
