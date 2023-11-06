import {
  Address,
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt } from 'viem'

import { MINICHEF_ABI } from '@/constants/abis/miniChef'

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
  const { request } = await prepareWriteContract({
    chainId,
    address: lpAddress,
    abi: MINICHEF_ABI,
    functionName: 'harvest',
    args: [poolId, address],
  })

  const { hash } = await writeContract(request)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
