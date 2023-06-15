import { ALL } from '@constants/withdrawTypes'
import {
  getSwapDepositContractFields,
  useSwapDepositContract,
} from '@hooks/useSwapDepositContract'
import ExplorerToastLink from '@components/ExplorerToastLink'
import { subtractSlippage } from '@utils/slippage'
import { txErrorHandler } from '@utils/txErrorHandler'
import { approveToken } from '@utils/approveToken'
import { Token } from '@types'
import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import toast from 'react-hot-toast'
import { useAnalytics } from '@/contexts/AnalyticsProvider'
import { getAccount } from '@wagmi/core'
import { shortenAddress } from '../shortenAddress'

export const approve = async (
  pool: Token,
  depositQuote: any,
  inputValue: BigNumber,
  chainId: number
) => {
  if (inputValue.isZero() || inputValue.lt(depositQuote.allowance)) {
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
  const analytics = useAnalytics()
  const account = getAccount()
  const address = account?.address

  pendingPopup = toast(`Starting your withdrawal...`, {
    id: 'withdraw-in-progress-popup',
    duration: Infinity,
  })

  try {
    analytics.track(
      `[Pool Withdrawal] ${shortenAddress(address)} Attempt for ${pool?.name}`,
      {
        inputAmount,
      },
      {
        context: { ip: '0.0.0.0' },
      }
    )
    if (withdrawType === ALL) {
      console.log(outputs[withdrawType])
      spendTransaction = await poolContract.removeLiquidity(
        inputAmount,
        pool.poolTokens?.map((t, index) =>
          subtractSlippage(
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

    analytics.track(
      `[Pool Withdrawal] ${shortenAddress(address)} Success for ${pool?.name}`,
      {
        inputAmount,
      },
      { context: { ip: '0.0.0.0' } }
    )

    return tx
  } catch (error) {
    analytics.track(
      `[Pool Withdrawal] ${shortenAddress(address)} Failure for ${pool?.name}`,
      {
        inputAmount,
        errorCode: error.code,
      },
      { context: { ip: '0.0.0.0' } }
    )
    toast.dismiss(pendingPopup)
    txErrorHandler(error)
    return error
  }
}
