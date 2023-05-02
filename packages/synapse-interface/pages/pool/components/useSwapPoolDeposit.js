import _ from 'lodash'
import { useState, useEffect } from 'react'

import { Zero } from '@ethersproject/constants'
import { parseUnits } from '@ethersproject/units'

import { sanitizeValue } from '@utils/sanitizeValue'

import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'
import { usePoolData } from '@hooks/pools/usePoolData'

import { useSwapDepositContract } from '@hooks/contracts/useContract'

import { usePoolTokenInfo } from '@hooks/pools/usePools'


import { calculatePriceImpact } from '@utils/priceImpact'




export function useSwapPoolDeposit(poolName) {
  const { account, chainId } = useActiveWeb3React()
  const { poolTokens } = usePoolTokenInfo(poolName)

  let defaultInputState = {}
  for (const token of poolTokens) {
    defaultInputState[token.symbol] = ""
  }
  const [inputState, setInputState] = useState(defaultInputState)
  const [depositAmount, setDepositAmount] = useState(Zero)

  const [priceImpact, setPriceImpact] = useState(Zero)

  const swapContract = useSwapDepositContract(poolName)

  const [poolData] = usePoolData(poolName)


  let sanitizedInputState = {}
  for (const [symbol, inputStr] of _.toPairs(inputState)) {
    sanitizedInputState[symbol] = sanitizeValue(inputStr)
  }

  let numericalInputState = {}
  for (const [symbol, inputStr] of _.toPairs(inputState)) {
    numericalInputState[symbol] = Number(sanitizeValue(inputStr))
  }

  let tokenInputSum = 0
  for (const poolToken of poolTokens) {
    tokenInputSum += numericalInputState[poolToken.symbol]
  }
  tokenInputSum = parseUnits(`${tokenInputSum}`)

  /**
   * evaluate if a new deposit will exceed the pool's per-user limit
   * */
  async function calculateMaxDeposits() {
    if (swapContract == null || poolData == null || account == null) {
      return
    }

    let depositLPTokenAmount
    if (poolData.totalLocked.gt(0) && tokenInputSum.gt(0)) {
      depositLPTokenAmount = await swapContract.calculateTokenAmount(
        // account,
        poolTokens.map(i =>
          parseUnits(sanitizedInputState[i.symbol], i.decimals[chainId])
        ),
        true // deposit boolean
      )
    } else {
      // when pool is empty, estimate the lptokens by just summing the input instead of calling contract
      depositLPTokenAmount = tokenInputSum
    }
    const calcedPriceImpact = calculatePriceImpact(
      tokenInputSum,
      depositLPTokenAmount,
      poolData.virtualPrice
    )

    setDepositAmount(depositLPTokenAmount)
    setPriceImpact(calcedPriceImpact)
  }

  useEffect(
    () => {
      // if (poolData.virtualPrice) {
        calculateMaxDeposits()
      // }
    },
    [inputState, swapContract, account]
  )


  function onChangeTokenInputValue(symbol, value) {
    setInputState({ ...inputState, [symbol]: value })
  }


  function clearInputs() {
    setInputState(defaultInputState)
  }


  return {
    onChangeTokenInputValue,
    clearInputs,
    priceImpact,
    poolTokens,
    inputState,
    tokenInputSum,
    depositAmount
  }
}

