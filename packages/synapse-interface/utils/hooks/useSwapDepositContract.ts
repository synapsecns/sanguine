import { useMemo } from 'react'
import { Token } from '@types'
import SWAP_ABI from '@abis/swap.json'
import AV_SWAP_WRAPPER_ABI from '@abis/avSwapWrapper.json'
import SWAP_ETH_WRAPPER_ABI from '@abis/swapEthWrapper.json'
import { Contract } from '@ethersproject/contracts'

export const useSwapDepositContract = (pool: Token, chainId: number) => {
  let address
  let abi
  if (pool?.swapEthAddresses?.[chainId]) {
    address = pool.swapAddresses[chainId]
    abi = SWAP_ETH_WRAPPER_ABI
  } else if (pool?.swapWrapperAddresses?.[chainId]) {
    address = pool.swapWrapperAddresses[chainId]
    abi = AV_SWAP_WRAPPER_ABI
  } else {
    address = pool.swapAddresses[chainId]
    abi = SWAP_ABI
  }

  const swapContract = new Contract(address, abi)

  return useMemo(() => swapContract, [pool])
}
