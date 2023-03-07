import MULTICALL_ABI from '@abis/multicall.json'
import MULTICALL2_ABI from '@abis/multicall2.json'
import { MULTICALL_ADDRESSES, MULTICALL2_ADDRESSES } from '@constants/multicall'

import { useContract, useGenericContract } from '@hooks/contracts/useContract'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'


export function useMulticallContract() {
  const { chainId } = useActiveWeb3React()
  return useContract(
    chainId && MULTICALL_ADDRESSES[chainId],
    MULTICALL_ABI,
    false
  )
}


export function useGenericMulticallContract(chainId) {
  return useGenericContract(
    chainId,
    chainId && MULTICALL_ADDRESSES[chainId],
    MULTICALL_ABI,
    false
  )
}

export function useMulticall2Contract() {
  const { chainId } = useActiveWeb3React()
  return useContract(
    chainId && MULTICALL2_ADDRESSES[chainId],
    MULTICALL2_ABI,
    false
  )
}

export function useGenericMulticall2Contract(chainId) {
  return useGenericContract(
    chainId,
    chainId && MULTICALL2_ADDRESSES[chainId],
    MULTICALL2_ABI,
    false
  )
}