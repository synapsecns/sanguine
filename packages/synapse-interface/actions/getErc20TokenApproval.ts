import {
  Address,
  erc20ABI,
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt } from 'viem'

export const getErc20TokenApproval = async ({
  chainId,
  amount,
  spender,
  tokenAddress,
}: {
  chainId: number
  amount: bigint
  spender: Address
  tokenAddress: Address
}) => {
  const config = await prepareWriteContract({
    chainId,
    address: tokenAddress,
    abi: erc20ABI,
    functionName: 'approve',
    args: [spender, amount],
  })

  const { hash } = await writeContract(config)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
