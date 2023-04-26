import _ from 'lodash'
import { useEffect, useState } from 'react'
import Slider from 'react-input-slider'

import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits } from '@ethersproject/units'

import { Transition } from '@headlessui/react'

import { getCoinTextColorCombined } from '@styles/tokens'

import { ETH } from '@constants/tokens/master'
import { WETH } from '@constants/tokens/swapMaster'
import { ALL } from '@constants/withdrawTypes'

import { useDebounce } from '@hooks/useDebounce'
import { usePendingTxWrapper } from '@hooks/usePendingTxWrapper'
// import { useApproveAndWithdraw } from '@hooks/actions/useApproveAndWithdraw'
// import { usePoolToken } from '@hooks/pools/usePools'
// import { useTokenBalance } from '@hooks/tokens/useTokenBalances'

import Grid from '@tw/Grid'

import TokenInput from '@components/TokenInput'
import RadioButton from '@components/buttons/RadioButton'
import ButtonLoadingSpinner from '@components/buttons/ButtonLoadingSpinner'
import RecievedTokenSection from './RecievedTokenSection'
import PriceImpactDisplay from './PriceImpactDisplay'

import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { Zero } from '@ethersproject/constants'
// need to add pending for deposit func
import { TransactionResponse } from '@ethersproject/providers'

import { OPTIMISM_ETH_SWAP_TOKEN } from '@constants/tokens/poolMaster'
export default function PoolManagementWithdraw({ pool, chainId, address }) {
  // const {
  //   onChangeTokenInputValue,
  //   clearInputs,
  //   priceImpact,

  //   depositTokens,
  //   poolTokens,
  //   inputState,
  //   setInputState,
  //   tokenInputSum,

  //   poolData,
  //   withdrawType,
  //   setWithdrawType,
  //   percentage,
  //   setPercentage,

  //   lpTokenValue,
  //   setLpTokenValue,
  //   lpTokenAmount,
  // } = useSwapPoolWithdraw(poolName)
  const onChangeTokenInputValue = ''
  const clearInputs = () => {}
  const priceImpact = ''

  const depositTokens = []
  const poolTokens = ''
  const inputState = ''
  const setInputState = ''
  const tokenInputSum = ''
  const withdrawType = ALL
  const poolData = ETH
  const setWithdrawType = (v: string) => undefined
  const percentage = ''
  const setPercentage = (v: Number) => {}

  const lpTokenValue = ''
  const setLpTokenValue = (v: string) => {}
  const lpTokenAmount = Zero
  const debouncedPoolData = useDebounce(poolData, 500)

  const lpToken = pool.poolTokens
  // const lpTokenBalance = useTokenBalance(lpToken) FUCK YOU
  const lpTokenBalance = Zero
  const checkPoolNameChange = poolData?.name === debouncedPoolData?.name
  const placeholder = async (): Promise<TransactionResponse> => {
    console.log('placeholder')
    return
  }
  // const approveAndWithdraw = useApproveAndWithdraw(poolName) TODO

  const [isPending, pendingTxWrapFunc] = usePendingTxWrapper()

  const [lastChangeField, setLastChangeField] = useState(undefined)

  useEffect(() => {
    if (withdrawType === ALL && lastChangeField == 'PERCENT') {
      const numericalOut = formatUnits(
        lpTokenBalance.mul(Number(percentage)).div(100),
        lpToken.decimals[chainId]
      )

      setLpTokenValue(`${numericalOut}`)
    }
  }, [
    withdrawType,
    percentage,
    checkPoolNameChange,
    lastChangeField, // the pool data displayer here is essential to update on initial load
  ])

  useEffect(() => {
    if (withdrawType === ALL && lastChangeField == 'TOKEN_INPUT') {
      if (lpTokenBalance.gt(Zero)) {
        const pn = lpTokenAmount.mul(100).div(lpTokenBalance).toNumber()
        if (pn > 100) {
          setPercentage(100)
        } else {
          setPercentage(pn)
        }
      }
    }
  }, [lastChangeField, lpTokenAmount, lpTokenValue])

  const percentageStr = percentage //`${Math.round(percentage)}`
  const error = {}

  function onPercentChange(percent) {
    let numPercent = Number(percent)
    if (numPercent > 100) {
      numPercent = 100
    }

    setPercentage(numPercent)

    // if (withdrawType === ALL) {
    const numericalOut = formatUnits(
      lpTokenBalance.mul(Number(numPercent)).div(100),
      lpToken.decimals[chainId]
    )

    setLpTokenValue(`${numericalOut}`)
    // }
    return
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
          placeholder="100"
          onChange={(e) => {
            setLastChangeField('PERCENT')
            onPercentChange(e.currentTarget.value)
          }}
          onFocus={(e) => e.target.select()}
          value={percentageStr ?? ''}
        />
        <div className="my-2">
          <Slider
            axis="x"
            xstep={10}
            xmin={0}
            xmax={100}
            x={percentageStr ?? '100'}
            onChange={(i) => {
              setLastChangeField('PERCENT')
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
        {error && (
          <div className="text-red-400 opacity-80">{error?.message}</div>
        )}
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
        {depositTokens.map((t) => {
          const checked = withdrawType === t.symbol

          return (
            <RadioButton
              radioClassName={getCoinTextColorCombined(t)}
              key={t.symbol}
              checked={checked}
              onChange={() => {
                setWithdrawType(t.symbol)
              }}
              labelClassName={
                checked && `${getCoinTextColorCombined(t)} opacity-90`
              }
              label={t.name}
            />
          )
        })}
      </Grid>
      <TokenInput
        token={lpToken}
        key={lpToken.symbol}
        inputValue={lpTokenValue}
        max={lpTokenBalance.toString()}
        onChange={(value: string) => {
          setLastChangeField('TOKEN_INPUT')
          if (value == '') {
            clearInputs()
          }
          setLpTokenValue(value)
        }}
        chainId={chainId}
        address={address}
      />
      <TransactionButton
        label="Withdraw"
        pendingLabel="Withdrawing"
        // OH GOD
        // onClick={async () => {
        //   await pendingTxWrapFunc(
        //     approveAndWithdraw({
        //       poolTokens: depositTokens,
        //       inputState,
        //       withdrawType,
        //       infiniteApproval: false,
        //       lpTokenAmountToSpend: lpTokenAmount,
        //     })
        //   )
        //   clearInputs()
        //   setLpTokenValue('')
        // }}
        onClick={placeholder}
      />
      <Transition
        appear={true}
        unmount={false}
        show={lpTokenAmount.gt(0)}
        enter="transition duration-100 ease-out"
        enterFrom="transform-gpu scale-y-0 "
        enterTo="transform-gpu scale-y-100 opacity-100"
        leave="transition duration-75 ease-out "
        leaveFrom="transform-gpu scale-y-100 opacity-100"
        leaveTo="transform-gpu scale-y-0 "
        className="-mx-6 origin-top "
      >
        <WithdrawCardFooter
          pool={pool}
          inputState={inputState}
          priceImpact={priceImpact}
        />
      </Transition>
    </div>
  )
}

function WithdrawCardFooter({ priceImpact, pool, inputState }) {
  return (
    <div
      className={`py-3.5 pr-6 pl-6 mt-2 rounded-b-2xl bg-bgBase transition-all`}
    >
      <Grid cols={{ xs: 2 }}>
        <div>
          <RecievedTokenSection
            poolTokens={pool.poolTokens}
            inputState={inputState}
          />
        </div>
        <div>
          <PriceImpactDisplay priceImpact={priceImpact} />
        </div>
      </Grid>
    </div>
  )
}
