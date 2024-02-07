import {
  Address,
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt } from 'viem'

import { MINICHEF_ABI } from '@/constants/abis/miniChef'

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
  const { request } = await prepareWriteContract({
    chainId,
    address: lpAddress,
    abi: MINICHEF_ABI,
    functionName: 'withdraw',
    args: [poolId, amount, address],
  })

  const { hash } = await writeContract(request)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
