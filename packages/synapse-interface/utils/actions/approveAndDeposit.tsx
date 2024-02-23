import _ from 'lodash'

import toast from 'react-hot-toast'

import { zeroAddress } from 'viem'

import type { Token } from '@types'

import { subtractSlippageBigInt } from '@/utils/slippage'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { approveToken } from '@/utils/approveToken'
import { txErrorHandler } from '@/utils/txErrorHandler'


import { CHAINS_BY_ID } from '@/constants/chains'
import { WETHE, WETH } from '@/constants/tokens/bridgeable'
import { AVWETH } from '@/constants/tokens/auxilliary'

import { swapPoolCalculateTokenAmount } from '@/actions/swapPoolCalculateTokenAmount'
import { swapPoolAddLiquidity } from '@/actions/swapPoolAddLiquidity'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import ExplorerToastLink from '@components/ExplorerToastLink'


export const approve = async (
  pool: Token,
  depositQuote: any,
  inputValue: Record<string, bigint>,
  chainId: number
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name

  const { poolAddress, swapType } = getSwapDepositContractFields(pool, chainId)

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
      toast.dismiss(requestingApprovalPopup)
      return
    }

    if (token.addresses[pool.chainId] === zeroAddress) {
      toast.dismiss(requestingApprovalPopup)
      return
    }

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
    segmentAnalyticsEvent(`[Pool Approval] Successful`, {
      poolName: pool?.name,
    })

    return approveTx
  }

  const tokens = swapType === 'AV_SWAP' ? pool.nativeTokens : pool.poolTokens

  for (let token of tokens) {
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
    segmentAnalyticsEvent(
      `[Pool Deposit] Attempt`,
      {
        poolName: pool?.name,
      },
      true
    )

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
    segmentAnalyticsEvent(`[Pool Deposit] Success`, {
      poolName: pool?.name,
      inputAmounts,
    })

    return tx
  } catch (error) {
    console.log('error from deposit: ', error)
    toast.dismiss(pendingPopup)
    segmentAnalyticsEvent(`[Pool Deposit] Failure`, {
      poolName: pool?.name,
      inputAmounts,
      errorCode: error.code,
    })
    txErrorHandler(error)
    return error
  }
}

export const emptyPoolDeposit = async (
  pool: Token,
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
    segmentAnalyticsEvent(
      `[Empty Pool Deposit] Attempt`,
      {
        poolName: pool?.name,
      },
      true
    )

    const result = Array.from(Object.values(inputAmounts), (value) => value)

    const wethIndex = _.findIndex(
      pool.poolTokens,
      (t) => t.symbol == WETH.symbol
    )

    let spendTransactionArgs = [
      result,
      0n as any,
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
    segmentAnalyticsEvent(`[Empty Pool Deposit] Success`, {
      poolName: pool?.name,
      inputAmounts,
    })

    return tx
  } catch (error) {
    console.log('error from deposit: ', error)
    toast.dismiss(pendingPopup)
    segmentAnalyticsEvent(`[Empty Pool Deposit] Failure`, {
      poolName: pool?.name,
      inputAmounts,
      errorCode: error.code,
    })
    txErrorHandler(error)
    return error
  }
}
