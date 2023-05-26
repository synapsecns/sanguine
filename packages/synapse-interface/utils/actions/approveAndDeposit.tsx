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

export const approve = async (
  pool: Token,
  depositQuote: any,
  inputValue: any,
  chainId: number
) => {
  const currentChainName = CHAINS_BY_ID[chainId].name
  let pendingPopup: any
  let successPopup: any

  pendingPopup = toast(`Requesting approval on ${currentChainName}`, {
    id: 'approve-in-progress-popup',
    duration: Infinity,
  })

  for (let token of pool.poolTokens) {
    const tokenAddr = token.addresses[chainId]
    if (
      (inputValue[tokenAddr] && inputValue[tokenAddr].isZero()) ||
      inputValue[tokenAddr].lt(depositQuote.allowances[tokenAddr])
    )
      continue

    if (token.symbol != WETH.symbol) {
      try {
        await approveToken(
          pool.swapAddresses[chainId],
          chainId,
          token.symbol === AVWETH.symbol
            ? WETHE.addresses[chainId]
            : token.addresses[chainId],
          inputValue[tokenAddr]
        ).then((approveTx) => {
          if (approveTx) {
            toast.dismiss(pendingPopup)

            const successToastContent = (
              <div>
                <div>Successfully approved on {currentChainName}</div>
                <ExplorerToastLink
                  transactionHash={approveTx?.hash ?? AddressZero}
                  chainId={chainId}
                />
              </div>
            )

            successPopup = toast.success(successToastContent, {
              id: 'approve-success-popup',
              duration: 10000,
            })
          }
        })
      } catch {
        toast.dismiss(pendingPopup)
      }
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
  let pendingPopup: any
  let successPopup: any

  pendingPopup = toast(`Starting your deposit...`, {
    id: 'deposit-in-progress-popup',
    duration: Infinity,
  })

  try {
    // get this from quote?
    let minToMint = await poolContract.calculateTokenAmount(
      Object.values(inputAmounts),
      true
    )
    minToMint = subtractSlippage(minToMint, slippageSelected, slippageCustom)

    const result = Array.from(Object.values(inputAmounts), (value) => value)

    let spendTransactionArgs = [
      result,
      minToMint,
      Math.round(new Date().getTime() / 1000 + 60 * 10),
    ]

    const spendTransaction = await poolContract.addLiquidity(
      ...spendTransactionArgs
    )

    const tx = await spendTransaction.wait()

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
  } catch (err) {
    toast.dismiss(pendingPopup)
    txErrorHandler(err)
  }
}
