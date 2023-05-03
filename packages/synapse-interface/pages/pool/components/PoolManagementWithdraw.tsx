import _ from 'lodash'
import { useEffect, useState } from 'react'
import Slider from 'react-input-slider'
import { stringToBigNum } from '@/utils/stringToBigNum'

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
import { useSwapDepositContract } from '@hooks/useSwapDepositContract'

import { OPTIMISM_ETH_SWAP_TOKEN } from '@constants/tokens/poolMaster'
import { Token } from '@synapsecns/sdk-router'
const PoolManagementWithdraw = ({
  pool,
  chainId,
  address,
  poolData,
  poolUserData,
}: {
  pool: any
  chainId: number
  address: string
  poolData: any
  poolUserData: any
}) => {
  const [inputValue, setInputValue] = useState<{
    bn: BigNumber
    str: string
  }>({ bn: Zero, str: '' })
  const [withdrawQuote, setWithdrawQuote] = useState<{
    priceImpact: BigNumber
    outputs: Record<string, BigNumber>
    allowance: BigNumber
    routerAddress: string
  }>({
    priceImpact: Zero,
    outputs: {},
    allowance: undefined,
    routerAddress: '',
  })

  const [withdrawType, setWithdrawType] = useState(ALL)
  const [percentage, setPercentage] = useState(100)

  const clearInputs = () => {}

  const lpTokenValue = ''
  const lpTokenAmount = Zero
  const debouncedPoolData = useDebounce(poolData, 500)

  const lpTokenBalance = Zero
  // const checkPoolNameChange = poolData?.name === debouncedPoolData?.name
  const placeholder = async (): Promise<TransactionResponse> => {
    console.log('placeholder')
    return
  }

  const [lastChangeField, setLastChangeField] = useState(undefined)

  // useEffect(() => {
  //   if (withdrawType === ALL && lastChangeField == 'PERCENT') {
  //     const numericalOut = formatUnits(
  //       lpTokenBalance.mul(Number(percentage)).div(100),
  //       pool.decimals[chainId]
  //     )

  //     setLpTokenValue(`${numericalOut}`)
  //   }
  // }, [
  //   withdrawType,
  //   percentage,
  //   checkPoolNameChange,
  //   lastChangeField, // the pool data displayer here is essential to update on initial load
  // ])

  const calculateMaxWithdraw = async () => {
    if (poolUserData == null || address == null) {
      return
    }
    const swapContract = await useSwapDepositContract(pool, chainId)
    const newInputState = {}
    const depositTokens = pool.poolTokens
    if (withdrawType == ALL) {
      const results = await swapContract.calculateRemoveLiquidity(inputValue.bn)
      console.log('results', results)
      //  for (const [token, amount] of _.zip(depositTokens, results) ) {
      //    newInputState[token.symbol] = formatUnits(amount, token.decimals[chainId])
      //  }
    } else {
      for (const token of pool.poolTokens) {
        newInputState[token.symbol] = ''
      }
      const tokenIndex = await swapContract.getTokenIndex(withdrawType)
      const amount = await swapContract.calculateRemoveLiquidityOneToken(
        inputValue.bn,
        tokenIndex
      )
      console.log('resultsamount', amount)

      // newInputState[token.symbol] = formatUnits(amount, token.decimals[chainId])
    }
    // TODO: DOUBLE CHECK THIS
    // setWithdrawQuote({
    //   priceImpact,
    //   allowance,
    //   outputs,
    //   routerAddress: pool.swapAddresses[chainId],
    // })
  }

  useEffect(() => {
    calculateMaxWithdraw()
  }, [inputValue])

  const onPercentChange = (percent: number) => {
    if (percent > 100) {
      percent = 100
    }
    setPercentage(percent)
    const numericalOut = formatUnits(
      poolUserData.lpTokenBalance.mul(Number(percent)).div(100),
      pool.decimals[chainId]
    )
    onChangeInputValue(pool, numericalOut)
  }

  const onChangeInputValue = (token: Token, value: string) => {
    const bigNum = stringToBigNum(value, token.decimals[chainId]) ?? Zero
    setInputValue({ bn: bigNum, str: value })
    const pn = bigNum.mul(100).div(poolUserData.lpTokenBalance).toNumber()
    console.log('poolUserData.lpTokenBalance', pn)

    if (pn > 100) {
      setPercentage(100)
    } else {
      setPercentage(pn)
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
          placeholder="100"
          onChange={(e) => {
            onPercentChange(Number(e.currentTarget.value))
          }}
          onFocus={(e) => e.target.select()}
          value={percentage ?? ''}
        />
        <div className="my-2">
          <Slider
            axis="x"
            xstep={10}
            xmin={0}
            xmax={100}
            x={percentage ?? '100'}
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
        {pool.poolTokens.map((token) => {
          const checked = withdrawType === token.addresses[chainId]

          return (
            <RadioButton
              radioClassName={getCoinTextColorCombined(token.color)}
              key={token.symbol}
              checked={checked}
              onChange={() => {
                setWithdrawType(token.addresses[chainId])
              }}
              labelClassName={
                checked && `${getCoinTextColorCombined(token.color)} opacity-90`
              }
              label={token.name}
            />
          )
        })}
      </Grid>
      <TokenInput
        token={pool}
        key={pool.symbol}
        inputValueStr={inputValue.str}
        balanceStr={poolUserData.lpTokenBalanceStr}
        onChange={(value) => onChangeInputValue(pool, value)}
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
        <div
          className={`py-3.5 pr-6 pl-6 mt-2 rounded-b-2xl bg-bgBase transition-all`}
        >
          <Grid cols={{ xs: 2 }}>
            <div>
              <RecievedTokenSection
                poolTokens={pool.poolTokens}
                withdrawQuote={withdrawQuote}
                chainId={chainId}
              />
            </div>
            <div>
              <PriceImpactDisplay priceImpact={withdrawQuote.priceImpact} />
            </div>
          </Grid>
        </div>
      </Transition>
    </div>
  )
}

export default PoolManagementWithdraw
