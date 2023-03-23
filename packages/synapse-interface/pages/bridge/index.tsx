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
  DEFAULT_FROM_CHAIN,
  DEFAULT_TO_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_TOKEN,
} from '@/constants/bridge'
console.log('BRIDGE_CHAINS_BY_TYPE', BRIDGE_CHAINS_BY_TYPE)
console.log('BRIDGABLE_TOKENS', BRIDGABLE_TOKENS)
console.log('BRIDGE_SWAPABLE_TOKENS_BY_TYPE', BRIDGE_SWAPABLE_TOKENS_BY_TYPE)
export default function BridgePage() {
  const tokenSymbolToToken = (chainId: number, symbol: string) => {
    const token = _.find(BRIDGABLE_TOKENS[chainId], (token) => {
      return token.symbol === symbol
    })
    return token
  }

  // Get data from wagmi.
  const { address } = useAccount()
  const { chain: fromChainIdRaw } = useNetwork()
  const { error: networkSwitchError, switchNetwork } = useSwitchNetwork()

  // Init ChainIds
  const [fromChainId, setFromChainId] = useState(DEFAULT_FROM_CHAIN)

  // Get data from query url.
  const router = useRouter()
  const {
    outputChain: toChainIdUrl,
    inputCurrency: fromTokenSymbolUrl,
    outputCurrency: toTokenSymbolUrl,
  } = router.query
  console.log(
    'toTokenSymbolUrltoTokenSymbolUrltoTokenSymbolUrltoTokenSymbolUrl',
    toChainIdUrl,
    fromTokenSymbolUrl,
    toTokenSymbolUrl
  )
  // let tempToChainId = DEFAULT_TO_CHAIN
  // if (toChainIdUrl) {
  //   tempToChainId = Number(toChainIdUrl)
  // }
  // if (tempToChainId === fromChainId) {
  //   tempToChainId =
  //     fromChainId === DEFAULT_FROM_CHAIN ? DEFAULT_TO_CHAIN : DEFAULT_FROM_CHAIN
  // }
  // console.log('tempToChainId', tempToChainId)

  // // Don't update the state if the symbol isn't valid
  // let tempFromToken = DEFAULT_FROM_TOKEN
  // if (fromTokenSymbolUrl) {
  //   let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
  //   if (token) {
  //     tempFromToken = token
  //   }
  // }

  // // Don't update the state if the symbol isn't valid
  // let tempToToken = DEFAULT_TO_TOKEN
  // if (toTokenSymbolUrl) {
  //   let token = tokenSymbolToToken(tempToChainId, String(toTokenSymbolUrl))
  //   if (token) {
  //     tempToToken = token
  //   }
  // }

  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  const [lastToChainId, setLastToChainId] = useState(DEFAULT_TO_CHAIN)

  // Init token
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)
  const [lastFromToken, setLastFromToken] = useState(DEFAULT_TO_TOKEN)

  // Handle entry/amounts in the card
  const [fromValue, setFromValue] = useState('')
  const [toValue, setToValue] = useState('')

  const [lastChangeType, setLastChangeType] = useState('from')
  const [swapableType, setSwapableType] = useState(DEFAULT_SWAPABLE_TYPE)

  // dependant on from chain changes
  const [fromBridgeableTokens, setFromBridgeableTokens] = useState(
    BRIDGABLE_TOKENS[DEFAULT_FROM_CHAIN]
  )

  // dependant on from chain changes, form token changes, to chain changes
  const [toBridgeableTokens, setToBridgeableTokens] = useState(
    BRIDGABLE_TOKENS[DEFAULT_TO_CHAIN]
  )

  const [toBridgeableChains, setToBridgeableChains] = useState(
    BRIDGE_CHAINS_BY_TYPE[String(DEFAULT_FROM_TOKEN.swapableType)]
  )
  console.log('toBridgeableChains', toBridgeableChains)
  // Auxiliary data
  const [priceImpact, setPriceImpact] = useState(Zero)
  const [exchangeRate, setExchangeRate] = useState(Zero)
  const [feeAmount, setFeeAmount] = useState(Zero)
  const [error, setError] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')
  const fromRef = useRef(null)
  const toRef = useRef(null)

  // Handle wagmi changes
  useEffect(() => {
    setFromChainId(Number(fromChainIdRaw?.id))
  }, [fromChainIdRaw])

  // // Handle url changes
  useEffect(() => {
    console.log('SDLKSLDJKLSDJLKSDJ', toChainIdUrl)
    if (toChainIdUrl) {
      console.log('toChainIdUrl', toChainIdUrl)
      setToChainId(Number(toChainIdUrl))
    }
  }, [toChainIdUrl])

  useEffect(() => {
    // Don't update the state if the symbol isn't valid
    if (fromTokenSymbolUrl) {
      let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
      if (token) {
        setFromToken(token)
      }
    }
  }, [])
  useEffect(() => {
    // Don't update the state if the symbol isn't valid
    if (toTokenSymbolUrl) {
      let token = tokenSymbolToToken(toChainId, String(toTokenSymbolUrl))
      if (token) {
        setToToken(token)
      }
    }
  }, [])

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

  // Keeps the url in sync with the state
  // useEffect(() => {
  //   router.push({
  //     pathname: BRIDGE_PATH,

  //     query: {
  //       outputChain: toChainId,
  //       inputCurrency: fromToken.symbol,
  //       outputCurrency: toToken.symbol,
  //     },
  //   })
  // }, [toChainId, fromToken, toToken])

  const updateUrlParams = ({
    outputChain,
    inputCurrency,
    outputCurrency,
  }: {
    outputChain: any
    inputCurrency: any
    outputCurrency: any
  }) => {
    router.push({
      pathname: BRIDGE_PATH,
      query: {
        outputChain: outputChain,
        inputCurrency: inputCurrency,
        outputCurrency: outputCurrency,
      },
    })
  }
  // handles the case if the user changes the fromChainId and the toChainId is the same
  useEffect(() => {
    if (fromChainId === toChainId) {
      setToChainId(lastToChainId)
      // updateUrlParams({
      //   outputChain: lastToChainId,
      //   inputCurrency: fromToken.symbol,
      //   outputCurrency: toToken.symbol,
      // })
    }
  }, [fromChainId])

  // Handles when chains are flipped or user creates toChainId == fromChainId condition
  const handleChainFlip = async () => {
    // let oldFromChainId = fromChainId
    // let oldToChainId = toChainId
    switchNetwork?.(toChainId)
    // setToChainId(fromChainId)
  }

  // Changes destination change when the user changes the toChainId
  const handleFromChainChange = (chainId: number) => {
    setLastChangeType('from')
    switchNetwork?.(chainId)
  }

  const handleToChainChange = (chainId: number) => {
    setLastChangeType('to')
    setToChainId(chainId)
    updateUrlParams({
      outputChain: chainId,
      inputCurrency: fromToken.symbol,
      outputCurrency: toToken.symbol,
    })
  }

  const getBridgeableTokens = (chainId: number, token: Token) => {
    let newToChain = toChainId

    // handle case where the token is not bridgeable to the current destination chain
    // get all chains where the token is bridgeable
    let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[String(token.swapableType)]
    console.log(
      'bridgeabbridgeableChainsbridgeableChainsleChains',
      !bridgeableChains.includes(String(newToChain)),
      bridgeableChains,
      chainId
    )
    if (!bridgeableChains.includes(String(newToChain))) {
      newToChain =
        Number(bridgeableChains[0]) === fromChainId
          ? Number(bridgeableChains[1])
          : Number(bridgeableChains[0])
    }

    let newTokens =
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId][String(token.swapableType)]
    return { newTokens, newToChain }
  }
  const handleTokenChange = (token: Token, type: 'from' | 'to') => {
    /*
1. set the new token
2. if setting the origin, set the swapable types for the destination
3. set last switch?
4. calculate the bridge amount for destination/do a quote call to the sdk
5. update the url

*/

    // set the new token
    if (type == 'from') {
      console.log('from token change', token)
      setFromToken(token)

      // set the bridgeable tokens
      let { newTokens: bridgeableTokens, newToChain } = getBridgeableTokens(
        fromChainId,
        token
      )
      console.log('bridgeableTokens', bridgeableTokens)
      let bridgeableToken: Token = bridgeableTokens[0]
      setToBridgeableTokens(bridgeableTokens)
      setToToken(bridgeableToken)
      setToBridgeableChains(BRIDGE_CHAINS_BY_TYPE[String(token.swapableType)])
      setToChainId(newToChain)
      updateUrlParams({
        outputChain: newToChain,
        inputCurrency: token.symbol,
        outputCurrency: bridgeableToken.symbol,
      })
    } else {
      setToToken(token)
      updateUrlParams({
        outputChain: toChainId,
        inputCurrency: fromToken.symbol,
        outputCurrency: token.symbol,
      })
    }
    // let bridgeableTokens = getBridgeableTokens(fromChainId, token)

    //
    // reset the

    // check swapable types

    // dont think i need this but will double check
    // if (toCoin.symbol == 'WETH' && toChainId == ChainId.KLAYTN) {
    //   setToCoin(KLAYTN_WETH)
    // }

    // if (fromCoin.symbol == 'WETH' && fromChainId == ChainId.KLAYTN) {
    //   setFromCoin(KLAYTN_WETH)
    // }
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
                      address={address}
                      fromChainId={fromChainId}
                      toChainId={toChainId}
                      onSelectFromChain={handleFromChainChange}
                      onSelectToChain={handleToChainChange}
                      swapFromToChains={handleChainFlip}
                      onChangeFromAmount={onChangeFromAmount}
                      onChangeToAmount={onChangeToAmount}
                      fromCoin={fromToken}
                      toCoin={toToken}
                      possibleChains={toBridgeableChains}
                      handleTokenChange={handleTokenChange}
                      toBridgeableTokens={toBridgeableTokens}
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
