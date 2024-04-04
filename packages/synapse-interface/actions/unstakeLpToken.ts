import {
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'
import { Address } from 'viem'

import { MINICHEF_ABI } from '@/constants/abis/miniChef'
import { wagmiConfig } from '@/wagmiConfig'

export const unstakeLpToken = async ({
  address,
  chainId,
  poolId,
  amount,
  lpAddress,
}: {
  address: Address
  chainId: number
  poolId: number
  amount: bigint
  lpAddress: Address
}) => {
  const { request } = await simulateContract(wagmiConfig, {
    // TODO: Fix any
    chainId: chainId as any,
    address: lpAddress,
    abi: MINICHEF_ABI,
    functionName: 'withdraw',
    args: [poolId, amount, address],
  })

  const hash = await writeContract(wagmiConfig, request)
  const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

  return txReceipt
}
