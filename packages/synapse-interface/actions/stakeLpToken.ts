import {
  type SimulateContractParameters,
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'
import { Address } from 'viem'

import { MINICHEF_ABI } from '@/constants/abis/miniChef'
import { wagmiConfig } from '@/wagmiConfig'

export const stakeLpToken = async ({
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
    chainId,
    address: lpAddress,
    abi: MINICHEF_ABI,
    functionName: 'deposit',
    args: [poolId, amount, address],
  } as SimulateContractParameters)

  const hash = await writeContract(wagmiConfig, request)
  const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

  return txReceipt
}
