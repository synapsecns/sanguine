import { Token } from '@types'
import { SWAP_ABI } from '@abis/swap'
import { SWAP_ETH_WRAPPER_ABI } from '@abis/swapEthWrapper'

import { AV_SWAP_WRAPPER_ABI } from '@/constants/abis/avSwapWrapper'

export const getSwapDepositContractFields = (pool: Token, chainId: number) => {
  let poolAddress
  let abi
  let swapType

  if (pool?.swapWrapperAddresses?.[chainId]) {
    poolAddress = pool.swapWrapperAddresses[chainId]
    abi = AV_SWAP_WRAPPER_ABI
    swapType = 'AV_SWAP'
  } else if (pool?.swapEthAddresses?.[chainId]) {
    poolAddress = pool.swapEthAddresses[chainId]
    abi = SWAP_ETH_WRAPPER_ABI
    swapType = 'SWAP_ETH'
  } else if (pool?.swapAddresses?.[chainId]) {
    poolAddress = pool?.swapAddresses[chainId]
    abi = SWAP_ABI
    swapType = 'SWAP'
  } else {
    poolAddress = undefined
    abi = undefined
    swapType = undefined
  }

  return { poolAddress, abi, swapType }
}
