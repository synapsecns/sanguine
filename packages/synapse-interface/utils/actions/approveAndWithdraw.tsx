import { ALL } from '@constants/withdrawTypes'
import {
  getSwapDepositContractFields,
  useSwapDepositContract,
} from '@hooks/useSwapDepositContract'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { subtractSlippage, subtractSlippageBigInt } from '@utils/slippage'
import { txErrorHandler } from '@utils/txErrorHandler'
import { approveToken } from '@utils/approveToken'
import { Token } from '@types'
import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import toast from 'react-hot-toast'

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

  return await approveToken(
    poolAddress,
    chainId,
    pool.addresses[chainId],
    inputValue
  )
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
  const poolContract = await useSwapDepositContract(pool, chainId)
  let spendTransaction
  let pendingPopup: any
  let successPopup: any

  pendingPopup = toast(`Starting your withdrawal...`, {
    id: 'withdraw-in-progress-popup',
    duration: Infinity,
  })

  try {
    if (withdrawType === ALL) {
      spendTransaction = await poolContract.removeLiquidity(
        inputAmount,
        pool.poolTokens?.map((t, index) =>
          subtractSlippageBigInt(
            outputs[withdrawType][index].value,
            slippageSelected,
            slippageCustom
          )
        ),
        Math.round(new Date().getTime() / 1000 + 60 * 10)
      )
    } else {
      const outputAmount = Object.values(outputs)[0]
      const poolTokenIndex = outputAmount.index

      spendTransaction = await poolContract.removeLiquidityOneToken(
        inputAmount,
        poolTokenIndex,
        subtractSlippageBigInt(
          outputAmount.value,
          slippageSelected,
          slippageCustom
        ),
        Math.round(new Date().getTime() / 1000 + 60 * 10)
      )
    }

    const tx = await spendTransaction.wait()

    toast.dismiss(pendingPopup)

    const successToastContent = (
      <div>
        <div>Completed Withdrawal: </div>
        <ExplorerToastLink
          transactionHash={tx?.transactionHash}
          chainId={chainId}
        />
      </div>
    )

    successPopup = toast.success(successToastContent, {
      id: 'withdraw-success-popup',
      duration: 10000,
    })

    return tx
  } catch (error) {
    toast.dismiss(pendingPopup)
    txErrorHandler(error)
    return error
  }
}
