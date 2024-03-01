import _ from 'lodash'
import { useEffect, useMemo, useState } from 'react'
import Slider from 'react-input-slider'
import { Address, waitForTransaction } from '@wagmi/core'
import { useDispatch, useSelector } from 'react-redux'
import { Token } from '@types'
import { RootState } from '@/store/store'

import Grid from '@tw/Grid'
import { getCoinTextColorCombined } from '@styles/tokens'
import { ALL } from '@constants/withdrawTypes'
import { WithdrawTokenInput } from '@components/TokenInput'
import RadioButton from '@components/buttons/RadioButton'
import ReceivedTokenSection from '../components/ReceivedTokenSection'
import PriceImpactDisplay from '../components/PriceImpactDisplay'

import { approve, withdraw } from '@/utils/actions/approveAndWithdraw'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { calculatePriceImpact } from '@/utils/priceImpact'
import { formatBigIntToString } from '@/utils/bigint/format'
import { stringToBigInt } from '@/utils/bigint/format'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { isTransactionReceiptError } from '@/utils/isTransactionReceiptError'

import {
  setInputValue,
  setWithdrawQuote,
  setWithdrawType,
  setIsLoading,
  resetPoolWithdraw,
} from '@/slices/poolWithdrawSlice'
import { fetchPoolUserData } from '@/slices/poolUserDataSlice'
import { fetchPoolData } from '@/slices/poolDataSlice'

import WithdrawButton from './WithdrawButton'

const Withdraw = ({ address }: { address: string }) => {
  const [percentage, setPercentage] = useState(0)
  const { pool, poolData } = useSelector((state: RootState) => state.poolData)
  const { poolUserData } = useSelector((state: RootState) => state.poolUserData)
  const { withdrawQuote, inputValue, withdrawType } = useSelector(
    (state: RootState) => state.poolWithdraw
  )
  const chainId = pool?.chainId
  const poolDecimals = pool?.decimals[pool?.chainId]
  const { poolAddress } = getSwapDepositContractFields(pool, chainId)
  const { synapseSDK } = useSynapseContext()

  const dispatch: any = useDispatch()

  // An ETH swap pool has nativeTokens vs. most other pools have poolTokens
  const poolSpecificTokens = pool ? pool.nativeTokens ?? pool.poolTokens : []

  const isApproved = useMemo(() => {
    return (
      withdrawQuote?.allowance &&
      stringToBigInt(inputValue, poolDecimals) <= withdrawQuote.allowance
    )
  }, [inputValue, withdrawQuote])

  const calculateMaxWithdraw = async () => {
    if (poolUserData === null || address === null) {
      return
    }
    dispatch(setIsLoading(true))
    try {
      const outputs: Record<
        string,
        {
          value: bigint
          index: number
        }
      > = {}

      const { virtualPrice } = poolData

      if (withdrawType === ALL) {
        const { amounts } = await synapseSDK.calculateRemoveLiquidity(
          chainId,
          poolAddress,
          stringToBigInt(inputValue, poolDecimals)
        )

        outputs[withdrawType] = amounts.map(transformAmount)
      } else {
        const { amount } = await synapseSDK.calculateRemoveLiquidityOne(
          chainId,
          poolAddress,
          stringToBigInt(inputValue, poolDecimals),
          Number(withdrawType)
        )

        outputs[withdrawType] = transformAmount(amount)
      }

      console.log(`outputs`, outputs)

      const outputTokensSum = sumBigInts(pool, outputs, withdrawType)

      const priceImpact = calculatePriceImpact(
        stringToBigInt(inputValue, poolDecimals),
        outputTokensSum,
        virtualPrice,
        true
      )

      const allowance = await getTokenAllowance(
        poolAddress,
        pool.addresses[chainId] as Address,
        address as Address,
        chainId
      )
      console.log(`allowance`, allowance)
      dispatch(
        setWithdrawQuote({
          priceImpact,
          allowance,
          outputs,
        })
      )
      dispatch(setIsLoading(false))
    } catch (e) {
      dispatch(setIsLoading(false))
      console.log(e)
    }
  }

  useEffect(() => {
    if (
      poolUserData &&
      poolData &&
      address &&
      pool &&
      stringToBigInt(inputValue, poolDecimals) > 0n
    ) {
      calculateMaxWithdraw()
    }
  }, [inputValue, withdrawType])

  const onPercentChange = (percent: number) => {
    if (percent > 100) {
      percent = 100
    }
    setPercentage(percent)
    const numericalOut: string = poolUserData.lpTokenBalance
      ? formatBigIntToString(
          (poolUserData.lpTokenBalance * BigInt(percent)) / BigInt(100),
          poolDecimals
        )
      : ''

    dispatch(setInputValue(numericalOut))
  }

  const onChangeInputValue = (token: Token, value: string) => {
    const bigInt = stringToBigInt(value, token.decimals[chainId])

    if (poolUserData.lpTokenBalance === 0n) {
      dispatch(setInputValue(value))

      setPercentage(0)
      return
    }
    const pn = bigInt
      ? Number(
          (bigInt * BigInt(100)) /
            BigInt(poolUserData.lpTokenBalance.toString())
        )
      : 0

    dispatch(setInputValue(value))

    if (pn > 100) {
      setPercentage(100)
    } else {
      setPercentage(pn)
    }
  }

  const approveTxn = async () => {
    try {
      const tx = approve(
        pool,
        withdrawQuote,
        stringToBigInt(inputValue, poolDecimals),
        chainId
      )

      try {
        await tx
        calculateMaxWithdraw()
      } catch (error) {
        txErrorHandler(error)
      }
    } catch (error) {
      txErrorHandler(error)
    }
  }

  const withdrawTxn = async () => {
    try {
      const tx = withdraw(
        pool,
        'ONE_TENTH',
        null,
        stringToBigInt(inputValue, poolDecimals),
        chainId,
        withdrawType,
        withdrawQuote.outputs
      )

      const resolvedTx = await tx

      const transactionReceipt = await waitForTransaction({
        hash: resolvedTx.transactionHash as Address,
        timeout: 60_000,
      })

      console.log('transactionReceipt:', transactionReceipt)

      dispatch(fetchPoolUserData({ pool, address: address as Address }))
      dispatch(fetchPoolData({ poolName: String(pool.routerIndex) }))
      dispatch(resetPoolWithdraw())
    } catch (error) {
      if (isTransactionReceiptError(error)) {
        dispatch(fetchPoolUserData({ pool, address: address as Address }))
        dispatch(fetchPoolData({ poolName: String(pool.routerIndex) }))
        dispatch(resetPoolWithdraw())
      }
      txErrorHandler(error)
    }
  }

  // dispatch(fetchPoolData({ poolName: pool.routerIndex }))

  return (
    pool && (
      <div>
        <div className="percentage">
          <span className="mr-2 text-white">Withdraw Percentage %</span>
          <input
            className={`
            px-2 py-1 w-1/5 rounded-md
            focus:ring-indigo-500 focus:outline-none focus:border-purple-700
            border border-transparent
            bg-[#111111]
            text-gray-300
          `}
            placeholder="0"
            onChange={(e) => onPercentChange(Number(e.currentTarget.value))}
            onFocus={(e) => e.target.select()}
            value={percentage ?? ''}
          />
          <div className="my-2">
            {/* @ts-ignore */}
            <Slider
              axis="x"
              xstep={10}
              xmin={0}
              xmax={100}
              x={percentage ?? 100}
              onChange={(i) => {
                onPercentChange(i.x)
              }}
              styles={{
                track: {
                  backgroundColor: '#E0E7FF',
                  width: '95%',
                },
                active: {
                  backgroundColor: '#B286FF',
                },
                thumb: {
                  backgroundColor: '#CE55FE',
                },
              }}
            />
          </div>
        </div>
        <Grid gap={2} cols={{ xs: 1 }} className="mt-2 mb-4">
          <RadioButton
            checked={withdrawType === ALL}
            onChange={() => {
              dispatch(setWithdrawType(ALL))
            }}
            label="Combo"
            labelClassName={withdrawType === ALL && 'text-indigo-500'}
          />
          {poolSpecificTokens &&
            poolSpecificTokens.map((poolSpecificToken, index) => {
              const checked = withdrawType === index.toString()

              return (
                <RadioButton
                  radioClassName={getCoinTextColorCombined(
                    poolSpecificToken.color
                  )}
                  key={poolSpecificToken?.symbol}
                  checked={checked}
                  onChange={() => {
                    dispatch(setWithdrawType(index.toString()))
                  }}
                  labelClassName={
                    checked &&
                    `${getCoinTextColorCombined(
                      poolSpecificToken.color
                    )} opacity-90`
                  }
                  label={poolSpecificToken.name}
                />
              )
            })}
        </Grid>
        <WithdrawTokenInput
          onChange={(value) => onChangeInputValue(pool, value)}
        />
        <div className="mb-4" />
        <WithdrawButton
          approveTxn={approveTxn}
          withdrawTxn={withdrawTxn}
          isApproved={isApproved}
        />

        {stringToBigInt(inputValue, poolDecimals) > 0n && (
          <div className={` mt-2  bg-bgBase `}>
            <Grid cols={{ xs: 2 }}>
              <div>
                <ReceivedTokenSection
                  poolTokens={poolSpecificTokens}
                  withdrawQuote={withdrawQuote}
                  chainId={chainId}
                />
              </div>
              <div>
                {withdrawQuote.priceImpact && (
                  <PriceImpactDisplay priceImpact={withdrawQuote.priceImpact} />
                )}
              </div>
            </Grid>
          </div>
        )}
      </div>
    )
  )
}

const sumBigInts = (
  pool: Token,
  bigIntMap: Record<string, { value: bigint; index: number }>,
  withdrawType: string
) => {
  if (!pool?.poolTokens) {
    return 0n
  }

  const chainId = pool.chainId

  const currentTokens =
    withdrawType === ALL ? bigIntMap[withdrawType] : bigIntMap

  return pool.poolTokens.reduce((sum, token, index) => {
    if (!currentTokens[index]) {
      return sum
    }

    // Compute the power of 10 using pow10BigInt function
    const scaleFactor = pow10BigInt(
      BigInt(18) - BigInt(token.decimals[chainId])
    )
    const valueToAdd = currentTokens[index].value * scaleFactor

    return sum + valueToAdd
  }, 0n)
}

const pow10BigInt = (n: bigint) => {
  let result = 1n
  for (let i = 0n; i < n; i++) {
    result *= 10n
  }
  return result
}

const transformAmount = (amount) => {
  return {
    value: BigInt(amount.value.toString()),
    index: amount.index,
  }
}

const poolTokenByIndex = (poolTokens: Token[], index: number) => {
  return poolTokens.find((_poolToken, idx) => index === idx)
}

export default Withdraw
