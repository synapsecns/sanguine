const MAX_UINT256 = BigInt(
  '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
)

import {
  Address,
  erc20ABI,
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt } from 'viem'

export const approveErc20Token = async ({
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
    args: [spender, amount ?? MAX_UINT256],
  })

  const { hash } = await writeContract(config)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
