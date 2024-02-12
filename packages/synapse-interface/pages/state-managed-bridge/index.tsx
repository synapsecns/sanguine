import { getWalletClient, waitForTransaction } from '@wagmi/core'
import { useAccount, useNetwork } from 'wagmi'
import { useSelector } from 'react-redux'
import { useEffect, useRef, useState, useMemo } from 'react'
import toast from 'react-hot-toast'
import { useRouter } from 'next/router'
import { Address, zeroAddress } from 'viem'
import { commify } from '@ethersproject/units'
import { isAddress } from '@ethersproject/address'

import { RootState } from '@/store/store'
import { useAppDispatch } from '@/store/hooks'

import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState } from '@/slices/bridge/reducer'
import {
  updateFromValue,
  setBridgeQuote,
  setIsLoading,
  setDestinationAddress,
} from '@/slices/bridge/reducer'
import {
  setShowDestinationAddress,
  setShowSettingsSlideOver,
} from '@/slices/bridgeDisplaySlice'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'
import {
  updatePendingBridgeTransaction,
  addPendingBridgeTransaction,
  removePendingBridgeTransaction,
} from '@/slices/transactions/actions'
import {
  fetchArbPrice,
  fetchEthPrice,
  fetchGmxPrice,
} from '@/slices/priceDataSlice'

import { EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'
import { AcceptedChainId, CHAINS_BY_ID } from '@/constants/chains'

import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'

import { formatBigIntToString, stringToBigInt } from '@/utils/bigint/format'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'

import { Token } from '@/utils/types'
import { getTimeMinutesFromNow } from '@/utils/time'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { approveToken } from '@/utils/approveToken'
import { SECTION_TRANSITION_PROPS } from '@/styles/transitions'


import Card from '@tw/Card'
import Button from '@tw/Button'
import { Transition } from '@headlessui/react'
import { PageHeader } from '@/components/PageHeader'
import { SettingsIcon } from '@/components/icons/SettingsIcon'

import ExplorerToastLink from '@/components/ExplorerToastLink'
import { Warning } from '@/components/Warning'

import { FromChainListOverlay } from '@/components/StateManagedBridge/FromChainListOverlay'
import { ToChainListOverlay } from '@/components/StateManagedBridge/ToChainListOverlay'
import { FromTokenListOverlay } from '@/components/StateManagedBridge/FromTokenListOverlay'
import { ToTokenListOverlay } from '@/components/StateManagedBridge/ToTokenListOverlay'
import { InputContainer } from '@/components/StateManagedBridge/InputContainer'
import { OutputContainer } from '@/components/StateManagedBridge/OutputContainer'
import { DestinationAddressInput } from '@/components/StateManagedBridge/DestinationAddressInput'
import { BridgeTransactionButton } from '@/components/StateManagedBridge/BridgeTransactionButton'
import SettingsSlideOver from '@/components/StateManagedBridge/SettingsSlideOver'
import BridgeExchangeRateInfo from '@/components/StateManagedBridge/BridgeExchangeRateInfo'

import { XIcon } from '@heroicons/react/outline'
import { AnimatedOverlay } from '@/components/AnimatedOverlay'


const StateManagedBridge = () => {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const { synapseSDK } = useSynapseContext()
  const bridgeDisplayRef = useRef(null)
  const currentSDKRequestID = useRef(0)
  const quoteToastRef = useRef({ id: '' })
  const router = useRouter()
  const { query, pathname } = router

  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    bridgeQuote,
    debouncedFromValue,
    destinationAddress,
  }: BridgeState = useBridgeState()
  const {
    showSettingsSlideOver,
    showDestinationAddress,
    showFromChainListOverlay,
    showToChainListOverlay,
    showFromTokenListOverlay,
    showToTokenListOverlay,
  } = useSelector((state: RootState) => state.bridgeDisplay)

  const [isApproved, setIsApproved] = useState(false)

  const dispatch = useAppDispatch()

  useEffect(() => {
    segmentAnalyticsEvent(`[Bridge page] arrives`, {
      fromChainId,
      query,
      pathname,
    })
  }, [query])

  useEffect(() => {
    if (
      fromToken &&
      toToken &&
      fromToken?.decimals[fromChainId] &&
      stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]) > 0n
    ) {
      console.log('trying to set bridge quote')
      getAndSetBridgeQuote()
    } else {
      dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
      dispatch(setIsLoading(false))
    }
  }, [fromChainId, toChainId, fromToken, toToken, debouncedFromValue])

  // don't like this, rewrite: could be custom hook
  useEffect(() => {
    if (fromToken) {
      if (fromToken?.addresses[fromChainId] === zeroAddress) {
        setIsApproved(true)
      } else {
        setIsApproved(
          bridgeQuote?.allowance &&
            stringToBigInt(
              debouncedFromValue, fromToken.decimals[fromChainId]
            ) <= bridgeQuote.allowance
        )
      }
    }
  }, [bridgeQuote, fromToken, debouncedFromValue, fromChainId, toChainId])

  const getAndSetBridgeQuote = async () => {
    currentSDKRequestID.current += 1
    const thisRequestId = currentSDKRequestID.current
    // will have to handle deadlineMinutes here at later time, gets passed as optional last arg in .bridgeQuote()

    /* clear stored bridge quote before requesting new bridge quote */
    dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))

    try {
      dispatch(setIsLoading(true))

      const allQuotes = await synapseSDK.allBridgeQuotes(
        fromChainId,
        toChainId,
        fromToken.addresses[fromChainId],
        toToken.addresses[toChainId],
        stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId])
      )

      if (allQuotes.length === 0) {
        const msg = `No route found for bridging ${debouncedFromValue} ${fromToken?.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken?.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
        throw new Error(msg)
      }

      const rfqQuote = allQuotes.find(
        (quote) => quote.bridgeModuleName === 'SynapseRFQ'
      )

      let quote

      if (rfqQuote) {
        quote = rfqQuote
      } else {
        /* allBridgeQuotes returns sorted quotes by maxAmountOut descending */
        quote = allQuotes[0]
      }

      const {
        feeAmount,
        routerAddress,
        maxAmountOut,
        originQuery,
        destQuery,
        estimatedTime,
        bridgeModuleName,
        gasDropAmount,
      } = quote

      if (!(originQuery && maxAmountOut && destQuery && feeAmount)) {
        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
        dispatch(setIsLoading(false))
        return
      }

      const toValueBigInt = BigInt(maxAmountOut.toString()) ?? 0n

      // Bridge Lifecycle: originToken -> bridgeToken -> destToken
      // debouncedFromValue is in originToken decimals
      // originQuery.minAmountOut and feeAmount is in bridgeToken decimals
      // Adjust feeAmount to be in originToken decimals
      const adjustedFeeAmount =
        (BigInt(feeAmount) *
          stringToBigInt(
            `${debouncedFromValue}`,
            fromToken?.decimals[fromChainId]
          )) /
        BigInt(originQuery.minAmountOut)

      const isUnsupported = AcceptedChainId[fromChainId] ? false : true

      const allowance =
        fromToken?.addresses[fromChainId] === zeroAddress ||
        address === undefined ||
        isUnsupported
          ? 0n
          : await getErc20TokenAllowance({
              address,
              chainId: fromChainId,
              tokenAddress: fromToken?.addresses[fromChainId] as Address,
              spender: routerAddress,
            })

      const {
        originQuery: originQueryWithSlippage,
        destQuery: destQueryWithSlippage,
      } = synapseSDK.applyBridgeSlippage(
        bridgeModuleName,
        originQuery,
        destQuery
      )

      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(
          setBridgeQuote({
            outputAmount: toValueBigInt,
            outputAmountString: commify(
              formatBigIntToString(
                toValueBigInt,
                toToken.decimals[toChainId],
                8
              )
            ),
            routerAddress,
            allowance,
            exchangeRate: calculateExchangeRate(
              stringToBigInt(
                debouncedFromValue,
                fromToken?.decimals[fromChainId]
              ) - BigInt(adjustedFeeAmount),
              fromToken?.decimals[fromChainId],
              toValueBigInt,
              toToken.decimals[toChainId]
            ),
            feeAmount,
            delta: BigInt(maxAmountOut.toString()),
            originQuery: originQueryWithSlippage,
            destQuery: destQueryWithSlippage,
            estimatedTime: estimatedTime,
            bridgeModuleName: bridgeModuleName,
            gasDropAmount: BigInt(gasDropAmount.toString()),
          })
        )

        toast.dismiss(quoteToastRef.current.id)

        dispatch(fetchEthPrice())
        dispatch(fetchArbPrice())
        dispatch(fetchGmxPrice())

        const message = `Route found for bridging ${debouncedFromValue} ${fromToken?.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
        console.log(message)

        quoteToastRef.current.id = toast(message, { duration: 3000 })
      }
    } catch (err) {
      console.log(err)
      if (thisRequestId === currentSDKRequestID.current) {
        toast.dismiss(quoteToastRef.current.id)

        let message: string
        if (!fromChainId) {
          message = 'Please select an origin chain'
        } else if (!toChainId) {
          message = 'Please select a destination chain'
        } else if (!fromToken) {
          message = 'Please select an origin token'
        } else if (!toToken) {
          message = 'Please select a destination token'
        } else {
          message = `No route found for bridging ${debouncedFromValue} ${fromToken?.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
        }
        console.log(message)

        quoteToastRef.current.id = toast(message, { duration: 3000 })
        dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))

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
        bridgeQuote?.routerAddress,
        fromChainId,
        fromToken?.addresses[fromChainId],
        stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId])
      )
      await tx
      /** Re-fetch bridge quote to re-check approval state */
      getAndSetBridgeQuote()
    } catch (error) {
      return txErrorHandler(error)
    }
  }

  const executeBridge = async () => {
    let pendingPopup: any
    segmentAnalyticsEvent(
      `[Bridge] initiates bridge`,
      {
        address,
        originChainId: fromChainId,
        destinationChainId: toChainId,
        inputAmount: debouncedFromValue,
        expectedReceivedAmount: bridgeQuote.outputAmountString,
        slippage: bridgeQuote.exchangeRate,
        originToken: fromToken?.routeSymbol,
        destinationToken: toToken?.routeSymbol,
        exchangeRate: BigInt(bridgeQuote.exchangeRate.toString()),
        routerAddress: bridgeQuote.routerAddress,
      },
      true
    )
    const currentTimestamp: number = getTimeMinutesFromNow(0)
    dispatch(
      addPendingBridgeTransaction({
        id: currentTimestamp,
        originChain: CHAINS_BY_ID[fromChainId],
        originToken: fromToken,
        originValue: debouncedFromValue,
        destinationChain: CHAINS_BY_ID[toChainId],
        destinationToken: toToken,
        transactionHash: undefined,
        timestamp: undefined,
        isSubmitted: false,
        estimatedTime: bridgeQuote.estimatedTime,
        bridgeModuleName: bridgeQuote.bridgeModuleName,
      })
    )
    try {
      const wallet = await getWalletClient({
        chainId: fromChainId,
      })
      const toAddress =
        destinationAddress && isAddress(destinationAddress)
          ? destinationAddress
          : address

      const fromTokenAddress = fromToken?.addresses[fromChainId as keyof Token['addresses']]

      const data = await synapseSDK.bridge(
        toAddress,
        bridgeQuote.routerAddress,
        fromChainId,
        toChainId,
        fromTokenAddress,
        stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]),
        bridgeQuote.originQuery,
        bridgeQuote.destQuery
      )

      const payload =
        [zeroAddress, ''].includes(fromTokenAddress)
          ? {
              data: data.data,
              to: data.to,
              value: stringToBigInt(
                debouncedFromValue,
                fromToken?.decimals[fromChainId]
              ),
            }
          : data

      const tx = await wallet.sendTransaction(payload)

      const originChainName = CHAINS_BY_ID[fromChainId]?.name
      const destinationChainName = CHAINS_BY_ID[toChainId]?.name
      pendingPopup = toast(
        `Bridging from ${fromToken?.symbol} on ${originChainName} to ${toToken.symbol} on ${destinationChainName}`,
        { id: 'bridge-in-progress-popup', duration: Infinity }
      )
      segmentAnalyticsEvent(`[Bridge] bridges successfully`, {
        address,
        originChainId: fromChainId,
        destinationChainId: toChainId,
        inputAmount: debouncedFromValue,
        expectedReceivedAmount: bridgeQuote.outputAmountString,
        slippage: bridgeQuote.exchangeRate,
        originToken: fromToken?.routeSymbol,
        destinationToken: toToken?.routeSymbol,
        exchangeRate: BigInt(bridgeQuote.exchangeRate.toString()),
        routerAddress: bridgeQuote.routerAddress,
      })
      dispatch(
        updatePendingBridgeTransaction({
          id: currentTimestamp,
          timestamp: undefined,
          transactionHash: tx,
          isSubmitted: false,
        })
      )
      dispatch(setBridgeQuote(EMPTY_BRIDGE_QUOTE_ZERO))
      dispatch(setDestinationAddress(null))
      dispatch(setShowDestinationAddress(false))
      dispatch(updateFromValue(''))

      const successToastContent = (
        <div>
          <div>
            Successfully initiated bridge from {fromToken?.symbol} on{' '}
            {originChainName} to {toToken.symbol} on {destinationChainName}
          </div>
          <ExplorerToastLink
            transactionHash={tx ?? zeroAddress}
            chainId={fromChainId}
          />
        </div>
      )

      toast.success(successToastContent, {
        id: 'bridge-success-popup',
        duration: 10000,
      })

      toast.dismiss(pendingPopup)

      const transactionReceipt = await waitForTransaction({
        hash: tx as Address,
        timeout: 30_000,
      })
      console.log('Transaction Receipt: ', transactionReceipt)

      /** Update Origin Chain token balances after resolved tx or timeout reached */
      /** Assume tx has been actually resolved if above times out */
      dispatch(
        fetchAndStoreSingleNetworkPortfolioBalances({
          address,
          chainId: fromChainId,
        })
      )

      return tx
    } catch (error) {
      segmentAnalyticsEvent(`[Bridge]  error bridging`, {
        address,
        errorCode: error.code,
      })
      dispatch(removePendingBridgeTransaction(currentTimestamp))
      console.log('Error executing bridge', error)
      toast.dismiss(pendingPopup)
      return txErrorHandler(error)
    }
  }

  const springClass =
    '-mt-4 fixed z-50 w-full h-full bg-opacity-50 bg-slate-400/10'

  return (
    <div className="flex flex-col w-full max-w-lg mx-auto lg:mx-0">
      <div className="flex flex-col">
        <div className="flex items-center justify-between">
          <PageHeader
            title="Bridge"
            subtitle="Send your assets across chains."
          />
          <div>
            <Button
              className="group flex items-center p-2 text-opacity-75 bg-bgBase/10 hover:bg-bgBase/20 ring-1 ring-white/10 hover:ring-white/30 text-secondaryTextColor hover:text-white"
              onClick={() => {
                dispatch(setShowSettingsSlideOver(showSettingsSlideOver !== true))
              }}
            >
              {!showSettingsSlideOver ? (
                <>
                  <SettingsIcon className="w-4 h-4 mr-2 group-hover:animate-spin" />
                  <span className='text-sm mr-1'>Settings</span>
                </>
              ) : (
                <>
                  <XIcon className="w-4 h-4 mr-2"/>
                  <span className='text-sm mr-1'>Close</span>
                </>

              )}
            </Button>
          </div>
        </div>
        <Card
          divider={false}
          className={`
            pb-3 mt-5 overflow-hidden
            transition-all duration-100 transform rounded-lg
            bg-bgBase/10
          `}
        >
          <div ref={bridgeDisplayRef}>
            <AnimatedOverlay show={showSettingsSlideOver}>
              <SettingsSlideOver key="settings" />
            </AnimatedOverlay>
            <AnimatedOverlay show={showFromChainListOverlay}>
              <FromChainListOverlay />
            </AnimatedOverlay>
            <AnimatedOverlay show={showFromTokenListOverlay}>
              <FromTokenListOverlay />
            </AnimatedOverlay>
            <AnimatedOverlay show={showToChainListOverlay}>
              <ToChainListOverlay />
            </AnimatedOverlay>
            <AnimatedOverlay show={showToTokenListOverlay}>
              <ToTokenListOverlay />
            </AnimatedOverlay>
            <InputContainer />
            <OutputContainer />
            <Warning />
            <Transition
              appear={true}
              unmount={false}
              show={true}
              {...SECTION_TRANSITION_PROPS}
            >
              <BridgeExchangeRateInfo />
            </Transition>
            {showDestinationAddress && (
              <DestinationAddressInput
                toChainId={toChainId}
                destinationAddress={destinationAddress}
              />
            )}
            <div className="md:my-3">
              <BridgeTransactionButton
                isApproved={isApproved}
                approveTxn={approveTxn}
                executeBridge={executeBridge}
              />
            </div>
          </div>
        </Card>
      </div>
    </div>
  )
}

export default StateManagedBridge
