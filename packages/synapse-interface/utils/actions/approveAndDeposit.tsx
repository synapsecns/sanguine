import _ from 'lodash'

import toast from 'react-hot-toast'

import { useSwapDepositContract } from '@hooks/useSwapDepositContract'
import { subtractSlippage } from '@utils/slippage'

import ExplorerToastLink from '@components/ExplorerToastLink'

import { txErrorHandler } from '@utils/txErrorHandler'
import { AVWETH, WETHE } from '@constants/tokens/master'
import { WETH } from '@constants/tokens/swapMaster'
import { approveToken } from '@utils/approveToken'
import { Token } from '@types'
import { getAddress } from '@ethersproject/address'

export const approve = async (
  pool: Token,
  depositQuote: any,
  inputValue: any,
  chainId: number
) => {
  const poolContract = await useSwapDepositContract(pool, chainId)

  for (let token of pool.poolTokens) {
    const tokenAddr = getAddress(token.addresses[chainId])
    if (
      inputValue[tokenAddr].isZero() ||
      inputValue[tokenAddr].lt(depositQuote.allowances[tokenAddr])
    )
      continue

    if (token.symbol != WETH.symbol) {
      await approveToken(
        poolContract.address,
        chainId,
        token.symbol === AVWETH.symbol
          ? WETHE.addresses[chainId]
          : token.addresses[chainId],
        inputValue[tokenAddr]
      )
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
  try {
    // get this from quote?
    let minToMint = await poolContract.calculateTokenAmount(
      Object.values(inputAmounts),
      true
    )
    minToMint = subtractSlippage(minToMint, slippageSelected, slippageCustom)

    toast('Starting your deposit...')

    let spendTransactionArgs = [
      Object.values(inputAmounts),
      minToMint,
      Math.round(new Date().getTime() / 1000 + 60 * 10),
    ]

    const spendTransaction = await poolContract.addLiquidity(
      ...spendTransactionArgs
    )

    const tx = await spendTransaction.wait()
    const toastContent = (
      <div>
        <div>Liquidity added!</div>
        <ExplorerToastLink {...tx} chainId={chainId} />
      </div>
    )

    toast.success(toastContent)

    return tx
  } catch (err) {
    txErrorHandler(err)
  }
}
