import _ from 'lodash'
import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import { useRouter } from 'next/router'
import { useRef, useState, useEffect } from 'react'
import { Zero } from '@ethersproject/constants'

// import BridgeCard from './BridgeCard'
// import BridgeWatcher from './BridgeWatcher'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { useAccount, useBalance, useNetwork } from 'wagmi'
import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import { ChainId } from '@constants/networks'
import {
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_TYPES_BY_CHAIN,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  BRIDGABLE_TOKENS,
  // BRIDGE_SWAPABLE_TOKENS_BY_CHAIN,
} from '@constants/tokens'

import {
  DEFAULT_FROM_CHAIN,
  DEFAULT_TO_CHAIN,
  DEFAULT_FROM_COIN,
  DEFAULT_TO_COIN,
} from '@/constants/bridge'
const sanitizeQueryData = (data: string) => {}
console.log(
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_TYPES_BY_CHAIN,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE
)
export default function BridgePage() {
  // Get data from query url.
  const router = useRouter()
  const { outputChain, inputCurrency, outputCurrency } = router.query
  const toChainQuery =
    outputChain && BRIDGABLE_TOKENS[Number(outputChain)]
      ? Number(outputChain)
      : DEFAULT_TO_CHAIN
  const fromQuery = inputCurrency ? String(inputCurrency) : DEFAULT_FROM_COIN
  const toQuery = outputCurrency ? String(outputCurrency) : DEFAULT_TO_COIN

  // Get data from wagmi.
  const { address } = useAccount()
  const { chain } = useNetwork()

  // set bridgeable tokens
  const fromChainTokens = BRIDGABLE_TOKENS[Number(chain?.id)] // BRIDGE_SWAPABLE_TOKENS_BY_CHAIN[fromChainId]
  const toChainTokens = BRIDGABLE_TOKENS[toChainQuery]

  // console.log('FROMCHAIN', fromChainTokens)

  // console.log('TO', toChainTokens, toChainQuery)

  // gets the from/to token objects from the url params.
  const defaultFrom = _.find(
    fromChainTokens,
    (token) => token.symbol === fromQuery
  )
  const defaultTo = _.find(toChainTokens, (token) => token.symbol === toQuery)

  console.log('defaultTo', defaultTo)

  // init state needed for bridge

  const [fromCoin, setFromCoin] = useState(defaultFrom ?? fromChainTokens?.[0])
  const [toCoin, setToCoin] = useState(defaultTo ?? toChainTokens?.[0])
  const [swapableType, setSwapableType] = useState(
    (defaultFrom ?? fromChainTokens?.[0])?.swapableType ?? 'USD'
  )
  const [fromValue, setFromValue] = useState('')
  const [toValue, setToValue] = useState('')
  const [priceImpact, setPriceImpact] = useState(Zero)
  const [exchangeRate, setExchangeRate] = useState(Zero)
  const [feeAmount, setFeeAmount] = useState(Zero)
  const [error, setError] = useState(null)
  const [destinationAddress, setDestinationAddress] = useState('')
  const fromRef = useRef(null)
  const toRef = useRef(null)
  // const [lastChangeType, setLastChangeType] = useState('from')

  const fromTokenSymbols = fromChainTokens?.map((i) => i.symbol)
  const toTokenSymbols = toChainTokens?.map((i) => i.symbol)

  // Set data from query url.
  console.log('FROM', fromCoin)
  console.log('TO', toCoin)
  console.log('Swapable Type', swapableType)
  // Initialize state variables needed for bridging
  // const [fromChainId, setFromChainId] = useState(DEFAULT_FROM_CHAIN)
  // const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  // const [fromCoin, setFromCoin] = useState(DEFAULT_FROM_COIN)
  // const [toCoin, setToCoin] = useState(DEFAULT_TO_COIN)

  const updateUrlParams = () => {
    router.push({
      pathname: BRIDGE_PATH,
      query: {
        outputChain: toChainQuery,
        inputCurrency: fromCoin.symbol,
        outputCurrency: toCoin.symbol,
      },
    })
  }

  useEffect(() => {
    updateUrlParams()
  }, [fromCoin, toCoin])

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
              <button
                onClick={() => {
                  console.log(BRIDGABLE_TOKENS[toChainQuery], toChainQuery)
                  setFromCoin(BRIDGABLE_TOKENS[toChainQuery][2])
                }}
              >
                asdsadsa
              </button>
              {/* <HarmonyCheck fromChainId={fromChainId} toChainId={toChainId} /> */}
              <div className="flex justify-center">
                <div className="pb-3 place-self-center">
                  {/* <BridgeCard
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
                /> */}
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
