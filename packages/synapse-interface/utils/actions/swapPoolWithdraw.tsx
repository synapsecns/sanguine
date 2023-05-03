import _ from 'lodash'
import { useState, useEffect, useMemo } from 'react'

import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import { formatUnits, parseUnits } from '@ethersproject/units'

import { sanitizeValue } from '@utils/sanitizeValue'

import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

import { usePoolData } from '@hooks/pools/usePoolData'
import { usePoolTokenInfo } from '@hooks/pools/usePools'

import { useSwapDepositContract } from '@hooks/contracts/useContract'

import { calculatePriceImpactWithdraw } from '@utils/priceImpact'

import { ALL } from '@constants/withdrawTypes'

export function useSwapPoolWithdraw(poolName) {
  const { account, chainId } = useActiveWeb3React()
  const { poolTokens, depositTokens, decimals } = usePoolTokenInfo(poolName)
  const [lpTokenValue, setLpTokenValue] = useState('')
  const lpTokenAmount = useMemo(
    () => parseUnits(sanitizeValue(lpTokenValue), decimals),
    [lpTokenValue]
  )
  let defaultInputState = {}
  for (const token of depositTokens) {
    defaultInputState[token.symbol] = ''
  }
  const [inputState, setInputState] = useState(defaultInputState)
  let defaultWithdrawAmountState = {}
  for (const token of depositTokens) {
    defaultWithdrawAmountState[token.symbol] = Zero
  }
  const [withdrawAmount, setWithdrawAmount] = useState(Zero)

  const [priceImpact, setPriceImpact] = useState(Zero)

  const swapContract = useSwapDepositContract(poolName)
  const [withdrawType, setWithdrawType] = useState(ALL)
  const [percentage, setPercentage] = useState(100)
  const [poolData, userShareData] = usePoolData(poolName)

  let sanitizedInputState = {}
  for (const [symbol, inputStr] of _.toPairs(inputState)) {
    sanitizedInputState[symbol] = sanitizeValue(inputStr)
  }

  let numericalInputState = {}
  for (const [symbol, inputStr] of _.toPairs(inputState)) {
    numericalInputState[symbol] = Number(sanitizeValue(inputStr))
  }

  let tokenInputSum = Zero
  for (const poolToken of depositTokens) {
    tokenInputSum = tokenInputSum.add(
      parseUnits(sanitizedInputState[poolToken.symbol])
    )
  }

  /**
   * evaluate if a new deposit will exceed the pool's per-user limit
   * */
  async function calculateMaxWithdraw() {
    if (swapContract == null || poolData == null || account == null) {
      return
    }

    const calcedPriceImpact = calculatePriceImpactWithdraw(
      lpTokenAmount,
      tokenInputSum,
      poolData.virtualPrice
    )

    setWithdrawAmount(lpTokenAmount)
    setPriceImpact(calcedPriceImpact)
  }

  useEffect(
    () => {
      calculateMaxWithdraw()
    },
    [inputState, swapContract, account] // poolData
  )

  /**
   * evaluate if a new withdraw will exceed the pool's per-user limit
   */
  async function calculateWithdrawBonus() {
    if (swapContract == null || poolData == null || account == null) {
      return
    }
    const newInputState = {}
    if (withdrawType == ALL) {
      const results = await swapContract.calculateRemoveLiquidity(lpTokenAmount)
      for (const [t, amt] of _.zip(depositTokens, results)) {
        newInputState[t.symbol] = formatUnits(amt, t.decimals[chainId])
      }
    } else {
      for (const t of depositTokens) {
        newInputState[t.symbol] = ''
      }
      const tokenIndex = depositTokens.findIndex(
        (i) => i.symbol === withdrawType
      )
      const token = depositTokens[tokenIndex]
      const amt = await swapContract.calculateRemoveLiquidityOneToken(
        lpTokenAmount,
        tokenIndex
      )
      newInputState[token.symbol] = formatUnits(amt, token.decimals[chainId])
    }
    setInputState(newInputState)
  }

  useEffect(
    () => {
      calculateWithdrawBonus()
    },
    [lpTokenAmount, swapContract, account, withdrawType]
    //, withdrawType
  )

  function onChangeTokenInputValue(symbol, value) {
    setInputState({ ...inputState, [symbol]: value })
  }

  function clearInputs() {
    setInputState(defaultInputState)
  }

  return {
    withdrawType,
    setWithdrawType,
    setInputState,
    onChangeTokenInputValue,
    clearInputs,
    priceImpact,

    depositTokens,
    poolTokens,
    inputState,
    tokenInputSum,

    poolData,
    percentage,
    setPercentage,

    lpTokenValue,
    setLpTokenValue,
    lpTokenAmount,
  }
}

// {
//   cardNav === 'addLiquidity' && (
//     <PoolManagementDeposit

//     />
//   )
// }
// {
//   cardNav === 'removeLiquidity' && (
//     <PoolManagementWithdraw
//       onFormChange={updateWithdrawFormState}
//       tokens={withdrawTokens}
//       formStateData={withdrawFormState}
//       onConfirmTransaction={onConfirmWithdrawTransaction}
//       priceImpact={estWithdrawBonus}
//       poolData={poolData}
//     />
//   )
// }
