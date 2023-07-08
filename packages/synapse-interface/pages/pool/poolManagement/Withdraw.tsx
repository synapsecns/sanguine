import _ from 'lodash'
import { useEffect, useState, useMemo } from 'react'
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
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { Token } from '@types'
import { approve, withdraw } from '@/utils/actions/approveAndWithdraw'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { PoolData, PoolUserData } from '@types'
import { getSwapDepositContractFields } from '@/utils/getSwapDepositContractFields'
import { calculatePriceImpact } from '@/utils/priceImpact'
import { formatBigIntToString } from '@/utils/bigint/format'

import { Address } from '@wagmi/core'

const DEFAULT_WITHDRAW_QUOTE = {
  priceImpact: 0n,
  outputs: {},
  allowance: undefined,
  routerAddress: '',
}

const Withdraw = ({
  pool,
  chainId,
  address,
  poolData,
  poolUserData,
  refetchCallback,
}: {
  pool: any
  chainId: number
  address: string
  poolData: PoolData
  poolUserData: PoolUserData
  refetchCallback: () => void
}) => {
  const [inputValue, setInputValue] = useState<{
    bi: bigint
    str: string
  }>({ bi: 0n, str: '' })

  const [withdrawQuote, setWithdrawQuote] = useState<{
    priceImpact: bigint
    outputs: Record<
      string,
      {
        value: bigint
        index: number
      }
    >
    allowance: bigint
    routerAddress: string
  }>(DEFAULT_WITHDRAW_QUOTE)

  const [withdrawType, setWithdrawType] = useState(ALL)
  const [percentage, setPercentage] = useState(0)
  const [time, setTime] = useState(Date.now())

  const [isApproved, setIsApproved] = useState(false)

  const resetInput = () => {
    setInputValue({ bi: 0n, str: '' })
  }
  const { synapseSDK } = useSynapseContext()

  const showTokens = pool ? pool.nativeTokens ?? pool.poolTokens : []
  const { poolAddress } = getSwapDepositContractFields(pool, chainId)

  const calculateMaxWithdraw = async () => {
    if (poolUserData == null || address == null) {
      return
    }
    try {
      const outputs: Record<
        string,
        {
          value: bigint
          index: number
        }
      > = {}
      const { virtualPrice } = poolData
      if (withdrawType == ALL) {
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

      const outputTokensSum = sumBigInts(pool, outputs, chainId, withdrawType)

      const priceImpact = calculatePriceImpact(
        inputValue.bi,
        outputTokensSum,
        virtualPrice,
        true
      )

      const allowance = await getTokenAllowance(
        poolAddress,
        pool.addresses[chainId],
        address as Address,
        chainId
      )
      console.log(`allowance`, allowance)
      setWithdrawQuote({
        priceImpact,
        allowance,
        outputs,
        routerAddress: poolAddress,
      })
    } catch (e) {
      console.log(e)
    }
  }

  useEffect(() => {
    if (poolUserData && poolData && address && pool && inputValue.bi > 0n) {
      calculateMaxWithdraw()
    }
  }, [inputValue, time, withdrawType])

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
    setInputValue({ bi: bigInt, str: numericalOut })
  }

  const onChangeInputValue = (token: Token, value: string) => {
    const bigInt = stringToBigInt(value, token.decimals[chainId])

    if (poolUserData.lpTokenBalance === 0n) {
      setInputValue({ bi: bigInt, str: value })

      setPercentage(0)
      return
    }
    const pn = bigInt
      ? Number(
          (bigInt * BigInt(100)) /
            BigInt(poolUserData.lpTokenBalance.toString())
        )
      : 0

    setInputValue({ bi: bigInt, str: value })

    if (pn > 100) {
      setPercentage(100)
    } else {
      setPercentage(pn)
    }
  }

  let isFromBalanceEnough = true
  let isAllowanceEnough = false

  const getButtonProperties = () => {
    let properties = {
      label: 'Withdraw',
      pendingLabel: 'Withdrawing funds...',
      className: '',
      disabled: false,
      buttonAction: () =>
        withdraw(
          pool,
          'ONE_TENTH',
          null,
          stringToBigInt(inputValue.str, pool.decimals[chainId]),
          chainId,
          withdrawType,
          withdrawQuote.outputs
        ),
      postButtonAction: () => {
        refetchCallback()
        setPercentage(0)
        setWithdrawQuote(DEFAULT_WITHDRAW_QUOTE)
        resetInput()
      },
    }

    if (inputValue.bi === 0n) {
      properties.label = `Enter amount`
      properties.disabled = true
      return properties
    }

    if (!isFromBalanceEnough) {
      properties.label = `Insufficient Balance`
      properties.disabled = true
      return properties
    }

    console.log(`isallowanceneough`, isAllowanceEnough)
    console.log(`isApproved`, isApproved)

    if (!isAllowanceEnough && !isApproved) {
      properties.label = `Approve Token(s)`
      properties.pendingLabel = `Approving Token(s)`
      properties.className = 'from-[#feba06] to-[#FEC737]'
      properties.disabled = false
      properties.buttonAction = () =>
        approve(pool, withdrawQuote, inputValue.bi, chainId).then((res) => {
          if (res && res.status === 'success') {
            setIsApproved(true)
          }
        })
      properties.postButtonAction = () => setTime(0)
      return properties
    }

    return properties
  }

  if (
    withdrawQuote.allowance &&
    inputValue.bi !== 0n &&
    inputValue.bi > withdrawQuote.allowance
  ) {
    isAllowanceEnough = false
  }

  if (
    inputValue.bi !== 0n &&
    inputValue.bi > BigInt(poolUserData.lpTokenBalance.toString())
  ) {
    isFromBalanceEnough = false
  }

  const {
    label: btnLabel,
    pendingLabel,
    className: btnClassName,
    buttonAction,
    postButtonAction,
    disabled,
  } = useMemo(getButtonProperties, [
    isFromBalanceEnough,
    isAllowanceEnough,
    address,
    inputValue,
    withdrawQuote,
    isApproved,
  ])

  const actionBtn = useMemo(
    () => (
      <TransactionButton
        className={btnClassName}
        disabled={disabled}
        onClick={() => buttonAction()}
        onSuccess={() => postButtonAction()}
        label={btnLabel}
        pendingLabel={pendingLabel}
      />
    ),
    [
      buttonAction,
      postButtonAction,
      btnLabel,
      pendingLabel,
      btnClassName,
      isFromBalanceEnough,
      isAllowanceEnough,
      isApproved,
    ]
  )

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
        {/* {error && (
          <div className="text-red-400 opacity-80">{error?.message}</div>
        )} */}
      </div>
      <Grid gap={2} cols={{ xs: 1 }} className="mt-2">
        <RadioButton
          checked={withdrawType === ALL}
          onChange={() => {
            setWithdrawType(ALL)
          }}
          label="Combo"
          labelClassName={withdrawType === ALL && 'text-indigo-500'}
        />
        {showTokens &&
          showTokens.map((token) => {
            // TODO: poolsToken.findIndex is too verbose and was a hacky solution to not have to refactor a lot of state passing. Needs to be fixed to handle indexes correctly.
            const checked =
              withdrawType ===
              (pool.nativeTokens ? pool.nativeTokens : pool.poolTokens)
                .findIndex(
                  (poolToken) =>
                    poolToken.addresses[chainId] === token.addresses[chainId]
                )
                .toString()
            return (
              <RadioButton
                radioClassName={getCoinTextColorCombined(token.color)}
                key={token?.symbol}
                checked={checked}
                onChange={() => {
                  // Determine the tokens array
                  const tokensArray = pool.nativeTokens
                    ? pool.nativeTokens
                    : pool.poolTokens

                  // Find the index
                  const index = tokensArray.findIndex(
                    (poolToken) =>
                      poolToken.addresses[chainId] === token.addresses[chainId]
                  )

                  // Convert the index to a string
                  const indexString = index.toString()
                  setWithdrawType(indexString)
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
      {actionBtn}

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
  chainId: number,
  withdrawType: string
) => {
  if (!pool?.poolTokens) {
    return 0n
  }

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
