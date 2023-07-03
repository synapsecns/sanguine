import { Contract } from 'ethers'
import { MAX_UINT256 } from './bigint/format'

type TransactionCallbackFunctions = {
  onTransactionStart?: () => any
  onTransactionSuccess?: (confirmedTransaction: any) => void
  onTransactionError?: (error: Error) => any
}

/**
 * Checks if a spender is allowed to spend some amount of a token.
 * Approves them to spend if they're not already allowed.
 * Won't make requests if spendingValue eq 0
 * @param {Contract} srcTokenContract
 * @param {string} swapAddress
 * @param {string} spenderAddress
 * @param {bigint} spendingValue
 * @param {boolean} [infiniteApproval]
 * @param {{
 *  onTransactionStart:   function,
 *  onTransactionSuccess: function,
 *  onTransactionError:   function
 * }} callbacks
 */

export async function checkAndApproveTokenForTrade(
  srcTokenContract: Contract,
  swapAddress: string,
  spenderAddress: string,
  spendingValue: bigint,
  infiniteApproval: boolean = false,
  callbacks: TransactionCallbackFunctions = {}
) {
  if (srcTokenContract == null) return
  if (spendingValue === 0n) return

  const [tokenName, existingAllowance] = await Promise.all([
    srcTokenContract.name(),
    srcTokenContract.allowance(spenderAddress, swapAddress),
  ])

  console.debug(
    `Existing ${tokenName} Allowance: ${existingAllowance.toString()}`
  )
  if (existingAllowance.gte(spendingValue)) return
  async function approve(amount) {
    try {
      const cleanupOnStart = callbacks.onTransactionStart?.()
      const approvalTransaction = await srcTokenContract.approve(
        swapAddress,
        amount
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
    await approve(0n)
  }

  // todo: create max uint256 bigint const
  await approve(infiniteApproval ? MAX_UINT256 : spendingValue)
  console.debug(`Approving ${tokenName} spend of ${spendingValue.toString()}`)
}
