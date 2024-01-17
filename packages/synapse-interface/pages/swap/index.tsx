import { useAccount, useNetwork } from 'wagmi'
import { useSelector } from 'react-redux'
import { RootState } from '../../store/store'
import toast from 'react-hot-toast'
import { animated } from 'react-spring'
import { useRouter } from 'next/router'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { setIsLoading } from '@/slices/swap/reducer'

import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { subtractSlippage } from '@/utils/slippage'
import { commify } from '@ethersproject/units'
import { formatBigIntToString } from '@/utils/bigint/format'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { useEffect, useRef, useState } from 'react'
import { Token } from '@/utils/types'
import { getWalletClient } from '@wagmi/core'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { AcceptedChainId, CHAINS_ARR, CHAINS_BY_ID } from '@/constants/chains'
import { approveToken } from '@/utils/approveToken'
import { PageHeader } from '@/components/PageHeader'
import Card from '@/components/ui/tailwind/Card'
import { Transition } from '@headlessui/react'
import {
  SECTION_TRANSITION_PROPS,
  TRANSITION_PROPS,
} from '@/styles/transitions'
import ExplorerToastLink from '@/components/ExplorerToastLink'
import { Address, zeroAddress } from 'viem'
import { stringToBigInt } from '@/utils/bigint/format'
import { useAppDispatch } from '@/store/hooks'
import { useFetchPortfolioBalances } from '@/slices/portfolio/hooks'
import { SwapTransactionButton } from '@/components/StateManagedSwap/SwapTransactionButton'
import SwapExchangeRateInfo from '@/components/StateManagedSwap/SwapExchangeRateInfo'
import { useSwapState } from '@/slices/swap/hooks'
import { SwapChainListOverlay } from '@/components/StateManagedSwap/SwapChainListOverlay'
import { SwapFromTokenListOverlay } from '@/components/StateManagedSwap/SwapFromTokenListOverlay'
import { SwapInputContainer } from '@/components/StateManagedSwap/SwapInputContainer'
import { SwapOutputContainer } from '@/components/StateManagedSwap/SwapOutputContainer'
import { setSwapQuote, updateSwapFromValue } from '@/slices/swap/reducer'
import { DEFAULT_FROM_CHAIN, EMPTY_SWAP_QUOTE_ZERO } from '@/constants/swap'
import { SwapToTokenListOverlay } from '@/components/StateManagedSwap/SwapToTokenListOverlay'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import useSyncQueryParamsWithSwapState from '@/utils/hooks/useSyncQueryParamsWithSwapState'

const StateManagedSwap = () => {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const { synapseSDK } = useSynapseContext()
  const swapDisplayRef = useRef(null)
  const currentSDKRequestID = useRef(0)
  const router = useRouter()
  const { query, pathname } = router

  useSyncQueryParamsWithSwapState()

  const { balances: portfolioBalances, status: portfolioStatus } =
    useFetchPortfolioBalances()

  const { swapChainId, swapFromToken, swapToToken, swapFromValue, swapQuote } =
    useSwapState()

  const {
    showSwapFromTokenListOverlay,
    showSwapChainListOverlay,
    showSwapToTokenListOverlay,
  } = useSelector((state: RootState) => state.swapDisplay)

  let pendingPopup
  let successPopup

  const [isApproved, setIsApproved] = useState(false)

  const dispatch = useAppDispatch()

  useEffect(() => {
    segmentAnalyticsEvent(`[Swap page] arrives`, {
      swapChainId,
      query,
      pathname,
    })
  }, [query])

  useEffect(() => {
    if (
      swapFromToken &&
      swapToToken &&
      swapFromToken?.decimals[swapChainId] &&
      stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId]) > 0n
    ) {
      console.log('trying to set swap quote')
      getAndSetSwapQuote()
    } else {
      dispatch(setSwapQuote(EMPTY_SWAP_QUOTE_ZERO))
      dispatch(setIsLoading(false))
    }
  }, [
    swapChainId,
    swapFromToken,
    swapToToken,
    swapFromValue,
    address,
    portfolioBalances,
  ])

  useEffect(() => {
    if (
      swapFromToken &&
      swapFromToken?.addresses[swapChainId] === zeroAddress
    ) {
      setIsApproved(true)
    } else {
      if (
        swapFromToken &&
        swapQuote?.allowance &&
        stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId]) <=
          swapQuote.allowance
      ) {
        setIsApproved(true)
      } else {
        setIsApproved(false)
      }
    }
  }, [swapQuote, swapFromToken, swapFromValue, swapChainId])

  let quoteToast

  const getAndSetSwapQuote = async () => {
    currentSDKRequestID.current += 1
    const thisRequestId = currentSDKRequestID.current
    try {
      dispatch(setIsLoading(true))

      const { routerAddress, maxAmountOut, query } = await synapseSDK.swapQuote(
        swapChainId,
        swapFromToken.addresses[swapChainId],
        swapToToken.addresses[swapChainId],
        stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId])
      )

      if (!(query && maxAmountOut)) {
        dispatch(setSwapQuote(EMPTY_SWAP_QUOTE_ZERO))
        dispatch(setIsLoading(true))
        return
      }

      const toValueBigInt = BigInt(maxAmountOut.toString()) ?? 0n

      const allowance =
        swapFromToken.addresses[swapChainId] === zeroAddress ||
        address === undefined
          ? 0n
          : await getErc20TokenAllowance({
              address,
              chainId: swapChainId,
              tokenAddress: swapFromToken.addresses[swapChainId] as Address,
              spender: routerAddress,
            })

      const originQueryWithSlippage = synapseSDK.applySwapSlippage(query)

      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(
          setSwapQuote({
            outputAmount: toValueBigInt,
            outputAmountString: commify(
              formatBigIntToString(
                toValueBigInt,
                swapToToken.decimals[swapChainId],
                8
              )
            ),
            routerAddress,
            allowance: BigInt(allowance.toString()),
            exchangeRate: calculateExchangeRate(
              stringToBigInt(
                swapFromValue,
                swapFromToken.decimals[swapChainId]
              ),
              swapFromToken.decimals[swapChainId],
              toValueBigInt,
              swapToToken.decimals[swapChainId]
            ),
            delta: toValueBigInt,
            quote: originQueryWithSlippage,
          })
        )

        dispatch(setIsLoading(false))
        if (quoteToast) {
          toast.dismiss(quoteToast)
        }
        const message = `Route found for swapping ${swapFromValue} ${swapFromToken.symbol} on ${CHAINS_BY_ID[swapChainId]?.name} to ${swapToToken.symbol}`
        console.log(message)
        quoteToast = toast(message, { duration: 3000 })
      }
    } catch (err) {
      console.log(err)
      if (thisRequestId === currentSDKRequestID.current) {
        toast.dismiss(quoteToast)
        let message
        if (!swapChainId) {
          message = 'Please select an origin chain'
        } else if (!swapFromToken) {
          message = 'Please select an origin token'
        } else if (!swapToToken) {
          message = 'Please select a destination token'
        } else {
          message = `No route found for swapping ${swapFromValue} ${swapFromToken.symbol} on ${CHAINS_BY_ID[swapChainId]?.name} to ${swapToToken.symbol}`
        }
        console.log(message)
        quoteToast = toast(message, { duration: 3000 })

        dispatch(setSwapQuote(EMPTY_SWAP_QUOTE_ZERO))
        return
      }
    } finally {
      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(setIsLoading(false))
      }
    }
  }

  const approveTxn = async () => {
    try {
      const tx = approveToken(
        swapQuote?.routerAddress,
        swapChainId,
        swapFromToken?.addresses[swapChainId]
      )
      await tx
      /** Re-fetch swap quote to re-check approval state */
      getAndSetSwapQuote()
    } catch (error) {
      return txErrorHandler(error)
    }
  }

  const executeSwap = async () => {
    const currentChainName = CHAINS_BY_ID[swapChainId]?.name

    pendingPopup = toast(
      `Initiating swap from ${swapFromToken.symbol} to ${swapToToken.symbol} on ${currentChainName}`,
      { id: 'swap-in-progress-popup', duration: Infinity }
    )
    segmentAnalyticsEvent(
      `[Swap] initiates swap`,
      {
        address,
        chainId: swapChainId,
        swapFromToken: swapFromToken.symbol,
        swapToToken: swapToToken.symbol,
        inputAmount: swapFromValue,
        expectedReceivedAmount: swapQuote.outputAmountString,
        exchangeRate: swapQuote.exchangeRate,
      },
      true
    )
    try {
      const wallet = await getWalletClient({
        chainId: swapChainId,
      })

      const data = await synapseSDK.swap(
        swapChainId,
        address,
        swapFromToken.addresses[swapChainId],
        stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId]),
        swapQuote.quote
      )

      const payload =
        swapFromToken.addresses[swapChainId as keyof Token['addresses']] ===
          zeroAddress ||
        swapFromToken.addresses[swapChainId as keyof Token['addresses']] === ''
          ? {
              data: data.data,
              to: data.to,
              value: stringToBigInt(
                swapFromValue,
                swapFromToken.decimals[swapChainId]
              ),
            }
          : data

      const tx = await wallet.sendTransaction(payload)

      const originChainName = CHAINS_BY_ID[swapChainId]?.name
      pendingPopup = toast(
        `Swapping ${swapFromToken.symbol} on ${originChainName} to ${swapToToken.symbol}`,
        { id: 'swap-in-progress-popup', duration: Infinity }
      )

      try {
        const successTx = await tx

        segmentAnalyticsEvent(`[Swap] swaps successfully`, {
          address,
          originChainId: swapChainId,
          inputAmount: swapFromValue,
          expectedReceivedAmount: swapQuote.outputAmountString,
          exchangeRate: swapQuote.exchangeRate,
        })

        toast.dismiss(pendingPopup)

        const successToastContent = (
          <div>
            <div>
              Successfully swapped from {swapFromToken.symbol} to{' '}
              {swapToToken.symbol} on {currentChainName}
            </div>
            <ExplorerToastLink
              transactionHash={tx ?? zeroAddress}
              chainId={swapChainId}
            />
          </div>
        )

        successPopup = toast.success(successToastContent, {
          id: 'swap-successful-popup',
          duration: 10000,
        })

        dispatch(setSwapQuote(EMPTY_SWAP_QUOTE_ZERO))
        dispatch(updateSwapFromValue())
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

  const springClass =
    '-mt-4 fixed z-50 w-full h-full bg-opacity-50 bg-[#343036]'

  return (
    <LandingPageWrapper>
      <div className="flex justify-center px-4 py-16 mx-auto lg:mx-0">
        <div className="flex flex-col">
          <div className="flex items-center justify-between">
            <PageHeader title="Swap" subtitle="Exchange assets on chain." />
          </div>
          <Card
            divider={false}
            className={`
              pb-3 mt-5 overflow-hidden
              transition-all duration-100 transform rounded-md
              bg-bgBase
            `}
          >
            <div ref={swapDisplayRef}>
              <Transition show={showSwapChainListOverlay} {...TRANSITION_PROPS}>
                <animated.div className={springClass}>
                  <SwapChainListOverlay />
                </animated.div>
              </Transition>
              <Transition
                show={showSwapFromTokenListOverlay}
                {...TRANSITION_PROPS}
              >
                <animated.div className={springClass}>
                  <SwapFromTokenListOverlay />
                </animated.div>
              </Transition>
              <Transition
                show={showSwapToTokenListOverlay}
                {...TRANSITION_PROPS}
              >
                <animated.div className={springClass}>
                  <SwapToTokenListOverlay />
                </animated.div>
              </Transition>
              <SwapInputContainer />
              <SwapOutputContainer />
              <Transition
                appear={true}
                unmount={false}
                show={true}
                {...SECTION_TRANSITION_PROPS}
              >
                <SwapExchangeRateInfo
                  fromAmount={
                    swapFromToken
                      ? stringToBigInt(
                          swapFromValue,
                          swapFromToken.decimals[swapChainId]
                        )
                      : 0n
                  }
                  toToken={swapToToken}
                  exchangeRate={swapQuote.exchangeRate}
                  toChainId={swapChainId}
                />
              </Transition>
              <div className="md:my-3">
                <SwapTransactionButton
                  isApproved={isApproved}
                  approveTxn={approveTxn}
                  executeSwap={executeSwap}
                />
              </div>
            </div>
          </Card>
        </div>
      </div>
    </LandingPageWrapper>
  )
}

export default StateManagedSwap
