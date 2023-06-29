import { useNetwork } from 'wagmi'
import { getContract } from 'wagmi/actions'
import { Address } from 'wagmi'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'
import {MINICHEF_ABI} from '@/constants/abis/miniChef'

export const useMiniChefContract = (): [any, Address] => {
  const { chain: connectedChain } = useNetwork()
  const miniChefAddress: Address = MINICHEF_ADDRESSES[connectedChain.id]

  const miniChefContract = getContract({
    address: MINICHEF_ADDRESSES[connectedChain.id],
    abi: MINICHEF_ABI,
  })

  return [miniChefContract, miniChefAddress]
}
