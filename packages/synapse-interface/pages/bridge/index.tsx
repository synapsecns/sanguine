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
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
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
console.log('BRIDGE_CHAINS_BY_TYPE', BRIDGE_CHAINS_BY_TYPE)
console.log('BRIDGABLE_TOKENS', BRIDGABLE_TOKENS)
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
  const { switchNetwork } = useSwitchNetwork()

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
  const onChangeFromAmount = (value: string) => {
    setLastChangeType('from')
    if (
      !(
        value.split('.')[1]?.length >
        fromToken[fromChainId as keyof Token['decimals']]
      )
    ) {
      setFromValue(value)
    }
  }

  const onChangeToAmount = (value: string) => {
    setLastChangeType('to')
    if (
      !(
        value.split('.')[1]?.length >
        toToken[toChainId as keyof Token['decimals']]
      )
    ) {
      setToValue(value)
    }
  }
  const tokenSymbolToToken = (chainId: number, symbol: string) => {
    const token = _.find(BRIDGABLE_TOKENS[chainId], (token) => {
      return token.symbol === symbol
    })
    return token
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

  // Handles when chains are flipped or user creates toChainId == fromChainId condition
  const handleChainFlip = () => {
    switchNetwork?.(toChainId)
    setToChainId(fromChainId)
  }

  // Changes destination change when the user changes the toChainId
  const handleFromChainChange = (chainId: number) => {
    setLastChangeType('from')
    if (chainId == toChainId) {
      handleChainFlip()
    } else {
      switchNetwork?.(chainId)
    }
  }

  const handleToChainChange = (chainId: number) => {
    setLastChangeType('to')
    if (chainId == fromChainId) {
      handleChainFlip()
    } else {
      setToChainId(chainId)
    }
  }

  const handleTokenChange = (token: Token, type: 'from' | 'to') => {
    const fromSwapableTypes =
      BRIDGE_CHAINS_BY_TYPE[String(fromToken.swapableType)]
    const toSwapableTypes = BRIDGE_CHAINS_BY_TYPE[String(toToken.swapableType)]
    const validSwapableTypes = _.intersection(
      fromSwapableTypes,
      toSwapableTypes
    )
    // dont think i need this but will double check
    // if (toCoin.symbol == 'WETH' && toChainId == ChainId.KLAYTN) {
    //   setToCoin(KLAYTN_WETH)
    // }

    // if (fromCoin.symbol == 'WETH' && fromChainId == ChainId.KLAYTN) {
    //   setFromCoin(KLAYTN_WETH)
    // }
    if (type == 'from') {
      setFromToken(token)
    } else {
      setToToken(token)
    }
  }
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
                      onSelectFromChain={handleFromChainChange}
                      onSelectToChain={handleToChainChange}
                      onSelectToCoin={() => null}
                      onSelectFromCoin={() => null}
                      swapFromToChains={handleChainFlip}
                      onChangeFromAmount={onChangeFromAmount}
                      onChangeToAmount={onChangeToAmount}
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
