import {
  prepareWriteContract,
  waitForTransaction,
  writeContract,
} from '@wagmi/core'
import type { TransactionReceipt } from 'viem'

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
  const { abi, poolAddress, swapType } = getSwapDepositContractFields(
    pool,
    chainId
  )

  const pwcBaseConfig = {
    chainId,
    address: poolAddress,
    abi,
    functionName: 'addLiquidity',
    args: [...spendTransactionArgs].slice(0, 3),
  }

  let pwcConfig

  if (swapType === 'SWAP') {
    pwcConfig = pwcBaseConfig
  } else if (swapType === 'SWAP_ETH') {
    pwcConfig = {
      ...pwcBaseConfig,
      value: [...spendTransactionArgs][3].value,
    }
  } else if (swapType === 'AV_SWAP') {
    pwcConfig = pwcBaseConfig
  }

  const { request } = await prepareWriteContract(pwcConfig)

  const { hash } = await writeContract(request)
  const txReceipt: TransactionReceipt = await waitForTransaction({ hash })

  return txReceipt
}
