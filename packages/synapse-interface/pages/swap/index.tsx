import { useAccount } from 'wagmi'
import toast from 'react-hot-toast'
import { useRouter } from 'next/router'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import {
  setIsLoading,
  setSwapFromToken,
  setSwapToToken,
} from '@/slices/swap/reducer'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { commify } from '@ethersproject/units'
import { formatBigIntToString } from '@/utils/bigint/format'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { useEffect, useRef, useState } from 'react'
import { Token } from '@/utils/types'
import { getWalletClient, waitForTransactionReceipt } from '@wagmi/core'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { CHAINS_BY_ID } from '@/constants/chains'
import { approveToken } from '@/utils/approveToken'
import { PageHeader } from '@/components/PageHeader'
import { BridgeCard } from '@/components/ui/BridgeCard'
import ExplorerToastLink from '@/components/ExplorerToastLink'
import { Address, zeroAddress } from 'viem'
import { stringToBigInt } from '@/utils/bigint/format'
import { useAppDispatch } from '@/store/hooks'
import {
  useFetchPortfolioBalances,
  fetchAndStoreSingleNetworkPortfolioBalances,
} from '@/slices/portfolio/hooks'
import { SwapTransactionButton } from '@/components/StateManagedSwap/SwapTransactionButton'
import SwapExchangeRateInfo from '@/components/StateManagedSwap/SwapExchangeRateInfo'
import { useSwapState } from '@/slices/swap/hooks'
import { SwapInputContainer } from '@/components/StateManagedSwap/SwapInputContainer'
import { SwapOutputContainer } from '@/components/StateManagedSwap/SwapOutputContainer'
import { setSwapQuote, updateSwapFromValue } from '@/slices/swap/reducer'
import { EMPTY_SWAP_QUOTE_ZERO } from '@/constants/swap'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import useSyncQueryParamsWithSwapState from '@/utils/hooks/useSyncQueryParamsWithSwapState'
import { isTransactionReceiptError } from '@/utils/isTransactionReceiptError'
import { wagmiConfig } from '@/wagmiConfig'
import { SwitchButton } from '@/components/buttons/SwitchButton'
import { useMaintenance } from '@/components/Maintenance/Maintenance'
import { useWalletState } from '@/slices/wallet/hooks'
import { setIsWalletPending } from '@/slices/wallet/reducer'

const StateManagedSwap = () => {
  const { address } = useAccount()
  const { synapseSDK } = useSynapseContext()
  const swapDisplayRef = useRef(null)
  const quoteToastRef = useRef({ id: '' })
  const currentSDKRequestID = useRef(0)
  const router = useRouter()
  const { query, pathname } = router

  const [isTyping, setIsTyping] = useState(false)

  useSyncQueryParamsWithSwapState()

  const { balances: portfolioBalances } = useFetchPortfolioBalances()

  const { swapChainId, swapFromToken, swapToToken, swapFromValue, swapQuote } =
    useSwapState()

  const { isWalletPending } = useWalletState()

  const {
    isSwapPaused,
    pausedChainsList,
    pausedModulesList,
    SwapMaintenanceProgressBar,
    SwapMaintenanceWarningMessage,
  } = useMaintenance()

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

        toast.dismiss(quoteToastRef.current.id)

        const message = `Route found for swapping ${swapFromValue} ${swapFromToken.symbol} on ${CHAINS_BY_ID[swapChainId]?.name} to ${swapToToken.symbol}`
        console.log(message)

        quoteToastRef.current.id = toast(message, { duration: 3000 })
      }
    } catch (err) {
      console.log(err)
      if (thisRequestId === currentSDKRequestID.current) {
        toast.dismiss(quoteToastRef.current.id)

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

        quoteToastRef.current.id = toast(message, { duration: 3000 })
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
      dispatch(setIsWalletPending(true))
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
    } finally {
      dispatch(setIsWalletPending(false))
    }
  }

  const onSuccessSwap = () => {
    dispatch(
      fetchAndStoreSingleNetworkPortfolioBalances({
        address,
        chainId: swapChainId,
      })
    )
    dispatch(setSwapQuote(EMPTY_SWAP_QUOTE_ZERO))
    dispatch(updateSwapFromValue(''))
  }
  const executeSwap = async () => {
    const currentChainName = CHAINS_BY_ID[swapChainId]?.name

    dispatch(setIsWalletPending(true))

    let pendingPopup: any
    pendingPopup = toast(
      `Initiating swap from ${swapFromToken.symbol} to ${swapToToken.symbol} on ${currentChainName}`,
      { id: 'swap-in-progress-popup', duration: Infinity }
    )
    segmentAnalyticsEvent(
      `[Swap] initiates swap`,
      {
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
      const wallet = await getWalletClient(wagmiConfig, {
        chainId: swapChainId,
      })

      const payload = await synapseSDK.swap(
        swapChainId,
        address,
        swapFromToken.addresses[swapChainId],
        stringToBigInt(swapFromValue, swapFromToken.decimals[swapChainId]),
        swapQuote.quote
      )

      const tx = await wallet.sendTransaction(payload)

      const originChainName = CHAINS_BY_ID[swapChainId]?.name
      pendingPopup = toast(
        `Swapping ${swapFromToken.symbol} on ${originChainName} to ${swapToToken.symbol}`,
        { id: 'swap-in-progress-popup', duration: Infinity }
      )

      const transactionReceipt = await waitForTransactionReceipt(wagmiConfig, {
        hash: tx as Address,
        timeout: 60_000,
      })
      console.log('Transaction Receipt: ', transactionReceipt)

      onSuccessSwap()

      segmentAnalyticsEvent(`[Swap] swaps successfully`, {
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

      return tx
    } catch (error) {
      console.log(`Swap Execution failed with error: ${error}`)

      /** Assume successful swap tx if await transaction receipt times out */
      if (isTransactionReceiptError(error)) {
        onSuccessSwap()
      }

      toast.dismiss(pendingPopup)
      txErrorHandler(error)
    } finally {
      dispatch(setIsWalletPending(false))
    }
  }

  return (
    <LandingPageWrapper>
      <div className="flex justify-center px-4 py-16 mx-auto lg:mx-0">
        <div className="flex flex-col">
          <div className="flex items-center justify-between">
            <PageHeader title="Swap" subtitle="Exchange assets on chain." />
          </div>
          <BridgeCard bridgeRef={swapDisplayRef}>
            <SwapMaintenanceProgressBar />
            <SwapInputContainer setIsTyping={setIsTyping} />
            <SwitchButton
              onClick={() => {
                dispatch(setSwapFromToken(swapToToken))
                dispatch(setSwapToToken(swapFromToken))
              }}
              disabled={isWalletPending}
            />
            <SwapOutputContainer />
            <SwapMaintenanceWarningMessage />
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
            <SwapTransactionButton
              isTyping={isTyping}
              isApproved={isApproved}
              approveTxn={approveTxn}
              executeSwap={executeSwap}
              isSwapPaused={isSwapPaused}
            />
          </BridgeCard>
        </div>
      </div>
    </LandingPageWrapper>
  )
}

export default StateManagedSwap
