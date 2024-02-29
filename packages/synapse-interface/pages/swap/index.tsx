import { useAccount, useNetwork } from 'wagmi'
import { getWalletClient } from '@wagmi/core'

import toast from 'react-hot-toast'
import { useRouter } from 'next/router'
import { Address, zeroAddress } from 'viem'
import { commify } from '@ethersproject/units'
import { useEffect, useRef, useState } from 'react'

import { segmentAnalyticsEvent } from '@/contexts/segmentAnalyticsEvent'

import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'

import { SECTION_TRANSITION_PROPS } from '@/styles/transitions'

import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import type { Token } from '@/utils/types'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { approveToken } from '@/utils/approveToken'

import { Transition } from '@headlessui/react'

import { useAppDispatch } from '@/store/hooks'

import { setIsLoading } from '@/slices/swap/reducer'
import { useFetchPortfolioBalances } from '@/slices/portfolio/hooks'
import { useSwapDisplayState, useSwapState } from '@/slices/swap/hooks'
import { setSwapQuote, updateSwapFromValue } from '@/slices/swap/reducer'

import { EMPTY_SWAP_QUOTE_ZERO } from '@/constants/swap'
import { CHAINS_BY_ID } from '@/constants/chains'

import useSyncQueryParamsWithSwapState from '@/utils/hooks/useSyncQueryParamsWithSwapState'

import Card from '@tw/Card'

import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { PageHeader } from '@/components/PageHeader'
import ExplorerToastLink from '@/components/ExplorerToastLink'

import { SwapChainListOverlay } from '@/components/StateManagedSwap/SwapChainListOverlay'
import { SwapFromTokenListOverlay } from '@/components/StateManagedSwap/SwapFromTokenListOverlay'
import { SwapInputContainer } from '@/components/StateManagedSwap/SwapInputContainer'
import { SwapOutputContainer } from '@/components/StateManagedSwap/SwapOutputContainer'
import { SwapToTokenListOverlay } from '@/components/StateManagedSwap/SwapToTokenListOverlay'
import { SwapTransactionButton } from '@/components/StateManagedSwap/SwapTransactionButton'
import { SwapExchangeRateInfo } from '@/components/StateManagedSwap/SwapExchangeRateInfo'
import { OverlayTransition } from '@/components/bridgeSwap/OverlayTransition'

const StateManagedSwap = () => {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const { synapseSDK } = useSynapseContext()
  const swapDisplayRef = useRef(null)
  const quoteToastRef = useRef('')
  const currentSDKRequestID = useRef(0)
  const router = useRouter()
  const { query, pathname } = router

  useSyncQueryParamsWithSwapState()

  const { balances: portfolioBalances } = useFetchPortfolioBalances()

  const { swapChainId, swapFromToken, swapToToken, swapFromValue, swapQuote } =
    useSwapState()

  const {
    showSwapFromTokenListOverlay,
    showSwapChainListOverlay,
    showSwapToTokenListOverlay,
  } = useSwapDisplayState()

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
    if (swapFromToken) {
      if (swapFromToken?.addresses[swapChainId] === zeroAddress) {
        setIsApproved(true)
      } else {
        setIsApproved(
          swapQuote?.allowance &&
          stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId]) <=
            swapQuote.allowance
        )
      }
    }
  }, [swapQuote, swapFromToken, swapFromValue, swapChainId])

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

        toast.dismiss(quoteToastRef.current)

        const message = `Route found for swapping ${swapFromValue} ${swapFromToken.symbol} on ${CHAINS_BY_ID[swapChainId]?.name} to ${swapToToken.symbol}`
        console.log(message)

        quoteToastRef.current = toast(message, { duration: 3000 })
      }
    } catch (err) {
      console.log(err)
      if (thisRequestId === currentSDKRequestID.current) {
        toast.dismiss(quoteToastRef.current)

        let message: string
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

        quoteToastRef.current = toast(message, { duration: 3000 })
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
        swapFromToken?.addresses[swapChainId],
        stringToBigInt(swapFromValue, swapFromToken?.decimals[swapChainId])
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

    let pendingPopup: any
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

        toast.success(successToastContent, {
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
              transition-all duration-100 transform rounded-lg
            `}
          >
            <div ref={swapDisplayRef}>
              <div className="-mt-4 pb-4"> {/** hackfix for making sure origin point is top of the card */}
                <OverlayTransition show={showSwapChainListOverlay}>
                  <SwapChainListOverlay />
                </OverlayTransition>
                <OverlayTransition show={showSwapFromTokenListOverlay}>
                  <SwapFromTokenListOverlay />
                </OverlayTransition>
                <OverlayTransition show={showSwapToTokenListOverlay}>
                  <SwapToTokenListOverlay />
                </OverlayTransition>
              </div>
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
