import _ from 'lodash'
import { useEffect, useRef, useState } from 'react'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits, parseUnits } from '@ethersproject/units'
import { Zero } from '@ethersproject/constants'
import { useAccount, useBalance, useNetwork } from 'wagmi'
import {
  DOGECHAIN_BUSD,
  KLAYTN_DAI,
  KLAYTN_USDC,
  KLAYTN_USDT,
  KLAYTN_WETH,
  SYN,
} from '@constants/tokens/basic'
import { CHAIN_INFO_MAP, ChainId } from '@constants/networks'
import { BRIDGABLE_TOKENS } from '@constants/bridge'
import {
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  BRIDGE_SWAPABLE_TYPES_BY_CHAIN,
} from '@constants/tokens/tokenGroups'
import { checkCleanedValue } from '@utils/checkCleanedValue'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import { useNetworkController } from '@hooks/wallet/useNetworkController'
import { usePrevious } from '@hooks/usePrevious'
import { useUrlQuery } from '@hooks/useUrlQuery'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useCalculateBridgeRate } from '@hooks/useCalculateBridgeRate'
import { useTerraWallet } from '@hooks/terra/useTerraWallet'
import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import StandardPageContainer from '@layouts/StandardPageContainer'
import AnyswapWarningCard from '@components/AnyswapWarningCard'
import { ActionCardFooter } from '@components/ActionCardFooter'

import BridgeCard from './BridgeCard'
import BridgeWatcher from './BridgeWatcher'
import { sanitizeValue } from '@/utils/stringToBigNum'
import { stringifyParams } from '@/constants/urls/stringifyParams'
import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'

export default function BridgePage() {
  // get stuff from the url
  const router = useRouter()
  const { outputChain, inputCurrency, outputCurrency } = router.query
  const toChainQuery = String(outputChain)
  const fromQuery = String(inputCurrency)
  const toQuery = String(outputCurrency)

  // get the current account and chainId from wagmi (nuke)
  const { account, chainId } = useActiveWeb3React()
  //nuke
  const { terraAddress } = useTerraWallet()

  // nuke lol
  let tempFromChainId
  if (terraAddress && !account) {
    tempFromChainId = ChainId.TERRA
  } else {
    tempFromChainId = chainId
  }

  // nuke
  const { address } = useAccount()
  const { chain } = useNetwork()

  // nuke, replace with redux, replace with change chain with rainbow or wagmi
  const {
    triggerChainSwitch,
    activeChainId: fromChainId,
    setActiveChainId: setFromChainId,
  } = useNetworkController()

  // nuke this shit for now lmfao
  // seems like it handles an edgecase where a url might be fucked, can handle that up
  const previousFromChainId = usePrevious(fromChainId)
  let defaultToChain
  if (toChainQuery?.length > 0 && parseInt(toChainQuery) != fromChainId) {
    // This specific if block is suspect. if state gets weird, drop url state
    defaultToChain = parseInt(toChainQuery)
  } else {
    // defaults the to chain to arb and if the current chain is arb then it does something random thats not arb (syntarded)
    if (fromChainId === ChainId.ARBITRUM) {
      defaultToChain = _.keys(CHAIN_INFO_MAP)
        .map((i) => parseInt(i))
        .filter((i) => i != fromChainId)[0]
    } else {
      defaultToChain = ChainId.ARBITRUM
    }
  }

  // move this up higher
  const [toChainId, setToChainId] = useState(defaultToChain) //(fromChainId !== 1) ? 1 : 56

  // need to get the bridge rate from sdk, move this to function i think
  const calculateBridgeRate = useCalculateBridgeRate({ fromChainId, toChainId })

  // ok so like, we want to have do this without hardcoding, need to figure out how to do this in some sort of util.
  const fromChainTokens = BRIDGABLE_TOKENS[fromChainId] // BRIDGE_SWAPABLE_TOKENS_BY_CHAIN[fromChainId]
  const toChainTokens = BRIDGABLE_TOKENS[toChainId] // BRIDGE_SWAPABLE_TOKENS_BY_CHAIN[toChainId]

  // gets the default from  tokens from the url
  const defaultFrom = _.find(
    fromChainTokens,
    (token) => token.symbol === fromQuery
  )
  const defaultTo = _.find(toChainTokens, (token) => token.symbol === toQuery)

  const [fromCoin, setFromCoin] = useState(defaultFrom ?? fromChainTokens[0])
  const [toCoin, setToCoin] = useState(defaultTo ?? toChainTokens[0])

  const [fromValue, setFromValue] = useState('')
  const [toValue, setToValue] = useState('')

  const fromRef = useRef(null)
  const toRef = useRef(null)

  const [priceImpact, setPriceImpact] = useState(Zero)
  const [exchangeRate, setExchangeRate] = useState(Zero)
  const [feeAmount, setFeeAmount] = useState(Zero)

  const [error, setError] = useState(null)

  // persist last change type
  const [lastChangeType, setLastChangeType] = useState('from')

  const defaultSwapableType = fromCoin.swapableType
  const [swapableType, setSwapableType] = useState(defaultSwapableType ?? 'USD')

  const fromTokenSymbols = fromChainTokens.map((i) => i.symbol)
  const toTokenSymbols = toChainTokens.map((i) => i.symbol)

  const [destinationAddress, setDestinationAddress] = useState('')

  // update with next router stuff
  //const router = useRouter();
  // router.query.NEWPARAMS = "VALUE"
  // router.push(router)
  function updateUrlParams(params) {
    const { pid } = router.query
    history.replace(`${BRIDGE_PATH}?${stringifyParams(params)}`)
  }

  // handle case when from/to chain ids are the same. sets the old from to the old.
  // how about handle this on the inital change, dont let them change to the same change if its selected.
  // so everything BUT arb shows up in the combo box for from when selected
  // OR could just check when cahnging if same and then just do a flip (use flip function)
  useEffect(() => {
    if (fromChainId == toChainId) {
      if (previousFromChainId) {
        setToChainId(previousFromChainId)
      }
    }
  }, [fromChainId, toChainId])

  // if the the swapable type changes, then we need to update the from/to coins
  // so if the to and from arent the same, it defaults to the first swapable type from the destination chain. vice versa for from
  // omg this is syntarded
  // so its listening for the most recent change, either its a from or to
  // and then it does the the below based on each case
  useEffect(() => {
    if (toCoin.swapableType != fromCoin.swapableType) {
      if (lastChangeType === 'from') {
        if (BRIDGE_SWAPABLE_TOKENS_BY_TYPE[toChainId][fromCoin.swapableType]) {
          setToCoin(
            BRIDGE_SWAPABLE_TOKENS_BY_TYPE[toChainId][fromCoin.swapableType][0]
          )
        }
      }
      if (lastChangeType === 'to') {
        if (BRIDGE_SWAPABLE_TOKENS_BY_TYPE[fromChainId][toCoin.swapableType]) {
          setFromCoin(
            BRIDGE_SWAPABLE_TOKENS_BY_TYPE[fromChainId][toCoin.swapableType][0]
          )
        }
      }
    }
  }, [swapableType])

  // handles chain changes
  useEffect(() => {
    let newFromCoin
    let newToCoin
    if (fromTokenSymbols.includes(fromCoin.symbol)) {
      newFromCoin = fromCoin
    } else if (
      BRIDGE_SWAPABLE_TYPES_BY_CHAIN[fromChainId].includes(
        fromCoin.swapableType
      )
    ) {
      newFromCoin =
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[fromChainId][fromCoin.swapableType][0]
    } else {
      newFromCoin = fromChainTokens[0]
    }

    if (toTokenSymbols.includes(toCoin.symbol)) {
      newToCoin = toCoin
    } else if (
      BRIDGE_SWAPABLE_TYPES_BY_CHAIN[toChainId].includes(toCoin.swapableType)
    ) {
      newToCoin =
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[toChainId][toCoin.swapableType][0]
    } else {
      newToCoin = toChainTokens[0]
    }

    if (newToCoin.swapableType != newFromCoin.swapableType) {
      if (lastChangeType === 'from') {
        setSwapableType(newFromCoin.swapableType)
      }
      if (lastChangeType === 'to') {
        setSwapableType(newToCoin.swapableType)
      }
    }

    setFromCoin(newFromCoin)
    setToCoin(newToCoin)
  }, [fromChainId, toChainId])

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

  // Handle a chain swap
  function swapFromToChains() {
    triggerChainSwitch(toChainId)
      .then(() => {
        setFromChainId(toChainId)
        setToChainId(fromChainId)
      })
      .catch((error) => console.error(error))

    if (lastChangeType === 'from') {
      setToValue('')
    } else {
      setFromValue('')
    }
    setPriceImpact(Zero)
    setExchangeRate(Zero)
    setDestinationAddress('')
  }

  function onSelectFromCoin(coin) {
    setLastChangeType('from')
    setError(null)
    setFromCoin(coin)
  }

  function resetRates() {
    setPriceImpact(Zero)
    setExchangeRate(Zero)
  }

  function onSelectToCoin(coin) {
    setLastChangeType('to')
    setError(null)
    setToCoin(coin)
    setToValue('')
    if (lastChangeType === 'to') {
      setFromValue('')
    }
    resetRates()
  }

  async function onSelectFromChain(itemChainId) {
    setLastChangeType('from')
    triggerChainSwitch(itemChainId)
      .then(() => {
        console.log({ itemChainId })
        setFromChainId(itemChainId)
        // if (itemChainId == ChainId.TERRA) {

        // } else
        if (itemChainId == toChainId) {
          setToChainId(fromChainId)
        }
      })
      .catch((e) => {
        console.error(e)
      })
  }

  async function onSelectToChain(itemChainId) {
    setLastChangeType('to')
    if (itemChainId == fromChainId) {
      triggerChainSwitch(toChainId).then(() => {
        setFromChainId(toChainId)
        setToChainId(itemChainId)
      })
    } else {
      setToChainId(itemChainId)
    }
  }

  // chenage with router stuff
  useEffect(() => {
    updateUrlParams({
      inputCurrency: fromCoin.symbol,
      outputCurrency: toCoin.symbol,
      outputChain: toChainId,
    })
  }, [fromCoin, toCoin, fromChainId, toChainId])

  function onChangeFromAmount(value) {
    setLastChangeType('from')
    if (!(value.split('.')[1]?.length > fromCoin.decimals[fromChainId])) {
      setFromValue(value)
    }
  }

  function onChangeToAmount(value) {
    setLastChangeType('to')
    if (!(value.split('.')[1]?.length > toCoin.decimals[toChainId])) {
      setToValue(value)
    }
  }

  function triggerRateAndImpact({ amountToGive, amountToReceive, bridgeFee }) {
    setFeeAmount(bridgeFee)
    setExchangeRate(
      calculateExchangeRate(
        amountToGive.sub(
          feeAmount.div(
            BigNumber.from(10).pow(18 - fromCoin.decimals[fromChainId])
          )
        ),
        fromCoin.decimals[fromChainId],
        amountToReceive,
        toCoin.decimals[toChainId]
      )
    )
  }

  // ADJUSTS THE UI
  useEffect(() => {
    const fromSwapableTypes = BRIDGE_CHAINS_BY_TYPE[fromCoin.swapableType]
    const toSwapableTypes = BRIDGE_CHAINS_BY_TYPE[toCoin.swapableType]
    const validSwapableTypes = _.intersection(
      fromSwapableTypes,
      toSwapableTypes
    )

    if (toCoin.symbol == 'WETH' && toChainId == ChainId.KLAYTN) {
      setToCoin(KLAYTN_WETH)
    }

    if (fromCoin.symbol == 'WETH' && fromChainId == ChainId.KLAYTN) {
      setFromCoin(KLAYTN_WETH)
    }

    if (lastChangeType === 'from') {
      if (
        fromCoin.swapableType == 'USD' &&
        ((toChainId == ChainId.KLAYTN && fromChainId == ChainId.ETH) ||
          (toChainId == ChainId.DOGECHAIN && fromChainId == ChainId.ETH))
      ) {
        if (fromCoin.symbol == 'USDT') {
          setFromCoin(KLAYTN_USDT)
          setToCoin(KLAYTN_USDT)
        }
        if (fromCoin.symbol == 'USDC') {
          setFromCoin(KLAYTN_USDC)
          setToCoin(KLAYTN_USDC)
        }
        if (fromCoin.symbol == 'DAI') {
          setFromCoin(KLAYTN_DAI)
          setToCoin(KLAYTN_DAI)
        }
      } else if (
        fromCoin.swapableType == 'USD' &&
        toCoin.swapableType == 'USD' &&
        toChainId == ChainId.DOGECHAIN &&
        fromChainId == ChainId.BSC
      ) {
        setFromCoin(DOGECHAIN_BUSD)
        setToCoin(DOGECHAIN_BUSD)
      } else {
        if (fromCoin.symbol === SYN.symbol) {
          setToCoin(SYN)
        }
        console.log(fromCoin.swapableType)

        // Finds which chains tokens can be bridged to (to set a new to chain by default)
        let newToChainId = BRIDGE_CHAINS_BY_TYPE[fromCoin.swapableType][0]
        // if the new chain chosen is the current from chain, switch to a random (2nd in array) chain
        if (newToChainId == fromChainId) {
          newToChainId = BRIDGE_CHAINS_BY_TYPE[fromCoin.swapableType][1]
        }

        // if token doesn't exist, it'll switch to a destination chain that has the token
        if (validSwapableTypes.length == 0) {
          onSelectToChain(newToChainId)
        } else {
          // else, continue with picking which token to display
          // if the token is not swapable in this chain combination, pick a new chain which can be bridged to
          if (!validSwapableTypes.includes(toCoin.swapableType)) {
            let targetChainId
            if (
              !BRIDGE_SWAPABLE_TYPES_BY_CHAIN[toChainId].includes(
                fromCoin.swapableType
              )
            ) {
              targetChainId = newToChainId
              onSelectToChain(newToChainId)
            } else {
              targetChainId = toChainId
            }
            if (fromCoin.swapableType != toCoin.swapableType) {
              setToCoin(
                BRIDGE_SWAPABLE_TOKENS_BY_TYPE[targetChainId][
                  fromCoin.swapableType
                ][0]
              )
            }
          }
        }
      }
    }

    if (lastChangeType === 'to') {
      if (
        toCoin.swapableType == 'USD' &&
        ((fromChainId == ChainId.KLAYTN && toChainId == ChainId.ETH) ||
          (fromChainId == ChainId.DOGECHAIN && toChainId == ChainId.ETH))
      ) {
        if (toCoin.symbol == 'USDT') {
          setFromCoin(KLAYTN_USDT)
          setToCoin(KLAYTN_USDT)
        }
        if (toCoin.symbol == 'USDC') {
          setFromCoin(KLAYTN_USDC)
          setToCoin(KLAYTN_USDC)
        }
        if (toCoin.symbol == 'DAI') {
          setFromCoin(KLAYTN_DAI)
          setToCoin(KLAYTN_DAI)
        }
      } else if (
        (fromCoin.swapableType == 'USD' &&
          toCoin.swapableType == 'USD' &&
          fromChainId == ChainId.DOGECHAIN &&
          toChainId == ChainId.BUSD) ||
        (fromChainId == ChainId.BSC && toChainId == ChainId.DOGECHAIN)
      ) {
        setFromCoin(DOGECHAIN_BUSD)
        setToCoin(DOGECHAIN_BUSD)
      } else {
        if (toCoin.symbol === SYN.symbol) {
          setFromCoin(SYN)
        }

        let newFromChainId = BRIDGE_CHAINS_BY_TYPE[toCoin.swapableType][0]
        if (newFromChainId == toChainId) {
          newFromChainId = BRIDGE_CHAINS_BY_TYPE[toCoin.swapableType][1]
        }
        if (validSwapableTypes.length == 0) {
          onSelectFromChain(newFromChainId)
        } else {
          if (!validSwapableTypes.includes(fromCoin.swapableType)) {
            if (
              !BRIDGE_SWAPABLE_TYPES_BY_CHAIN[fromChainId].includes(
                toCoin.swapableType
              )
            ) {
              onSelectFromChain(newFromChainId)
            }
            if (toCoin.swapableType != fromCoin.swapableType) {
              setFromCoin(
                BRIDGE_SWAPABLE_TOKENS_BY_TYPE[newFromChainId][
                  toCoin.swapableType
                ][0]
              )
            }
          }
        }
      }
    }
  }, [fromCoin, toCoin, toChainId, fromChainId])

  // REDO WITH SDK
  function calculateBridgeAmount() {
    const cleanedFromValue = sanitizeValue(fromValue)
    if (checkCleanedValue(cleanedFromValue)) {
      setToValue('')
      return
    }

    const amountToGive = parseUnits(
      cleanedFromValue,
      fromCoin.decimals[fromChainId]
    )

    calculateBridgeRate({ fromCoin, toCoin, amountToGive }).then(
      ({ amountToReceive, bridgeFee }) => {
        if (sanitizeValue(fromRef.current?.value) == sanitizeValue(fromValue)) {
          setToValue(formatUnits(amountToReceive, toCoin.decimals[toChainId]))
          triggerRateAndImpact({ amountToGive, amountToReceive, bridgeFee })
        }
      }
    )
  }

  useEffect(() => {
    if (fromCoin && toCoin) {
      calculateBridgeAmount()
    }
  }, [
    fromCoin.symbol,
    toCoin.symbol,
    fromValue,
    lastChangeType,
    fromChainId,
    toChainId,
    feeAmount,
  ])

  return (
    <LandingPageWrapper>
      <StandardPageContainer>
        <div>
          <Grid
            cols={{ xs: 1 }}
            gap={6}
            className="justify-center px-2 py-16 sm:px-6 md:px-8"
          >
            <HarmonyCheck fromChainId={fromChainId} toChainId={toChainId} />
            <div className="flex justify-center">
              <div className="pb-3 place-self-center">
                <BridgeCard
                  {...{
                    fromChainTokens,
                    toChainTokens,
                    fromChainId,
                    toChainId,
                    fromCoin,
                    fromValue,
                    toCoin,
                    toValue,
                    onSelectFromCoin,
                    onSelectToCoin,
                    onSelectFromChain,
                    onSelectToChain,
                    swapFromToCoins,
                    swapFromToChains,
                    onChangeFromAmount,
                    onChangeToAmount,
                    error,
                    priceImpact,
                    exchangeRate,
                    feeAmount,
                    fromRef,
                    toRef,
                    destinationAddress,
                    setDestinationAddress,
                  }}
                />
                <ActionCardFooter link={HOW_TO_BRIDGE_URL} />
              </div>
            </div>
            <div>
              <BridgeWatcher />
            </div>
            <div>
              <AnyswapWarningCard />
            </div>
          </Grid>
        </div>
      </StandardPageContainer>
    </LandingPageWrapper>
  )
}

export function HarmonyCheck({
  fromChainId,
  toChainId,
}: {
  fromChainId: number
  toChainId: number
}) {
  return (
    <>
      {(toChainId === ChainId.HARMONY || fromChainId === ChainId.HARMONY) && (
        <div
          className={`bg-gray-800 shadow-lg pt-3 px-6 pb-6 rounded-lg text-white`}
        >
          The native Harmony bridge has been exploited, which lead to a hard
          depeg of the following Harmony-specific tokens: 1DAI, 1USDC, 1USDT,
          1ETH.
          <br /> Please see the{' '}
          <a
            className="underline"
            href="https://twitter.com/harmonyprotocol/status/1540110924400324608"
          >
            official Harmony Twitter
          </a>{' '}
          for status updates and exercise caution when interacting with Harmony.
        </div>
      )}
    </>
  )
}
