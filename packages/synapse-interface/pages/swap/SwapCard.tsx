import Grid from '@tw/Grid'
import { useRouter } from 'next/router'
import { useEffect, useState, useMemo } from 'react'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import { fetchSigner, switchNetwork } from '@wagmi/core'
import { sortByTokenBalance, sortByVisibilityRank } from '@utils/sortTokens'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import ExchangeRateInfo from '@components/ExchangeRateInfo'
import { TransactionButton } from '@components/buttons/SubmitTxButton'
import BridgeInputContainer from '../../components/input/TokenAmountInput/index'
import { approveToken } from '@/utils/approveToken'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { formatBNToString } from '@utils/bignumber/format'
import { commify } from '@ethersproject/units'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'
import { ChainSlideOver } from '@/components/misc/ChainSlideOver'
import { TokenSlideOver } from '@/components/misc/TokenSlideOver'
import { Token } from '@/utils/types'
import { SWAP_PATH } from '@/constants/urls'
import { stringToBigNum } from '@/utils/stringToBigNum'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { Transition } from '@headlessui/react'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import Card from '@tw/Card'
import { SwapQuote } from '@types'
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
  const SynapseSDK = useSynapseContext()
  const [time, setTime] = useState(Date.now())
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [fromTokens, setFromTokens] = useState([])
  const [fromInput, setFromInput] = useState({ string: '', bigNum: Zero })
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)
  const [toTokens, setToTokens] = useState<Token[]>() //add default
  const [error, setError] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')
  const [swapQuote, setSwapQuote] = useState<SwapQuote>(EMPTY_SWAP_QUOTE)
  const [displayType, setDisplayType] = useState(undefined)
  const [fromTokenBalance, setFromTokenBalance] = useState<BigNumber>(Zero)
  const [validChainId, setValidChainId] = useState(true)
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
  }, [connectedChainId])

  /*
  useEffect Triggers: toToken, fromInput, toChainId, time
  - Gets a quote when the polling function is executed or any of the bridge attributes are altered.
  */
  useEffect(() => {
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
  }, [toToken, fromInput, time])

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
    resetRates()
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
  Function: handleChainChange
  - Produces and alert if chain not connected (upgrade to toaster)
  - Handles flipping to and from chains if flag is set to true
  - Handles altering the chain state for origin or destination depending on the type specified.
  */
  const handleChainChange = async (
    chainId: number,
    flip: boolean,
    type: 'from' | 'to'
  ) => {
    if (address === undefined) {
      return alert('Please connect your wallet')
    }
    const desiredChainId = Number(chainId)
    const res = switchNetwork({ chainId: desiredChainId })
      .then((res) => {
        return res
      })
      .catch(() => {
        return undefined
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
      swapableToken.symbol
    )
    return
  }
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
        return
      case 'to':
        resetRates()
        setToToken(token)
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
    const { routerAddress, maxAmountOut, query } = await SynapseSDK.swapQuote(
      connectedChainId,
      fromToken.addresses[connectedChainId],
      toToken.addresses[connectedChainId],
      fromInput.bigNum
    )
    if (!(query && maxAmountOut)) {
      setSwapQuote(EMPTY_SWAP_QUOTE_ZERO)
      return
    }
    const toValueBigNum = maxAmountOut ?? Zero

    const allowance =
      fromToken.addresses[connectedChainId] === AddressZero
        ? Zero
        : await getCurrentTokenAllowance(routerAddress)

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
      quote: query,
    })
    return
  }

  /*
  Function: approveToken
  - Gets raw unsigned tx data from sdk and then execute it with ethers.
  - Only executes if token has already been approved.
   */
  const executeSwap = async () => {
    const wallet = await fetchSigner({
      chainId: connectedChainId,
    })

    const data = await SynapseSDK.swap(
      connectedChainId,
      address,
      fromToken.addresses[connectedChainId as keyof Token['addresses']],
      fromInput.bigNum.mul(1000).div(999), // TODO Get rid of harcoded slippage
      swapQuote.quote
    )
    const tx = await wallet.sendTransaction(data)
    try {
      await tx.wait()
      console.log(`Transaction mined successfully: ${tx.hash}`)
      return tx
    } catch (error) {
      console.log(`Transaction failed with error: ${error}`)
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
  // TODO make this a function
  const ActionButton = useMemo(() => {
    let destAddrNotValid
    let btnLabel
    let btnClassName = ''
    let pendingLabel = 'Swapping funds...'
    let buttonAction = () => executeSwap()
    let postButtonAction = () => resetRates()
    const isFromBalanceEnough = fromTokenBalance?.gte(fromInput.bigNum)

    if (error) {
      btnLabel = error
    } else if (!isFromBalanceEnough) {
      btnLabel = `Insufficient ${fromToken.symbol} Balance`
    } else if (fromInput.bigNum.eq(0)) {
      btnLabel = `Amount must be greater than fee`
    } else if (
      swapQuote?.allowance &&
      swapQuote?.allowance?.lt(fromInput.bigNum)
    ) {
      buttonAction = () =>
        approveToken(
          swapQuote.routerAddress,
          connectedChainId,
          fromToken.addresses[connectedChainId]
        )
      btnLabel = `Approve ${fromToken.symbol}`
      pendingLabel = `Approving ${fromToken.symbol}`
      btnClassName = 'from-[#feba06] to-[#FEC737]'
      postButtonAction = () => setTime(0)
    } else if (
      destinationAddress &&
      !validateAndParseAddress(destinationAddress)
    ) {
      destAddrNotValid = true
      btnLabel = 'Invalid Destination Address'
    } else {
      btnLabel = swapQuote.outputAmount.eq(0)
        ? 'Enter amount to swap'
        : 'Swap your funds'

      const numExchangeRate = Number(
        formatBNToString(swapQuote.exchangeRate, 18, 4)
      )

      if (
        !fromInput.bigNum.eq(0) &&
        (numExchangeRate < 0.95 || numExchangeRate > 1.05)
      ) {
        btnClassName = 'from-[#fe064a] to-[#fe5281]'
        btnLabel = 'Slippage High - Swap Anyway?'
      }
    }

    return (
      <TransactionButton
        className={btnClassName}
        disabled={
          swapQuote.outputAmount.eq(0) ||
          !isFromBalanceEnough ||
          error != null ||
          destAddrNotValid
        }
        onClick={() => buttonAction()}
        onSuccess={() => {
          postButtonAction()
        }}
        label={btnLabel}
        pendingLabel={pendingLabel}
      />
    )

    //   <TransactionButton
    //   onClick={approveToken}
    //   label={`Approve ${displaySymbol(chainId, fromCoin)}`}
    //   pendingLabel={`Approving ${displaySymbol(chainId, fromCoin)}  `}
    // />
  }, [fromInput, time, swapQuote, error])

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
            tokens={fromTokens}
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
            tokens={toTokens}
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
          />
        </Grid>

        <ExchangeRateInfo
          fromAmount={fromInput.bigNum}
          toToken={toToken}
          exchangeRate={swapQuote.exchangeRate}
          toChainId={connectedChainId}
        />
        <div className="px-2 py-2 md:px-0 md:py-4">{ActionButton}</div>
      </div>
    </Card>
  )
}

export default SwapCard
