import toast from 'react-hot-toast'
import { useEffect, useRef, useState } from 'react'
import { commify } from '@ethersproject/units'
import { Address, zeroAddress, isAddress } from 'viem'
import { polygon } from 'viem/chains'
import { useAccount } from 'wagmi'
import { useSelector } from 'react-redux'
import { useRouter } from 'next/router'
import {
  getWalletClient,
  getPublicClient,
  waitForTransactionReceipt,
} from '@wagmi/core'
import { InputContainer } from '@/components/StateManagedBridge/InputContainer'
import { OutputContainer } from '@/components/StateManagedBridge/OutputContainer'
import { BridgeExchangeRateInfo } from '@/components/StateManagedBridge/BridgeExchangeRateInfo'
import { BridgeTransactionButton } from '@/components/StateManagedBridge/BridgeTransactionButton'
import ExplorerToastLink from '@/components/ExplorerToastLink'
import { Warning } from '@/components/Warning'
import { SwitchButton } from '@/components/buttons/SwitchButton'
import { PageHeader } from '@/components/PageHeader'
import SettingsSlideOver from '@/components/StateManagedBridge/SettingsSlideOver'
import Button from '@/components/ui/tailwind/Button'
import { SettingsToggle } from '@/components/StateManagedBridge/SettingsToggle'
import { BridgeCard } from '@/components/ui/BridgeCard'
import { ConfirmDestinationAddressWarning } from '@/components/StateManagedBridge/BridgeWarnings'
import { EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'
import { AcceptedChainId, CHAINS_BY_ID } from '@/constants/chains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import {
  BridgeState,
  setFromChainId,
  setFromToken,
  setToChainId,
  setToToken,
  updateFromValue,
  setBridgeQuote,
  setIsLoading,
  setIsWalletPending,
  setDestinationAddress,
} from '@/slices/bridge/reducer'
import {
  setShowDestinationAddress,
  setShowSettingsSlideOver,
} from '@/slices/bridgeDisplaySlice'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'
import { formatBigIntToString } from '@/utils/bigint/format'
import { calculateExchangeRate } from '@/utils/calculateExchangeRate'
import { Token } from '@/utils/types'
import { txErrorHandler } from '@/utils/txErrorHandler'
import { approveToken } from '@/utils/approveToken'
import { stringToBigInt } from '@/utils/bigint/format'
import { fetchAndStoreSingleNetworkPortfolioBalances } from '@/slices/portfolio/hooks'
import {
  updatePendingBridgeTransaction,
  addPendingBridgeTransaction,
  removePendingBridgeTransaction,
} from '@/slices/transactions/actions'
import { useAppDispatch } from '@/store/hooks'
import { RootState } from '@/store/store'
import { getTimeMinutesFromNow } from '@/utils/time'
import { isTransactionReceiptError } from '@/utils/isTransactionReceiptError'
import { isTransactionUserRejectedError } from '@/utils/isTransactionUserRejectedError'
import { useMaintenance } from '@/components/Maintenance/Maintenance'
import { getBridgeModuleNames } from '@/utils/getBridgeModuleNames'
import { wagmiConfig } from '@/wagmiConfig'
import { useStaleQuoteUpdater } from '@/utils/hooks/useStaleQuoteUpdater'
import { screenAddress } from '@/utils/screenAddress'

const StateManagedBridge = () => {
  const { address } = useAccount()
  const { synapseSDK } = useSynapseContext()
  const router = useRouter()
  const { query, pathname } = router

  const bridgeDisplayRef = useRef(null)
  const currentSDKRequestID = useRef(0)
  const quoteToastRef = useRef({ id: '' })

  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    bridgeQuote,
    debouncedFromValue,
    destinationAddress,
    isLoading: isQuoteLoading,
    isWalletPending,
  }: BridgeState = useBridgeState()
  const { showSettingsSlideOver, showDestinationAddress } = useSelector(
    (state: RootState) => state.bridgeDisplay
  )

  const {
    isBridgePaused,
    pausedModulesList,
    BridgeMaintenanceProgressBar,
    BridgeMaintenanceWarningMessage,
  } = useMaintenance()

  const [isApproved, setIsApproved] = useState<boolean>(false)

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
    if (fromToken && fromToken?.addresses[fromChainId] === zeroAddress) {
      setIsApproved(true)
    } else {
      if (
        fromToken &&
        bridgeQuote?.allowance &&
        stringToBigInt(debouncedFromValue, fromToken.decimals[fromChainId]) <=
          bridgeQuote.allowance
      ) {
        setIsApproved(true)
      } else {
        setIsApproved(false)
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
      const currentTimestamp: number = getTimeMinutesFromNow(0)

      const allQuotes = await synapseSDK.allBridgeQuotes(
        fromChainId,
        toChainId,
        fromToken.addresses[fromChainId],
        toToken.addresses[toChainId],
        stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]),
        {
          originUserAddress: address,
        }
      )

      const pausedBridgeModules = new Set(
        pausedModulesList
          .filter((module) =>
            module.chainId ? module.chainId === fromChainId : true
          )
          .flatMap(getBridgeModuleNames)
      )

      const activeQuotes = allQuotes.filter(
        (quote) => !pausedBridgeModules.has(quote.bridgeModuleName)
      )

      if (activeQuotes.length === 0) {
        const msg = `No route found for bridging ${debouncedFromValue} ${fromToken?.symbol} on ${CHAINS_BY_ID[fromChainId]?.name} to ${toToken?.symbol} on ${CHAINS_BY_ID[toChainId]?.name}`
        throw new Error(msg)
      }

      const rfqQuote = activeQuotes.find(
        (quote) => quote.bridgeModuleName === 'SynapseRFQ'
      )

      const nonRfqQuote = activeQuotes.find(
        (quote) => quote.bridgeModuleName !== 'SynapseRFQ'
      )

      let quote

      if (rfqQuote && nonRfqQuote) {
        const rfqMaxAmountOut = BigInt(rfqQuote.maxAmountOut.toString())
        const nonRfqMaxAmountOut = BigInt(nonRfqQuote.maxAmountOut.toString())

        const allowedPercentileDifference = 30n
        const maxDifference =
          (nonRfqMaxAmountOut * allowedPercentileDifference) / 100n

        if (rfqMaxAmountOut > nonRfqMaxAmountOut - maxDifference) {
          quote = rfqQuote
        } else {
          quote = nonRfqQuote

          segmentAnalyticsEvent(`[Bridge] use non-RFQ quote over RFQ`, {
            bridgeModuleName: nonRfqQuote.bridgeModuleName,
            originChainId: fromChainId,
            originToken: fromToken.symbol,
            originTokenAddress: fromToken.addresses[fromChainId],
            destinationChainId: toChainId,
            destinationToken: toToken.symbol,
            destinationTokenAddress: toToken.addresses[toChainId],
            rfqQuoteAmountOut: rfqQuote.maxAmountOut.toString(),
            nonRfqMaxAmountOut: nonRfqQuote.maxAmountOut.toString(),
          })
        }
      } else {
        quote = rfqQuote ?? nonRfqQuote
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
        originChainId,
        destChainId,
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
            timestamp: currentTimestamp,
            originChainId,
            destChainId,
          })
        )

        toast.dismiss(quoteToastRef.current.id)

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

  useStaleQuoteUpdater(
    bridgeQuote,
    getAndSetBridgeQuote,
    isQuoteLoading,
    isWalletPending
  )

  const approveTxn = async () => {
    try {
      dispatch(setIsWalletPending(true))
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
    } finally {
      dispatch(setIsWalletPending(false))
    }
  }

  const executeBridge = async () => {
    let pendingPopup: any

    if (destinationAddress) {
      const isRisky = await screenAddress(destinationAddress)
      if (isRisky) {
        return
      }
    }

    segmentAnalyticsEvent(
      `[Bridge] initiates bridge`,
      {
        originChainId: fromChainId,
        destinationChainId: toChainId,
        inputAmount: debouncedFromValue,
        expectedReceivedAmount: bridgeQuote.outputAmountString,
        slippage: bridgeQuote.exchangeRate,
        originToken: fromToken?.routeSymbol,
        destinationToken: toToken?.routeSymbol,
        exchangeRate: BigInt(bridgeQuote.exchangeRate.toString()),
        routerAddress: bridgeQuote.routerAddress,
        bridgeQuote,
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
        destinationAddress: destinationAddress,
      })
    )
    try {
      dispatch(setIsWalletPending(true))
      const wallet = await getWalletClient(wagmiConfig, {
        chainId: fromChainId,
      })
      const toAddress =
        destinationAddress && isAddress(destinationAddress)
          ? destinationAddress
          : address

      const data = await synapseSDK.bridge(
        toAddress,
        bridgeQuote.routerAddress,
        fromChainId,
        toChainId,
        fromToken?.addresses[fromChainId as keyof Token['addresses']],
        stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]),
        bridgeQuote.originQuery,
        bridgeQuote.destQuery
      )

      const payload =
        fromToken?.addresses[fromChainId as keyof Token['addresses']] ===
          zeroAddress ||
        fromToken?.addresses[fromChainId as keyof Token['addresses']] === ''
          ? {
              data: data.data,
              to: data.to,
              value: stringToBigInt(
                debouncedFromValue,
                fromToken?.decimals[fromChainId]
              ),
            }
          : data

      /** Setting custom gas limit for only Polygon transactions */
      let gasEstimate = undefined

      if (fromChainId === polygon.id) {
        const publicClient = getPublicClient(wagmiConfig, {
          chainId: fromChainId,
        })
        gasEstimate = await publicClient.estimateGas({
          value: payload.value,
          to: payload.to,
          account: address,
          data: payload.data,
        })
        gasEstimate = (gasEstimate * 3n) / 2n
      }

      const tx = await wallet.sendTransaction({
        ...payload,
        gas: gasEstimate,
      })

      const originChainName = CHAINS_BY_ID[fromChainId]?.name
      const destinationChainName = CHAINS_BY_ID[toChainId]?.name
      pendingPopup = toast(
        `Bridging from ${fromToken?.symbol} on ${originChainName} to ${toToken.symbol} on ${destinationChainName}`,
        { id: 'bridge-in-progress-popup', duration: Infinity }
      )
      segmentAnalyticsEvent(`[Bridge] bridges successfully`, {
        originChainId: fromChainId,
        destinationChainId: toChainId,
        inputAmount: debouncedFromValue,
        expectedReceivedAmount: bridgeQuote.outputAmountString,
        slippage: bridgeQuote.exchangeRate,
        originToken: fromToken?.routeSymbol,
        destinationToken: toToken?.routeSymbol,
        exchangeRate: BigInt(bridgeQuote.exchangeRate.toString()),
        routerAddress: bridgeQuote.routerAddress,
        bridgeQuote,
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

      const transactionReceipt = await waitForTransactionReceipt(wagmiConfig, {
        hash: tx as Address,
        timeout: 60_000,
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
        errorCode: error.code,
      })
      dispatch(removePendingBridgeTransaction(currentTimestamp))
      console.log('Error executing bridge', error)
      toast.dismiss(pendingPopup)

      if (isTransactionUserRejectedError(error)) {
        getAndSetBridgeQuote()
      }

      /** Fetch balances if await transaction receipt times out */
      if (isTransactionReceiptError(error)) {
        dispatch(
          fetchAndStoreSingleNetworkPortfolioBalances({
            address,
            chainId: fromChainId,
          })
        )
      }

      return txErrorHandler(error)
    } finally {
      dispatch(setIsWalletPending(false))
    }
  }

  return (
    <div className="flex flex-col w-full max-w-lg mx-auto lg:mx-0">
      <div className="flex flex-col">
        <div className="flex items-center justify-between">
          <PageHeader
            title="Bridge"
            subtitle="Send your assets across chains."
          />
          <Button
            className="flex items-center p-3 text-opacity-75 bg-bgLight hover:bg-bgLighter text-secondaryTextColor hover:text-white"
            onClick={() =>
              dispatch(setShowSettingsSlideOver(!showSettingsSlideOver))
            }
          >
            <SettingsToggle showSettingsToggle={!showSettingsSlideOver} />
          </Button>
        </div>
        <BridgeCard bridgeRef={bridgeDisplayRef}>
          <BridgeMaintenanceProgressBar />

          {showSettingsSlideOver ? (
            <div className="min-h-[472px]">
              <SettingsSlideOver key="settings" />
            </div>
          ) : (
            <>
              <InputContainer />
              <SwitchButton
                onClick={() => {
                  dispatch(setFromChainId(toChainId))
                  dispatch(setFromToken(toToken))
                  dispatch(setToChainId(fromChainId))
                  dispatch(setToToken(fromToken))
                }}
              />
              <OutputContainer />
              <Warning />
              <BridgeMaintenanceWarningMessage />
              <BridgeExchangeRateInfo />
              <ConfirmDestinationAddressWarning />
              <BridgeTransactionButton
                isApproved={isApproved}
                approveTxn={approveTxn}
                executeBridge={executeBridge}
                isBridgePaused={isBridgePaused}
              />
            </>
          )}
        </BridgeCard>
      </div>
    </div>
  )
}

export default StateManagedBridge
