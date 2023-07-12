import _ from 'lodash'

import { WETH } from '@constants/tokens/swapMaster'
import { AVWETH, ETH, WETHE } from '@constants/tokens/master'
import { stringToBigInt } from '@/utils/stringToBigNum'
import { DepositTokenInput } from '@components/TokenInput'
import PriceImpactDisplay from '../components/PriceImpactDisplay'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { Token } from '@types'
import { useState, useEffect, useMemo } from 'react'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { approve, deposit } from '@/utils/actions/approveAndDeposit'
import LoadingTokenInput from '@components/loading/LoadingTokenInput'
import { Address, fetchBalance } from '@wagmi/core'
import { formatBNToString } from '@/utils/bignumber/format'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { calculatePriceImpact } from '@/utils/priceImpact'
import { transformCalculateLiquidityInput } from '@/utils/transformCalculateLiquidityInput'
import { formatBigIntToString } from '@/utils/bigint/format'

import { getAddress } from '@ethersproject/address'
import { BigNumber } from '@ethersproject/bignumber'
import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'

import {
  resetPoolDeposit,
  setDepositQuote,
  setInputValue,
  setIsLoading,
  setPool,
} from '@/slices/poolDepositSlice'

import { useDispatch } from 'react-redux'
import DepositButton from './DepositButton'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { fetchPoolUserData } from '@/slices/poolUserDataSlice'

export const DEFAULT_DEPOSIT_QUOTE = {
  priceImpact: 0n,
  allowances: {},
  routerAddress: '',
}

const Deposit = ({
  chainId,
  address,
}: {
  chainId: number
  address: string
}) => {
  const { synapseSDK } = useSynapseContext()
  const dispatch: any = useDispatch()

  const { pool, poolData } = useSelector((state: RootState) => state.poolData)
  const { poolUserData } = useSelector((state: RootState) => state.poolUserData)
  const { depositQuote, inputValue, inputSum, filteredInputValue } =
    useSelector((state: RootState) => state.poolDeposit)

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

        // NOTE: Doing this as SDK requires BigNumber due to comparisons in
        // calculateAddLiquidity()
        // We can't fully remove BigNumber here until SDK supports
        let convertedInput = {}
        for (let key in input) {
          convertedInput[key] = BigNumber.from(input[key].toString())
        }

        const { amount } = await synapseSDK.calculateAddLiquidity(
          chainId,
          pool.swapAddresses[chainId],
          convertedInput
        )

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
          BigInt(amount),
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
      const tx = deposit(
        pool,
        'ONE_TENTH',
        null,
        filteredInputValue.bi,
        chainId
      )

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
      <div className="px-2 pt-1 pb-4 bg-bgLight rounded-xl">
        {pool && poolUserData.tokens && poolData ? (
          poolUserData.tokens.map((tokenObj, i) => {
            return (
              <SerializedDepositInput
                key={i}
                tokenObj={tokenObj}
                address={address}
                chainId={chainId}
                inputValue={inputValue}
                onChangeInputValue={onChangeInputValue}
              />
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
  tokenObj,
  address,
  chainId,
  inputValue,
  onChangeInputValue,
}) => {
  const [serializedToken, setSerializedToken] = useState(undefined)
  const balanceToken = correctToken(tokenObj.token)

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
          4
        )}
        inputValueStr={inputValue.str[serializedToken.addresses[chainId]]}
        onChange={(value) => onChangeInputValue(serializedToken, value)}
        chainId={chainId}
        address={address}
      />
    )
  )
}

const correctToken = (token: Token) => {
  let balanceToken: Token | undefined
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

const serializeToken = async (
  address: string,
  chainId: number,
  balanceToken: Token,
  tokenObj: any
) => {
  let fetchedBalance

  if (balanceToken === ETH) {
    fetchedBalance = await fetchBalance({
      address: address as Address,
      chainId,
    })

    return {
      ...balanceToken,
      rawBalance: fetchedBalance.value,
      balanceStr: formatBNToString(
        fetchedBalance.value,
        balanceToken.decimals[chainId],
        4
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

export default Deposit
