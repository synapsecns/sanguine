import _ from 'lodash'

import { parseUnits } from '@ethersproject/units'
import toast from 'react-hot-toast'

import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useBlockNumber } from '@hooks/useBlockNumber'
import { useSwapDepositContract } from '@hooks/contracts/useContract'
import { useAllContracts } from '@hooks/contracts/useAllContracts'
import { usePool } from '@hooks/pools/usePools'

import { checkAndApproveTokenForTrade } from '@utils/checkAndApproveTokenForTrade'
import { subtractSlippage } from '@utils/slippage'

import ExplorerToastLink from '@components/ExplorerToastLink'

import { sanitizeValue } from '@utils/sanitizeValue'
import { txErrorHandler } from '@utils/txErrorHandler'
import { AVWETH, WETH, WETHE } from '@constants/tokens/basic'

/**
 * @param {string} poolName
 */
export const approveAndDeposit = (poolName) => {
  const swapContract = useSwapDepositContract(poolName) //? DIFF

  const tokenContracts = useAllContracts()
  const { account, chainId } = useActiveWeb3React()
  const poolTokens = usePool(poolName)
  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)

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
          account,
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
      setBlockNumber(tx.blockNumber)

      return tx
    } catch (err) {
      txErrorHandler(err)
    }
  }
}
