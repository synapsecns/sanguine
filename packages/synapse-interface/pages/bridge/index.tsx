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
  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  const [lastToChainId, setLastToChainId] = useState(DEFAULT_TO_CHAIN)

  // Init token
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)

  // Handle entry/amounts in the card
  const [fromValue, setFromValue] = useState('')
  const [toValue, setToValue] = useState('')

  const [toBridgeableTokens, setToBridgeableTokens] = useState(
    BRIDGABLE_TOKENS[DEFAULT_TO_CHAIN]
  )

  const [toBridgeableChains, setToBridgeableChains] = useState(
    BRIDGE_CHAINS_BY_TYPE[String(DEFAULT_FROM_TOKEN.swapableType)].filter(
      (chain) => Number(chain) !== DEFAULT_FROM_CHAIN
    )
  )

  const [lastChangeType, setLastChangeType] = useState('from')
  // const [swapableType, setSwapableType] = useState(DEFAULT_SWAPABLE_TYPE)

  // Handle wagmi changes
  useEffect(() => {
    if (address === undefined) {
      setFromChainId(DEFAULT_FROM_CHAIN)
    } else {
      setFromChainId(Number(fromChainIdRaw?.id))
    }
  }, [fromChainIdRaw])
  // Get data from query url.
  const router = useRouter()

  // Handle url changes
  useEffect(() => {
    const {
      outputChain: toChainIdUrl,
      inputCurrency: fromTokenSymbolUrl,
      outputCurrency: toTokenSymbolUrl,
    } = router.query

    let fromChainTokensByType = Object.values(
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[fromChainId]
    )
    let maxTokenLength = 0
    let token = fromChainTokensByType[0][0]
    fromChainTokensByType.map((tokenArr, i) => {
      if (tokenArr.length > maxTokenLength) {
        maxTokenLength = tokenArr.length
        token = tokenArr[0]
      }
    })

    let tempFromToken: Token = token
    if (fromTokenSymbolUrl) {
      let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
      if (token) {
        tempFromToken = token
      }
    }
    // if (tempToChainId === fromChainId) {
    //   tempToChainId =
    //     fromChainId === DEFAULT_FROM_CHAIN
    //       ? DEFAULT_TO_CHAIN
    //       : DEFAULT_FROM_CHAIN
    // }
    // let tempToChainId = DEFAULT_TO_CHAIN
    // if (toChainIdUrl) {
    //   tempToChainId = Number(toChainIdUrl)
    // }

    // let tempToToken = undefined
    // if (toTokenSymbolUrl) {
    //   let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
    //   if (token) {
    //     tempFromToken = token
    //   }
    // }
    handleNewFromToken(
      tempFromToken,
      toChainIdUrl ? Number(toChainIdUrl) : undefined,
      toTokenSymbolUrl ? String(toTokenSymbolUrl) : undefined,
      false
    )
    setFromToken(tempFromToken)
    // console.log('fromChainTosadasdaskjdhaskkjhkens', fromChainTokensByType)

    // let maxTokenLength = 0
    // let token = fromChainTokensByType[0][0]
    // fromChainTokensByType.map((tokenArr, i) => {
    //   if (tokenArr.length > maxTokenLength) {
    //     maxTokenLength = tokenArr.length
    //     token = tokenArr[0]
    //   }
    // })

    // let tempFromToken: Token = token
    // if (fromTokenSymbolUrl) {
    //   let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
    //   if (token) {
    //     tempFromToken = token
    //   }
    // }
    // setFromToken(tempFromToken)

    //BRIDGE_SWAPABLE_TOKENS_BY_TYPE[toChainId][String(toToken.swapableType)]

    // let tempToToken: Token = BRIDGE_SWAPABLE_TOKENS_BY_TYPE[toChainId][tempFromToken.swapableType][0]
    // if (toTokenSymbolUrl) {
    //   let token = tokenSymbolToToken(toChainId, String(toTokenSymbolUrl))
    //   if (token) {
    //     tempToToken = token
    //   }
    // }

    // setToToken(tempToToken)
  }, [router.query])

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

  // // dependant on from chain changes
  // const [fromBridgeableTokens, setFromBridgeableTokens] = useState(
  //   BRIDGABLE_TOKENS[DEFAULT_FROM_CHAIN]
  // )

  // dependant on from chain changes, form token changes, to chain changes

  // console.log('toBridgeableChains', toBridgeableChains)
  // Auxiliary data
  const [priceImpact, setPriceImpact] = useState(Zero)
  const [exchangeRate, setExchangeRate] = useState(Zero)
  const [feeAmount, setFeeAmount] = useState(Zero)
  const [error, setError] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')
  const fromRef = useRef(null)
  const toRef = useRef(null)

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

  // //Keeps the url in sync with the state
  // useEffect(() => {
  //   // console.log("SLKdhsdhashdCUMMM")
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

  const handleNewFromToken = (
    token: Token,
    positedToChain: number | undefined,
    positedToSymbol: string | undefined,
    updateUrl: boolean
  ) => {
    console.log(
      'CHECK DESS INSPUT',
      token,

      positedToChain,
      positedToSymbol,
      updateUrl
    )

    let newToChain = positedToChain ? Number(positedToChain) : DEFAULT_TO_CHAIN

    let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[
      String(token.swapableType)
    ].filter((chainId) => Number(chainId) !== fromChainId)
    let swapExceptionsArr: number[] =
      token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]
    console.log(
      'LOVE YOU BABY1',
      swapExceptionsArr,
      token?.swapExceptions,
      newToChain
    )
    if (swapExceptionsArr?.length > 0) {
      bridgeableChains = swapExceptionsArr.map((chainId) => String(chainId))
      console.log('LOVE YOU BABY', bridgeableChains)
    }

    // TODO filter above
    if (!bridgeableChains.includes(String(newToChain))) {
      newToChain =
        Number(bridgeableChains[0]) === fromChainId
          ? Number(bridgeableChains[1])
          : Number(bridgeableChains[0])
    }
    let positedToToken = positedToSymbol
      ? tokenSymbolToToken(newToChain, positedToSymbol)
      : undefined

    let bridgeableTokens: Token[] =
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[newToChain][String(token.swapableType)]

    if (swapExceptionsArr?.length > 0) {
      bridgeableTokens = bridgeableTokens.filter(
        (toToken) => toToken.symbol === token.symbol
      )
    }
    console.log(
      'CHECK DESSS',
      positedToToken,
      bridgeableTokens,
      newToChain,
      positedToToken && token.swapableType === positedToToken.swapableType
    )
    let bridgeableToken: Token =
      positedToToken && token.swapableType === positedToToken.swapableType
        ? positedToToken
        : bridgeableTokens[0]

    setToToken(bridgeableToken)
    setToBridgeableTokens(bridgeableTokens)
    setToBridgeableChains(bridgeableChains)
    if (updateUrl) {
      updateUrlParams({
        outputChain: newToChain,
        inputCurrency: token.symbol,
        outputCurrency: bridgeableToken.symbol,
      })
    } else {
      setToChainId(Number(newToChain))
    }
  }

  // useEffect(() => {
  //   handleNewFromToken(fromToken, toChainId, toToken, true)
  // }, [fromToken])

  // handles the case if the user changes the fromChainId and the toChainId is the same
  useEffect(() => {
    if (fromChainId === toChainId) {
      setToChainId(lastToChainId)
      updateUrlParams({
        outputChain: lastToChainId,
        inputCurrency: fromToken.symbol,
        outputCurrency: toToken.symbol,
      })
    }
  }, [fromChainId])

  // Handles when chains are flipped or user creates toChainId == fromChainId condition
  const handleChainFlip = async () => {
    // let oldFromChainId = fromChainId
    // let oldToChainId = toChainId
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      switchNetwork?.(toChainId)
    }
    // setToChainId(fromChainId)
  }

  // Changes destination change when the user changes the toChainId
  const handleFromChainChange = (chainId: number) => {
    setLastChangeType('from')
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      switchNetwork?.(chainId)
    }
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

  const handleTokenChange = (token: Token, type: 'from' | 'to') => {
    /*
1. set the new token
2. if setting the origin, set the swapable types for the destination
3. set last switch?
4. calculate the bridge amount for destination/do a quote call to the sdk
5. update the url

*/
    console.log('start start start start start start start start start start')
    // set the new token
    if (type == 'from') {
      console.log('from token change', token, token.swapableType, token.symbol)

      setFromToken(token)
      handleNewFromToken(token, toChainId, toToken.symbol, true)
      // setToToken(bridgeableToken)
      // setToChainId(newToChain)
    } else {
      console.log('to token change', token)
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
