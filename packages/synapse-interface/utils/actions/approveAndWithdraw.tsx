import type { Token } from '@types'
import type { Address } from 'wagmi'

import toast from 'react-hot-toast'

import { ALL } from '@constants/withdrawTypes'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { txErrorHandler } from '@utils/txErrorHandler'

import { approveErc20Token } from '@/actions/approveErc20Token'
import { swapPoolRemoveLiquidity } from '@/actions/swapPoolRemoveLiquidity'
import { swapPoolRemoveLiquidityOneToken } from '@/actions/swapPoolRemoveLiquidityOneToken'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import ExplorerToastLink from '@components/ExplorerToastLink'

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

  return await approveErc20Token({
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
    segmentAnalyticsEvent(
      `[Pool Withdrawal] Attempt`,
      {
        poolName: pool?.name,
        inputAmount,
      },
      true
    )
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
      segmentAnalyticsEvent(`[Pool Withdrawal] Success`, {
        poolName: pool?.name,
        inputAmount,
      })

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
    segmentAnalyticsEvent(`[Pool Withdrawal] Failure`, {
      poolName: pool?.name,
      inputAmount,
      errorCode: error.code,
    })
    txErrorHandler(error)
    return error
  }
}
