import {
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'
import { Address } from 'viem'

import { MINICHEF_ABI } from '@/constants/abis/miniChef'
import { wagmiConfig } from '@/wagmiConfig'

export const harvestLpPool = async ({
  address,
  chainId,
  poolId,
  lpAddress,
}: {
  address: Address
  chainId: number
  poolId: number
  lpAddress: Address
}) => {
  const { request } = await simulateContract(wagmiConfig, {
    chainId,
    address: lpAddress,
    abi: MINICHEF_ABI,
    functionName: 'harvest',
    args: [poolId, address],
  })

  const hash = await writeContract(wagmiConfig, request)
  const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

  return txReceipt
}
