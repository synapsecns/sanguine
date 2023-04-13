import _ from 'lodash'
import { useRef, useState, useEffect } from 'react'
import { useHistory, Link } from 'react-router-dom'
import { parseUnits, formatUnits } from '@ethersproject/units'
import { Zero, One } from '@ethersproject/constants'

import { checkCleanedValue } from '@utils/checkCleanedValue'
import { sanitizeValue } from '@utils/sanitizeValue'

import { HOW_TO_SWAP_URL, POOLS_PATH, SWAP_PATH } from '@urls'
import { stringifyParams } from '@urls/stringifyParams'

import { WETH, ETH, SYNJEWEL } from '@constants/tokens/basic'
import { SWAPABLE_TOKENS } from '@constants/tokens/swap'
import { PRIORITY_RANKING } from '@constants/tokens/priorityRanking'

import { useUrlQuery } from '@hooks/useUrlQuery'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useSwapContract } from '@hooks/contracts/useContract'
import { useTokenBalance } from '@hooks/tokens/useTokenBalances'
import { usePool } from '@hooks/pools/usePools'
import { useDebounce } from '@hooks/useDebounce'

import Grid from '@tw/Grid'

import StandardPageContainer from '@layouts/StandardPageContainer'

import { matchSymbolWithinPool } from '@utils/matchSymbolWithinPool'

import { getInfoMultiCoin } from './useMultiCoinInfo'

import {
  estimateAmountToGive,
  calcAmountToRecieve,
  calculateExchangeRate,
} from './funcs'

import SwapCard from './SwapCard'
import NoSwapCard from './NoSwapCard'

import { LandingPageWrapper } from '@layouts/LandingPageWrapper'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { useNetworkController } from '@hooks/wallet/useNetworkController'
import { PageHeader } from '../../components/PageHeader'

import { HarmonyCheck } from '@pages/Bridge'

export default function SwapPage() {
  const { chainId } = useActiveWeb3React()

  let targetCard
  if (PRIORITY_RANKING[chainId]) {
    targetCard = <SwapPageCard />
  } else {
    targetCard = <NoSwapCard />
  }

  return (
    <LandingPageWrapper>
      <StandardPageContainer>
        <div>
          <Grid
            cols={{ xs: 1 }}
            gap={6}
            className="justify-center px-2 py-16 sm:px-6 md:px-8"
          >
            <HarmonyCheck fromChainId={chainId} toChainId={chainId} />
            <div className="pb-3 place-self-center">{targetCard}</div>
          </Grid>
        </div>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

function SwapPageCard() {
  const urlQuery = useUrlQuery()
  const history = useHistory()
  const { chainId } = useActiveWeb3React()
  const debouncedChainId = useDebounce(chainId, 50)

  const priorityRanking = PRIORITY_RANKING[chainId]
  const swapableTokens = SWAPABLE_TOKENS[chainId]

  const fromQuery = urlQuery.get('inputCurrency')
  const toQuery = urlQuery.get('outputCurrency')

  const defaultFrom = _.find(
    swapableTokens,
    (token) => token.symbol === fromQuery
  )
  const defaultTo = _.find(swapableTokens, (token) => token.symbol === toQuery)

  let [fromCoin, setFromCoin] = useState(defaultFrom ?? swapableTokens[0])
  let [toCoin, setToCoin] = useState(defaultTo ?? swapableTokens[1])

  let poolName
  if (chainId === debouncedChainId) {
    poolName = getInfoMultiCoin(fromCoin, toCoin, priorityRanking).poolName
  }

  const swapContract = useSwapContract(poolName)
  // const tokenBalances = usePoolTokenBalances(poolName)
  const poolTokens = usePool(poolName)

  let balanceCoin
  if (fromCoin.symbol == WETH.symbol) {
    balanceCoin = ETH
  } else {
    balanceCoin = fromCoin
  }
  const fromBalance = useTokenBalance(balanceCoin) ?? Zero

  const [fromValue, setFromValue] = useState('')
  const [toValue, setToValue] = useState('')

  const fromRef = useRef(null)
  const toRef = useRef(null)

  const [priceImpact, setPriceImpact] = useState(Zero)
  const [exchangeRate, setExchangeRate] = useState(Zero)

  const [error, setError] = useState(null)

  const [lastChangeType, setLastChangeType] = useState('from')

  const swapableTokenSymbols = swapableTokens.map((i) => i.symbol)

  function updateUrlParams(params) {
    history.replace(`${SWAP_PATH}?${stringifyParams(params)}`)
  }

  const {
    triggerChainSwitch,
    activeChainId: fromChainId,
    setActiveChainId: setFromChainId,
  } = useNetworkController()

  useEffect(() => {
    let newFromCoin
    let newToCoin
    if (swapableTokenSymbols.includes(fromCoin.symbol)) {
      newFromCoin = fromCoin
    } else {
      newFromCoin = swapableTokens[0]
    }

    if (swapableTokenSymbols.includes(toCoin.symbol)) {
      newToCoin = toCoin
    } else {
      newToCoin = swapableTokens[1]
    }

    if (newFromCoin.symbol === newToCoin.symbol) {
      newToCoin = swapableTokens.filter(
        (t) => t.symbol !== newFromCoin.symbol
      )[0]
    }

    setFromCoin(newFromCoin)
    setToCoin(newToCoin)
  }, [chainId])

  function swapFromToCoins() {
    setFromCoin(toCoin)
    setToCoin(fromCoin)
    if (lastChangeType === 'from') {
      setToValue('')
    } else {
      setFromValue('')
    }
    setPriceImpact(Zero)
    setExchangeRate(Zero)
  }

  function onSelectFromCoin(coin, checkPool = true) {
    if (checkPool) {
      const info = getInfoMultiCoin(coin, toCoin, priorityRanking)
      // setPoolName(info.poolName)
      onSelectToCoin(info.otherCoin, false)
    }

    if (coin.symbol === toCoin.symbol) {
      swapFromToCoins()
    } else {
      setError(null)
      setFromCoin(coin)
    }
  }

  function resetRates() {
    setPriceImpact(Zero)
    setExchangeRate(Zero)
  }

  function onSelectToCoin(coin, checkPool = true) {
    if (checkPool) {
      const info = getInfoMultiCoin(coin, fromCoin, priorityRanking)
      // setPoolName(info.poolName)
      onSelectFromCoin(info.otherCoin, false)
    }

    if (coin.symbol === fromCoin.symbol) {
      swapFromToCoins()
    } else {
      setError(null)
      setToCoin(coin)
      setToValue('')
      if (lastChangeType === 'to') {
        setFromValue('')
      }
      resetRates()
    }
  }

  const tokenIndexFrom = poolTokens.findIndex((i) =>
    matchSymbolWithinPool(i, fromCoin)
  )
  const tokenIndexTo = poolTokens.findIndex((i) =>
    matchSymbolWithinPool(i, toCoin)
  )

  function triggerRateAndImpact({ amountToGive, amountToReceive }) {
    setExchangeRate(
      calculateExchangeRate(
        amountToGive,
        fromCoin.decimals[chainId],
        amountToReceive,
        toCoin.decimals[chainId]
      )
    )
  }

  function checkIfBalanceSufficient({ amountToGive }) {
    if (amountToGive.gt(fromBalance)) {
      setError('Insufficent Balance')
    } else {
      setError(null)
    }
  }

  async function calculateSwapAmount() {
    if (swapContract == null) return

    let cleanedFromValue = sanitizeValue(fromValue)
    if (checkCleanedValue(cleanedFromValue)) {
      setToValue('')
      return
    }

    const amountToGive = parseUnits(
      cleanedFromValue,
      fromCoin.decimals[chainId]
    )

    checkIfBalanceSufficient({ amountToGive })

    const amountToReceive = await calcAmountToRecieve({
      swapContract,
      tokenIndexFrom,
      tokenIndexTo,
      amountToGive,
    })

    if (sanitizeValue(fromRef.current.value) == sanitizeValue(fromValue)) {
      setToValue(formatUnits(amountToReceive, toCoin.decimals[chainId]))
      triggerRateAndImpact({ amountToGive, amountToReceive })
    }
  }

  async function calculateInverseSwapAmount() {
    if (swapContract == null) return
    const cleanedToValue = sanitizeValue(toValue)
    if (checkCleanedValue(cleanedToValue)) {
      setFromValue('')
      return
    }

    const amountToReceive =
      parseUnits(cleanedToValue, toCoin.decimals[chainId]) ?? One

    const amountToGive = await estimateAmountToGive({
      targetAmountToRecieve: amountToReceive,
      swapContract,
      tokenIndexFrom,
      tokenIndexTo,
      fromCoin,
      toCoin,
      chainId,
    })

    checkIfBalanceSufficient({ amountToGive })
    if (sanitizeValue(toRef.current.value) == sanitizeValue(toValue)) {
      setFromValue(formatUnits(amountToGive, fromCoin.decimals[chainId]))
      triggerRateAndImpact({ amountToGive, amountToReceive })
    }
  }

  useEffect(() => {
    updateUrlParams({
      inputCurrency: fromCoin.symbol,
      outputCurrency: toCoin.symbol,
    })
  }, [fromCoin, toCoin, chainId])

  useEffect(() => {
    if (lastChangeType == 'from' && chainId == debouncedChainId) {
      calculateSwapAmount()
    }
  }, [fromCoin, toCoin, fromValue, lastChangeType, chainId, debouncedChainId])

  useEffect(() => {
    if (lastChangeType == 'to' && chainId == debouncedChainId) {
      calculateInverseSwapAmount()
    }
  }, [fromCoin, toCoin, toValue, lastChangeType, chainId, debouncedChainId])

  function onChangeFromAmount(value) {
    setLastChangeType('from')
    if (!(value.split('.')[1]?.length > fromCoin.decimals[chainId])) {
      setFromValue(value)
    }
  }

  function onChangeToAmount(value) {
    setLastChangeType('to')
    if (!(value.split('.')[1]?.length > toCoin.decimals[chainId])) {
      setToValue(value)
    }
  }

  async function onChangeChain(itemChainId) {
    setLastChangeType('from')
    triggerChainSwitch(itemChainId)
      .then(() => {
        console.log({ itemChainId })
        setFromChainId(itemChainId)
        if (itemChainId == toChainId) {
          setToChainId(fromChainId)
        }
      })
      .catch((e) => {
        console.error(e)
      })
  }

  return (
    <>
      <div className="flex justify-between mb-5 ml-5 mr-5">
        <PageHeader title="Swap" subtitle="Exchange stablecoins on-chain." />
      </div>
      <SwapCard
        {...{
          swapableTokens,
          fromCoin,
          fromValue,
          toCoin,
          toValue,
          onChangeChain,
          onSelectFromCoin,
          onSelectToCoin,
          swapFromToCoins,
          poolName,
          onChangeFromAmount,
          onChangeToAmount,
          error,
          priceImpact,
          exchangeRate,
          fromRef,
          toRef,
        }}
      />
      <ActionCardFooter link={HOW_TO_SWAP_URL} />
    </>
  )
}
