import _ from 'lodash'
import { useEffect, useState } from 'react'
import Slider from 'react-input-slider'
import { stringToBigNum } from '@/utils/stringToBigNum'

import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits } from '@ethersproject/units'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

import { Transition } from '@headlessui/react'

import { getCoinTextColorCombined } from '@styles/tokens'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'

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
import RecievedTokenSection from '../components/RecievedTokenSection'
import PriceImpactDisplay from '../components/PriceImpactDisplay'

import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { Zero } from '@ethersproject/constants'
// need to add pending for deposit func
import { TransactionResponse } from '@ethersproject/providers'
import { useSwapDepositContract } from '@hooks/useSwapDepositContract'
import { calculatePriceImpactWithdraw } from '@utils/priceImpact'
import { OPTIMISM_ETH_SWAP_TOKEN } from '@constants/tokens/poolMaster'
import { Token } from '@types'
import { approve, withdraw } from '@/utils/actions/approveAndWithdraw'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'

const Withdraw = ({
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
  const SynapseSDK = useSynapseContext()

  const lpTokenValue = ''
  const lpTokenAmount = Zero
  const debouncedPoolData = useDebounce(poolData, 500)

  const lpTokenBalance = Zero
  const placeholder = async (): Promise<TransactionResponse> => {
    console.log('placeholder')
    return
  }

  const sumBigNumbers = (pool: Token, bigNumMap: any) => {
    let sum = Zero
    pool.poolTokens.map((token) => {
      if (bigNumMap[token.addresses[chainId]]) {
        sum = sum.add(
          bigNumMap[token.addresses[chainId]].mul(
            BigNumber.from(10).pow(18 - token.decimals[chainId])
          )
        )
      }
    })
    return sum
  }
  const calculateMaxWithdraw = async () => {
    if (poolUserData == null || address == null) {
      return
    }
    const swapContract = await useSwapDepositContract(pool, chainId)
    const outputs = {}
    const depositTokens = pool.poolTokens
    if (withdrawType == ALL) {
      const { amounts } = await SynapseSDK.calculateRemoveLiquidity(
        chainId,
        pool.swapAddresses[chainId],
        inputValue.bn
      )
      // const results = await swapContract.calculateRemoveLiquidity(inputValue.bn)
      console.log('results', amounts)
      for (const tokenAddr in amounts) {
        outputs[tokenAddr] = amounts[tokenAddr]
      }
    } else {
      const tokenIndex = await swapContract.getTokenIndex(withdrawType)
      console.log('tokenIndex', tokenIndex)
      const amount = await swapContract.calculateRemoveLiquidityOneToken(
        inputValue.bn,
        tokenIndex
      )
      outputs[withdrawType] = amount

      // newInputState[token.symbol] = formatUnits(amount, token.decimals[chainId])
    }
    console.log('newWithdrawQuotenewWithdrawQuote', withdrawType, outputs)
    const tokenSum = sumBigNumbers(pool, outputs)
    const priceImpact = calculateExchangeRate(
      inputValue.bn,
      18,
      inputValue.bn.sub(tokenSum),
      18
    )
    const allowance = await getTokenAllowance(
      pool.swapAddresses[chainId],
      pool.addresses[chainId],
      address,
      chainId
    )
    setWithdrawQuote({
      priceImpact,
      allowance,
      outputs,
      routerAddress: pool.swapAddresses[chainId],
    })
  }

  useEffect(() => {
    if (poolUserData && poolData && address && pool && inputValue.bn.gt(Zero)) {
      calculateMaxWithdraw()
    }
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

  // some messy button gen stuff (will re-write)
  let isFromBalanceEnough = true
  let isAllowanceEnough = true
  let btnLabel = 'Withdraw'
  let pendingLabel = 'Withdrawing funds...'
  let btnClassName = ''
  let buttonAction = () =>
    withdraw(pool, 'ONE_TENTH', null, inputValue.bn, chainId)
  let postButtonAction = () => {
    console.log('JHK')
    resetInputs()
  }

  for (const [tokenAddr, amount] of Object.entries(inputValue.bn)) {
    if (
      withdrawQuote.allowance &&
      !amount.isZero() &&
      amount.gt(withdrawQuote.allowance)
    ) {
      isAllowanceEnough = false
    }
    poolUserData.tokens.map((tokenObj, i) => {
      if (
        tokenObj.token.addresses[chainId] === tokenAddr &&
        amount.gt(tokenObj.balance)
      ) {
        isFromBalanceEnough = false
      }
    })
  }

  if (!isFromBalanceEnough) {
    btnLabel = `Insufficient Balance`
  } else if (!isAllowanceEnough) {
    buttonAction = () => approve(pool, withdrawQuote, inputValue.bn, chainId)
    btnLabel = `Approve Token(s)`
    pendingLabel = `Approving Token(s)`
    btnClassName = 'from-[#feba06] to-[#FEC737]'
    postButtonAction = () => setTime(0)
  }
  const actionBtn = (
    <TransactionButton
      className={btnClassName}
      disabled={inputValue.bn.eq(0)}
      onClick={() => buttonAction()}
      onSuccess={() => postButtonAction()}
      label={btnLabel}
      pendingLabel={pendingLabel}
    />
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
      {/*
      TODO FIX THIS TRANSITION
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
      > */}
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
      {/* </Transition> */}
    </div>
  )
}

export default Withdraw
