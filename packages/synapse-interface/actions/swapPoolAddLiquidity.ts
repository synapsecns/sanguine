import {
  type SimulateContractParameters,
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'

import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { wagmiConfig } from '@/wagmiConfig'

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

  const { request } = await simulateContract(
    wagmiConfig,
    pwcConfig as SimulateContractParameters
  )

  const hash = await writeContract(wagmiConfig, request)
  const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

  return txReceipt
}
