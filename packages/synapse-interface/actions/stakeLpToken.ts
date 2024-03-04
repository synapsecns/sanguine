import {
  Address,
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import type { TransactionReceipt } from 'viem'

import { MINICHEF_ABI } from '@/constants/abis/miniChef'

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
  const { request } = await prepareWriteContract({
    chainId,
    address: lpAddress,
    abi: MINICHEF_ABI,
    functionName: 'deposit',
    args: [poolId, amount, address],
  })

  const { hash } = await writeContract(request)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
