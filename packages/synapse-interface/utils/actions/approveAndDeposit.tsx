import _ from 'lodash'

import toast from 'react-hot-toast'

import { subtractSlippageBigInt } from '@utils/slippage'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'

import ExplorerToastLink from '@components/ExplorerToastLink'

import { CHAINS_BY_ID } from '@/constants/chains'
import { txErrorHandler } from '@utils/txErrorHandler'
import { AVWETH, WETHE } from '@constants/tokens/master'
import { WETH } from '@constants/tokens/swapMaster'
import { approveToken } from '@utils/approveToken'
import { Token } from '@types'
import { zeroAddress } from 'viem'
import { swapPoolCalculateTokenAmount } from '@/actions/swapPoolCalculateTokenAmount'
import { swapPoolAddLiquidity } from '@/actions/swapPoolAddLiquidity'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

export const approve = async (
  pool: Token,
  depositQuote: any,
  inputValue: Record<string, bigint>,
  chainId: number
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name

  const { poolAddress } = getSwapDepositContractFields(pool, chainId)

  const requestingApprovalPopup = toast(
    `Requesting approval on ${currentChainName}`,
    {
      id: 'approve-in-progress-popup',
      duration: Infinity,
    }
  )

  const handleApproval = async (token, tokenAddr) => {
    if (
      inputValue[tokenAddr] &&
      (inputValue[tokenAddr] === 0n ||
        inputValue[tokenAddr] <= depositQuote.allowances[tokenAddr])
    ) {
      return
    }

    if (token.symbol === WETH.symbol) return

    const tokenToApprove =
      token.symbol === AVWETH.symbol
        ? WETHE.addresses[chainId]
        : token.addresses[chainId]

    const approveTx = await approveToken(
      poolAddress,
      chainId,
      tokenToApprove,
      inputValue[tokenAddr]
    )

    if (!approveTx) return

    toast.dismiss(requestingApprovalPopup)
    const successToastContent = (
      <div>
        <div>Successfully approved on {currentChainName}</div>
        <ExplorerToastLink
          transactionHash={approveTx ?? zeroAddress}
          chainId={chainId}
        />
      </div>
    )

    toast.success(successToastContent, {
      id: 'approve-success-popup',
      duration: 10000,
    })
    segmentAnalyticsEvent(`[Pool Approval] Successful for ${pool?.name}`, {})

    return approveTx
  }

  for (let token of pool.poolTokens) {
    try {
      const value = inputValue[token.addresses[chainId]]
      const hasNonZeroValue = !!value && value !== 0n

      if (hasNonZeroValue) {
        await handleApproval(token, token.addresses[chainId])
      }
    } catch (error) {
      toast.dismiss(requestingApprovalPopup)
      txErrorHandler(error)
      return error
    }
  }
}

export const deposit = async (
  pool: Token,
  slippageSelected: any,
  slippageCustom: any,
  inputAmounts: any,
  chainId: number
) => {
  let pendingPopup: any
  let successPopup: any

  pendingPopup = toast(`Starting your deposit...`, {
    id: 'deposit-in-progress-popup',
    duration: Infinity,
  })

  try {
    // get this from quote?
    segmentAnalyticsEvent(`[Pool Deposit] Attempt for ${pool?.name}`, {})
    let minToMint: any = await swapPoolCalculateTokenAmount({
      chainId,
      pool,
      inputAmounts,
    })

    minToMint = subtractSlippageBigInt(
      minToMint,
      slippageSelected,
      slippageCustom
    )

    const result = Array.from(Object.values(inputAmounts), (value) => value)

    const wethIndex = _.findIndex(
      pool.poolTokens,
      (t) => t.symbol == WETH.symbol
    )

    let spendTransactionArgs = [
      result,
      minToMint,
      Math.round(new Date().getTime() / 1000 + 60 * 10),
    ]

    const liquidityAmounts = Object.values(inputAmounts)

    if (wethIndex >= 0) {
      spendTransactionArgs.push({ value: liquidityAmounts[wethIndex] })
    }

    const tx = await swapPoolAddLiquidity({
      chainId,
      pool,
      spendTransactionArgs,
    })

    toast.dismiss(pendingPopup)

    const successToastContent = (
      <div>
        <div>Liquidity added!</div>
        <ExplorerToastLink
          transactionHash={tx?.transactionHash}
          chainId={chainId}
        />
      </div>
    )

    successPopup = toast.success(successToastContent, {
      id: 'deposit-success-popup',
      duration: 10000,
    })
    segmentAnalyticsEvent(`[Pool Deposit] Success for ${pool?.name}`, {
      inputAmounts,
    })

    return tx
  } catch (error) {
    console.log('error from deposit: ', error)
    toast.dismiss(pendingPopup)
    segmentAnalyticsEvent(`[Pool Deposit] Failure for ${pool?.name}`, {
      inputAmounts,
      errorCode: error.code,
    })
    txErrorHandler(error)
    return error
  }
}