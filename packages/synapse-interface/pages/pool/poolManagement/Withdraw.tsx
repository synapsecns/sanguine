import _ from 'lodash'
import { useEffect, useState } from 'react'
import Slider from 'react-input-slider'
import { stringToBigInt } from '@/utils/stringToBigNum'

import { useSynapseContext } from '@/utils/providers/SynapseProvider'

import { getCoinTextColorCombined } from '@styles/tokens'
import { ALL } from '@constants/withdrawTypes'
import Grid from '@tw/Grid'
import { WithdrawTokenInput } from '@components/TokenInput'
import RadioButton from '@components/buttons/RadioButton'
import ReceivedTokenSection from '../components/ReceivedTokenSection'
import PriceImpactDisplay from '../components/PriceImpactDisplay'

import { Transition } from '@headlessui/react'
import { Token } from '@types'
import { approve, withdraw } from '@/utils/actions/approveAndWithdraw'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { calculatePriceImpact } from '@/utils/priceImpact'
import { formatBigIntToString } from '@/utils/bigint/format'

import { Address } from '@wagmi/core'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '@/store/store'

import {
  setInputValue,
  setWithdrawQuote,
  setWithdrawType,
  setIsLoading,
} from '@/slices/poolWithdrawSlice'
import { WithdrawButton } from './WithdrawButton'
import { txErrorHandler } from '@/utils/txErrorHandler'

const Withdraw = ({
  chainId,
  address,
}: {
  chainId: number
  address: string
}) => {
  const [percentage, setPercentage] = useState(0)
  const [isApproved, setIsApproved] = useState<boolean>(false)
  const { pool, poolData } = useSelector((state: RootState) => state.poolData)
  const { poolUserData } = useSelector((state: RootState) => state.poolUserData)
  const { withdrawQuote, inputValue, withdrawType } = useSelector(
    (state: RootState) => state.poolWithdraw
  )
  const { poolAddress } = getSwapDepositContractFields(pool, chainId)
  const { synapseSDK } = useSynapseContext()

  const dispatch = useDispatch()

  const showTokens = pool ? pool.nativeTokens ?? pool.poolTokens : []

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
          inputValue.bi
        )
        outputs[withdrawType] = amounts.map((amount) => {
          return {
            value: BigInt(amount.value.toString()),
            index: amount.index,
          }
        })
      } else {
        const { amount } = await synapseSDK.calculateRemoveLiquidityOne(
          chainId,
          poolAddress,
          inputValue.bi,
          Number(withdrawType)
        )
        outputs[withdrawType] = {
          value: BigInt(amount.value.toString()),
          index: amount.index,
        }
      }

      const outputTokensSum = sumBigInts(pool, outputs, withdrawType)

      const priceImpact = calculatePriceImpact(
        inputValue.bi,
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
          routerAddress: poolAddress,
        })
      )
      dispatch(setIsLoading(false))
    } catch (e) {
      dispatch(setIsLoading(false))
      console.log(e)
    }
  }

  useEffect(() => {
    if (poolUserData && poolData && address && pool && inputValue.bi > 0n) {
      calculateMaxWithdraw()
    }
  }, [inputValue, withdrawType])

  useEffect(() => {
    if (withdrawQuote?.allowance && inputValue.bi <= withdrawQuote.allowance) {
      setIsApproved(true)
    } else {
      setIsApproved(false)
    }
  }, [inputValue, withdrawQuote])

  const onPercentChange = (percent: number) => {
    if (percent > 100) {
      percent = 100
    }
    setPercentage(percent)
    const numericalOut = poolUserData.lpTokenBalance
      ? formatBigIntToString(
          (poolUserData.lpTokenBalance * BigInt(percent)) / BigInt(100),
          pool.decimals[chainId]
        )
      : ''
    const bigInt = stringToBigInt(numericalOut, pool.decimals[chainId])
    dispatch(setInputValue({ bi: bigInt, str: numericalOut }))
  }

  const onChangeInputValue = (token: Token, value: string) => {
    const bigInt = stringToBigInt(value, token.decimals[chainId])

    if (poolUserData.lpTokenBalance === 0n) {
      dispatch(setInputValue({ bi: bigInt, str: value }))

      setPercentage(0)
      return
    }
    const pn = bigInt
      ? Number(
          (bigInt * BigInt(100)) /
            BigInt(poolUserData.lpTokenBalance.toString())
        )
      : 0

    dispatch(setInputValue({ bi: bigInt, str: value }))

    if (pn > 100) {
      setPercentage(100)
    } else {
      setPercentage(pn)
    }
  }

  const approveTxn = async () => {
    try {
      const tx = approve(pool, withdrawQuote, inputValue.bi, chainId)

      try {
        await tx
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
        stringToBigInt(inputValue.str, pool.decimals[chainId]),
        chainId,
        withdrawType,
        withdrawQuote.outputs
      )

      try {
        await tx
      } catch (error) {
        txErrorHandler(error)
      }
    } catch (error) {
      txErrorHandler(error)
    }
  }

  return (
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
      <Grid gap={2} cols={{ xs: 1 }} className="mt-2">
        <RadioButton
          checked={withdrawType === ALL}
          onChange={() => {
            dispatch(setWithdrawType(ALL))
          }}
          label="Combo"
          labelClassName={withdrawType === ALL && 'text-indigo-500'}
        />
        {showTokens &&
          showTokens.map((token, index) => {
            const checked = withdrawType === index.toString()

            return (
              <RadioButton
                radioClassName={getCoinTextColorCombined(token.color)}
                key={token?.symbol}
                checked={checked}
                onChange={() => {
                  dispatch(setWithdrawType(index.toString()))
                }}
                labelClassName={
                  checked &&
                  `${getCoinTextColorCombined(token.color)} opacity-90`
                }
                label={token.name}
              />
            )
          })}
      </Grid>
      <WithdrawTokenInput
        poolUserData={poolUserData}
        token={pool}
        key={pool?.symbol}
        inputValueStr={inputValue.str}
        balanceStr={
          poolUserData?.lpTokenBalance
            ? formatBigIntToString(
                poolUserData?.lpTokenBalance,
                pool.decimals[chainId],
                4
              )
            : '0.0000'
        }
        onChange={(value) => onChangeInputValue(pool, value)}
        chainId={chainId}
        address={address}
      />
      <WithdrawButton
        approveTxn={approveTxn}
        withdrawTxn={withdrawTxn}
        isApproved={isApproved}
      />

      <Transition
        appear={true}
        unmount={false}
        show={inputValue.bi > 0n}
        enter="transition duration-100 ease-out"
        enterFrom="transform-gpu scale-y-0 "
        enterTo="transform-gpu scale-y-100 opacity-100"
        leave="transition duration-75 ease-out "
        leaveFrom="transform-gpu scale-y-100 opacity-100"
        leaveTo="transform-gpu scale-y-0 "
        className="-mx-6 origin-top "
      >
        <div
          className={`py-3.5 pr-6 pl-6 mt-2 rounded-b-2xl bg-bgBase transition-all`}
        >
          <Grid cols={{ xs: 2 }}>
            <div>
              <ReceivedTokenSection
                poolTokens={showTokens}
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
      </Transition>
    </div>
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

function pow10BigInt(n) {
  let result = 1n
  for (let i = 0n; i < n; i++) {
    result *= 10n
  }
  return result
}

export default Withdraw
