import {
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt } from 'viem'

import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { subtractSlippageBigInt } from '@/utils/slippage'
import { Token } from '@/utils/types'

export const swapPoolRemoveLiquidityOneToken = async ({
  chainId,
  pool,
  amount,
  slippageSelected,
  slippageCustom,
  poolTokenIndex,
  outputAmount,
}: {
  chainId: number
  pool: Token
  amount: bigint
  slippageSelected: any
  slippageCustom: any
  poolTokenIndex: number
  outputAmount: { value: any; index: number }
}) => {
  const { abi, poolAddress } = getSwapDepositContractFields(pool, chainId)

  const config = await prepareWriteContract({
    chainId,
    address: poolAddress,
    abi,
    functionName: 'removeLiquidityOneToken',
    args: [
      amount,
      poolTokenIndex,
      subtractSlippageBigInt(
        outputAmount.value,
        slippageSelected,
        slippageCustom
      ),
      Math.round(new Date().getTime() / 1000 + 60 * 10),
    ],
  })

  const { hash } = await writeContract(config)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
