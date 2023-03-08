import _ from 'lodash'
import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import { useRouter } from 'next/router'
import { useRef, useState, useEffect } from 'react'
import { Zero } from '@ethersproject/constants'
import { Token } from '@utils/classes/Token'

// import BridgeCard from './BridgeCard'
// import BridgeWatcher from './BridgeWatcher'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { useAccount, useNetwork } from 'wagmi'
import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import { ChainId } from '@constants/networks'
import BridgeCard from './BridgeCard'
import {
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_TYPES_BY_CHAIN,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  BRIDGABLE_TOKENS,
  // BRIDGE_SWAPABLE_TOKENS_BY_CHAIN,
} from '@constants/tokens'

import {
  DEFAULT_SWAPABLE_TYPE,
  DEFAULT_TO_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_TOKEN,
} from '@/constants/bridge'

export default function BridgePage() {
  // Get data from query url.
  const router = useRouter()
  const {
    outputChain: toChainIdUrl,
    inputCurrency: fromTokenSymbolUrl,
    outputCurrency: toTokenSymbolUrl,
  } = router.query

  // Get data from wagmi.
  const { address } = useAccount()
  const { chain: fromChainIdRaw } = useNetwork()

  // Init ChainIds
  const fromChainId = Number(fromChainIdRaw?.id)
  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)

  // Init token
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)

  // Handle entry/amounts in the card
  const [fromValue, setFromValue] = useState('')
  const [toValue, setToValue] = useState('')

  const [lastChangeType, setLastChangeType] = useState('from')
  const [swapableType, setSwapableType] = useState(DEFAULT_SWAPABLE_TYPE)

  // Auxiliary data
  const [priceImpact, setPriceImpact] = useState(Zero)
  const [exchangeRate, setExchangeRate] = useState(Zero)
  const [feeAmount, setFeeAmount] = useState(Zero)
  const [error, setError] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')
  const fromRef = useRef(null)
  const toRef = useRef(null)

  // Handle url changes
  useEffect(() => {
    if (toChainIdUrl) {
      setToChainId(Number(toChainIdUrl))
    }

    // Don't update the state if the symbol isn't valid
    if (fromTokenSymbolUrl) {
      let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
      if (token) {
        setFromToken(token)
      }
    }

    // Don't update the state if the symbol isn't valid
    if (toTokenSymbolUrl) {
      let token = tokenSymbolToToken(toChainId, String(toTokenSymbolUrl))
      if (token) {
        setFromToken(token)
      }
    }
  }, [toChainIdUrl, fromTokenSymbolUrl, toTokenSymbolUrl])

  // Helpers
  const resetRates = () => {
    setPriceImpact(Zero)
    setExchangeRate(Zero)
  }

  // Keeps the url in sync with the state
  const updateUrlParams = () => {
    router.push({
      pathname: BRIDGE_PATH,
      query: {
        outputChain: toChainId,
        inputCurrency: fromToken.symbol,
        outputCurrency: toToken.symbol,
      },
    })
  }

  const handleChainChange = (to: boolean, chainId: number) => {}

  // const toChainId =
  //   outputChain && BRIDGABLE_TOKENS[Number(outputChain)]
  //     ? Number(outputChain)
  //     : DEFAULT_TO_CHAIN
  // const fromQuery = inputCurrency ? String(inputCurrency) : DEFAULT_FROM_COIN
  // const toQuery = outputCurrency ? String(outputCurrency) : DEFAULT_TO_COIN

  // const fromChainId =
  //   fromChainIdUrl?.id && BRIDGABLE_TOKENS[Number(fromChainIdUrl?.id)]
  //     ? fromChainIdUrl?.id
  //     : DEFAULT_FROM_CHAIN

  // // set bridgeable tokens
  // const fromChainTokens = BRIDGABLE_TOKENS[fromChainId]
  // const toChainTokens = BRIDGABLE_TOKENS[toChainQuery]

  // console.log('FROMCHAIN', fromChainTokens)

  // console.log('TO', toChainTokens, toChainQuery)

  // gets the from/to token objects from the url params.

  // init state needed for bridge

  // const [fromChainId, setFromChainId] = useState(DEFAULT_FROM_CHAIN)
  // const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)

  // const [fromCoin, setFromCoin] = useState(defaultFrom ?? fromChainTokens?.[0])
  // const [toCoin, setToCoin] = useState(defaultTo ?? toChainTokens?.[0])

  // const [lastChangeType, setLastChangeType] = useState('from')

  // const fromTokenSymbols = fromChainTokens?.map((i) => i.symbol)
  // const toTokenSymbols = toChainTokens?.map((i) => i.symbol)

  // Set data from query url.

  // Initialize state variables needed for bridging

  // useEffect(() => {
  //   updateUrlParams()
  // }, [fromCoin, toCoin])
  // console.log(
  //   'GUMM',
  //   fromChainTokens,
  //   toChainTokens,
  //   fromCoin,
  //   fromValue,
  //   toCoin,
  //   toValue,

  //   error,
  //   priceImpact,
  //   exchangeRate,
  //   feeAmount,
  //   fromRef,
  //   toRef,
  //   destinationAddress,
  //   setDestinationAddress
  // )

  // const onSelectToCoin = (token: Token) => {
  //   setLastChangeType('to')
  //   setError('')
  //   setToCoin(token)
  //   setToValue('')
  //   if (lastChangeType === 'to') {
  //     setFromValue('')
  //   }
  //   resetRates()
  // }
  // async function onSelectFromChain(itemChainId: number) {
  //   setLastChangeType('from')
  //   itemChainId
  //   if (itemChainId == toChainId) {
  //     setToChainId(ch)
  //   }
  //   triggerChainSwitch(itemChainId)
  //     .then(() => {
  //       console.log({ itemChainId })
  //       setFromChainId(itemChainId)
  //       // if (itemChainId == ChainId.TERRA) {

  //       // } else
  //     })
  //     .catch((e) => {
  //       console.error(e)
  //     })
  // }

  // async function onSelectToChain(itemChainId) {
  //   setLastChangeType('to')
  //   if (itemChainId == fromChainId) {
  //     triggerChainSwitch(toChainId).then(() => {
  //       setFromChainId(toChainId)
  //       setToChainId(itemChainId)
  //     })
  //   } else {
  //     setToChainId(itemChainId)
  //   }
  // }
  return (
    <LandingPageWrapper>
      <main className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none">
        <div className="items-center px-4 py-8 mx-auto mt-4 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <div>
            <Grid
              cols={{ xs: 1 }}
              gap={6}
              className="justify-center px-2 py-16 sm:px-6 md:px-8"
            >
              {/* <button
                onClick={() => {
                  console.log(BRIDGABLE_TOKENS[toChainQuery], toChainQuery)
                  setFromCoin(BRIDGABLE_TOKENS[toChainQuery][2])
                }}
              >
                asdsadsa
              </button> */}
              {/* <HarmonyCheck fromChainId={fromChainId} toChainId={toChainId} /> */}
              <div className="flex justify-center">
                <div className="pb-3 place-self-center">
                  {fromToken && fromChainId ? (
                    <BridgeCard
                      fromChainId={fromChainId}
                      toChainId={toChainId}
                      onSelectFromCoin={() => null}
                      onSelectToCoin={() => null}
                      onSelectFromChain={() => null}
                      onSelectToChain={() => null}
                      swapFromToChains={() => null}
                      onChangeFromAmount={() => null}
                      onChangeToAmount={() => null}
                      fromCoin={fromToken}
                      toCoin={toToken}
                      {...{
                        fromValue,
                        toValue,
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
                  ) : null}
                  <ActionCardFooter link={HOW_TO_BRIDGE_URL} />
                </div>
              </div>
              <div>{/* <BridgeWatcher /> */}</div>
            </Grid>
          </div>
        </div>
      </main>
    </LandingPageWrapper>
  )
}

// export function HarmonyCheck({ fromChainId, toChainId }: { fromChainId: number; toChainId: number }) {
//   return (
//     <>
//       {(toChainId === ChainId.HARMONY || fromChainId === ChainId.HARMONY) && (
//         <div
//           className={`bg-gray-800 shadow-lg pt-3 px-6 pb-6 rounded-lg text-white`}
//         >
//           The native Harmony bridge has been exploited, which lead to a hard depeg of the following Harmony-specific tokens: 1DAI, 1USDC, 1USDT, 1ETH.
//           <br /> Please see the{' '}
//           <a
//             className="underline"
//             href="https://twitter.com/harmonyprotocol/status/1540110924400324608"
//           >
//             official Harmony Twitter
//           </a>{' '}
//           for status updates and exercise caution when interacting with Harmony.
//         </div>
//       )}
//     </>
//   )
// }

const tokenSymbolToToken = (chainId: number, symbol: string) => {
  const token = _.find(BRIDGABLE_TOKENS[chainId], (token) => {
    return token.symbol === symbol
  })
  return token
}
