import {
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import type { TransactionReceipt } from 'viem'
import type { Token } from '@/utils/types'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { subtractSlippageBigInt } from '@/utils/slippage'


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

  const { request } = await prepareWriteContract({
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

  const { hash } = await writeContract(request)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
