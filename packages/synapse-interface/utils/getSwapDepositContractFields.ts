import { Token } from '@types'
import { SWAP_ABI } from '@abis/swap'
import AV_SWAP_WRAPPER_ABI from '@abis/avSwapWrapper.json'
import SWAP_ETH_WRAPPER_ABI from '@abis/swapEthWrapper.json'

export const getSwapDepositContractFields = (pool: Token, chainId: number) => {
  let poolAddress
  let abi

  if (pool?.swapEthAddresses?.[chainId]) {
    poolAddress = pool.swapEthAddresses[chainId]
    abi = SWAP_ETH_WRAPPER_ABI
  } else if (pool?.swapWrapperAddresses?.[chainId]) {
    poolAddress = pool.swapWrapperAddresses[chainId]
    abi = AV_SWAP_WRAPPER_ABI
  } else if (pool?.swapAddresses?.[chainId]) {
    poolAddress = pool?.swapAddresses[chainId]
    abi = SWAP_ABI
  } else {
    poolAddress = undefined
    abi = undefined
  }

  return { poolAddress, abi }
}
