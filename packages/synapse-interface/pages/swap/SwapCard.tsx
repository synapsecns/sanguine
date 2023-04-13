import { useState } from 'react'

import { parseUnits } from '@ethersproject/units'

import { Transition } from '@headlessui/react'

import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useComboSwapContract } from '@hooks/contracts/useContract'
import { useApproveAndSwap } from '@hooks/actions/useApproveAndSwap'
import { APPROVAL_STATE, useApproveToken } from '@hooks/actions/useApproveToken'

import { sanitizeValue } from '@utils/sanitizeValue'

import { getSwapCardShadowStyleForCoin } from '@styles/coins'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'

import Grid from '@tw/Grid'
import Card from '@tw/Card'

import ExchangeRateInfo from '@components/ExchangeRateInfo'
import { TransactionButton } from '@components/buttons/SubmitTxButton'

import CoreSwapContainer from './CoreSwapContainer'
import { WETH } from '@constants/tokens/basic'
import { displaySymbol } from '@utils/displaySymbol'

import { ChainLabel } from '@pages/Bridge/BridgeInputContainer/ChainLabel'
import { NetworkSlideOver } from '@components/misc/NetworkSlideOver'
import { CoinSlideOver } from '@components/misc/CoinSlideOver'

import { cleanNumberInput } from '@utils/cleanNumberInput'

export default function SwapCard({
  poolName,
  swapableTokens,
  fromCoin,
  fromValue,
  toCoin,
  toValue,
  swapFromToCoins,
  onChangeChain,
  onSelectFromCoin,
  onSelectToCoin,
  onChangeFromAmount,
  onChangeToAmount,
  error,
  priceImpact,
  exchangeRate,
  fromRef,
  toRef,
}) {
  const { chainId } = useActiveWeb3React()
  const approveAndSwap = useApproveAndSwap(poolName)
  const swapContract = useComboSwapContract(poolName)
  const [displayType, setDisplayType] = useState(undefined)

  const toAmount = parseUnits(sanitizeValue(toValue), toCoin.decimals[chainId])
  const fromAmount = parseUnits(
    sanitizeValue(fromValue),
    fromCoin.decimals[chainId]
  )

  const [approvalState, approveToken] = useApproveToken(
    fromCoin,
    swapContract?.address,
    fromAmount
  )

  const fromArgs = {
    isSwapFrom: true,
    selected: fromCoin,
    onChangeSelected: onSelectFromCoin,
    onChangeAmount: onChangeFromAmount,
    inputValue: fromValue,
    inputRef: fromRef,
    tokens: swapableTokens,
    chainId,
    setDisplayType,
    onChangeChain,
    selectedChainId: chainId,
  }

  const toArgs = {
    isSwapFrom: false,
    selected: toCoin,
    onChangeSelected: onSelectToCoin,
    onChangeAmount: onChangeToAmount,
    inputValue: toValue,
    swapFromToCoins: swapFromToCoins,
    inputRef: toRef,
    tokens: swapableTokens,
    chainId,
    setDisplayType,
    onChangeChain,
  }

  const approvalBtn = (
    <TransactionButton
      onClick={approveToken}
      label={`Approve ${displaySymbol(chainId, fromCoin)}`}
      pendingLabel={`Approving ${displaySymbol(chainId, fromCoin)}  `}
    />
  )

  let swapButtonLabel
  if (error) {
    swapButtonLabel = error
  } else {
    swapButtonLabel = fromAmount.eq(0)
      ? 'Enter amount to swap'
      : 'Swap your funds'
  }
  const swapBtn = (
    <TransactionButton
      disabled={toAmount.eq(0) || error}
      onClick={() => {
        return approveAndSwap({
          fromAmount: fromAmount,
          fromCoin,
          toAmount: toAmount,
          toCoin,
        })
      }}
      onSuccess={() => {
        onChangeFromAmount('')
      }}
      label={swapButtonLabel}
      pendingLabel={`Swapping...`}
    />
  )

  let actionBtn
  if (
    approvalState === APPROVAL_STATE.NOT_APPROVED &&
    fromCoin.symbol != WETH.symbol
  ) {
    actionBtn = approvalBtn
  } else {
    actionBtn = swapBtn
  }

  const fromChainCardContent = (
    <NetworkSlideOver key="fromChainBlock" {...fromArgs} />
  )

  const swapCardMainContent = (
    <>
      <Grid cols={{ xs: 1 }} gap={4} className="place-content-center">
        <div className="pt-3 pb-3 pl-4 pr-4 mt-2 border-none bg-bgLight rounded-xl">
          <ChainLabel
            isSwapFrom={true}
            chainId={chainId}
            setDisplayType={setDisplayType}
            onChangeChain={onChangeChain}
            titleText="Chain"
          />
        </div>
        <CoreSwapContainer {...fromArgs} />
        <CoreSwapContainer {...toArgs} />
      </Grid>
      {/* <Transition
        appear={true}
        unmount={false}
        show={!fromAmount.eq(0)}
        enter="transition duration-100 ease-out"
        enterFrom="transform-gpu scale-y-0 "
        enterTo="transform-gpu scale-y-100 opacity-100"
        leave="transition duration-75 ease-out "
        leaveFrom="transform-gpu scale-y-100 opacity-100"
        leaveTo="transform-gpu scale-y-0 "
        className="-mx-6 origin-top "
      > */}
      <ExchangeRateInfo
        fromAmount={fromAmount}
        fromCoin={fromCoin}
        toCoin={toCoin}
        exchangeRate={exchangeRate}
        priceImpact={priceImpact}
        toChainId={chainId}
      />
      {/* </Transition> */}
      <div className="px-2 py-2 md:px-0 md:py-4">{actionBtn}</div>
    </>
  )

  const fromCardContent = <CoinSlideOver key="fromBlock" {...fromArgs} />

  const toCardContent = <CoinSlideOver key="toBlock" {...toArgs} />

  const transitionProps = {
    ...COIN_SLIDE_OVER_PROPS,
    className: `
      origin-bottom absolute
      w-full h-full
      md:w-[95%] md:h-[95%]
      -ml-0 md:-ml-3
      md:mt-3 
      bg-bgBase
      z-20 rounded-3xl
    `,
  }

  let swapCardShadow
  if (displayType === 'from') {
    swapCardShadow = getSwapCardShadowStyleForCoin(fromCoin)
  } else if (displayType === 'to') {
    swapCardShadow = getSwapCardShadowStyleForCoin(toCoin)
  } else {
    swapCardShadow = 'shadow-indigo-xl hover:shadow-purple-2xl'
  }

  return (
    <Card
      divider={false}
      className="max-w-lg px-1 pb-0 -mb-3 transition-all duration-100 transform rounded-xl bg-bgBase md:px-6 lg:px-6"
    >
      <div className="mb-8">
        <Transition show={displayType === 'from'} {...transitionProps}>
          {fromCardContent}
        </Transition>
        <Transition show={displayType === 'to'} {...transitionProps}>
          {toCardContent}
        </Transition>
        <Transition show={displayType === 'fromChain'} {...transitionProps}>
          {fromChainCardContent}
        </Transition>
        {swapCardMainContent}
      </div>
    </Card>
  )
}
