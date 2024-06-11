import { ALL } from '@constants/withdrawTypes'
import {
  type SimulateContractParameters,
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'

import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { subtractSlippageBigInt } from '@/utils/slippage'
import { wagmiConfig } from '@/wagmiConfig'

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

  const { request } = await simulateContract(wagmiConfig, {
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
  } as SimulateContractParameters)

  const hash = await writeContract(wagmiConfig, request)
  const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

  return txReceipt
}
