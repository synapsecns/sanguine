import _ from 'lodash'

import toast from 'react-hot-toast'

import { useSwapDepositContract } from '@hooks/useSwapDepositContract'
import { subtractSlippage } from '@utils/slippage'

import ExplorerToastLink from '@components/ExplorerToastLink'

import { AddressZero } from '@ethersproject/constants'
import { CHAINS_BY_ID } from '@/constants/chains'
import { txErrorHandler } from '@utils/txErrorHandler'
import { AVWETH, WETHE } from '@constants/tokens/master'
import { WETH } from '@constants/tokens/swapMaster'
import { approveToken } from '@utils/approveToken'
import { Token } from '@types'

import { parseUnits } from '@ethersproject/units'

export const approve = async (
  pool: Token,
  depositQuote: any,
  inputValue: any,
  chainId: number
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name

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
      (inputValue[tokenAddr].isZero() ||
        inputValue[tokenAddr].lte(depositQuote.allowances[tokenAddr]))
    )
      return

    if (token.symbol === WETH.symbol) return

    const tokenToApprove =
      token.symbol === AVWETH.symbol
        ? WETHE.addresses[chainId]
        : token.addresses[chainId]

    const approveTx = await approveToken(
      pool.swapAddresses[chainId],
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
          transactionHash={approveTx?.hash ?? AddressZero}
          chainId={chainId}
        />
      </div>
    )

    toast.success(successToastContent, {
      id: 'approve-success-popup',
      duration: 10000,
    })

    return approveTx
  }

  for (let token of pool.poolTokens) {
    try {
      await handleApproval(token, token.addresses[chainId])
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
  const poolContract = await useSwapDepositContract(pool, chainId)
  console.log(`[approveAndDeposit] pool`, pool)
  console.log(`[approveAndDeposit] poolContract`, poolContract)
  console.log(`[approveAndDeposit deposit()] inputAmounts`, inputAmounts)
  let pendingPopup: any
  let successPopup: any

  pendingPopup = toast(`Starting your deposit...`, {
    id: 'deposit-in-progress-popup',
    duration: Infinity,
  })

  try {
    // get this from quote?
    console.log('[approveAndDeposit] here 1')
    console.log(
      `[approveAndDeposit] Object.values(inputAmounts)`,
      Object.values(inputAmounts)
    )
    let minToMint = await poolContract.calculateTokenAmount(
      Object.values(inputAmounts),
      true
    )
    console.log('[approveAndDeposit] here 2')
    console.log(`[approveAndDeposit] minToMint`, minToMint)
    console.log(`[approveAndDeposit] slippageSelected`, slippageSelected)
    console.log(`[approveAndDeposit] slippageCustom`, slippageCustom)
    minToMint = subtractSlippage(minToMint, slippageSelected, slippageCustom)

    const result = Array.from(Object.values(inputAmounts), (value) => value)
    console.log(`[approveAndDeposit] result`, result)
    console.log(`[approveAndDeposit] minToMint`, minToMint)

    const wethIndex = _.findIndex(
      pool.poolTokens,
      (t) => t.symbol == WETH.symbol
    )

    let spendTransactionArgs = [
      result,
      minToMint,
      Math.round(new Date().getTime() / 1000 + 60 * 10),
    ]

    //   const liquidityAmounts = pool.poolTokens.map((i) =>
    //   parseUnits(sanitizeValue(inputState[i.symbol]), i.decimals[chainId])
    // )

    // if (wethIndex >= 0) {
    //   spendTransactionArgs.push({ value: liquidityAmounts[wethIndex] })
    // }

    console.log(
      `[approveAndDeposit] spendTransactionArgs`,
      spendTransactionArgs
    )

    console.log('[approveAndDeposit] here 3')
    const spendTransaction = await poolContract.addLiquidity(
      ...spendTransactionArgs
    )
    console.log('[approveAndDeposit] here 4')

    const tx = await spendTransaction.wait()
    console.log('[approveAndDeposit] here 5')

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

    return tx
  } catch (error) {
    console.log('error from deposit: ', error)
    toast.dismiss(pendingPopup)
    txErrorHandler(error)
    return error
  }
}
