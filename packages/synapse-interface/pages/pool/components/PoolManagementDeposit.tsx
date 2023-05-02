import _ from 'lodash'

import { WETH } from '@constants/tokens/swapMaster'
import { AVWETH, ETH, WETHE } from '@constants/tokens/master'
import { stringToBigNum } from '@/utils/stringToBigNum'
import { getAddress } from '@ethersproject/address'
import { approveAndDeposit } from '@utils/actions/approveAndDeposit'
import TokenInput from '@components/TokenInput'
import PriceImpactDisplay from './PriceImpactDisplay'
import { TransactionResponse } from '@ethersproject/providers'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { formatBNToString } from '@utils/bignumber/format'
import { calculatePriceImpact } from '@utils/priceImpact'
import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { Zero, One } from '@ethersproject/constants'
import { Token } from '@types'
import { useState, useEffect } from 'react'
import { BigNumber } from '@ethersproject/bignumber'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
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
  const [inputValue, setInputValue] = useState<{
    bn: Record<string, BigNumber>
    str: Record<string, string>
  }>({ bn: {}, str: {} })
  const [depositQuote, setDepositQuote] = useState<{
    priceImpact: BigNumber
    allowances: Record<string, BigNumber>
    routerAddress: string
  }>({ priceImpact: undefined, allowances: {}, routerAddress: '' })

  const SynapseSDK = useSynapseContext()

  // TODO move this to utils
  const sumBigNumbersFromState = () => {
    let sum = Zero
    pool.poolTokens.map((token) => {
      if (inputValue.bn[getAddress(token.addresses[chainId])]) {
        sum = sum.add(
          inputValue.bn[getAddress(token.addresses[chainId])].mul(
            BigNumber.from(10).pow(18 - token.decimals[chainId])
          )
        )
      }
    })
    console.log('sumBigNumbersFromState', sum.toString())
    return sum
  }

  const calculateMaxDeposits = async () => {
    if (poolUserData == null || address == null) {
      return
    }

    let inputSum = sumBigNumbersFromState()
    console.log(
      poolData,
      poolData.totalLocked,
      inputSum.toString(),
      poolData.totalLocked.toString()
    )
    if (poolData.totalLocked.gt(0) && inputSum.gt(0)) {
      console.log(
        'iii',
        chainId,
        pool.swapAddresses[chainId],
        inputValue.bn,
        Object.values(inputValue.bn).map((v) => v.toString())
      )
      const { amount, routerAddress } = await SynapseSDK.calculateAddLiquidity(
        chainId,
        pool.swapAddresses[chainId],
        inputValue.bn
      )

      let allowances: Record<string, BigNumber> = {}
      for (const [key, value] of Object.entries(inputValue.bn)) {
        allowances[key] = await getTokenAllowance(
          routerAddress,
          key,
          address,
          chainId
        )
      }

      const priceImpact = calculateExchangeRate(
        inputSum,
        18,
        inputSum.sub(amount),
        18
      )
      // TODO: DOUBLE CHECK THIS
      setDepositQuote({
        priceImpact,
        allowances,
        routerAddress: pool.swapAddresses[chainId],
      })
    }
  }

  useEffect(() => {
    calculateMaxDeposits()
  }, [inputValue])

  const onChangeTokenInputValue = (token: Token, value: string) => {
    const bigNum = stringToBigNum(value, token.decimals[chainId]) ?? Zero
    if (chainId && token) {
      console.log
      setInputValue({
        bn: {
          ...inputValue.bn,
          [getAddress(token.addresses[chainId])]: bigNum,
        },
        str: {
          ...inputValue.str,
          [getAddress(token.addresses[chainId])]: value,
        },
      })
    }
  }

  useEffect(() => {
    if (poolData && poolUserData) {
      let initInputValue: {
        bn: Record<string, BigNumber>
        str: Record<string, string>
      } = { bn: {}, str: {} }
      poolUserData.tokens.map((tokenObj, i) => {
        initInputValue.bn[getAddress(tokenObj.token.addresses[chainId])] = Zero
        initInputValue.str[getAddress(tokenObj.token.addresses[chainId])] =
          undefined
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
                  onChangeTokenInputValue(balanceToken, value)
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
      <PriceImpactDisplay priceImpact={depositQuote.priceImpact} />
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
