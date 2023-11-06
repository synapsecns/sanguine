import { ALL } from '@constants/withdrawTypes'
import {
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt } from 'viem'

import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { subtractSlippageBigInt } from '@/utils/slippage'

export const swapPoolRemoveLiquidity = async ({
  chainId,
  pool,
  amount,
  outputs,
  slippageSelected,
  slippageCustom,
}: {
  chainId: number
  pool: any
  amount: bigint
  outputs: Record<string, { value: any; index: number }>
  slippageSelected: any
  slippageCustom: any
}) => {
  const { abi, poolAddress } = getSwapDepositContractFields(pool, chainId)

  const { request } = await prepareWriteContract({
    chainId,
    address: poolAddress,
    abi,
    functionName: 'removeLiquidity',
    args: [
      amount,
      pool.poolTokens?.map((t, index) =>
        subtractSlippageBigInt(
          outputs[ALL][index].value,
          slippageSelected,
          slippageCustom
        )
      ),
      Math.round(new Date().getTime() / 1000 + 60 * 10),
    ],
  })

  const { hash } = await writeContract(request)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
