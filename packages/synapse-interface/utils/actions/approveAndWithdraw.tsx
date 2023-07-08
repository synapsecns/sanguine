import { ALL } from '@constants/withdrawTypes'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { txErrorHandler } from '@utils/txErrorHandler'
import { Token } from '@types'
import toast from 'react-hot-toast'
import { getErc20TokenApproval } from '@/actions/getErc20TokenApproval'

import { Address } from 'wagmi'
import { swapPoolRemoveLiquidity } from '@/actions/swapPoolRemoveLiquidity'
import { swapPoolRemoveLiquidityOneToken } from '@/actions/swapPoolRemoveLiquidityOneToken'

export const approve = async (
  pool: Token,
  depositQuote: any,
  inputValue: bigint,
  chainId: number
) => {
  if (inputValue === 0n || inputValue < depositQuote.allowance) {
    return
  }

  const { poolAddress } = getSwapDepositContractFields(pool, chainId)

  return await getErc20TokenApproval({
    chainId,
    tokenAddress: pool.addresses[chainId] as Address,
    spender: poolAddress,
    amount: inputValue,
  })
}

export const withdraw = async (
  pool: Token,
  slippageSelected: any,
  slippageCustom: any,
  inputAmount: bigint,
  chainId: number,
  withdrawType: string,
  outputs: Record<
    string,
    {
      value: any
      index: number
    }
  >
) => {
  let spendTransaction
  let pendingPopup: any
  let successPopup: any

  pendingPopup = toast(`Starting your withdrawal...`, {
    id: 'withdraw-in-progress-popup',
    duration: Infinity,
  })

  try {
    if (withdrawType === ALL) {
      spendTransaction = await swapPoolRemoveLiquidity({
        chainId,
        pool,
        amount: inputAmount,
        outputs,
        slippageSelected,
        slippageCustom,
      })
    } else {
      const outputAmount = Object.values(outputs)[0]
      const poolTokenIndex = outputAmount.index

      spendTransaction = await swapPoolRemoveLiquidityOneToken({
        chainId,
        pool,
        amount: inputAmount,
        outputAmount,
        slippageSelected,
        slippageCustom,
        poolTokenIndex,
      })
    }

    if (spendTransaction.status === 'success') {
      toast.dismiss(pendingPopup)

      const successToastContent = (
        <div>
          <div>Completed Withdrawal: </div>
          <ExplorerToastLink
            transactionHash={spendTransaction.transactionHash}
            chainId={chainId}
          />
        </div>
      )

      successPopup = toast.success(successToastContent, {
        id: 'withdraw-success-popup',
        duration: 10000,
      })
    }

    return spendTransaction
  } catch (error) {
    toast.dismiss(pendingPopup)
    txErrorHandler(error)
    return error
  }
}
