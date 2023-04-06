import Grid from '@tw/Grid'
import { LandingPageWrapper } from '@components/layouts/LandingPageWrapper'
import { useRouter } from 'next/router'
import { useNetwork } from 'wagmi'
import { useEffect, useState } from 'react'
import { AddressZero, Zero } from '@ethersproject/constants'
import { BigNumber } from '@ethersproject/bignumber'
import { ActionCardFooter } from '@components/ActionCardFooter'
import { fetchSigner, getNetwork, switchNetwork } from '@wagmi/core'
import { sortByTokenBalance, sortByVisibilityRank } from '@utils/sortTokens'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import {
  BRIDGABLE_TOKENS,
  BRIDGE_CHAINS_BY_TYPE,
  BRIDGE_SWAPABLE_TOKENS_BY_TYPE,
  tokenSymbolToToken,
} from '@constants/tokens'
import { formatBNToString } from '@utils/bignumber/format'
import { commify } from '@ethersproject/units'
import { erc20ABI } from 'wagmi'
import { Contract } from 'ethers'

import { BridgeQuote } from '@/utils/types'
import { Token } from '@/utils/types'
import { BRIDGE_PATH, HOW_TO_BRIDGE_URL } from '@/constants/urls'
import { stringToBigNum } from '@/utils/stringToBigNum'
import BridgeCard from './BridgeCard'
import { useSynapseContext } from '@/utils/SynapseProvider'
import {
  DEFAULT_FROM_CHAIN,
  DEFAULT_FROM_TOKEN,
  DEFAULT_TO_CHAIN,
  DEFAULT_TO_TOKEN,
  EMPTY_BRIDGE_QUOTE,
  EMPTY_BRIDGE_QUOTE_ZERO,
  QUOTE_POLLING_INTERVAL,
} from '@/constants/bridge'
// import BridgeWatcher from './BridgeWatcher'

const BridgePage = ({ address }: { address: `0x${string}` }) => {
  const router = useRouter()
  const SynapseSDK = useSynapseContext()
  const { chain: connectedChain } = useNetwork()
  const [time, setTime] = useState(Date.now())
  const [fromChainId, setFromChainId] = useState(DEFAULT_FROM_CHAIN)
  const [fromToken, setFromToken] = useState(DEFAULT_FROM_TOKEN)
  const [fromTokens, setFromTokens] = useState([])
  const [fromInput, setFromInput] = useState({ string: '', bigNum: Zero })
  const [toChainId, setToChainId] = useState(DEFAULT_TO_CHAIN)
  const [toToken, setToToken] = useState(DEFAULT_TO_TOKEN)
  const [needsApproval, setNeedsApproval] = useState(false)
  const [error, setError] = useState('')
  const [destinationAddress, setDestinationAddress] = useState('')
  const [toOptions, setToOptions] = useState({
    tokens: BRIDGABLE_TOKENS[DEFAULT_TO_CHAIN],
    chains: BRIDGE_CHAINS_BY_TYPE[
      String(DEFAULT_FROM_TOKEN.swapableType)
    ].filter((chain) => Number(chain) !== DEFAULT_FROM_CHAIN),
  })
  const [bridgeQuote, setBridgeQuote] =
    useState<BridgeQuote>(EMPTY_BRIDGE_QUOTE)

  useEffect(() => {
    const { chain: fromChainIdRaw } = getNetwork()
    setFromChainId(fromChainIdRaw ? fromChainIdRaw?.id : DEFAULT_FROM_CHAIN)
    const interval = setInterval(
      () => setTime(Date.now()),
      QUOTE_POLLING_INTERVAL
    )
    return () => {
      clearInterval(interval)
    }
  }, [])
  useEffect(() => {
    if (connectedChain?.id && connectedChain?.id !== fromChainId) {
      setFromChainId(connectedChain?.id)
    }
  }, [connectedChain])
  useEffect(() => {
    if (fromChainId === undefined || address === undefined) {
      return
    }
    sortByTokenBalance(
      BRIDGABLE_TOKENS[fromChainId],
      fromChainId,
      address
    ).then((tokens) => {
      setFromTokens(tokens)
    })
  }, [fromChainId])
  // useEffect(() => {}, [bridgeQuote])
  useEffect(() => {
    console.log('TIME')
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
  }, [fromToken, toToken, fromInput, fromChainId, toChainId, time])

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
    resetRates()
    updateUrlParams({
      outputChain: newToChain,
      inputCurrency: newFromTokenSymbol,
      outputCurrency: newBridgeableTokenSymbol,
    })
  }

  const resetRates = () => {
    setBridgeQuote(EMPTY_BRIDGE_QUOTE)
    setFromInput({ string: '', bigNum: Zero })
  }

  const onChangeFromAmount = (value: string) => {
    if (
      !(
        value.split('.')[1]?.length >
        fromToken[fromChainId as keyof Token['decimals']]
      )
    ) {
      setFromInput({
        string: value,
        bigNum: stringToBigNum(value, fromToken.decimals[fromChainId]) ?? Zero,
      })
    }
  }
  const getMostCommonSwapableType = (chainId: number) => {
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
  }
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

  const handleNewFromToken = (
    token: Token,
    positedToChain: number | undefined,
    positedToSymbol: string | undefined,
    fromChainId: number
  ) => {
    let newToChain = positedToChain ? Number(positedToChain) : DEFAULT_TO_CHAIN
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
  }

  // Handles when chains are flipped or user creates toChainId == fromChainId condition
  const handleChainFlip = async () => {
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      const oldFromChain = fromChainId
      const res = switchNetwork({ chainId: Number(toChainId) })
        .then((res) => {
          return res
        })
        .catch((err) => {
          console.log("can't switch network ser", err)
          return undefined
        })
      if (res === undefined) {
        return
      }
      const bridgeableFromTokens: Token[] = sortByVisibilityRank(
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[fromChainId][
          String(fromToken.swapableType)
        ]
      )
      let tempFromToken: Token = fromToken

      if (bridgeableFromTokens?.length > 0) {
        tempFromToken = getMostCommonSwapableType(fromChainId)
      }
      const {
        bridgeableToken,
        newToChain,
        bridgeableTokens,
        bridgeableChains,
      } = handleNewFromToken(
        tempFromToken,
        oldFromChain,
        toToken.symbol,
        toChainId
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
    }
  }

  // Changes destination change when the user changes the toChainId
  const handleFromChainChange = async (chainId: number) => {
    if (address === undefined) {
      alert('Please connect your wallet')
    } else {
      const res = switchNetwork({ chainId: Number(chainId) })
        .then((res) => {
          if (res === undefined) {
            console.log("can't switch network", chainId)
            return
          }
          return res
        })
        .catch((err) => {
          console.log("can't switch network sir", err)
          return undefined
        })
      if (res === undefined) {
        console.log("can't switch network chainId", chainId)
        return
      }

      const bridgeableFromTokens: Token[] = sortByVisibilityRank(
        BRIDGE_SWAPABLE_TOKENS_BY_TYPE[chainId][String(fromToken.swapableType)]
      )
      let tempFromToken: Token = fromToken

      if (bridgeableFromTokens?.length > 0) {
        tempFromToken = getMostCommonSwapableType(chainId)
      }
      const {
        bridgeableToken,
        newToChain,
        bridgeableTokens,
        bridgeableChains,
      } = handleNewFromToken(tempFromToken, chainId, toToken.symbol, chainId)
      resetTokenPermutation(
        tempFromToken,
        newToChain,
        bridgeableToken,
        bridgeableChains,
        bridgeableTokens,
        tempFromToken.symbol,
        bridgeableToken.symbol
      )
    }
  }

  const handleToChainChange = (chainId: number) => {
    setToChainId(chainId)
    updateUrlParams({
      outputChain: chainId,
      inputCurrency: fromToken.symbol,
      outputCurrency: toToken.symbol,
    })
  }

  const handleTokenChange = (token: Token, type: 'from' | 'to') => {
    if (type === 'from') {
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
    } else {
      setToToken(token)
      updateUrlParams({
        outputChain: toChainId,
        inputCurrency: fromToken.symbol,
        outputCurrency: token.symbol,
      })
    }
  }

  const getQuote = async () => {
    console.log('repoll', fromInput)
    const { feeAmount, routerAddress, maxAmountOut, originQuery, destQuery } =
      await SynapseSDK.bridgeQuote(
        fromChainId,
        toChainId,
        fromToken.addresses[fromChainId].toLowerCase(),
        toToken.addresses[toChainId].toLowerCase(),
        fromInput.bigNum
      )
    if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
      setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO)
      return
    }
    const toValueBigNum = maxAmountOut ?? Zero
    const adjustedFeeAmount = feeAmount.lt(fromInput.bigNum)
      ? feeAmount
      : feeAmount.div(BigNumber.from(10).pow(18 - toToken.decimals[toChainId]))
    const allowance =
      fromToken.addresses[fromChainId] === AddressZero
        ? -1
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
  }

  const getCurrentTokenAllowance = async (routerAddress: string) => {
    // TODO store this erc20 and signer retrieval in a state in a parent component
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
                    destinationAddress={destinationAddress}
                    handleChainFlip={handleChainFlip}
                    handleTokenChange={handleTokenChange}
                    onChangeFromAmount={onChangeFromAmount}
                    onSelectFromChain={handleFromChainChange}
                    onSelectToChain={handleToChainChange}
                    setDestinationAddress={setDestinationAddress}
                  />
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
export default BridgePage
