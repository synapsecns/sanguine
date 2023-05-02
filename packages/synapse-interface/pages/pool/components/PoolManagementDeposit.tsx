import _ from 'lodash'

import { WETH } from '@constants/tokens/swapMaster'
import { AVWETH, ETH, WETHE } from '@constants/tokens/master'

import { approveAndDeposit } from '@utils/actions/approveAndDeposit'
import TokenInput from '@components/TokenInput'
import PriceImpactDisplay from './PriceImpactDisplay'
import { TransactionResponse } from '@ethersproject/providers'
import { useRouter } from 'next/router'

import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { Zero } from '@ethersproject/constants'
import { Token } from '@types'
import { useState, useEffect } from 'react'
import { BigNumber } from '@ethersproject/bignumber'

const PoolManagementDeposit = ({
  pool,
  chainId,
  address,
  poolData,
  poolUserData,
}: {
  pool: Token
  chainId: number
  address: string
  poolData: any
  poolUserData: any
}) => {
  // todo store sum in here?
  const [inputValue, setInputValue] = useState<Record<string, BigNumber>>({})
  const router = useRouter()

  // TODO move this to utils
  const sumBigNumbersFromState = () => {
    let sum = Zero
    for (let key in inputValue) {
      sum = sum.add(inputValue[key])
    }
    return sum
  }

  const calculateMaxDeposits = async () => {
    if (poolUserData == null || address == null) {
      return
    }

    let inputSum = sumBigNumbersFromState()
    let depositLPTokenAmount
    if (poolData.totalLocked.gt(0) && inputSum.gt(0)) {
      depositLPTokenAmount = await swapContract.calculateTokenAmount(
        // account,
        poolTokens.map((i) =>
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

  const onChangeTokenInputValue = (tokenSymbol: string, value: BigNumber) => {
    setInputValue({ ...inputValue, [tokenSymbol]: value })
  }
  useEffect(() => {
    if (poolUserData) {
      let initInputValue = {}
      poolUserData.tokens.map((tokenObj, i) => {
        initInputValue[tokenObj.token.symbol] = undefined
      })
      setInputValue(initInputValue)
    }
  }, [poolUserData])

  // const {
  //   onChangeTokenInputValue,
  //   clearInputs,
  //   priceImpact,
  //   poolTokens,
  //   inputState,
  //   tokenInputSum,
  //   depositAmount,
  // } = useSwapPoolDeposit(poolName)
  const clearInputs = ''
  const priceImpact = Zero
  const poolTokens = []
  const inputState = {}
  const tokenInputSum = Zero
  const depositAmount = ''
  console.log('poolpoolpool', pool)
  const placeholder = async (): Promise<TransactionResponse> => {
    console.log('placeholder')
    return
  }
  console.log('poolUserData', poolUserData)
  return (
    <div className="flex-col">
      <div className="px-2 pt-1 pb-4 bg-bgLight rounded-xl">
        {pool &&
          poolUserData &&
          poolUserData.tokens.map((tokenObj, i) => {
            const balanceToken = correctToken(tokenObj.token)
            return (
              <TokenInput
                token={balanceToken}
                key={balanceToken.symbol}
                balanceStr={String(tokenObj.balanceStr)}
                inputValue={inputValue}
                onChange={(value) =>
                  onChangeTokenInputValue(tokenObj.token.symbol, value)
                }
                chainId={chainId}
                address={address}
              />
            )
          })}
      </div>
      <TransactionButton
        label="Add Liquidity"
        pendingLabel="Adding Liquidity"
        disabled={tokenInputSum.eq(0)}
        className="items-center w-full px-6 py-3 mt-6 text-md rounded-xl"
        // FIX
        // onClick={async () => {
        //   const appAndDeposit = await approveAndDeposit({
        //     slippageCustom: null,
        //     slippageSelected: 'ONE_TENTH',
        //     infiniteApproval: true,
        //     inputState,
        //     depositAmount,
        //   })
        //   // Clear input after deposit
        //   clearInputs()
        // }}
        onClick={placeholder}
      />
      <PriceImpactDisplay priceImpact={priceImpact} />
    </div>
  )
}
const correctToken = (token: Token) => {
  let balanceToken
  if (token.symbol == WETH.symbol) {
    balanceToken = ETH
  } else if (token.symbol == AVWETH.symbol) {
    // token = WETHE
    balanceToken = WETHE
  } else {
    balanceToken = token
  }
  return balanceToken
}

export default PoolManagementDeposit
