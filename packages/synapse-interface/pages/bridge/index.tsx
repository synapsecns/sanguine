import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import { useRouter } from 'next/router'
import { useNetwork, useAccount } from 'wagmi'
import { useEffect, useState, useCallback, useMemo } from 'react'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { fetchSigner, getNetwork, switchNetwork } from '@wagmi/core'
import { sortByTokenBalance, sortByVisibilityRank } from '@utils/sortTokens'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import { subtractSlippage } from '@utils/slippage'
import toast from 'react-hot-toast'
import Popup from '@components/Popup'

import {
  BRIDGABLE_TOKENS,
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  tokenSymbolToToken,
} from '@constants/tokens'
import { formatBNToString } from '@utils/bignumber/format'
import { commify } from '@ethersproject/units'
import { isAddress } from '@ethersproject/address'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'
import BridgeWatcher from './BridgeWatcher'
import { BridgeQuote } from '@/utils/types'
import { Token } from '@/utils/types'
import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import { stringToBigNum } from '@/utils/stringToBigNum'
import BridgeCard from './BridgeCard'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { checkStringIfOnlyZeroes } from '@/utils/regex'
import { timeout } from '@/utils/timeout'
import {
  DEFAULT_FROM_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_CHAIN,
  DEFAULT_TO_TOKEN,
  EMPTY_BRIDGE_QUOTE,
  EMPTY_BRIDGE_QUOTE_ZERO,
  QUOTE_POLLING_INTERVAL,
} from '@/constants/bridge'

/* TODO
  - look into getting rid of fromChainId state and just using wagmi hook (ran into problems when trying this but forgot why)
*/

const BridgePage = ({
  address,
  fromChainId,
}: {
  address: `0x${string}`
  fromChainId: number
}) => {
  const { address: currentAddress, isDisconnected } = useAccount()
  const router = useRouter()
  const { synapseSDK } = useSynapseContext()
  const [time, setTime] = useState(Date.now())
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [fromTokens, setFromTokens] = useState([])
  const [fromInput, setFromInput] = useState({ string: '', bigNum: Zero })
  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)
  const [isQuoteLoading, setIsQuoteLoading] = useState<boolean>(false)
  const [error, setError] = useState('')
  const [bridgeTxHash, setBridgeTxHash] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')
  const [toOptions, setToOptions] = useState({
    tokens: BRIDGABLE_TOKENS[DEFAULT_TO_CHAIN],
    chains: BRIDGE_CHAINS_BY_TYPE[
      String(DEFAULT_FROM_TOKEN.swapableType)
    ].filter((chain) => Number(chain) !== DEFAULT_FROM_CHAIN),
  })
  const [bridgeQuote, setBridgeQuote] =
    useState<BridgeQuote>(EMPTY_BRIDGE_QUOTE)

  let popup: string
  /*
  useEffect Trigger: onMount
  - Gets current network connected and sets it as the state.
  - Initializes polling (setInterval) func to re-retrieve quotes.
  */
  useEffect(() => {
    sortByTokenBalance(
      BRIDGABLE_TOKENS[fromChainId],
      fromChainId,
      address
    ).then((tokens) => {
      setFromTokens(tokens)
    })
    const interval = setInterval(
      () => setTime(Date.now()),
      QUOTE_POLLING_INTERVAL
    )

    return () => {
      clearInterval(interval)
    }
  }, [bridgeTxHash, fromChainId])

  useEffect(() => {
    if (!router.isReady) {
      return
    }
    const {
      outputChain: toChainIdUrl,
      inputCurrency: fromTokenSymbolUrl,
      outputCurrency: toTokenSymbolUrl,
    } = router.query

    let tempFromToken: Token = getMostCommonSwapableType(fromChainId)

    if (fromTokenSymbolUrl) {
      let token = tokenSymbolToToken(fromChainId, String(fromTokenSymbolUrl))
      if (token) {
        tempFromToken = token
      }
    }
    const { bridgeableToken, newToChain, bridgeableTokens, bridgeableChains } =
      handleNewFromToken(
        tempFromToken,
        toChainIdUrl ? Number(toChainIdUrl) : undefined,
        toTokenSymbolUrl ? String(toTokenSymbolUrl) : undefined,
        fromChainId
      )
    resetTokenPermutation(
      tempFromToken,
      newToChain,
      bridgeableToken,
      bridgeableChains,
      bridgeableTokens,
      tempFromToken.symbol,
      bridgeableToken.symbol
    )
    updateUrlParams({
      outputChain: newToChain,
      inputCurrency: fromToken.symbol,
      outputCurrency: bridgeableToken.symbol,
    })
  }, [router.isReady])

  /*
  useEffect Triggers: toToken, fromInput, toChainId, time
  - Gets a quote when the polling function is executed or any of the bridge attributes are altered.
  - Debounce quote call by calling quote price AFTER user has stopped typing for 500ms
  */
  useEffect(() => {
    let isCancelled = false

    const handleChange = async () => {
      await timeout(400) // debounce by 400ms
      if (!isCancelled) {
        if (
          fromChainId &&
          toChainId &&
          String(fromToken.addresses[fromChainId]) &&
          String(toToken.addresses[toChainId]) &&
          fromInput &&
          fromInput.bigNum.gt(Zero)
        ) {
          getQuote()
        } else {
          setBridgeQuote(EMPTY_BRIDGE_QUOTE)
        }
      }
    }
    handleChange()

    return () => {
      isCancelled = true
    }
  }, [toToken, fromInput, time, fromChainId, toChainId])

  /*
  useEffect Triggers: fromInput
  - Checks that user input is not zero. When input changes,
  - isQuoteLoading state is set to true for loading state interactions
  */
  useEffect(() => {
    const { string, bigNum } = fromInput
    const isInvalid = checkStringIfOnlyZeroes(string)
    isInvalid ? () => null : setIsQuoteLoading(true)

    return () => {
      setIsQuoteLoading(false)
    }
  }, [fromInput, fromChainId])

  /*
  Helper Function: resetTokenPermutation
  - Handles when theres a new from token/chain and all other parts of the bridge arrangement needs to be updated
  - Updates url params.
  */
  const resetTokenPermutation = (
    newFromToken: Token,
    newToChain: number,
    newToToken: Token,
    newBridgeableChains: string[],
    newBridgeableTokens: Token[],
    newFromTokenSymbol: string,
    newBridgeableTokenSymbol: string
  ) => {
    setFromToken(newFromToken)
    setToChainId(newToChain)
    setToToken(newToToken)
    setToOptions({ tokens: newBridgeableTokens, chains: newBridgeableChains })
    updateUrlParams({
      outputChain: newToChain,
      inputCurrency: newFromTokenSymbol,
      outputCurrency: newBridgeableTokenSymbol,
    })
  }

  /*
  Helper Function: resetRates
  - Called when switching from chain/token so that the from input isn't populated with stale data.
  */
  const resetRates = () => {
    setBridgeQuote(EMPTY_BRIDGE_QUOTE)
    setFromInput({ string: '', bigNum: Zero })
  }

  /*
  Helper Function: onChangeFromAmount
  - Ensures inputted data isn't too long and then sets state with the input.
  - Calculates BigNum from the input and stores in state as well (for quotes)
  */
  const onChangeFromAmount = (value: string) => {
    if (
      !(
        value.split('.')[1]?.length >
        fromToken[fromChainId as keyof Token['decimals']]
      )
    ) {
      let bigNum =
        stringToBigNum(value, fromToken.decimals[fromChainId]) ?? Zero

      setFromInput({
        string: value,
        bigNum: bigNum,
      })
    }
  }

  /*
  Helper Function: getMostCommonSwapableType
  - Returns the default token to display when switching chains. Usually returns stables or eth/wrapped eth.
  */
  const getMostCommonSwapableType = useCallback(
    (chainId: number) => {
      const fromChainTokensByType = Object.values(
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId]
      )
      let maxTokenLength = 0
      let mostCommonSwapableType: Token[] = fromChainTokensByType[0]
      fromChainTokensByType.map((tokenArr, i) => {
        if (tokenArr.length > maxTokenLength) {
          maxTokenLength = tokenArr.length
          mostCommonSwapableType = tokenArr
        }
      })

      return sortByVisibilityRank(mostCommonSwapableType)[0]
    },
    [currentAddress, isDisconnected]
  )

  /*
  Helper Function: updateUrlParams
  - Pushes chain and token changes to url
  NOTE: did not alter any variable names in case previous users have saved links of different bridging permutations.
  */
  const updateUrlParams = ({
    outputChain,
    inputCurrency,
    outputCurrency,
  }: {
    outputChain: number
    inputCurrency: string
    outputCurrency: string
  }) => {
    router.replace(
      {
        pathname: BRIDGE_PATH,
        query: {
          outputChain,
          inputCurrency,
          outputCurrency,
        },
      },
      undefined,
      { shallow: true }
    )
  }

  /*
   Helper Function: getCurrentTokenAllowance
  - Gets quote data from the Synapse SDK (from the imported provider)
  - Calculates slippage by subtracting fee from input amount (checks to ensure proper num of decimals are in use - ask someone about stable swaps if you want to learn more)
  TODO store this erc20 and signer retrieval in a state in a parent component? add to utils + use memo?
  */
  const getCurrentTokenAllowance = async (routerAddress: string) => {
    const wallet = await fetchSigner({
      chainId: fromChainId,
    })

    const erc20 = new Contract(
      fromToken.addresses[fromChainId],
      erc20ABI,
      wallet
    )
    const allowance = await erc20.allowance(address, routerAddress)
    return allowance
  }

  /*
  Function: handleNewFromToken
  - Handles all the changes that occur when selecting a new "from token", such as generating lists of potential chains/tokens
   to bridge to and handling if the current "to chain/token" are incompatible.
  */
  const handleNewFromToken = useCallback(
    (
      token: Token,
      positedToChain: number | undefined,
      positedToSymbol: string | undefined,
      fromChainId: number
    ) => {
      let newToChain =
        positedToChain && positedToChain !== fromChainId
          ? Number(positedToChain)
          : DEFAULT_TO_CHAIN
      let bridgeableChains = BRIDGE_CHAINS_BY_TYPE[
        String(token.swapableType)
      ].filter((chainId) => Number(chainId) !== fromChainId)
      const swapExceptionsArr: number[] =
        token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]
      if (swapExceptionsArr?.length > 0) {
        bridgeableChains = swapExceptionsArr.map((chainId) => String(chainId))
      }

      if (!bridgeableChains.includes(String(newToChain))) {
        newToChain =
          Number(bridgeableChains[0]) === fromChainId
            ? Number(bridgeableChains[1])
            : Number(bridgeableChains[0])
      }
      const positedToToken = positedToSymbol
        ? tokenSymbolToToken(newToChain, positedToSymbol)
        : tokenSymbolToToken(newToChain, token.symbol)

      let bridgeableTokens: Token[] = sortByVisibilityRank(
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[newToChain][String(token.swapableType)]
      )

      if (swapExceptionsArr?.length > 0) {
        bridgeableTokens = bridgeableTokens.filter(
          (toToken) => toToken.symbol === token.symbol
        )
      }
      let bridgeableToken: Token = positedToToken
      if (!bridgeableTokens.includes(positedToToken)) {
        bridgeableToken = bridgeableTokens[0]
      }

      return {
        bridgeableToken,
        newToChain,
        bridgeableTokens,
        bridgeableChains,
      }
    },
    [fromToken, fromChainId, toToken, toChainId]
  )

  /*
  useEffect triggers: currentAddress, isDisconnected, popup
  - will dismiss toast asking user to connect wallet once wallet has been connected
  */
  useEffect(() => {
    if (currentAddress) {
      toast.dismiss(popup)
    }
  }, [currentAddress, isDisconnected, popup])

  /*
  Function: handleChainChange
  - Produces and alert if chain not connected (upgrade to toaster)
  - Handles flipping to and from chains if flag is set to true
  - Handles altering the chain state for origin or destination depending on the type specified.
  */
  const handleChainChange = useCallback(
    async (chainId: number, flip: boolean, type: 'from' | 'to') => {
      if (currentAddress === undefined || isDisconnected) {
        popup = toast.error('Please connect your wallet', {
          id: 'bridge-connect-wallet',
          duration: 20000,
        })
        return popup
      }

      if (flip || type === 'from') {
        const positedToChain = flip ? fromChainId : undefined
        const desiredChainId = flip ? Number(toChainId) : Number(chainId)

        const res = await switchNetwork({ chainId: desiredChainId })
          .then((res) => {
            if (fromInput.string !== '') {
              setIsQuoteLoading(true)
            }
            return res
          })
          .catch((error) => {
            return error && undefined
          })

        if (res === undefined) {
          console.log("can't switch chain, chainId: ", chainId)
          return
        }

        const bridgeableFromTokens: Token[] = sortByVisibilityRank(
          BRIDGE_SWAPABLE_TOKENS_BY_TYPE[desiredChainId][
            String(fromToken.swapableType)
          ]
        )
        let tempFromToken: Token = fromToken

        if (bridgeableFromTokens?.length > 0) {
          tempFromToken = getMostCommonSwapableType(desiredChainId)
        }
        const {
          bridgeableToken,
          newToChain,
          bridgeableTokens,
          bridgeableChains,
        } = handleNewFromToken(
          tempFromToken,
          positedToChain,
          toToken.symbol,
          desiredChainId
        )
        resetTokenPermutation(
          tempFromToken,
          newToChain,
          bridgeableToken,
          bridgeableChains,
          bridgeableTokens,
          tempFromToken.symbol,
          bridgeableToken.symbol
        )
        sortByTokenBalance(
          BRIDGABLE_TOKENS[desiredChainId],
          desiredChainId,
          address
        ).then((tokens) => {
          setFromTokens(tokens)
        })
        return
      } else if (type === 'to') {
        const {
          bridgeableToken: toBridgeableToken,
          newToChain: toNewToChain,
          bridgeableTokens: toBridgeableTokens,
          bridgeableChains: toBridgeableChains,
        } = handleNewFromToken(fromToken, chainId, toToken.symbol, fromChainId)
        resetTokenPermutation(
          fromToken,
          toNewToChain,
          toBridgeableToken,
          toBridgeableChains,
          toBridgeableTokens,
          fromToken.symbol,
          toBridgeableToken.symbol
        )
        if (fromInput.string !== '') {
          setIsQuoteLoading(true)
        }
        return
      }
    },
    [
      currentAddress,
      isDisconnected,
      fromToken,
      fromChainId,
      toToken,
      toChainId,
      isQuoteLoading,
      handleNewFromToken,
      switchNetwork,
    ]
  )

  /*
    Function:handleTokenChange
  - Handles when the user selects a new token from either the origin or destination
  */
  const handleTokenChange = (token: Token, type: 'from' | 'to') => {
    switch (type) {
      case 'from':
        const {
          bridgeableToken,
          newToChain,
          bridgeableTokens,
          bridgeableChains,
        } = handleNewFromToken(token, toChainId, toToken.symbol, fromChainId)
        resetTokenPermutation(
          token,
          newToChain,
          bridgeableToken,
          bridgeableChains,
          bridgeableTokens,
          token.symbol,
          bridgeableToken.symbol
        )
        if (fromInput.string !== '') {
          setIsQuoteLoading(true)
        }
        return
      case 'to':
        setToToken(token)
        if (fromInput.string !== '') {
          setIsQuoteLoading(true)
        }
        updateUrlParams({
          outputChain: toChainId,
          inputCurrency: fromToken.symbol,
          outputCurrency: token.symbol,
        })
        return
    }
  }

  /*
   Function: getQuote
  - Gets quote data from the Synapse SDK (from the imported provider)
  - Calculates slippage by subtracting fee from input amount (checks to ensure proper num of decimals are in use - ask someone about stable swaps if you want to learn more)
  */
  const getQuote = async () => {
    try {
      if (bridgeQuote === EMPTY_BRIDGE_QUOTE) {
        setIsQuoteLoading(true)
      }
      const { feeAmount, routerAddress, maxAmountOut, originQuery, destQuery } =
        await synapseSDK.bridgeQuote(
          fromChainId,
          toChainId,
          fromToken.addresses[fromChainId],
          toToken.addresses[toChainId],
          fromInput.bigNum
        )

      if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
        setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO)
        setIsQuoteLoading(false)
        return
      }
      const toValueBigNum = maxAmountOut ?? Zero
      const adjustedFeeAmount = feeAmount.lt(fromInput.bigNum)
        ? feeAmount
        : feeAmount.div(
            BigNumber.from(10).pow(18 - toToken.decimals[toChainId])
          )

      const allowance =
        fromToken.addresses[fromChainId] === AddressZero ||
        address === undefined
          ? Zero
          : await getCurrentTokenAllowance(routerAddress)
      setBridgeQuote({
        outputAmount: toValueBigNum,
        outputAmountString: commify(
          formatBNToString(toValueBigNum, toToken.decimals[toChainId], 8)
        ),
        routerAddress,
        allowance,
        exchangeRate: calculateExchangeRate(
          fromInput.bigNum.sub(adjustedFeeAmount),
          fromToken.decimals[fromChainId],
          toValueBigNum,
          toToken.decimals[toChainId]
        ),
        feeAmount,
        delta: maxAmountOut,
        quotes: {
          originQuery,
          destQuery,
        },
      })
      setIsQuoteLoading(false)
      return
    } catch (error) {
      console.log(error)
      setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO)
      setIsQuoteLoading(false)
      return
    }
  }

  /*
  Function: executeBridge
  - Gets raw unsigned tx data from sdk and then execute it with ethers.
  - Only executes if token has already been approved.
   */
  const executeBridge = async () => {
    try {
      const wallet = await fetchSigner({
        chainId: fromChainId,
      })
      var newAddress =
        destinationAddress && isAddress(destinationAddress)
          ? destinationAddress
          : address
      const data = await synapseSDK.bridge(
        newAddress,
        fromChainId,
        toChainId,
        fromToken.addresses[fromChainId as keyof Token['addresses']],
        fromInput.bigNum,
        bridgeQuote.quotes.originQuery,
        bridgeQuote.quotes.destQuery
      )
      const payload =
        fromToken.addresses[fromChainId as keyof Token['addresses']] ===
          AddressZero ||
        fromToken.addresses[fromChainId as keyof Token['addresses']] === ''
          ? { data: data.data, to: data.to, value: fromInput.bigNum }
          : data
      const tx = await wallet.sendTransaction(payload)
      try {
        await tx.wait()
        console.log(`Transaction mined successfully: ${tx.hash}`)
        setBridgeTxHash(tx.hash)
        return tx
      } catch (error) {
        console.log(`Transaction failed with error: ${error}`)
      }
    } catch (error) {
      console.log('Error executing bridge', error)
      return
    }
  }

  return (
    <LandingPageWrapper>
      <main
        data-test-id="bridge-page"
        className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none"
      >
        <div className="items-center px-4 py-8 mx-auto mt-4 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <div>
            <Grid
              cols={{ xs: 1 }}
              gap={6}
              className="justify-center px-2 py-16 sm:px-6 md:px-8"
            >
              <Popup chainId={fromChainId} />
              <div className="flex justify-center">
                <div className="pb-3 place-self-center">
                  <BridgeCard
                    error={error}
                    address={address}
                    bridgeQuote={bridgeQuote}
                    fromInput={fromInput}
                    fromToken={fromToken}
                    fromTokens={fromTokens}
                    fromChainId={fromChainId}
                    toToken={toToken}
                    toChainId={toChainId}
                    toOptions={toOptions}
                    isQuoteLoading={isQuoteLoading}
                    setIsQuoteLoading={setIsQuoteLoading}
                    destinationAddress={destinationAddress}
                    handleChainChange={handleChainChange}
                    handleTokenChange={handleTokenChange}
                    onChangeFromAmount={onChangeFromAmount}
                    setDestinationAddress={setDestinationAddress}
                    executeBridge={executeBridge}
                    resetRates={resetRates}
                    setTime={setTime}
                    bridgeTxnHash={bridgeTxHash}
                  />
                  <ActionCardFooter link={HOW_TO_BRIDGE_URL} />
                </div>
              </div>
              <div>
                <BridgeWatcher
                  fromChainId={fromChainId}
                  toChainId={toChainId}
                  address={address}
                  destinationAddress={destinationAddress}
                  bridgeTxHash={bridgeTxHash}
                />
              </div>
            </Grid>
          </div>
        </div>
      </main>
    </LandingPageWrapper>
  )
}

export default BridgePage
