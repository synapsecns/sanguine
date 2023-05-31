import { ALL } from '@constants/withdrawTypes'
import { useSwapDepositContract } from '@hooks/useSwapDepositContract'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { subtractSlippage } from '@utils/slippage'
import { txErrorHandler } from '@utils/txErrorHandler'
import { approveToken } from '@utils/approveToken'
import { Token } from '@types'
import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import toast from 'react-hot-toast'

export const approve = async (
  pool: Token,
  depositQuote: any,
  inputValue: BigNumber,
  chainId: number
) => {
  if (inputValue.isZero() || inputValue.lt(depositQuote.allowance)) {
    return
  }
  await approveToken(
    pool.swapAddresses[chainId],
    chainId,
    pool.addresses[chainId],
    inputValue
  )
}

export const withdraw = async (
  pool: Token,
  slippageSelected: any,
  slippageCustom: any,
  inputAmount: BigNumber,
  chainId: number,
  withdrawType: string,
  outputs: Record<
    string,
    {
      value: BigNumber
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
      const outputMinArr = pool.poolTokens.map(() => Zero)
      for (let poolToken of pool.poolTokens) {
        const outputAmount = outputs[poolToken.addresses[chainId]]
        outputMinArr[outputAmount.index] = subtractSlippage(
          outputAmount.value,
          slippageSelected,
          slippageCustom
        )
      }
      spendTransaction = await poolContract.removeLiquidity(
        inputAmount,
        outputMinArr,
        Math.round(new Date().getTime() / 1000 + 60 * 10)
      )
    } else {
      const outputAmount = Object.values(outputs)[0]
      const poolTokenIndex = outputAmount.index
      spendTransaction = await poolContract.removeLiquidityOneToken(
        inputAmount,
        poolTokenIndex,
        subtractSlippage(outputAmount.value, slippageSelected, slippageCustom),
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
  }
}
