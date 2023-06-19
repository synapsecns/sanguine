import { readContracts, ReadContractResult, Address } from '@wagmi/core'
import { Zero } from '@ethersproject/constants'
import MINICHEF_ABI from '@abis/miniChef.json'

import { MINICHEF_ADDRESSES } from '@/constants/minichef'

export const getStakedBalance = async (
  address: Address,
  chainId: number,
  poolId: number
) => {
  const miniChefContractAddress: Address = MINICHEF_ADDRESSES[chainId]
  try {
    const data: ReadContractResult = await readContracts({
      contracts: [
        {
          address: miniChefContractAddress,
          abi: MINICHEF_ABI,
          functionName: 'userInfo',
          args: [poolId, address],
        },
        {
          address: miniChefContractAddress,
          abi: MINICHEF_ABI,
          functionName: 'pendingSynapse',
          args: [poolId, address],
        },
      ],
    })

    return { amount: data[0]?.amount ?? Zero, reward: data[1] ?? Zero }
  } catch (error) {
    console.error('Error from useStakedBalance: ', error)
    return { amount: Zero, reward: Zero }
  }
}
