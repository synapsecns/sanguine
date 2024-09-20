import toast from 'react-hot-toast'
import { useEffect, useRef, useState } from 'react'
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
import { useTranslations } from 'next-intl'

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
import { CHAINS_BY_ID } from '@/constants/chains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import {
  BridgeState,
  setFromChainId,
  setFromToken,
  setToChainId,
  setToToken,
  setDestinationAddress,
  updateDebouncedFromValue,
} from '@/slices/bridge/reducer'
import { setIsWalletPending } from '@/slices/wallet/reducer'
import {
  setShowDestinationAddress,
  setShowSettingsSlideOver,
} from '@/slices/bridgeDisplaySlice'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
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
import { getUnixTimeMinutesFromNow } from '@/utils/time'
import { isTransactionReceiptError } from '@/utils/isTransactionReceiptError'
import { wagmiConfig } from '@/wagmiConfig'
import { useStaleQuoteUpdater } from '@/utils/hooks/useStaleQuoteUpdater'
import { useMaintenance } from '@/components/Maintenance/Maintenance'
import { screenAddress } from '@/utils/screenAddress'
import { useWalletState } from '@/slices/wallet/hooks'
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks'
import { resetBridgeQuote } from '@/slices/bridgeQuote/reducer'
import { fetchBridgeQuote } from '@/slices/bridgeQuote/thunks'
import { useIsBridgeApproved } from '@/utils/hooks/useIsBridgeApproved'

const StateManagedBridge = () => {
  const dispatch = useAppDispatch()
  const { address } = useAccount()
  const { synapseSDK } = useSynapseContext()
  const router = useRouter()
  const { query, pathname } = router

  const bridgeDisplayRef = useRef(null)
  const currentSDKRequestID = useRef(0)
  const quoteToastRef = useRef({ id: '' })
  const quoteTimeout = 15000

  const t = useTranslations('Bridge')

  const [isTyping, setIsTyping] = useState(false)

  const {
    fromChainId,
    toChainId,
    fromToken,
    toToken,
    debouncedFromValue,
    destinationAddress,
  }: BridgeState = useBridgeState()

  const { bridgeQuote, isLoading } = useBridgeQuoteState()

  const isApproved = useIsBridgeApproved()

  const { isWalletPending } = useWalletState()

  const { showSettingsSlideOver } = useSelector(
    (state: RootState) => state.bridgeDisplay
  )

  const {
    isBridgePaused,
    pausedModulesList,
    BridgeMaintenanceProgressBar,
    BridgeMaintenanceWarningMessage,
  } = useMaintenance()

  useEffect(() => {
    segmentAnalyticsEvent(
      `[Bridge page] arrives`,
      {
        fromChainId,
        query,
        pathname,
      },
      true
    )
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
      dispatch(resetBridgeQuote())
    }
  }, [fromChainId, toChainId, fromToken, toToken, debouncedFromValue])

  const getAndSetBridgeQuote = async () => {
    currentSDKRequestID.current += 1
    const thisRequestId = currentSDKRequestID.current

    // will have to handle deadlineMinutes here at later time, gets passed as optional last arg in .bridgeQuote()

    /* clear stored bridge quote before requesting new bridge quote */
    dispatch(resetBridgeQuote())
    const currentTimestamp: number = getUnixTimeMinutesFromNow(0)

    try {
      if (thisRequestId === currentSDKRequestID.current) {
        const result = await dispatch(
          fetchBridgeQuote({
            synapseSDK,
            fromChainId,
            toChainId,
            fromToken,
            toToken,
            debouncedFromValue,
            requestId: thisRequestId,
            currentTimestamp,
            address,
            pausedModulesList,
          })
        )

        toast.dismiss(quoteToastRef.current.id)

        if (fetchBridgeQuote.fulfilled.match(result)) {
          const message = t(
            'Route found for bridging {debouncedFromValue} {fromToken} on {fromChainId} to {toToken} on {toChainId}',
            {
              debouncedFromValue: debouncedFromValue,
              fromToken: fromToken?.symbol,
              fromChainId: CHAINS_BY_ID[fromChainId]?.name,
              toToken: toToken?.symbol,
              toChainId: CHAINS_BY_ID[toChainId]?.name,
            }
          )

          quoteToastRef.current.id = toast(message, { duration: 3000 })
        }

        if (fetchBridgeQuote.rejected.match(result)) {
          const message = t(
            'No route found for bridging {debouncedFromValue} {fromToken} on {fromChainId} to {toToken} on {toChainId}',
            {
              debouncedFromValue: debouncedFromValue,
              fromToken: fromToken?.symbol,
              fromChainId: CHAINS_BY_ID[fromChainId]?.name,
              toToken: toToken?.symbol,
              toChainId: CHAINS_BY_ID[toChainId]?.name,
            }
          )

          quoteToastRef.current.id = toast(message, { duration: 3000 })
        }
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
        dispatch(resetBridgeQuote())

        return
      }
    }
  }

  useStaleQuoteUpdater(
    bridgeQuote,
    getAndSetBridgeQuote,
    isLoading,
    isWalletPending,
    quoteTimeout
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

    const currentTimestamp: number = getUnixTimeMinutesFromNow(0)

    if (destinationAddress) {
      const isRisky = await screenAddress(destinationAddress)
      if (isRisky) {
        return
      }
    }

    segmentAnalyticsEvent(
      `[Bridge] initiates bridge`,
      {
        id: bridgeQuote.id,
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

      const payload = await synapseSDK.bridge(
        toAddress,
        bridgeQuote.routerAddress,
        fromChainId,
        toChainId,
        fromToken?.addresses[fromChainId as keyof Token['addresses']],
        stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]),
        bridgeQuote.originQuery,
        bridgeQuote.destQuery
      )

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
        id: bridgeQuote.id,
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
      dispatch(resetBridgeQuote())
      dispatch(setDestinationAddress(null))
      dispatch(setShowDestinationAddress(false))
      dispatch(updateDebouncedFromValue(''))

      const successToastContent = (
        <div>
          <div>
            {t(
              'Succesfully initiated bridge from {fromToken} on {originChainName} to {toToken} on {destinationChainName}',
              {
                fromToken: fromToken?.symbol,
                originChainName: originChainName,
                toToken: toToken?.symbol,
                destinationChainName: destinationChainName,
              }
            )}
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

      await waitForTransactionReceipt(wagmiConfig, {
        hash: tx as Address,
        timeout: 60_000,
      })

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
      console.error('Error executing bridge: ', error)
      toast.dismiss(pendingPopup)

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
            title={t('Bridge')}
            subtitle={t('Send your assets across chains')}
          />
          <Button
            className="flex items-center p-3 text-opacity-75 bg-bgLight hover:bg-bgLighter text-secondaryTextColor hover:text-white"
            onClick={() =>
              dispatch(setShowSettingsSlideOver(!showSettingsSlideOver))
            }
            disabled={isWalletPending}
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
              <InputContainer setIsTyping={setIsTyping} />
              <SwitchButton
                onClick={() => {
                  dispatch(setFromChainId(toChainId))
                  dispatch(setFromToken(toToken))
                  dispatch(setToChainId(fromChainId))
                  dispatch(setToToken(fromToken))
                }}
                disabled={isWalletPending}
              />
              <OutputContainer />
              <Warning />
              <BridgeMaintenanceWarningMessage />
              <BridgeExchangeRateInfo />
              <ConfirmDestinationAddressWarning />
              <BridgeTransactionButton
                isTyping={isTyping}
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
