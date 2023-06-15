import Grid from '@tw/Grid'
import { useEffect, useState, useMemo, useCallback } from 'react'
import { useRouter } from 'next/router'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import { fetchSigner, switchNetwork } from '@wagmi/core'
import { useWatchPendingTransactions } from 'wagmi'
import { sortByTokenBalance, sortByVisibilityRank } from '@utils/sortTokens'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import ExchangeRateInfo from '@components/ExchangeRateInfo'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import BridgeInputContainer from '../../components/input/TokenAmountInput/index'
import { approveToken } from '@/utils/approveToken'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { formatBNToString } from '@utils/bignumber/format'
import { commify } from '@ethersproject/units'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'
import { subtractSlippage } from '@utils/slippage'
import { ChainSlideOver } from '@/components/misc/ChainSlideOver'
import { TokenSlideOver } from '@/components/misc/TokenSlideOver'
import { Token } from '@/utils/types'
import { SWAP_PATH } from '@/constants/urls'
import { stringToBigNum } from '@/utils/stringToBigNum'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { checkStringIfOnlyZeroes } from '@/utils/regex'
import { timeout } from '@/utils/timeout'
import { Transition } from '@headlessui/react'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import Card from '@tw/Card'
import { SwapQuote, Query } from '@types'
import { IMPAIRED_CHAINS } from '@/constants/impairedChains'
import { CHAINS_BY_ID } from '@constants/chains'
import { toast } from 'react-hot-toast'
import { txErrorHandler } from '@/utils/txErrorHandler'
import ExplorerToastLink from '@/components/ExplorerToastLink'

import {
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_TOKEN,
  EMPTY_SWAP_QUOTE,
  QUOTE_POLLING_INTERVAL,
  EMPTY_SWAP_QUOTE_ZERO,
} from '@/constants/swap'
import {
  SWAPABLE_TOKENS,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  SWAPABLE_TOKENS_BY_TYPE,
  tokenSymbolToToken,
} from '@constants/tokens'

const SwapCard = ({
  address,
  connectedChainId,
}: {
  address: `0x${string}` | undefined
  connectedChainId: number
}) => {
  const router = useRouter()
  const { synapseSDK } = useSynapseContext()
  const [time, setTime] = useState(Date.now())
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [fromTokens, setFromTokens] = useState([])
  const [fromInput, setFromInput] = useState({ string: '', bigNum: Zero })
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)
  const [toTokens, setToTokens] = useState<Token[]>([]) //add default
  const [isQuoteLoading, setIsQuoteLoading] = useState<boolean>(false)
  const [error, setError] = useState(undefined)
  const [destinationAddress, setDestinationAddress] = useState('')
  const [swapQuote, setSwapQuote] = useState<SwapQuote>(EMPTY_SWAP_QUOTE)
  const [displayType, setDisplayType] = useState(undefined)
  const [fromTokenBalance, setFromTokenBalance] = useState<BigNumber>(Zero)
  const [validChainId, setValidChainId] = useState(true)
  const [swapTxnHash, setSwapTxnHash] = useState<string>('')
  const [approveTx, setApproveTx] = useState<string>(null)

  let pendingPopup: any
  let successPopup: any
  let errorPopup: string

  /*
  useEffect Trigger: onMount
  - Gets current network connected and sets it as the state.
  - Initializes polling (setInterval) func to re-retrieve quotes.
  */
  useEffect(() => {
    const interval = setInterval(
      () => setTime(Date.now()),
      QUOTE_POLLING_INTERVAL
    )
    return () => {
      clearInterval(interval)
    }
  }, [])

  /*
  useEffect Trigger: fromInput
  - Resets approve txn status if user input changes after amount is approved
  */

  useEffect(() => {
    if (approveTx) {
      setApproveTx(null)
    }
  }, [fromInput])

  /*
  useEffect Trigger: fromToken, fromTokens
  - When either the from token or list of from tokens are mutated, the selected token's balance is set in state
  this is for checking max bridge possible as well as for producing the option to select max bridge
  */
  useEffect(() => {
    if (fromTokens && fromToken) {
      setFromTokenBalance(
        fromTokens.filter((token) => token.token === fromToken)[0]?.balance
          ? fromTokens.filter((token) => token.token === fromToken)[0]?.balance
          : Zero
      )
    }
  }, [fromToken, fromTokens])

  useEffect(() => {
    if (!router.isReady || !SWAPABLE_TOKENS[connectedChainId]) {
      return
    }
    const {
      inputCurrency: fromTokenSymbolUrl,
      outputCurrency: toTokenSymbolUrl,
    } = router.query

    let tempFromToken: Token = getMostCommonSwapableType(connectedChainId)

    if (fromTokenSymbolUrl) {
      let token = tokenSymbolToToken(
        connectedChainId,
        String(fromTokenSymbolUrl)
      )
      if (token) {
        tempFromToken = token
      }
    }
    const { swapableToken, swapableTokens } = handleNewFromToken(
      tempFromToken,
      toTokenSymbolUrl ? String(toTokenSymbolUrl) : undefined,
      connectedChainId
    )
    resetTokenPermutation(
      tempFromToken,
      swapableToken,
      swapableTokens,
      tempFromToken.symbol,
      swapableToken.symbol
    )
    updateUrlParams({
      inputCurrency: fromToken.symbol,
      outputCurrency: swapableToken.symbol,
    })
  }, [router.isReady])

  /*
  useEffect Trigger: connectedChain
  - when the connected chain changes (wagmi hook), update the state
  */
  useEffect(() => {
    if (address === undefined) {
      return
    }
    handleChainChange(connectedChainId, undefined, undefined)

    sortByTokenBalance(
      SWAPABLE_TOKENS[connectedChainId],
      connectedChainId,
      address
    ).then((tokens) => {
      setFromTokens(tokens)
    })
    return
  }, [connectedChainId, swapTxnHash, address])

  /*
  useEffect Triggers: toToken, fromInput, toChainId, time
  - Gets a quote when the polling function is executed or any of the bridge attributes are altered.
    - Debounce quote call by calling quote price AFTER user has stopped typing for 1s or 1000ms
  */
  useEffect(() => {
    let isCancelled = false

    const handleChange = async () => {
      // await timeout(1000)
      if (
        connectedChainId &&
        String(fromToken.addresses[connectedChainId]) &&
        fromInput &&
        fromInput.bigNum.gt(Zero)
      ) {
        // TODO this needs to be debounced or throttled somehow to prevent spam and lag in the ui
        getQuote()
      } else {
        setSwapQuote(EMPTY_SWAP_QUOTE)
      }
    }
    handleChange()

    return () => {
      isCancelled = true
    }
  }, [toToken, fromInput, time, connectedChainId])

  /*
  Helper Function: resetTokenPermutation
  - Handles when theres a new from token/chain and all other parts of the bridge arrangement needs to be updated
  - Updates url params.
  */
  const resetTokenPermutation = (
    newFromToken: Token,
    newToToken: Token,
    newSwapableTokens: Token[],
    newFromTokenSymbol: string,
    newSwapableTokenSymbol: string
  ) => {
    setFromToken(newFromToken)
    setToToken(newToToken)
    setToTokens(newSwapableTokens)
    // resetRates()
    updateUrlParams({
      inputCurrency: newFromTokenSymbol,
      outputCurrency: newSwapableTokenSymbol,
    })
  }

  /*
  Helper Function: resetRates
  - Called when switching from chain/token so that the from input isn't populated with stale data.
  */
  const resetRates = () => {
    setSwapQuote(EMPTY_SWAP_QUOTE)
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
        fromToken[connectedChainId as keyof Token['decimals']]
      )
    ) {
      let bigNum =
        stringToBigNum(value, fromToken.decimals[connectedChainId]) ?? Zero
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
  const getMostCommonSwapableType = (chainId: number) => {
    const fromChainTokensByType = Object.values(
      SWAPABLE_TOKENS_BY_TYPE[chainId]
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
  }

  /*
  Helper Function: updateUrlParams
  - Pushes chain and token changes to url
  NOTE: did not alter any variable names in case previous users have saved links of different bridging permutations.
  */
  const updateUrlParams = ({
    inputCurrency,
    outputCurrency,
  }: {
    inputCurrency: string
    outputCurrency: string
  }) => {
    router.replace(
      {
        pathname: SWAP_PATH,
        query: {
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
      chainId: connectedChainId,
    })

    const erc20 = new Contract(
      fromToken.addresses[connectedChainId],
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
  const handleNewFromToken = (
    token: Token,
    positedToSymbol: string | undefined,
    fromChainId: number
  ) => {
    const swapExceptionsArr: number[] =
      token?.swapExceptions?.[fromChainId as keyof Token['swapExceptions']]

    const positedToToken = positedToSymbol
      ? tokenSymbolToToken(fromChainId, positedToSymbol)
      : tokenSymbolToToken(fromChainId, token.symbol)

    let swapableTokens: Token[] = sortByVisibilityRank(
      BRIDGE_SWAPABLE_TOKENS_BY_TYPE[fromChainId][String(token.swapableType)]
    ).filter((toToken) => toToken !== token)
    if (swapExceptionsArr?.length > 0) {
      swapableTokens = swapableTokens.filter(
        (toToken) => toToken.symbol === token.symbol
      )
    }
    let swapableToken: Token = positedToToken
    if (!swapableTokens.includes(positedToToken)) {
      swapableToken = swapableTokens[0]
    }

    return {
      swapableToken,
      swapableTokens,
    }
  }

  /*
  useEffect triggers: address, popup
  - will dismiss toast asking user to connect wallet once wallet has been connected
  */
  useEffect(() => {
    if (address && errorPopup) {
      toast.dismiss(errorPopup)
    }
  }, [address, errorPopup])

  /*
  Function: handleChainChange
  - Produces and alert if chain not connected (upgrade to toaster)
  - Handles flipping to and from chains if flag is set to true
  - Handles altering the chain state for origin or destination depending on the type specified.
  */
  const handleChainChange = useCallback(
    async (chainId: number, flip: boolean, type: 'from' | 'to') => {
      if (address === undefined) {
        errorPopup = toast.error('Please connect your wallet', {
          id: 'bridge-connect-wallet',
          duration: 20000,
        })
        return errorPopup
      }
      const desiredChainId = Number(chainId)

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
        console.log("can't switch network, chainId: ", chainId)
        return
      }
      if (!SWAPABLE_TOKENS[desiredChainId]) {
        return
      }
      setValidChainId(true)

      const swapableFromTokens: Token[] = sortByVisibilityRank(
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId][String(fromToken.swapableType)]
      )
      let tempFromToken: Token = fromToken

      if (swapableFromTokens?.length > 0) {
        tempFromToken = getMostCommonSwapableType(chainId)
      }
      const { swapableToken, swapableTokens } = handleNewFromToken(
        tempFromToken,
        toToken.symbol,
        desiredChainId
      )
      resetTokenPermutation(
        tempFromToken,
        swapableToken,
        swapableTokens,
        tempFromToken.symbol,
        swapableToken?.symbol
      )
      return
    },
    [
      fromToken,
      toToken,
      connectedChainId,
      address,
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
        const { swapableToken, swapableTokens } = handleNewFromToken(
          token,
          toToken.symbol,
          connectedChainId
        )
        resetTokenPermutation(
          token,
          swapableToken,
          swapableTokens,
          token.symbol,
          swapableToken.symbol
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
      if (swapQuote === EMPTY_SWAP_QUOTE) {
        setIsQuoteLoading(true)
      }
      const { routerAddress, maxAmountOut, query } = await synapseSDK.swapQuote(
        connectedChainId,
        fromToken.addresses[connectedChainId],
        toToken.addresses[connectedChainId],
        fromInput.bigNum
      )
      if (!(query && maxAmountOut)) {
        setSwapQuote(EMPTY_SWAP_QUOTE_ZERO)
        setIsQuoteLoading(false)
        return
      }
      const toValueBigNum = maxAmountOut ?? Zero

      const allowance =
        fromToken.addresses[connectedChainId] === AddressZero ||
        address === undefined
          ? Zero
          : await getCurrentTokenAllowance(routerAddress)

      const minWithSlippage = subtractSlippage(
        query?.minAmountOut ?? Zero,
        'ONE_TENTH',
        null
      )
      // TODO 1) make dynamic 2) clean up
      let newOriginQuery = { ...query }
      newOriginQuery.minAmountOut = minWithSlippage

      setSwapQuote({
        outputAmount: toValueBigNum,
        outputAmountString: commify(
          formatBNToString(toValueBigNum, toToken.decimals[connectedChainId], 8)
        ),
        routerAddress,
        allowance,
        exchangeRate: calculateExchangeRate(
          fromInput.bigNum.sub(Zero), // this needs to be changed once we can get fee data from router.
          fromToken.decimals[connectedChainId],
          toValueBigNum,
          toToken.decimals[connectedChainId]
        ),
        delta: maxAmountOut,
        quote: newOriginQuery,
      })
      setIsQuoteLoading(false)
      return
    } catch (error) {
      setIsQuoteLoading(false)
      console.log(`Quote failed with error: ${error}`)
      return
    }
  }

  /*
  Function: approveToken
  - Gets raw unsigned tx data from sdk and then execute it with ethers.
  - Only executes if token has already been approved.
   */
  const executeSwap = async () => {
    const currentChainName = CHAINS_BY_ID[connectedChainId]?.name
    pendingPopup = toast(
      `Initiating swap from ${fromToken.symbol} to ${toToken.symbol} on ${currentChainName}`,
      { id: 'swap-in-progress-popup', duration: Infinity }
    )

    try {
      const wallet = await fetchSigner({
        chainId: connectedChainId,
      })

      const data = await synapseSDK.swap(
        connectedChainId,
        address,
        fromToken.addresses[connectedChainId as keyof Token['addresses']],
        fromInput.bigNum,
        swapQuote.quote
      )
      const payload =
        fromToken.addresses[connectedChainId as keyof Token['addresses']] ===
          AddressZero ||
        fromToken.addresses[connectedChainId as keyof Token['addresses']] === ''
          ? { data: data.data, to: data.to, value: fromInput.bigNum }
          : data
      const tx = await wallet.sendTransaction(payload)

      try {
        const successTx = await tx.wait()

        setSwapTxnHash(successTx?.transactionHash)

        toast.dismiss(pendingPopup)

        console.log(`Transaction mined successfully: ${tx.hash}`)

        const successToastContent = (
          <div>
            <div>
              Successfully swapped from {fromToken.symbol} to {toToken.symbol}{' '}
              on {currentChainName}
            </div>
            <ExplorerToastLink
              transactionHash={tx?.hash ?? AddressZero}
              chainId={connectedChainId}
            />
          </div>
        )

        successPopup = toast.success(successToastContent, {
          id: 'swap-successful-popup',
          duration: 10000,
        })

        resetRates()
        return tx
      } catch (error) {
        toast.dismiss(pendingPopup)
        console.log(`Transaction failed with error: ${error}`)
      }
    } catch (error) {
      console.log(`Swap Execution failed with error: ${error}`)
      toast.dismiss(pendingPopup)
      txErrorHandler(error)
    }
  }

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

  const isFromBalanceEnough = fromTokenBalance?.gte(fromInput.bigNum)
  let destAddrNotValid: boolean

  const getButtonProperties = () => {
    let properties = {
      label: `Enter amount to swap`,
      pendingLabel: 'Swapping funds...',
      className: '',
      disabled: true,
      buttonAction: () => executeSwap(),
      postButtonAction: () => resetRates(),
    }

    const isInputZero = checkStringIfOnlyZeroes(fromInput?.string)

    if (error) {
      properties.label = error
      properties.disabled = true
      return properties
    }

    if (isInputZero || fromInput?.bigNum?.eq(0)) {
      properties.label = `Enter amount to swap`
      properties.disabled = true
      return properties
    }

    if (!isFromBalanceEnough) {
      properties.label = `Insufficient ${fromToken.symbol} Balance`
      properties.disabled = true
      return properties
    }

    if (IMPAIRED_CHAINS[connectedChainId]?.disabled) {
      properties.label = `${CHAINS_BY_ID[connectedChainId]?.name} is currently paused`
      properties.disabled = true
      return properties
    }

    if (fromInput.bigNum.eq(0)) {
      properties.label = `Amount must be greater than fee`
      properties.disabled = true
      return properties
    }

    if (
      !fromInput?.bigNum?.eq(0) &&
      fromToken?.addresses[connectedChainId] !== '' &&
      fromToken?.addresses[connectedChainId] !== AddressZero &&
      swapQuote?.allowance &&
      swapQuote?.allowance?.lt(fromInput.bigNum) &&
      !approveTx
    ) {
      properties.buttonAction = () =>
        approveToken(
          swapQuote.routerAddress,
          connectedChainId,
          fromToken.addresses[connectedChainId]
        )
      properties.label = `Approve ${fromToken.symbol}`
      properties.pendingLabel = `Approving ${fromToken.symbol}`
      properties.className = 'from-[#feba06] to-[#FEC737]'
      properties.disabled = false
      properties.postButtonAction = () => {
        setApproveTx('approved')
      }
      return properties
    }

    if (destinationAddress && !validateAndParseAddress(destinationAddress)) {
      destAddrNotValid = true
      properties.label = 'Invalid Destination Address'
      properties.disabled = true
      return properties
    }

    // default case
    properties.label = 'Swap your funds'
    properties.disabled = false

    const numExchangeRate = swapQuote?.exchangeRate
      ? Number(formatBNToString(swapQuote.exchangeRate, 18, 4))
      : 0

    if (
      !fromInput.bigNum.eq(0) &&
      numExchangeRate !== 0 &&
      (numExchangeRate < 0.95 || numExchangeRate > 1.05)
    ) {
      properties.className = 'from-[#fe064a] to-[#fe5281]'
      properties.label = 'Slippage High - Swap Anyway?'
    }

    return properties
  }

  const {
    label: btnLabel,
    pendingLabel,
    className: btnClassName,
    buttonAction,
    postButtonAction,
    disabled,
  } = useMemo(getButtonProperties, [
    isFromBalanceEnough,
    address,
    fromInput,
    fromToken,
    swapQuote,
    isQuoteLoading,
    destinationAddress,
    error,
    approveTx,
  ])

  const ActionButton = useMemo(() => {
    return (
      <TransactionButton
        onClick={() => buttonAction()}
        disabled={disabled || destAddrNotValid}
        className={btnClassName}
        label={btnLabel}
        pendingLabel={pendingLabel}
        chainId={connectedChainId}
        onSuccess={() => {
          postButtonAction()
        }}
      />
    )
  }, [fromInput, time, swapQuote, error, approveTx])

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
  }, [fromInput])

  return (
    <Card
      divider={false}
      className="max-w-lg px-1 pb-0 -mb-3 transition-all duration-100 transform rounded-xl bg-bgBase md:px-6 lg:px-6"
    >
      <div className="mb-8">
        <Transition show={displayType === 'from'} {...transitionProps}>
          <TokenSlideOver
            key="fromBlock"
            isOrigin={true}
            tokens={fromTokens ?? SWAPABLE_TOKENS[connectedChainId] ?? []}
            chainId={connectedChainId}
            selectedToken={fromToken}
            setDisplayType={setDisplayType}
            handleTokenChange={handleTokenChange}
          />
        </Transition>
        <Transition show={displayType === 'to'} {...transitionProps}>
          <TokenSlideOver
            key="toBlock"
            isOrigin={false}
            tokens={toTokens ?? SWAPABLE_TOKENS[connectedChainId] ?? []}
            chainId={connectedChainId}
            selectedToken={toToken}
            setDisplayType={setDisplayType}
            handleTokenChange={handleTokenChange}
          />
        </Transition>
        <Transition show={displayType === 'fromChain'} {...transitionProps}>
          <ChainSlideOver
            key="fromChainBlock"
            isOrigin={true}
            chains={Object.keys(SWAPABLE_TOKENS)}
            chainId={connectedChainId}
            onChangeChain={handleChainChange}
            setDisplayType={setDisplayType}
          />
        </Transition>
        <Grid cols={{ xs: 1 }} gap={4} className="place-content-center">
          <div className="pt-3 "></div>
          <BridgeInputContainer
            address={address}
            isOrigin={true}
            isSwap={true}
            chains={Object.keys(SWAPABLE_TOKENS)}
            chainId={connectedChainId}
            inputString={fromInput.string}
            selectedToken={fromToken}
            connectedChainId={connectedChainId}
            onChangeChain={handleChainChange}
            onChangeAmount={onChangeFromAmount}
            setDisplayType={setDisplayType}
            fromTokenBalance={fromTokenBalance}
          />
          <BridgeInputContainer
            address={address}
            isOrigin={false}
            isSwap={true}
            chains={Object.keys(SWAPABLE_TOKENS)}
            chainId={connectedChainId}
            inputString={swapQuote.outputAmountString}
            selectedToken={toToken}
            connectedChainId={connectedChainId}
            onChangeChain={handleChainChange}
            onChangeAmount={onChangeFromAmount}
            setDisplayType={setDisplayType}
            isQuoteLoading={isQuoteLoading}
          />
        </Grid>

        <ExchangeRateInfo
          fromAmount={fromInput.bigNum}
          toToken={toToken}
          exchangeRate={swapQuote.exchangeRate}
          toChainId={connectedChainId}
          showGasDrop={false}
        />
        <div className="px-2 py-2 md:px-0 md:py-4">{ActionButton}</div>
      </div>
    </Card>
  )
}

export default SwapCard
