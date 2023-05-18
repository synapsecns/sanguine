import { useContract, useNetwork, Address } from 'wagmi'
import { Contract } from 'ethers'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import MINI_CHEF_ABI from '@/constants/abis/miniChef.json'

export const useMiniChefContract = (): [Contract, Address] => {
  const { chain: connectedChain } = useNetwork()
  const miniChefAddress: Address = MINICHEF_ADDRESSES[connectedChain.id]

  const miniChefContract = useContract({
    address: MINICHEF_ADDRESSES[connectedChain.id],
    abi: MINI_CHEF_ABI,
  })

  return [miniChefContract, miniChefAddress]
}
