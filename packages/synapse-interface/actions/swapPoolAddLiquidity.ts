import {
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt } from 'viem'

import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'

export const swapPoolAddLiquidity = async ({
  chainId,
  pool,
  spendTransactionArgs,
}: {
  chainId: number
  pool: any
  spendTransactionArgs: any
}) => {
  const { abi, poolAddress } = getSwapDepositContractFields(pool, chainId)

  console.log(`spendTransactionArgs`, spendTransactionArgs)

  const config = await prepareWriteContract({
    chainId,
    address: poolAddress,
    abi,
    functionName: 'addLiquidity',
    args: [...spendTransactionArgs],
  })

  const { hash } = await writeContract(config)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
