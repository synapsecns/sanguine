import _ from 'lodash'

import { parseUnits } from '@ethersproject/units'
import toast from 'react-hot-toast'

import { useSwapDepositContract } from '@hooks/contracts/useContract'
import { useAllContracts } from '@hooks/contracts/useAllContracts'
import { usePool } from '@hooks/pools/usePools'

import { checkAndApproveTokenForTrade } from '@utils/checkAndApproveTokenForTrade'
import { subtractSlippage } from '@utils/slippage'

import ExplorerToastLink from '@components/ExplorerToastLink'

import { sanitizeValue } from '@utils/sanitizeValue'
import { txErrorHandler } from '@utils/txErrorHandler'
import { AVWETH, WETH, WETHE } from '@constants/tokens/basic'

import {Token} from '@types'
export const approveAndDeposit = (pool: Token, address: string, chainId: number) => {
  const swapContract = useSwapDepositContract(pool, chainId)




}
/**
 * @param {string} poolName
 */
export const approveAndDeposiwt = (pool, address, chainId) => {
  // Approve token
  const tokenTx = await approveToken(address, chainId, pool.)

  const swapContract = useSwapDepositContract(poolName) //? DIFF

  const tokenContracts = useAllContracts()
  const poolTokens = pool.poolTokens

  if (!poolTokens)
    throw new Error('useApproveAndDeposit requires a valid pool name')

  let poolTokenObj = {}
  for (const poolToken of poolTokens) {
    poolTokenObj[poolToken.symbol] = poolToken
  }

  const wethIndex = _.findIndex(poolTokens, (t) => t.symbol == WETH.symbol)

  return async function approveAndDeposit({
    inputState,
    depositAmount,
    infiniteApproval,
    slippageSelected,
    slippageCustom,
  }) {
    try {
      if (!account) throw new Error('Wallet must be connected')
      if (!swapContract) throw new Error('Swap contract is not loaded')
      if (!poolTokens)
        throw new Error('useApproveAndDeposit requires a valid pool name')

      async function approveSingleToken(token, spendingValue) {
        // const spendingValue = parseUnits(sanitizeValue(inputState[token.symbol]), token.decimals[chainId])

        if (spendingValue.isZero()) return
        const tokenContract = tokenContracts?.[token.symbol]
        if (tokenContract == null) return
        await checkAndApproveTokenForTrade(
          tokenContract,
          swapContract.address,
          address,
          spendingValue,
          infiniteApproval,
          {
            onTransactionStart: () => {},
            onTransactionSuccess: () => {},
            onTransactionError: () => {
              throw new Error('Your transaction could not be completed')
            },
          }
        )
      }

      const liquidityAmounts = poolTokens.map((i) =>
        parseUnits(sanitizeValue(inputState[i.symbol]), i.decimals[chainId])
      )

      for (const [token, spendingValue] of _.zip(
        poolTokens,
        liquidityAmounts
      )) {
        console.log(token) // leave log in otherwise inconsistent behavior (sometimes)
        if (token.symbol === AVWETH.symbol) {
          await approveSingleToken(WETHE, spendingValue)
        } else if (token.symbol != WETH.symbol) {
          await approveSingleToken(token, spendingValue)
        }
      }

      // if (pool is empty, minToMint should be 0. else calculate normally)
      // let minToMint = 0

      let minToMint = await swapContract.calculateTokenAmount(
        liquidityAmounts,
        true
      )

      minToMint = subtractSlippage(minToMint, slippageSelected, slippageCustom)

      toast('Starting your deposit...')

      let spendTransactionArgs = [
        liquidityAmounts,
        minToMint,
        Math.round(new Date().getTime() / 1000 + 60 * 10),
      ]
      if (wethIndex >= 0) {
        spendTransactionArgs.push({ value: liquidityAmounts[wethIndex] })
      }

      const spendTransaction = await swapContract.addLiquidity(
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
}


export async function checkAndApproveTokenForTrade(
  srcTokenContract,
  swapAddress,
  spenderAddress,
  spendingValue, // max is MaxUint256
  infiniteApproval = false,
  callbacks={}
) {
  if (srcTokenContract == null) return
  if (spendingValue.eq(0)) return

  const [tokenName, existingAllowance] = await Promise.all([
    srcTokenContract.name(),
    srcTokenContract.allowance(spenderAddress, swapAddress)
  ])

  console.debug(
    `Existing ${tokenName} Allowance: ${existingAllowance.toString()}`,
  )
  if (existingAllowance.gte(spendingValue)) return
  async function approve(amount) {
    try {
      const cleanupOnStart = callbacks.onTransactionStart?.()
      const approvalTransaction = await srcTokenContract.approve(
        swapAddress,
        amount,
      )
      const confirmedTransaction = await approvalTransaction.wait()
      cleanupOnStart?.()
      callbacks.onTransactionSuccess?.(confirmedTransaction)
    } catch (error) {
      callbacks.onTransactionError?.(error)
      throw error
    }
  }
  if (existingAllowance.gt('0')) {
    // Reset to 0 before updating approval
    await approve(Zero)
  }
  await approve(infiniteApproval ? MaxUint256 : spendingValue)
  console.debug(`Approving ${tokenName} spend of ${spendingValue.toString()}`)
}
