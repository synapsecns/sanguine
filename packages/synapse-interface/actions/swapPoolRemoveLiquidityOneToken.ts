import {
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'

import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { subtractSlippageBigInt } from '@/utils/slippage'
import { Token } from '@/utils/types'
import { wagmiConfig } from '@/wagmiConfig'

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

  const { request } = await simulateContract(wagmiConfig, {
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

  const hash = await writeContract(wagmiConfig, request)
  const txReceipt = await waitForTransactionReceipt(wagmiConfig, {
    hash,
  })

  return txReceipt
}
