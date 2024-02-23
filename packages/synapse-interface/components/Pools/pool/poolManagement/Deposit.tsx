import _ from 'lodash'
import { useState, useEffect } from 'react'
import { useDispatch } from 'react-redux'
import { getAddress } from '@ethersproject/address'
import { Address, fetchBalance } from '@wagmi/core'
import { zeroAddress } from 'viem'

import {
  usePoolDataState,
  usePoolUserDataState,
  usePoolDepositState,
} from '@/slices/pool/hooks'
import { fetchPoolUserData } from '@/slices/poolUserDataSlice'
import {
  resetPoolDeposit,
  setDepositQuote,
  setInputValue,
  setIsLoading,
  setPool,
} from '@/slices/poolDepositSlice'

import { swapPoolCalculateAddLiquidity } from '@/actions/swapPoolCalculateAddLiquidity'

import { ETH, WETHE, WETH } from '@/constants/tokens/bridgeable'
import { AVWETH } from '@/constants/tokens/auxilliary'
import type { Token } from '@types'

import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import {
  approve,
  deposit,
  emptyPoolDeposit,
} from '@/utils/actions/approveAndDeposit'

import { stringToBigInt, formatBigIntToString } from '@/utils/bigint/format'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { calculatePriceImpact } from '@/utils/priceImpact'
import { transformCalculateLiquidityInput } from '@/utils/transformCalculateLiquidityInput'


import LoadingTokenInput from '@/components/loading/LoadingTokenInput'
import { DepositTokenInput } from '@/components/TokenInput'
import PriceImpactDisplay from '../components/PriceImpactDisplay'

import DepositButton from './DepositButton'


export const DEFAULT_DEPOSIT_QUOTE = {
  priceImpact: 0n,
  allowances: {},
  routerAddress: '',
}

export default function Deposit({
  chainId,
  address,
}: {
  chainId: number
  address: string
}) {
  const dispatch: any = useDispatch()

  const { pool, poolData } = usePoolDataState()
  const { poolUserData } = usePoolUserDataState()
  const { depositQuote, inputValue, inputSum, filteredInputValue } = usePoolDepositState()

  const { poolAddress } = getSwapDepositContractFields(pool, chainId)

  dispatch(setPool(pool))

  const calculateMaxDeposits = async () => {
    try {
      if (poolUserData == null || address == null) {
        return
      }
      dispatch(setIsLoading(true))
      const { totalLocked, virtualPrice } = poolData

      if (totalLocked > 0 && inputSum > 0n) {
        const input = transformCalculateLiquidityInput(
          chainId,
          pool,
          filteredInputValue.bi
        )

        const inputs: bigint[] = []

        Object.keys(input).forEach((key) => inputs.push(input[key]))

        const amount = await swapPoolCalculateAddLiquidity({
          chainId,
          pool,
          inputs,
        })

        let allowances: Record<string, bigint> = {}
        for (const [tokenAddress, value] of Object.entries(
          filteredInputValue.bi
        )) {
          allowances[tokenAddress] = await getTokenAllowance(
            poolAddress,
            tokenAddress as Address,
            address as Address,
            chainId
          )
        }

        const priceImpact = calculatePriceImpact(
          inputSum,
          amount as bigint,
          virtualPrice
        )

        dispatch(
          setDepositQuote({
            priceImpact: priceImpact,
            allowances,
            routerAddress: poolAddress,
          })
        )
        dispatch(setIsLoading(false))
      } else {
        dispatch(setDepositQuote(DEFAULT_DEPOSIT_QUOTE))
        dispatch(setIsLoading(false))
      }
    } catch (e) {
      dispatch(setIsLoading(false))
      console.log(e)
    }
  }

  useEffect(() => {
    calculateMaxDeposits()
  }, [inputValue, filteredInputValue, pool, chainId, address])

  const onChangeInputValue = (token: Token, value: string) => {
    const bigInt = stringToBigInt(value, token.decimals[chainId]) ?? 0n
    if (chainId && token) {
      const tokenAddr = getAddress(token.addresses[chainId])
      dispatch(
        setInputValue({
          bi: {
            ...inputValue.bi,
            [tokenAddr]: bigInt,
          },
          str: {
            ...inputValue.str,
            [tokenAddr]: value,
          },
        })
      )
    }
  }

  useEffect(() => {
    if (poolData && poolUserData && pool && chainId && address) {
      resetInputs()
    }
  }, [poolUserData])

  const resetInputs = () => {
    dispatch(resetPoolDeposit())
  }

  const approveTxn = async () => {
    try {
      const tx = approve(pool, depositQuote, inputValue.bi, chainId)
      try {
        await tx
        calculateMaxDeposits()
      } catch (error) {
        return txErrorHandler
      }
    } catch (error) {
      return txErrorHandler(error)
    }
  }

  const depositTxn = async () => {
    try {
      let tx

      if (poolData.totalLocked === 0) {
        tx = emptyPoolDeposit(pool, filteredInputValue.bi, chainId)
      } else {
        tx = deposit(pool, 'ONE_TENTH', null, filteredInputValue.bi, chainId)
      }

      try {
        await tx
        dispatch(fetchPoolUserData({ pool, address: address as Address }))
        dispatch(resetPoolDeposit())
      } catch (error) {
        txErrorHandler(error)
      }
    } catch (error) {
      txErrorHandler(error)
    }
  }

  return (
    <div className="flex-col">
      <div className="mb-4">
        {pool && poolUserData.tokens && poolData ? (
          poolUserData.tokens.map((tokenObj, i) => {
            return (
              <div className='py-0.5'>
                <SerializedDepositInput
                  key={i}
                  tokenObj={tokenObj}
                  pool={pool}
                  address={address}
                  chainId={chainId}
                  inputValue={inputValue}
                  onChangeInputValue={onChangeInputValue}
                />
              </div>
            )
          })
        ) : (
          <>
            <LoadingTokenInput />
            <LoadingTokenInput />
          </>
        )}
      </div>
      <DepositButton approveTxn={approveTxn} depositTxn={depositTxn} />
      {depositQuote.priceImpact &&
        depositQuote.priceImpact !== 0n &&
        inputSum !== 0n && (
          <PriceImpactDisplay priceImpact={depositQuote.priceImpact} />
        )}
    </div>
  )
}

const SerializedDepositInput = ({
  pool,
  tokenObj,
  address,
  chainId,
  inputValue,
  onChangeInputValue,
}) => {
  const [serializedToken, setSerializedToken] = useState(undefined)
  const balanceToken = getBalanceToken(tokenObj.token, pool)

  useEffect(() => {
    const fetchSerializedData = async () => {
      try {
        const token = await serializeToken(
          address,
          chainId,
          balanceToken,
          tokenObj
        )
        setSerializedToken(token)
      } catch (error) {
        console.log(`error`, error)
      }
    }

    fetchSerializedData()
  }, [])

  return (
    serializedToken && (
      <DepositTokenInput
        token={serializedToken}
        key={serializedToken.symbol}
        rawBalance={serializedToken.rawBalance}
        balanceStr={formatBigIntToString(
          serializedToken.rawBalance,
          serializedToken.decimals[chainId],
          5
        )}
        inputValueStr={inputValue.str[serializedToken.addresses[chainId]]}
        onChange={(value) => onChangeInputValue(serializedToken, value)}
        chainId={chainId}
        address={address}
      />
    )
  )
}

const getBalanceToken = (token: Token, pool: Token) => {
  let balanceToken: Token | undefined
  if (token.symbol == WETH.symbol && !pool.nativeTokens.includes(WETH)) {
    balanceToken = ETH
  } else if (token.symbol == AVWETH.symbol) {
    balanceToken = WETHE
  } else {
    balanceToken = token
  }
  return balanceToken
}

const serializeToken = async (
  address: string,
  chainId: number,
  balanceToken: Token,
  tokenObj: any
) => {
  let fetchedBalance

  if (balanceToken.addresses[chainId] === zeroAddress) {
    fetchedBalance = await fetchBalance({
      address: address as Address,
      chainId,
    })

    return {
      ...balanceToken,
      rawBalance: fetchedBalance.value,
      balanceStr: formatBigIntToString(
        fetchedBalance.value,
        balanceToken.decimals[chainId],
        5
      ),
    }
  } else if (balanceToken === WETHE) {
    fetchedBalance = await fetchBalance({
      address: address as Address,
      chainId,
      token: balanceToken.addresses[chainId] as Address,
    })

    return {
      ...balanceToken,
      rawBalance: fetchedBalance.value,
      balanceStr: fetchedBalance.formatted,
    }
  } else {
    return {
      ...balanceToken,
      rawBalance: tokenObj.rawBalance,
      balanceStr: tokenObj.balanceStr,
    }
  }
}

